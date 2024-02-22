package key

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"testing"

	"github.com/stretchr/testify/require"
	jose "gopkg.in/square/go-jose.v2"
)

func TestReadKey(t *testing.T) {
	pkcsKey := `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvW7HRieO22sL8COf9O+1
KUcoU42AL35kPHoaogqLWQoqJ32o1gRydrICCvWlUlDp6VbnOGpIbAqhk0EyOg3+
dCf2DLcCYDL1C57M8Q/iwmq13rS9Ulp5ioKK68CCV+whZOZdVvgpck35szm5fCZy
aPWMTZKfevbdDz+orjHCP0NltAaiqHw6+C2nlawb7HkjiRLU/3hQ6xXOWh7WK3jQ
WugdT+fSm9DiSK7zz4Q6M0dEtcvRhcQr/UfnOrDiq2js+QyPBjwewGAZmafytxeF
vtaqVEtf7EI7hRwgPOyiYaGAgUZDzWPwh8to4n27Ofv/S91OLyrzvZ6PXcgjQEfc
uQIDAQAB
-----END PUBLIC KEY-----	
`
	block, _ := pem.Decode([]byte(pkcsKey))
	require.NotNil(t, block)

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	require.NoError(t, err)

	rsaPub, ok := pub.(*rsa.PublicKey)
	require.True(t, ok)

	keyResponse, err := ReadKey([]byte(pkcsKey))
	require.Nil(t, err)

	expected := KeyResponse{
		Keys: []jose.JSONWebKey{
			{
				Use:       "sig",
				Key:       rsaPub,
				KeyID:     "EJxmVCQAjdwyqOeQW9X1CpKJaTqGCwqH8EVB34UjsJk",
				Algorithm: "RS256",
			},
			{
				Use:       "sig",
				Key:       rsaPub,
				KeyID:     "",
				Algorithm: "RS256",
			},
		},
	}
	require.Equal(t, expected, keyResponse)
}
