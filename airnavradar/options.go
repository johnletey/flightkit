package airnavradar

import (
	"context"
	"net/http"
)

// WithAuthentication returns a ClientOption that authenticates each request by
// setting the Authorization header to "Bearer <apiKey>".
//
// https://www.airnavradar.com/api/documentation
func WithAuthentication(apiKey string) ClientOption {
	return WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+apiKey)

		return nil
	})
}
