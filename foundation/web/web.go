// Package web.
package web

import (
	"context"
	"net/http"
)

type HandleFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	*http.ServeMux
}

func NewApp() *App {
	return &App{
		ServeMux: http.NewServeMux(),
	}
}

// HandleFunc IS MY OWN VERSION.
func (app *App) HandleFunc(pattern string, handler HandleFunc) {
	h := func(w http.ResponseWriter, r *http.Request) {

		// I CAN DO WHAT I WAN'T HERE

		handler(r.Context(), w, r)

		// I CAN DO WHAT I WAN'T HERE
	}

	app.ServeMux.HandleFunc(pattern, h)
}
