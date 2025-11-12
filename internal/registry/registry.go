package registry

import (
	"fmt"

	"github.com/ryanmab/rdap-go/internal/model"
	"github.com/ryanmab/rdap-go/internal/registry/internal/dns"
	"github.com/ryanmab/rdap-go/internal/registry/internal/ipv4"
	"github.com/ryanmab/rdap-go/internal/registry/internal/ipv6"
)

// GetServers returns the RDAP servers for the given query type and identifier.
func GetServers(query model.RdapQuery, identifier string) ([]string, error) {
	switch query {
	case model.DomainQuery:
		return dns.GetServers(identifier)
	case model.IPv4Query:
		return ipv4.GetServers(identifier)
	case model.IPv6Query:
		return ipv6.GetServers(identifier)
	}

	return nil, fmt.Errorf("unknown query type: %s", query)
}
