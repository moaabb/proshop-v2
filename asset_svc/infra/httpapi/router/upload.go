package router

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/moaabb/proshop-v2/asset_svc/infra/httpapi/handlers"
	"github.com/moaabb/proshop-v2/asset_svc/infra/httpapi/middleware"
	"go.uber.org/zap"
)

func Load(r *gin.Engine, uh *handlers.UploadHandler, l *zap.Logger, am *middleware.AuthMiddleware) {
	r.Use(func(c *gin.Context) {
		l.Info(fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL))
		c.Next()
	})
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://10.0.0.9:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/v1/upload", uh.UploadImage)

}
