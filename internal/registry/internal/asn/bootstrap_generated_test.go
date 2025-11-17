package asn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratedBootstrapHasAllServersEndingWithTailingSlash(t *testing.T) {
	for _, servers := range Bootstrap {
		assert.NotEmpty(t, servers, "Generated bootstrap has an empty server list for one of the IPs")

		for _, server := range servers {
			assert.Equal(t, "/", string(server[len(server)-1]), "Generated bootstrap server %q does not end with a trailing slash", server)
		}
	}
}

func TestGeneratedBootstrapASNRangeIsValid(t *testing.T) {
	for asn, servers := range Bootstrap {
		start, end := asn[0], asn[1]

		assert.GreaterOrEqualf(t, end, start, "Generated bootstrap has an ASN range which ends (%d) before it start (%d)", end, start)

		assert.NotEmpty(t, servers, "Generated bootstrap has an empty server list for ASN %d", asn)
	}
}
