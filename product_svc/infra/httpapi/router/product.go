package router

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/moaabb/ecommerce/product_svc/infra/httpapi/handlers"
	"github.com/moaabb/ecommerce/product_svc/infra/httpapi/middleware"
	"go.uber.org/zap"
)

func Load(r *gin.Engine, ph *handlers.ProductHandler, l *zap.Logger, am *middleware.AuthMiddleware) {
	r.Use(func(c *gin.Context) {
		l.Info(fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL))
		c.Next()
	})
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PATCH", "PUT", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/v1/products", ph.GetAll)
	r.POST("/v1/products", am.Authenticate(), am.Admin(), ph.Create)
	r.GET("/v1/products/top", ph.GetTopProducts)
	r.GET("/v1/products/:id", ph.GetById)
	r.PUT("/v1/products/:id", am.Authenticate(), am.Admin(), ph.Update)
	r.DELETE("/v1/products/:id", am.Authenticate(), am.Admin(), ph.Delete)
}
