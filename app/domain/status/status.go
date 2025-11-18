// Package status asdlk ja.
package status

import (
	"context"
	"encoding/json"
	"net/http"
)

func Status(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return json.NewEncoder(w).Encode(status)
}
