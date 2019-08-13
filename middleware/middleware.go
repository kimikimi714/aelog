package middleware

import (
	"net/http"

	"github.com/emahiro/ae-plain-logger/spancontext"
)

// MwAEPlainLogger is middleware for setting stackdrvier logging.
// In this log middleware, label is required.
// If you don't set label, this log middleware return panic.
func MwAEPlainLogger(label string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if label == "" {
				panic("label is required")
			}

			ctx, done := spancontext.Set(r, label)
			defer done()
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
