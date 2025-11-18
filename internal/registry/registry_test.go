package registry

import (
	"testing"

	"github.com/ryanmab/rdap-go/internal/query"
	"github.com/stretchr/testify/assert"
)

func TestResolvingIpV4ToServers(t *testing.T) {
	ip := "8.8.8.8"
	servers, err := GetServers(query.IPv4Query, ip)

	assert.Nil(t, err)

	assert.Equal(t, []string{"https://rdap.arin.net/registry/", "http://rdap.arin.net/registry/"}, servers)
}

func TestResolvingTldToServers(t *testing.T) {
	domain := "com"
	servers, err := GetServers(query.DomainQuery, domain)

	assert.Nil(t, err)

	assert.Equal(t, []string{"https://rdap.verisign.com/com/v1/"}, servers)
}

func TestResolvingIPv6AddressToServers(t *testing.T) {
	ip := "2001:0db8:85a3:0000:0000:8a2e:0370:7334"

	servers, err := GetServers(query.IPv6Query, ip)

	assert.Nil(t, err)

	assert.Equal(t, []string{"https://rdap.apnic.net/"}, servers)
}

func TestResolvingASNToServers(t *testing.T) {
	asn := "394241"

	servers, err := GetServers(query.AsnQuery, asn)

	assert.Nil(t, err)

	assert.Equal(
		t,
		[]string{
			"https://rdap.arin.net/registry/",
			"http://rdap.arin.net/registry/",
		},
		servers,
	)
}
