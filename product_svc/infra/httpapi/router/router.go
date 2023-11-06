package router

import (
	"github.com/gin-gonic/gin"
	"github.com/moaabb/ecommerce/product_svc/infra/httpapi/handlers"
	"go.uber.org/zap"
)

func NewRouter(ph *handlers.ProductHandler, z *zap.Logger) *gin.Engine {
	app := gin.New()

	Load(app, ph, z)

	return app
}
