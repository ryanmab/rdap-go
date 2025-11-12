package dns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolvingTldToServers(t *testing.T) {
	domain := "com"
	servers, err := GetServers(domain)

	assert.Nil(t, err)

	assert.Equal(t, []string{"https://rdap.verisign.com/com/v1/"}, servers)
}

func TestResolvingInvalidTldToServersReturnsAnError(t *testing.T) {
	domain := "notatld"
	_, err := GetServers(domain)

	assert.NotNil(t, err)
}
