package router

import (
	"github.com/gin-gonic/gin"
	"github.com/moaabb/ecommerce/order_svc/infra/httpapi/handlers"
	"go.uber.org/zap"
)

func NewRouter(oh *handlers.OrderHandler, z *zap.Logger) *gin.Engine {
	app := gin.New()

	Load(app, oh, z)

	return app
}
