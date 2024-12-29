package mongodb

import (
	"go.uber.org/fx"
)

var Module = fx.Module("mongodb",
	fx.Provide(NewMongoDB),
	fx.Invoke(func(lc fx.Lifecycle, db *MongoDB) {
		lc.Append(fx.Hook{
			OnStart: db.onStart,
			OnStop:  db.onStop,
		})
	}),
)
