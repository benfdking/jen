package jwks

import "github.com/lestrrat-go/jwx/jwk"

type Set struct {
	Keys []jwk.Key `json:"keys"`
}
