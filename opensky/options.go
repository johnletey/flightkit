package opensky

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const tokenURL = "https://auth.opensky-network.org/auth/realms/opensky-network/protocol/openid-connect/token"

// WithAuthentication returns a ClientOption that authenticates each request by
// setting the Authorization header to "Bearer <token>". It exchanges the given
// API client credentials for a token, caches it, and fetches a new one shortly
// before the 30-minute expiry.
//
// https://openskynetwork.github.io/opensky-api/rest.html
func WithAuthentication(clientID, clientSecret string) ClientOption {
	source := &tokenSource{clientID: clientID, clientSecret: clientSecret}

	return WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		token, err := source.token(ctx)
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", "Bearer "+token)

		return nil
	})
}

type tokenSource struct {
	clientID     string
	clientSecret string

	mu     sync.Mutex
	cached string
	expiry time.Time
}

func (s *tokenSource) token(ctx context.Context) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cached != "" && time.Now().Before(s.expiry) {
		return s.cached, nil
	}

	form := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {s.clientID},
		"client_secret": {s.clientSecret},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("opensky: token request returned %s", resp.Status)
	}

	var body struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return "", err
	}

	// Renew a minute early so we never send an almost-expired token.
	s.cached = body.AccessToken
	s.expiry = time.Now().Add(time.Duration(body.ExpiresIn)*time.Second - time.Minute)

	return s.cached, nil
}
