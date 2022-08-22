package initialize

import (
	"go.uber.org/fx"
)

// Module provides common application dependencies
var Module = fx.Options(
	fx.Provide(
		// newLogger,
		InitLogrus,
	),
	fx.Invoke(),
)
