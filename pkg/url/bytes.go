package url

import (
	"context"
	"crypto"
	"encoding/json"
	"github.com/benfking/jen/pkg/rsapem"
	"github.com/lestrrat-go/jwx/jwk"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"net/http"
	"time"
)

func ReturnJWKSAndPrivateKey(version string) (jwk.Set, crypto.PrivateKey, error) {
	jwksURL, cryptoURL, err := ReturnJWKSAndPrivatePEMURL(version)
	if err != nil {
		return jwk.Set{}, nil, err
	}

	g, _ := errgroup.WithContext(context.Background())
	client := &http.Client{Timeout: 5 * time.Second}

	var set jwk.Set
	g.Go(func() error {
		response, err := client.Get(jwksURL)
		if err != nil {
			return err
		}
		defer response.Body.Close()
		return json.NewDecoder(response.Body).Decode(&set)
	})

	var private crypto.PrivateKey
	g.Go(func() error {
		response, err := client.Get(cryptoURL)
		if err != nil {
			return err
		}
		defer response.Body.Close()
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		private, err = rsapem.RSAPrivateFromPen(data)
		if err != nil {
			return err
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return jwk.Set{}, nil, err
	}
	return set, private, nil
}
