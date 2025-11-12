package dns

import (
	"fmt"
)

// GetServers returns the RDAP servers for a given TLD from the IANA bootstrap data.
//
// See: https://data.iana.org/rdap/
func GetServers(tld string) ([]string, error) {
	if servers, ok := Bootstrap[tld]; ok {
		return servers, nil
	}

	return nil, fmt.Errorf("no RDAP servers found for TLD: %s", tld)
}
