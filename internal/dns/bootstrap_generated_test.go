package dns

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
