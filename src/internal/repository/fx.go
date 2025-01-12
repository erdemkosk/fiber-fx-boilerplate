package repository

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(
			NewFooRepository,
			fx.As(new(IFooRepository)),
		),
	),
)
