package asn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolvingASNToServers(t *testing.T) {
	servers, err := GetServers(24575)

	assert.Nil(t, err)

	assert.Equal(t, []string{"https://rdap.apnic.net/"}, servers)
}

func TestResolvingInvalidASNToServersReturnsAnError(t *testing.T) {
	_, err := GetServers(00000)

	assert.NotNil(t, err)
}
