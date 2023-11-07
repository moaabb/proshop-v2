package router

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/moaabb/ecommerce/order_svc/infra/httpapi/handlers"
	"github.com/moaabb/ecommerce/order_svc/infra/httpapi/middleware"
	"go.uber.org/zap"
)

func Load(r *gin.Engine, oh *handlers.OrderHandler, l *zap.Logger) {
	r.Use(func(c *gin.Context) {
		l.Info(fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL))
		c.Next()
	})

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

	r.GET("/v1/orders", middleware.Authenticate(l), oh.GetAll)
	r.GET("/v1/users/orders", middleware.Authenticate(l), oh.GetByUserId)
	r.POST("/v1/orders", middleware.Authenticate(l), oh.Create)
	r.GET("/v1/orders/:id", middleware.Authenticate(l), oh.GetById)
	r.PUT("/v1/orders/:id", middleware.Authenticate(l), oh.Update)
	r.DELETE("/v1/orders/:id", middleware.Authenticate(l), oh.Delete)
}
