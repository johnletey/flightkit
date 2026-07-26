package flightaware

import (
	"context"
	"net/http"
)

// WithAuthentication returns a ClientOption that authenticates each request by
// setting the x-apikey header to the given API key.
//
// https://www.flightaware.com/aeroapi/portal/documentation#auth
func WithAuthentication(apiKey string) ClientOption {
	return WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("x-apikey", apiKey)

		return nil
	})
}
