// Package status asdlk ja.
package status

import (
	"context"
	"math/rand/v2"
	"net/http"

	"github.com/ardanlabs/service/app/sdk/errs"
	"github.com/ardanlabs/service/foundation/web"
)

func Status(ctx context.Context, r *http.Request) web.Encoder {
	// DECODE AND VALIDATE INPUT
	// CALL INTO THE BUSINESS LAYER
	// RETURN ERROR OR AN OBJECT FOR 200

	if n := rand.IntN(100); n%2 == 0 {
		return errs.Newf(errs.InvalidArgument, "error example: %s", "NOT GREAT")
	}

	status := status{
		Status: "OK",
	}

	return status
}
