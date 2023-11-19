package router

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/moaabb/ecommerce/auth_svc/infra/httpapi/handlers"
	"go.uber.org/zap"
)

func Load(r *gin.Engine, ah *handlers.AuthHandler, l *zap.Logger) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PATCH", "PUT", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://172.21.193.94:8080"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.Use(func(c *gin.Context) {
		l.Info(fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL))
		c.Next()
	})

	r.POST("/v1/auth/login", ah.Login)
	r.POST("/v1/auth/logout", ah.Logout)
	r.POST("/v1/auth", ah.ValidateRequest)
	r.GET("/v1/config/paypal", ah.GetPaypalConfig)

}
