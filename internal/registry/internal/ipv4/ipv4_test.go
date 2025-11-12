package ipv4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolvingIpV4ToServers(t *testing.T) {
	ip := "8.8.8.8"
	servers, err := GetServers(ip)

	assert.Nil(t, err)

	assert.Equal(t, []string{"https://rdap.arin.net/registry/", "http://rdap.arin.net/registry/"}, servers)
}

func TestResolvingInvalidIpV4ToServersReturnsAnError(t *testing.T) {
	t.Run("greater than 255 octet", func(t *testing.T) {
		ipv4 := "256.256.256.256"
		_, err := GetServers(ipv4)

		assert.NotNil(t, err)
	})

	t.Run("smaller than 0", func(t *testing.T) {
		ipv4 := "0.8.8.8"
		_, err := GetServers(ipv4)

		assert.NotNil(t, err)
	})

}

func TestResolvingNonIpV4ToServersReturnsAnError(t *testing.T) {
	ipv4 := "clearly-not-an-ipv4"
	_, err := GetServers(ipv4)

	assert.NotNil(t, err)
}
