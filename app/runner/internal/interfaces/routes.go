package interfaces

import (
	"context"
	"galileo/app/runner/internal/conf"
	"galileo/app/runner/internal/data"
	"galileo/pkg/ctxdata"
	"galileo/pkg/errResponse"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwt2 "github.com/golang-jwt/jwt/v4"
	stdHttp "net/http"
	"strings"
	"time"
)

func setHeaderInfo() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			c, _ := kgin.FromGinContext(ctx)
			if tr, ok := transport.FromServerContext(ctx); ok {
				if ht, ok := tr.(*http.Transport); ok {
					ctx = context.WithValue(ctx, "RemoteAddr", ht.Request().RemoteAddr)
				}
				auth := strings.SplitN(tr.RequestHeader().Get("Authorization"), " ", 2)
				if len(auth) != 2 || !strings.EqualFold(auth[0], "Bearer") {
					c.JSON(stdHttp.StatusUnauthorized,
						gin.H{
							"code":     stdHttp.StatusUnauthorized,
							"reason":   errResponse.ReasonUnauthorizedUser,
							"message":  errResponse.GetErrInfoByReason(errResponse.ReasonUnauthorizedUser),
							"metadata": map[string]interface{}{},
						})
					c.Abort()
					return
				}
				jwtToken := auth[1]
				token, _ := data.RedisCli.Get(ctx, "token:"+jwtToken).Result()
				if token == "" {
					c.JSON(stdHttp.StatusUnauthorized,
						gin.H{
							"code":     stdHttp.StatusUnauthorized,
							"reason":   errResponse.ReasonUnauthorizedUser,
							"message":  errResponse.GetErrInfoByReason(errResponse.ReasonUnauthorizedUser),
							"metadata": map[string]interface{}{},
						})
					c.Abort()
					return
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
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			c, _ := kgin.FromGinContext(ctx)
			claim, _ := jwt.FromContext(ctx)
			if claim == nil {
				c.JSON(stdHttp.StatusUnauthorized,
					gin.H{
						"code":     stdHttp.StatusUnauthorized,
						"reason":   errResponse.ReasonUnauthorizedUser,
						"message":  errResponse.GetErrInfoByReason(errResponse.ReasonUnauthorizedUser),
						"metadata": map[string]interface{}{},
					})
				c.Abort()
				return
			}
			claimInfo := claim.(jwt2.MapClaims)
			userId := uint32(claimInfo["ID"].(float64))
			ctx = context.WithValue(ctx, ctxdata.UserIdKey, userId)
			ctx = context.WithValue(ctx, ctxdata.Username, claimInfo["Username"])
			return handler(ctx, req)
		}
	}
}

func RegisterHTTPServer(ac *conf.Auth, us *RunnerUseCase) *gin.Engine {
	router := gin.New()

	// cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Upgrade", "Origin", "Connection", "Accept-Encoding", "Accept-Language", "Host", "x-requested-with"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(gin.Logger())

	router.Use(kgin.Middlewares(recovery.Recovery(),
		selector.Server(
			setHeaderInfo(),
			jwt.Server(
				func(token *jwt2.Token) (interface{}, error) {
					return []byte(ac.JwtKey), nil
				}, jwt.WithSigningMethod(jwt2.SigningMethodES256), jwt.WithClaims(func() jwt2.Claims {
					return jwt2.MapClaims{}
				}),
			),
			setUserInfo(),
		).Match(NewWhiteListMatcher()).Build(),
	))

	v1 := router.Group("/v1")
	{
		rootGrp := v1.Group("/api")
		{
			rootGrp.GET("/sayhi", us.SayHi)
		}
	}
	return router
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
