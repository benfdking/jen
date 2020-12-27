package url

import "fmt"

// urls formattable strings where to fetch jwks and and private keys respectively
const (
	jwksURL       = "https://raw.githubusercontent.com/benfdking/jen/master/keys/%s/jwks.json"
	privatePEMURL = "https://raw.githubusercontent.com/benfdking/jen/master/keys/%s/private.pem"
)

func ReturnJWKSAndPrivatePEMURL(version string) (jwks string, pem string, err error) {
	in, err := returnVersion(version)
	if err != nil {
		return "", "", nil
	}
	return fmt.Sprintf(jwksURL, in), fmt.Sprintf(privatePEMURL, in), nil
}

func returnVersion(version string) (string, error) {
	switch version {
	case a:
		return a, nil
	case b:
		return b, nil
	case c:
		return c, nil
	default:
		return "", fmt.Errorf("invalid version %s", version)
	}
}
