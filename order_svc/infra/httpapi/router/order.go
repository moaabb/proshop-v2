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

func Load(r *gin.Engine, oh *handlers.OrderHandler, l *zap.Logger, am *middleware.AuthMiddleware) {
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
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/v1/orders", am.Authenticate(), oh.GetAll)
	r.GET("/v1/users/orders", am.Authenticate(), oh.GetByUserId)
	r.POST("/v1/orders", am.Authenticate(), oh.Create)
	r.GET("/v1/orders/:id", am.Authenticate(), oh.GetById)
	r.PUT("/v1/orders/:id/pay", am.Authenticate(), oh.UpdateToPaid)
	r.PUT("/v1/orders/:id/deliver", am.Authenticate(), am.Admin(), oh.UpdateToDelivered)
	r.DELETE("/v1/orders/:id", am.Authenticate(), am.Admin(), oh.Delete)
}
