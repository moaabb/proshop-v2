package main

import (
	"net/http"

	"github.com/moaabb/proshop-v2/asset_svc/infra/config"
	"github.com/moaabb/proshop-v2/asset_svc/infra/httpapi"
	"github.com/moaabb/proshop-v2/asset_svc/infra/httpapi/handlers"
	"github.com/moaabb/proshop-v2/asset_svc/infra/httpapi/middleware"
	"github.com/moaabb/proshop-v2/asset_svc/infra/httpapi/router"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	app := fx.New(
		handlers.Module,
		middleware.Module,
		router.Module,
		fx.Provide(
			zap.NewExample,
		),
		httpapi.Module,
		config.Module,
		fx.Invoke(func(*http.Server) {}),
	)
	app.Run()
}
