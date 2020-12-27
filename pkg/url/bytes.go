package url

import (
	"context"
	"crypto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/benfdking/jen/pkg/keys"
	"github.com/benfdking/jen/pkg/rsapem"
	"github.com/lestrrat-go/jwx/jwk"
	"golang.org/x/sync/errgroup"
)

func ReturnJWKSAndPrivateKeyFromData(version string) (jwk.Set, crypto.PrivateKey, error) {
	if !IsVersion(version) {
		return jwk.Set{}, nil, fmt.Errorf("version %s is invalid", version)
	}

	g, _ := errgroup.WithContext(context.Background())
	var set jwk.Set
	g.Go(func() error {
		data, err := keys.Asset(version + "/jwks.json")
		if err != nil {
			return err
		}
		return json.Unmarshal(data, &set)
	})

	var private crypto.PrivateKey
	g.Go(func() error {
		data, err := keys.Asset(version + "/private.pem")
		if err != nil {
			return err
		}
		private, err = rsapem.RSAPrivateFromPen(data)
		if err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return jwk.Set{}, nil, err
	}
	return set, private, nil
}

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
