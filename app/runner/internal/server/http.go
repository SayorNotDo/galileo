package server

import (
	"context"
	"fmt"
	"galileo/app/runner/internal/conf"
	"galileo/app/runner/internal/data"
	"galileo/app/runner/internal/router"
	"galileo/app/runner/internal/service"
	"galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"
	"strings"
)

func customMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		if tr, ok := transport.FromServerContext(ctx); ok {
			fmt.Println("operation:", tr.Operation())
		}
		reply, err = handler(ctx, req)
		return
	}
}

func setHeaderInfo() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				if ht, ok := tr.(*http.Transport); ok {
					ctx = context.WithValue(ctx, "RemoteAddr", ht.Request().RemoteAddr)
				}
				auth := strings.SplitN(tr.RequestHeader().Get("Authorization"), " ", 2)
				if len(auth) != 2 || !strings.EqualFold(auth[0], "Bearer") {
					return nil, errResponse.SetErrByReason(errResponse.ReasonUnauthorizedUser)
				}
				jwtToken := auth[1]
				token, _ := data.RedisCli.Get(ctx, "token:"+jwtToken).Result()
				if token == "" {
					return nil, errResponse.SetErrByReason(errResponse.ReasonUnauthorizedUser)
				}
				// set Authorization header
				tr.RequestHeader().Set("Authorization", "Bearer "+token)
			}
			return handler(ctx, req)
		}
	}
}

func NewHTTPServer(c *conf.Server, ac *conf.Auth, runner *service.RunnerService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			metrics.Server(),
			validate.Validator(),
			selector.Server(
				setHeaderInfo(),
				jwt.Server(func(token *jwt2.Token) (interface{}, error) {
					return []byte(ac.JwtKey), nil
				}, jwt.WithSigningMethod(jwt2.SigningMethodHS256),
					jwt.WithClaims(func() jwt2.Claims {
						return jwt2.MapClaims{}
					})),
			).Match(NewWhiteListMatcher()).Build(),
			metadata.Server(),
		),
		http.Filter(
			handlers.CORS(
				handlers.AllowedHeaders([]string{"Authorization", "Content-Type", "Upgrade", "Origin", "Connection", "Accept-Encoding", "Accept-Language", "Host", "x-requested-with", "CurrentTeamID"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "PATCH", "DELETE", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}),
				handlers.ExposedHeaders([]string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"}),
			),
		),
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
	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", router.RegisterHTTPServer(runner))
	return srv
}

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	//whiteList["/api.core.v1.Core/Register"] = struct{}{}
	//whiteList["/api.core.v1.Core/Login"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
