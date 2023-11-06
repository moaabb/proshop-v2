package orderdb

import "go.uber.org/fx"

var Module = fx.Provide(NewRepository)
