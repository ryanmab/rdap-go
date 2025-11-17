package asn

import "fmt"

// GetServers returns the RDAP servers for a given ASN from the IANA bootstrap data.
//
// See: https://data.iana.org/rdap/
func GetServers(asn uint32) ([]string, error) {
	for asnRange, servers := range Bootstrap {
		if asn >= asnRange[0] && asn <= asnRange[1] {
			return servers, nil
		}
	}

	return nil, fmt.Errorf("no RDAP servers found for TLD: %d", asn)
}
