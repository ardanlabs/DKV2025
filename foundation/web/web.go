// Package web.
package web

import (
	"context"
	"net/http"
)

type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	*http.ServeMux
	mw []MidFunc
}

func NewApp(mw ...MidFunc) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		mw:       mw,
	}
}

// HandleFunc IS MY OWN VERSION.
func (app *App) HandleFunc(pattern string, handler HandlerFunc, mw ...MidFunc) {
	handler = wrapMiddleware(app.mw, handler)
	handler = wrapMiddleware(mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {

		// I CAN DO WHAT I WAN'T HERE

		handler(r.Context(), w, r)

		// I CAN DO WHAT I WAN'T HERE
	}

	app.ServeMux.HandleFunc(pattern, h)
}
