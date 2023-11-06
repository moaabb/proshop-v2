package httpapi

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2/log"
	"github.com/moaabb/ecommerce/product_svc/infra/config"
	"go.uber.org/fx"
)

func NewServer(lc fx.Lifecycle, r *gin.Engine, cfg *config.Config) *http.Server {
	srv := http.Server{
		Handler: r,
		Addr:    cfg.Port,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Info(fmt.Sprintf("Application running on port %s", strings.Split(cfg.Port, ":")[1]))
				srv.ListenAndServe()
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("stopping the server...")
			return srv.Close()
		},
	})

	return &srv
}
