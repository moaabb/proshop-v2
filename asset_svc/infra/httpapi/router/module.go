package router

import "go.uber.org/fx"

var Module = fx.Provide(NewRouter)
