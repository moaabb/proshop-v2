package router

import (
	"github.com/gin-gonic/gin"
	"github.com/moaabb/ecommerce/order_svc/infra/httpapi/handlers"
	"github.com/moaabb/ecommerce/order_svc/infra/httpapi/middleware"
	"go.uber.org/zap"
)

func NewRouter(oh *handlers.OrderHandler, z *zap.Logger, am *middleware.AuthMiddleware) *gin.Engine {
	app := gin.New()

	Load(app, oh, z, am)

	return app
}
