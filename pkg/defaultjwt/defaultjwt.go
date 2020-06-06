package defaultjwt

import (
	"crypto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

const defaultSubject = "fa645455-f280-4de4-b010-bf7e28310c66"

// NewToken returns a new token
func NewToken() *jwt.Token {
	return jwt.New()
}

// SignToken signs a token and returns a string
func SignToken(token *jwt.Token, key crypto.PrivateKey) (string, error) {
	output, err := token.Sign(jwa.RS256, key)
	if err != nil {
		return "", fmt.Errorf("failed to generate signed payload: %w", err)
	}
	return string(output), nil
}

// AddDefaultClaims adds default claims to jwt
func AddDefaultClaims(token *jwt.Token) (*jwt.Token, error) {
	cs := map[string]interface{}{
		jwt.IssuerKey:     "https://github.com/benfdking/jen",
		jwt.IssuedAtKey:   time.Now(),
		jwt.SubjectKey:    defaultSubject,
		jwt.AudienceKey:   "AuthenticationGurus",
		jwt.ExpirationKey: time.Now().Add(1 * time.Hour).Unix(),
		jwt.NotBeforeKey:  time.Now(),
	}
	for k, v := range cs {
		err := token.Set(k, v)
		if err != nil {
			return nil, err
		}
	}
	return token, nil
}

// addMapClaimsToToken adds all the properties of a map to the token
func addMapClaimsToToken(t *jwt.Token, cs map[string]string) (*jwt.Token, error) {
	for k, v := range cs {
		err := t.Set(k, v)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

// AddJSONFileClaimsToToken adds content of json file to token
func AddJSONFileClaimsToToken(t *jwt.Token, path string) (*jwt.Token, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var values map[string]string
	err = json.Unmarshal(bytes, &values)
	if err != nil {
		return nil, err
	}
	return addMapClaimsToToken(t, values)
}
