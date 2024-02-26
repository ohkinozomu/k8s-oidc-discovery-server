package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ohkinozomu/cloudrunurlfetcher"
	"github.com/ohkinozomu/k8s-oidc-discovery-server/pkg/handlers"
	"github.com/ohkinozomu/k8s-oidc-discovery-server/pkg/key"
)

// When creating Cloud Run from the GUI, the public key may not be formatted with line breaks.
func formatPublicKey(pubKey string) string {
	const pemHeader = "-----BEGIN PUBLIC KEY-----"
	const pemFooter = "-----END PUBLIC KEY-----"

	body := strings.TrimPrefix(pubKey, pemHeader)
	body = strings.TrimSuffix(body, pemFooter)
	body = strings.ReplaceAll(body, " ", "")

	var formattedBody strings.Builder
	for i := 0; i < len(body); i += 64 {
		if i+64 < len(body) {
			formattedBody.WriteString(body[i:i+64] + "\n")
		} else {
			formattedBody.WriteString(body[i:] + "\n")
		}
	}

	return pemHeader + "\n" + formattedBody.String() + pemFooter
}

func generateKeysHandlerResponse() (key.KeyResponse, error) {
	pkcsKey := os.Getenv("PKCS_KEY")
	jwks := os.Getenv("JWKS")
	if jwks != "" && pkcsKey != "" {
		return key.KeyResponse{}, errors.New("both PKCS_KEY and JWKS are set")
	}
	if pkcsKey == "" && jwks == "" {
		return key.KeyResponse{}, errors.New("PKCS_KEY or JWKS is not set")
	}

	if jwks != "" {
		var keyResponse key.KeyResponse
		if err := json.Unmarshal([]byte(jwks), &keyResponse); err != nil {
			return key.KeyResponse{}, fmt.Errorf("failed to unmarshal JWKS: %w", err)
		}
		return keyResponse, nil
	}
	regex := regexp.MustCompile(`^-----BEGIN PUBLIC KEY-----(\s|\S)+-----END PUBLIC KEY-----$`)
	if regex.MatchString(pkcsKey) {
		pkcsKey = formatPublicKey(pkcsKey)
	}

	keyResponse, err := key.ReadKey([]byte(pkcsKey))
	if err != nil {
		return key.KeyResponse{}, nil
	}
	return keyResponse, nil
}

func main() {
	keysHandlerResponse, err := generateKeysHandlerResponse()
	if err != nil {
		panic(err)
	}

	serviceURL, err := cloudrunurlfetcher.GetServiceURL()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("keysHandlerResponse", keysHandlerResponse)
		c.Set("serviceURL", serviceURL)
		c.Next()
	})
	r.GET("/.well-known/openid-configuration", handlers.OIDCHandler)
	r.GET("/keys.json", handlers.KeysHandler)
	r.GET("/", handlers.MainHandler)
	r.Run()
}
