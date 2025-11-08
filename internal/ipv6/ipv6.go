package ipv6

import (
	"fmt"
	"net"
	"strings"

	"github.com/ryanmab/rdap-go/internal/model"
)

// Response represents the RDAP response structure for ipv6 queries.
//
// See: https://datatracker.ietf.org/doc/rfc9083/
type Response struct {
	// An array of strings each providing a hint as to the
	// specifications used in the construction of the response.
	Conformance []string `json:"rdapConformance" validate:"dive,required"`

	ObjectType   string `json:"objectClassName" validate:"required,eq=ip network"`
	Handle       string `json:"handle" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Type         string `json:"type" validate:"required"`
	ParentHandle string `json:"parentHandle,omitempty" validate:"omitempty,required"`
	Country      string `json:"country,omitempty" validate:"omitempty,required,len=2"`

	StartAddress string `json:"startAddress" validate:"required,ipv6"`
	EndAddress   string `json:"endAddress" validate:"required,ipv6"`
	IPVersion    string `json:"ipVersion" validate:"required,eq=v6"`

	Events []model.Event  `json:"events" validate:"dive,required"`
	Status []model.Status `json:"status" validate:"dive,required"`

	Entities []model.Entity `json:"entities,omitempty" validate:"dive,required"`

	Links []model.Link `json:"links,omitempty" validate:"dive,required"`
}

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
