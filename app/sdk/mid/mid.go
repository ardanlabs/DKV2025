// Package mid
package mid

import (
	"github.com/ardanlabs/service/foundation/web"
)

func checkIsError(e web.Encoder) error {
	err, hasError := e.(error)
	if hasError {
		return err
	}

	return nil
}
