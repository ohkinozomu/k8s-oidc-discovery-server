package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

func OIDCHandler(c *gin.Context) {
	if serviceURL, exists := c.Get("serviceURL"); exists {
		conf := oidc.DiscoveryConfiguration{
			Issuer:                 serviceURL.(string),
			JwksURI:                serviceURL.(string) + "/keys.json",
			AuthorizationEndpoint:  "urn:kubernetes:programmatic_authorization",
			ResponseTypesSupported: []string{"id_token"},
			SubjectTypesSupported:  []string{"public"},
			IDTokenSigningAlgValuesSupported: []string{
				"RS256",
			},
			ClaimsSupported: []string{
				"sub",
				"iss",
			},
		}
		c.JSON(http.StatusOK, conf)
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
}
