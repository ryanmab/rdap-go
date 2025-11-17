package registry

import (
	"fmt"
	"strconv"

	"github.com/ryanmab/rdap-go/internal/query"
	"github.com/ryanmab/rdap-go/internal/registry/internal/asn"
	"github.com/ryanmab/rdap-go/internal/registry/internal/dns"
	"github.com/ryanmab/rdap-go/internal/registry/internal/ipv4"
	"github.com/ryanmab/rdap-go/internal/registry/internal/ipv6"
)

// GetServers returns the RDAP servers for the given query type and identifier.
func GetServers(queryType query.RdapQuery, identifier string) ([]string, error) {
	switch queryType {
	case query.DomainQuery:
		return dns.GetServers(identifier)
	case query.IPv4Query:
		return ipv4.GetServers(identifier)
	case query.IPv6Query:
		return ipv6.GetServers(identifier)
	case query.AsnQuery:
		identifierAsInt, err := strconv.ParseUint(identifier, 10, 32)

		if err != nil {
			return nil, fmt.Errorf("expected autnum to be a 32 bit unsigned int: %s", identifier)
		}

		return asn.GetServers(uint32(identifierAsInt))
	}

	return nil, fmt.Errorf("unknown query type: %s", queryType)
}
