package defaultjwt

import (
	"strconv"
	"time"
)

// DefaultClaims is a structure that contains the default claims as well as a Description of how the Value
// was computed
type DefaultClaims struct {
	Key         string
	Value       string
	Description string
}

// GetDefaultClaims returns a list of default claims. Default claims have been taken from rfc 7519
func GetDefaultClaims() []DefaultClaims {
	return []DefaultClaims{
		{
			Key:         "iss",
			Value:       "jen",
			Description: "sample issuer",
		},
		{
			Key:         "sub",
			Value:       "4a70bf3c-2bb9-4a64-bc4e-300d94296d23",
			Description: "constant uuid",
		},
		{
			Key:         "aud",
			Value:       "jen-users",
			Description: "sample audiance",
		},
		{
			Key:         "exp",
			Value:       strconv.Itoa(int(time.Now().Add(1 * time.Hour).Unix())),
			Description: "current time + 1 hour",
		},
		{
			Key:         "nbf",
			Value:       strconv.Itoa(int(time.Now().Unix())),
			Description: "current time",
		},
		{
			Key:         "iat",
			Value:       strconv.Itoa(int(time.Now().Unix())),
			Description: "current time",
		},
		{
			Key:         "jti",
			Value:       "83e86e9d-afbd-4b23-a7dc-e1a29cada4c2",
			Description: "constant uuid",
		},
	}
}
