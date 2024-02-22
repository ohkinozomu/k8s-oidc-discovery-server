// originally copied from https://github.com/aws/amazon-eks-pod-identity-webhook/blob/d4269f48c6d6427f83e31ca52957b8c67d3c2fcf/hack/self-hosted/main.go
package key

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	jose "github.com/go-jose/go-jose/v3"
	"github.com/pkg/errors"
)

// originally copied from kubernetes/kubernetes#78502
func keyIDFromPublicKey(publicKey interface{}) (string, error) {
	publicKeyDERBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", fmt.Errorf("failed to serialize public key to DER format: %v", err)
	}

	hasher := crypto.SHA256.New()
	hasher.Write(publicKeyDERBytes)
	publicKeyDERHash := hasher.Sum(nil)

	keyID := base64.RawURLEncoding.EncodeToString(publicKeyDERHash)

	return keyID, nil
}

type KeyResponse struct {
	Keys []jose.JSONWebKey `json:"keys"`
}

func ReadKey(content []byte) (KeyResponse, error) {
	block, _ := pem.Decode(content)
	if block == nil {
		return KeyResponse{}, errors.Errorf("Error decoding PEM")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return KeyResponse{}, errors.Wrapf(err, "Error parsing key content")
	}
	switch pubKey.(type) {
	case *rsa.PublicKey:
	default:
		return KeyResponse{}, errors.New("Public key was not RSA")
	}

	var alg jose.SignatureAlgorithm
	switch pubKey.(type) {
	case *rsa.PublicKey:
		alg = jose.RS256
	default:
		return KeyResponse{}, fmt.Errorf("invalid public key type %T, must be *rsa.PrivateKey", pubKey)
	}

	kid, err := keyIDFromPublicKey(pubKey)
	if err != nil {
		return KeyResponse{}, err
	}

	var keys []jose.JSONWebKey
	keys = append(keys, jose.JSONWebKey{
		Key:       pubKey,
		KeyID:     kid,
		Algorithm: string(alg),
		Use:       "sig",
	})
	// jq '.keys += [.keys[0]] | .keys[1].kid = ""'
	keys = append(keys, jose.JSONWebKey{
		Key:       pubKey,
		KeyID:     "",
		Algorithm: string(alg),
		Use:       "sig",
	})

	keyResponse := KeyResponse{Keys: keys}
	return keyResponse, nil
}
