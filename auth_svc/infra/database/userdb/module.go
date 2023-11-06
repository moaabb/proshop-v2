package userdb

import "go.uber.org/fx"

var Module = fx.Provide(NewRepository)
