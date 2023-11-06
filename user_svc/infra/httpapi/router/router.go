package router

import (
	"github.com/gin-gonic/gin"
	"github.com/moaabb/ecommerce/user_svc/infra/httpapi/handlers"
	"go.uber.org/zap"
)

func NewRouter(uh *handlers.UserHandler, z *zap.Logger) *gin.Engine {
	app := gin.New()

	Load(app, uh, z)

	return app
}
