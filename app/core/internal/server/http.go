package server

import (
	"context"
	"encoding/json"
	v1 "galileo/api/core/v1"
	"galileo/app/core/internal/conf"
	"galileo/app/core/internal/data"
	"galileo/app/core/internal/pkg/ctxdata"
	"galileo/app/core/internal/service"
	"galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"
	stdHttp "net/http"
	"strings"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, s *service.CoreService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			tracing.Server(),
			// logging record
			logging.Server(logger),
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
	opts = append(opts, http.ResponseEncoder(responseEncoder))
	// add error custom json response
	opts = append(opts, http.ErrorEncoder(errorEncoder))
	srv := http.NewServer(opts...)
	v1.RegisterCoreHTTPServer(srv, s)
	return srv
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func responseEncoder(w stdHttp.ResponseWriter, r *stdHttp.Request, v interface{}) error {
	reply := &Response{}
	reply.Code = 20000
	reply.Message = "success"

	codec, _ := http.CodecForRequest(r, "Accept")
	marshalData, err := codec.Marshal(v)
	_ = json.Unmarshal(marshalData, &reply.Data)
	if err != nil {
		return err
	}
	resp, err := codec.Marshal(reply)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", codec.Name())
	w.WriteHeader(stdHttp.StatusOK)
	_, _ = w.Write(resp)
	return nil
}

func errorEncoder(w stdHttp.ResponseWriter, r *stdHttp.Request, err error) {
	codec, _ := http.CodecForRequest(r, "Accept")
	w.Header().Set("Content-Type", "application/"+codec.Name())
	// status code set 200
	err = errResponse.SetCustomizeErrInfo(err)
	se := errors.FromError(err)
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	_, _ = w.Write(body)
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
	whiteList["/api.core.v1.Core/Register"] = struct{}{}
	whiteList["/api.core.v1.Core/Login"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
