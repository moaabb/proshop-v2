package router

import (
	"github.com/gin-gonic/gin"
	"github.com/moaabb/ecommerce/auth_svc/infra/httpapi/handlers"
	"go.uber.org/zap"
)

func NewRouter(ah *handlers.AuthHandler, z *zap.Logger) *gin.Engine {
	app := gin.New()

	Load(app, ah, z)

	return app
}
