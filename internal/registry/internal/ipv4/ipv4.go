package ipv4

import (
	"fmt"
	"strconv"
	"strings"
)

// GetServers returns the RDAP servers for a given IPv4 from the IANA bootstrap data.
//
// See: https://data.iana.org/rdap/
func GetServers(ip string) ([]string, error) {
	firstOctet := strings.Split(ip, ".")[0]

	firstOctetAsInt, err := strconv.ParseInt(firstOctet, 10, 8)

	if err != nil {
		return nil, fmt.Errorf("expected first octet (%s) of IPv4 to be an integer: %s", firstOctet, ip)
	}

	if firstOctetAsInt < 0 || firstOctetAsInt > 255 {
		return nil, fmt.Errorf("out of range first octet (%s) of IPv4: %s", firstOctet, ip)
	}

	if servers, ok := Bootstrap[uint8(firstOctetAsInt)]; ok {
		return servers, nil
	}

	return nil, fmt.Errorf("no RDAP servers found for first octet (%s) of IPv4: %s", firstOctet, ip)
}
