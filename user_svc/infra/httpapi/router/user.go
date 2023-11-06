package router

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/moaabb/ecommerce/user_svc/infra/httpapi/handlers"
	"go.uber.org/zap"
)

func Load(r *gin.Engine, uh *handlers.UserHandler, l *zap.Logger) {
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

	r.GET("/v1/users", uh.GetAll)
	r.POST("/v1/users", uh.Create)
	r.GET("/v1/users/:id", uh.GetById)
	r.PUT("/v1/users/:id", uh.Update)
	r.DELETE("/v1/users/:id", uh.Delete)
}
