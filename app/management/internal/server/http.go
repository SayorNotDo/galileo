package server

import (
	"context"
	projectV1 "galileo/api/management/project/v1"
	taskV1 "galileo/api/management/task/v1"
	testCaseV1 "galileo/api/management/testcase/v1"
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
	"github.com/go-kratos/kratos/v2/transport"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

const (
	tokenKey = "token"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, project *service.ProjectService, testcase *service.TestcaseService, task *service.TaskService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			logging.Server(logger),
			recovery.Recovery(),
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
			metadata.Server(),
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
	taskV1.RegisterTaskHTTPServer(srv, task)
	projectV1.RegisterProjectHTTPServer(srv, project)
	testCaseV1.RegisterTestcaseHTTPServer(srv, testcase)
	return srv
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
			ctx = context.WithValue(ctx, ctxdata.UserIdKey, userId)
			ctx = context.WithValue(ctx, ctxdata.Username, claimInfo["Username"])
			return handler(ctx, req)
		}
	}
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
