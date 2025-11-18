// Package mux does whatever.
package mux

import (
	"github.com/ardanlabs/service/app/domain/status"
	"github.com/ardanlabs/service/foundation/web"
)

func WebAPI() *web.App {
	app := web.NewApp()

	app.HandleFunc("GET /status", status.Status)

	return app
}
