package main

import (
	"net/http"

	"github.com/moaabb/ecommerce/order_svc/infra/config"
	"github.com/moaabb/ecommerce/order_svc/infra/database"
	"github.com/moaabb/ecommerce/order_svc/infra/database/orderdb"
	"github.com/moaabb/ecommerce/order_svc/infra/httpapi"
	"github.com/moaabb/ecommerce/order_svc/infra/httpapi/handlers"
	"github.com/moaabb/ecommerce/order_svc/infra/httpapi/router"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	app := fx.New(
		httpapi.Module,
		database.Module,
		orderdb.Module,
		config.Module,
		router.Module,
		handlers.Module,
		fx.Provide(
			zap.NewExample,
		),
		fx.Invoke(func(*http.Server) {}),
	)
	app.Run()
}
