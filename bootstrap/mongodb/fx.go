package mongodb

import (
	"context"

	"go.uber.org/fx"
)

var Module = fx.Module("mongodb",
	fx.Provide(NewMongoDB),
	fx.Invoke(func(lc fx.Lifecycle, db *MongoDB) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				return db.onStart()
			},
			OnStop: func(ctx context.Context) error {
				return db.onStop(ctx)
			},
		})
	}),
)
