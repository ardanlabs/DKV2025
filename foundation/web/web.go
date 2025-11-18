// Package web.
package web

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type Encoder interface {
	Encode() (data []byte, contentType string, err error)
}

type HandlerFunc func(ctx context.Context, r *http.Request) Encoder

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
		ctx := setTraceID(r.Context(), uuid.New())
		ctx = setWriter(ctx, w)

		resp := handler(ctx, r)
		if err := Respond(ctx, w, resp); err != nil {
			return
		}
	}

	app.ServeMux.HandleFunc(pattern, h)
}
