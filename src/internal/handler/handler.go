package handler

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewFooHandler),
	fx.Provide(NewHealthHandler),
)
