package ipv6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolvingIPv6AddressToServers(t *testing.T) {
	ip := "2001:0db8:85a3:0000:0000:8a2e:0370:7334"

	servers, err := GetServers(ip)

	assert.Nil(t, err)

	assert.Equal(t, []string{"https://rdap.apnic.net/"}, servers)
}

func TestResolvingInvalidIpV6ReturnsAnError(t *testing.T) {
	ip := "256.256.256.256"

	_, err := GetServers(ip)

	assert.NotNil(t, err)
}

func TestResolvingIpV6ReturnsServersFromMostSpecificRange(t *testing.T) {
	ip := "2001:4c00:0000:0000:0000:0000:0000:0001"

	servers, err := GetServers(ip)

	assert.Nil(t, err)

	assert.Equal(t, []string{"https://rdap.db.ripe.net/"}, servers)

}
