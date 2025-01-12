package handler

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(
			NewFooHandler,
			fx.As(new(IFooHandler)),
		),
		fx.Annotate(
			NewHealthHandler,
			fx.As(new(IHealthHandler)),
		),
	),
)
