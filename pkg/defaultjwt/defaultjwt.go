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

// AddDefaultsToToken adds default properties to jwt
func AddDefaultsToToken(token *jwt.Token) *jwt.Token {
	token.Set(jwt.IssuerKey, "https://github.com/benfking/jen")
	token.Set(jwt.IssuedAtKey, time.Now())
	token.Set(jwt.SubjectKey, defaultSubject)
	token.Set(jwt.AudienceKey, "AuthenticationGurus")
	token.Set(jwt.IssuerKey, "https://github.com/benfking/jen ")
	token.Set(jwt.ExpirationKey, time.Now().Add(1*time.Hour).Unix())
	token.Set(jwt.NotBeforeKey, time.Now())
	return token
}

// addMapToToken adds all the properties of a map to the token
func addMapToToken(token *jwt.Token, values map[string]string) *jwt.Token {
	for key, value := range values {
		token.Set(key, value)
	}
	return token
}

// AddJSONFileToToken adds content of json file to token
func AddJSONFileToToken(token *jwt.Token, filePath string) (*jwt.Token, error) {
	file, err := os.Open(filePath)
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
	token = addMapToToken(token, values)
	return token, nil
}
