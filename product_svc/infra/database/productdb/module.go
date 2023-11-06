package productdb

import "go.uber.org/fx"

var Module = fx.Provide(NewRepository)
