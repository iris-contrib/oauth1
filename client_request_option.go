package oauth1

import "net/http"

// RequestOption returns the iris/x/client request option to be attached per request.
func RequestOption(config *Config, accessToken, tokenSecret string) func(*http.Request) error {
	a := newAuther(config)
	tok := NewToken(accessToken, tokenSecret)
	return func(req *http.Request) error {
		return a.setRequestAuthHeader(req, tok)
	}
}
