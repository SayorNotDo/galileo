package server

import (
	"context"
	managementV1 "galileo/api/management/v1"
	"galileo/app/management/internal/conf"
	"galileo/app/management/internal/data"
	"galileo/app/management/internal/service"
	"galileo/pkg/ctxdata"
	"galileo/pkg/errResponse"
	"galileo/pkg/responseEncoder"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

const (
	tokenKey = "token"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(tr *conf.Trace, c *conf.Server, ac *conf.Auth, management *service.ManagementService, logger log.Logger) *http.Server {
	err := initTracer(tr.Endpoint)
	if err != nil {
		panic(err)
	}
	///* 初始化 WebSocket 连接池 */
	pool := NewConnectionPool(1000)
	///* 注入连接池到 WebSocketHandler */
	webSocketHandler := NewWebSocketHandler(logger, pool)
	var opts = []http.ServerOption{
		http.Middleware(
			logging.Server(logger),
			recovery.Recovery(),
			validate.Validator(),
			tracing.Server(),
			selector.Server(
				// set header information
				setHeaderInfo(),
				jwt.Server(func(token *jwt2.Token) (interface{}, error) {
					return []byte(ac.JwtKey), nil
				}, jwt.WithSigningMethod(jwt2.SigningMethodHS256),
					jwt.WithClaims(func() jwt2.Claims {
						return jwt2.MapClaims{}
					})),
				// set global ctx
				setUserInfo(),
			).Match(NewWhiteListMatcher()).Build(),
			metadata.Server(
				metadata.WithPropagatedPrefix("x-md-global", "userId"),
				metadata.WithPropagatedPrefix("x-md-global", "username"),
			),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Accept"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "PATCH", "DELETE", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	// add success custom json response
	opts = append(opts, http.ResponseEncoder(responseEncoder.ResponseEncoder))
	// add error custom json response
	opts = append(opts, http.ErrorEncoder(responseEncoder.ErrorEncoder))
	srv := http.NewServer(opts...)
	route := srv.Route("/")
	route.POST("v1/api/management/testcase/upload", management.UploadTestcaseFile)
	route.POST("v1/api/management/interface/upload", management.UploadApiFile)
	srv.HandlePrefix("/ws", webSocketHandler)
	managementV1.RegisterManagementHTTPServer(srv, management)
	return srv
}

// 设置全局trace
func initTracer(url string) error {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return err
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String("galileo-trace"),
			attribute.String("exporter", "jaeger"),
			attribute.Float64("float", 312.23),
		)),
	)
	otel.SetTracerProvider(tp)
	return nil
}

func setHeaderInfo() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				if ht, ok := tr.(*http.Transport); ok {
					ctx = context.WithValue(ctx, "RemoteAddr", ht.Request().RemoteAddr)
				}
				auth := strings.SplitN(tr.RequestHeader().Get(ctxdata.AuthorizationKey), " ", 2)
				if len(auth) != 2 || !strings.EqualFold(auth[0], ctxdata.AuthorizationType) {
					return nil, errResponse.SetErrByReason(errResponse.ReasonUnauthorizedUser)
				}
				jwtToken := auth[1]
				token, _ := data.RedisCli.Get(ctx, tokenKey+":"+jwtToken).Result()
				if token == "" {
					return nil, errResponse.SetErrByReason(errResponse.ReasonUnauthorizedUser)
				}
				// set Authorization header
				tr.RequestHeader().Set(ctxdata.AuthorizationKey, ctxdata.AuthorizationType+" "+token)
			}
			return handler(ctx, req)
		}
	}
}

func setUserInfo() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			claim, _ := jwt.FromContext(ctx)
			if claim == nil {
				return nil, errResponse.SetErrByReason(errResponse.ReasonUnauthorizedInfoMissing)
			}
			claimInfo := claim.(jwt2.MapClaims)
			userId := uint32(claimInfo["ID"].(float64))
			username := claimInfo["Username"].(string)
			ctx = context.WithValue(ctx, ctxdata.UserIdKey, userId)
			ctx = context.WithValue(ctx, ctxdata.Username, username)
			return handler(ctx, req)
		}
	}
}

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/api.task.v1.Task/Test"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
