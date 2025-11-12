package ipv6

import (
	"fmt"
	"net"
	"strings"
)

// GetServers returns the RDAP servers for a given IPv6 from the IANA bootstrap data.
//
// See: https://data.iana.org/rdap/
func GetServers(ip string) ([]string, error) {
	for _, cidr := range BootstrapAccessOrder {
		_, cidrNet, err := net.ParseCIDR(cidr)
		if err != nil {
			return nil, fmt.Errorf("invalid CIDR notation: %s", cidr)
		}

		if !strings.Contains(ip, "/") {
			ip = fmt.Sprintf("%s/128", ip)
		}

		_, ipNet, err := net.ParseCIDR(ip)

		if err != nil {
			return nil, fmt.Errorf("invalid IP address: %s", ip)
		}

		if cidrNet.Contains(ipNet.IP) {
			// The bootstrap entries are ordered from most specific to least specific
			// (i.e. highest subnet mask to lowest subnet mask, and highest number of hextets
			// specified to lowest number of hextets specified). Therefore, we can return the
			// first match we find, as this will be the most specific matching range.
			return Bootstrap[cidr], nil
		}
	}

	return nil, fmt.Errorf("no RDAP servers found for IPv4: %s", ip)
}
