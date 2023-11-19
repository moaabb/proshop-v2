package router

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/moaabb/ecommerce/user_svc/infra/httpapi/handlers"
	"github.com/moaabb/ecommerce/user_svc/infra/httpapi/middleware"
	"go.uber.org/zap"
)

func Load(r *gin.Engine, uh *handlers.UserHandler, l *zap.Logger, am *middleware.AuthMiddleware) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PATCH", "PUT", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(func(c *gin.Context) {
		l.Info(fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL))
		c.Next()
	})

	r.GET("/v1/users", am.Authenticate(), am.Admin(), uh.GetAll)
	r.POST("/v1/users", uh.Create)
	r.GET("/v1/users/:id", am.Authenticate(), uh.GetById)
	r.PUT("/v1/users/:id", am.Authenticate(), uh.Update)
	r.DELETE("/v1/users/:id", am.Authenticate(), am.Admin(), uh.Delete)
}
