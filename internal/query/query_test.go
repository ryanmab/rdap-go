package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// See: https://datatracker.ietf.org/doc/html/rfc7482#section-3
func TestQueriesMapToCorrectRDAPPaths(t *testing.T) {
	t.Run("IPv4", func(t *testing.T) {
		assert.Equal(t, "ip", IPv4Query.String())
	})

	t.Run("IPv6", func(t *testing.T) {
		assert.Equal(t, "ip", IPv6Query.String())
	})

	t.Run("Domain", func(t *testing.T) {
		assert.Equal(t, "domain", DomainQuery.String())
	})

	t.Run("ASN", func(t *testing.T) {
		assert.Equal(t, "autnum", AsnQuery.String())
	})
}
