// Package mux does whatever.
package mux

import (
	"github.com/ardanlabs/service/app/domain/status"
	"github.com/ardanlabs/service/app/sdk/mid"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

func WebAPI(log *logger.Logger) *web.App {
	app := web.NewApp(mid.Logger(log), mid.Errors(log))

	app.HandleFunc("GET /status", status.Status)

	return app
}
