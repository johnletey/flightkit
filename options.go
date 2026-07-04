package fr24api

import (
	"context"
	"net/http"
)

// WithAuthentication returns a ClientOption that sets the Authorization
// header to "Bearer <apiKey>" on each request.
func WithAuthentication(apiKey string) ClientOption {
	return WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+apiKey)

		return nil
	})
}
