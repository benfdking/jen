package defaultjwt

import (
	"crypto"
	"fmt"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"time"
)

const defaultSubject = "fa645455-f280-4de4-b010-bf7e28310c66"

// Default returns a sample key
func Default(key crypto.PrivateKey) (string, error) {
	token := jwt.New()
	token.Set(jwt.IssuerKey, "https://github.com/benfking/jen")
	token.Set(jwt.SubjectKey, defaultSubject)
	token.Set(jwt.AudienceKey, "AuthenticationGurus")
	token.Set(jwt.IssuerKey, "https://github.com/benfking/jen ")
	token.Set(jwt.ExpirationKey, time.Now().Add(1 * time.Hour).Unix())
	token.Set(jwt.NotBeforeKey, time.Now())
	token.Set(jwt.IssuedAtKey, time.Now())
	output, err := token.Sign(jwa.RS256, key)
	if err != nil {
		return "", fmt.Errorf("failed to generate signed payload: %w\n", err)
	}
	return string(output), nil
}

