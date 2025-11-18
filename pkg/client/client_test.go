package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookingUpDomain(t *testing.T) {
	client := New()

	t.Run("Lowercase domain", func(t *testing.T) {
		client.ClearCache()

		response, err := client.LookupDomain("ryanmaber.com")

		assert.NoError(t, err)

		assert.NotNil(t, response)

		assert.Equal(t, "RYANMABER.COM", response.LdhName)
	})

	t.Run("Uppercase domain", func(t *testing.T) {
		client.ClearCache()

		response, err := client.LookupDomain("RYANMABER.COM")

		assert.NoError(t, err)

		assert.NotNil(t, response)

		assert.Equal(t, "RYANMABER.COM", response.LdhName)
	})

	t.Run("Mixed case domain", func(t *testing.T) {
		client.ClearCache()

		response, err := client.LookupDomain("rYanMaBeR.cOm")

		assert.NoError(t, err)

		assert.NotNil(t, response)

		assert.Equal(t, "RYANMABER.COM", response.LdhName)
	})

	t.Run("Fully qualified URL", func(t *testing.T) {
		client.ClearCache()

		response, err := client.LookupDomain("    https://rYanMaBeR.cOm/  ")

		assert.NoError(t, err)

		assert.NotNil(t, response)

		assert.Equal(t, "RYANMABER.COM", response.LdhName)
	})

	t.Run("Domain with port", func(t *testing.T) {
		client.ClearCache()

		response, err := client.LookupDomain("ryanmaber.com:8080")

		assert.NoError(t, err)

		assert.NotNil(t, response)

		assert.Equal(t, "RYANMABER.COM", response.LdhName)
	})
}

func TestLookingUpIpV4(t *testing.T) {
	client := New()

	t.Run("Valid IPv4", func(t *testing.T) {
		client.ClearCache()

		response, err := client.LookupIPv4("8.8.8.8")

		assert.NoError(t, err)

		assert.NotNil(t, response)

		assert.Equal(t, "GOGL", response.Name)
	})

	t.Run("Valid IpV4 with whitespace", func(t *testing.T) {
		client.ClearCache()

		response, err := client.LookupIPv4("   8.8.8.8  ")

		assert.NoError(t, err)

		assert.NotNil(t, response)

		assert.Equal(t, "GOGL", response.Name)
	})
}

func TestLookingUpIpV6(t *testing.T) {
	client := New()

	t.Run("Valid IPv4", func(t *testing.T) {
		client.ClearCache()

		response, err := client.LookupIPv6("2001:4860:4860::6464")

		assert.NoError(t, err)

		assert.NotNil(t, response)

		assert.Equal(t, "GOOGLE-IPV6", response.Name)
	})

	t.Run("Valid IpV4 with whitespace", func(t *testing.T) {
		client.ClearCache()

		response, err := client.LookupIPv6(" 2001:4860:4860::6464 ")

		assert.NoError(t, err)

		assert.NotNil(t, response)

		assert.Equal(t, "GOOGLE-IPV6", response.Name)
	})
}

func TestLookingUpASN(t *testing.T) {
	client := New()

	t.Run("Valid ASN", func(t *testing.T) {
		client.ClearCache()

		response, err := client.LookupASN(63489)

		assert.NoError(t, err)

		assert.NotNil(t, response)

		assert.Equal(t, "IDNIC-BPRTIK-AS-ID", response.Name)
	})

	t.Run("Invalid ASN", func(t *testing.T) {
		client.ClearCache()

		response, err := client.LookupASN(000000)

		assert.Error(t, err)

		assert.Nil(t, response)
	})

}
