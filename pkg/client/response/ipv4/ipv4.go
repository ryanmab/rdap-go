package ipv4

import "github.com/ryanmab/rdap-go/pkg/client/response"

// Response represents the RDAP response structure for ipv4 queries.
// See: https://datatracker.ietf.org/doc/rfc9083/
type Response struct {
	// An array of strings each providing a hint as to the
	// specifications used in the construction of the
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

	Events []response.Event  `json:"events" validate:"dive,required"`
	Status []response.Status `json:"status" validate:"dive,required"`

	Entities []response.Entity `json:"entities,omitempty" validate:"dive,required"`

	Links []response.Link `json:"links,omitempty" validate:"dive,required"`
}
