package asn

import "github.com/ryanmab/rdap-go/pkg/client/response"

// Response represents the RDAP response structure for autnum queries.
// See: https://datatracker.ietf.org/doc/rfc9083/
type Response struct {
	// An array of strings each providing a hint as to the
	// specifications used in the construction of the
	Conformance []string `json:"rdapConformance" validate:"dive,required"`

	ObjectType string `json:"objectClassName" validate:"required,eq=autnum"`
	Handle     string `json:"handle" validate:"required"`

	Name string `json:"name" validate:"required"`

	Lang string `json:"lang" validate:"required"`

	Events []response.Event  `json:"events" validate:"dive,required"`
	Status []response.Status `json:"status" validate:"dive,required"`
	Links  []response.Link   `json:"links,omitempty" validate:"dive,required"`

	Entities []response.Entity `json:"entities,omitempty" validate:"dive,required"`

	StartAsn uint32 `json:"startAutnum" validate:"required"`
	EndAsn   uint32 `json:"endAutnum" validate:"required"`
}
