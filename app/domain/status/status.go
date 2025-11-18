// Package status asdlk ja.
package status

import (
	"context"
	"net/http"

	"github.com/ardanlabs/service/foundation/web"
)

func Status(ctx context.Context, r *http.Request) web.Encoder {
	// DECODE AND VALIDATE INPUT
	// CALL INTO THE BUSINESS LAYER
	// RETURN ERROR OR AN OBJECT FOR 200

	status := status{
		Status: "OK",
	}

	return status
}
