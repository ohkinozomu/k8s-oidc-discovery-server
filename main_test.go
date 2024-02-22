package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatPublicKey(t *testing.T) {
	pkcsKey := "-----BEGIN PUBLIC KEY-----MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvW7HRieO22sL8COf9O+1KUcoU42AL35kPHoaogqLWQoqJ32o1gRydrICCvWlUlDp6VbnOGpIbAqhk0EyOg3+dCf2DLcCYDL1C57M8Q/iwmq13rS9Ulp5ioKK68CCV+whZOZdVvgpck35szm5fCZyaPWMTZKfevbdDz+orjHCP0NltAaiqHw6+C2nlawb7HkjiRLU/3hQ6xXOWh7WK3jQWugdT+fSm9DiSK7zz4Q6M0dEtcvRhcQr/UfnOrDiq2js+QyPBjwewGAZmafytxeFvtaqVEtf7EI7hRwgPOyiYaGAgUZDzWPwh8to4n27Ofv/S91OLyrzvZ6PXcgjQEfcuQIDAQAB-----END PUBLIC KEY-----"

	expected := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvW7HRieO22sL8COf9O+1
KUcoU42AL35kPHoaogqLWQoqJ32o1gRydrICCvWlUlDp6VbnOGpIbAqhk0EyOg3+
dCf2DLcCYDL1C57M8Q/iwmq13rS9Ulp5ioKK68CCV+whZOZdVvgpck35szm5fCZy
aPWMTZKfevbdDz+orjHCP0NltAaiqHw6+C2nlawb7HkjiRLU/3hQ6xXOWh7WK3jQ
WugdT+fSm9DiSK7zz4Q6M0dEtcvRhcQr/UfnOrDiq2js+QyPBjwewGAZmafytxeF
vtaqVEtf7EI7hRwgPOyiYaGAgUZDzWPwh8to4n27Ofv/S91OLyrzvZ6PXcgjQEfc
uQIDAQAB
-----END PUBLIC KEY-----`

	require.Equal(t, expected, formatPublicKey(pkcsKey))
}
