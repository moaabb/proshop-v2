package router

import (
	"github.com/gin-gonic/gin"
	"github.com/moaabb/ecommerce/product_svc/infra/httpapi/handlers"
	"github.com/moaabb/ecommerce/product_svc/infra/httpapi/middleware"
	"go.uber.org/zap"
)

func NewRouter(ph *handlers.ProductHandler, z *zap.Logger, am *middleware.AuthMiddleware) *gin.Engine {
	app := gin.New()

	Load(app, ph, z, am)

	return app
}
