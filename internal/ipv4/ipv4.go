package ipv4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ryanmab/rdap-go/internal/model"
)

// Response represents the RDAP response structure for ipv4 queries.
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

	StartAddress string `json:"startAddress" validate:"required,ipv4"`
	EndAddress   string `json:"endAddress" validate:"required,ipv4"`
	IPVersion    string `json:"ipVersion" validate:"required,eq=v4"`

	Events []model.Event  `json:"events" validate:"dive,required"`
	Status []model.Status `json:"status" validate:"dive,required"`

	Entities []model.Entity `json:"entities,omitempty" validate:"dive,required"`

	Links []model.Link `json:"links,omitempty" validate:"dive,required"`
}

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

	if servers, ok := Bootstrap[firstOctet]; ok {
		return servers, nil
	}

	return nil, fmt.Errorf("no RDAP servers found for first octet (%s) of IPv4: %s", firstOctet, ip)
}
