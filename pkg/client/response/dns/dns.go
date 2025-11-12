package dns

import "github.com/ryanmab/rdap-go/pkg/client/response"

// Response represents the RDAP response structure for domain queries.
// See: https://datatracker.ietf.org/doc/rfc9083/
type Response struct {
	// An array of strings each providing a hint as to the
	// specifications used in the construction of the
	Conformance []string `json:"rdapConformance" validate:"dive,required"`

	ObjectType string `json:"objectClassName" validate:"required,eq=domain"`
	Handle     string `json:"handle" validate:"required"`

	// A string describing a domain name in LDH form
	LdhName string `json:"ldhName" validate:"required"`

	// A string describing a domain name in Unicode form
	UnicodeName *string `json:"unicodeName,omitempty"`

	Events []response.Event  `json:"events" validate:"dive,required"`
	Status []response.Status `json:"status" validate:"dive,required"`
	Links  []response.Link   `json:"links,omitempty" validate:"dive,required"`

	Nameservers []response.Nameserver `json:"nameservers" validate:"dive"`

	SecureDNS *struct {
		ZoneSigned       *bool `json:"zoneSigned,omitempty"`
		DelegationSigned *bool `json:"delegationSigned,omitempty"`
		MaxSignatureLife *int  `json:"maxSigLife,omitempty" validate:"omitempty,min=0"`

		DsData []struct {
			KeyTag     int              `json:"keyTag" validate:"required,min=0"`
			Algorithm  int              `json:"algorithm" validate:"required,min=0"`
			DigestType int              `json:"digestType" validate:"required,min=0"`
			Digest     string           `json:"digest" validate:"required"`
			Events     []response.Event `json:"events,omitempty" validate:"dive,required"`
		} `json:"dsData,omitempty" validate:"dive"`

		KeyData []struct {
			Flags     int              `json:"flags" validate:"required,min=0"`
			Protocol  int              `json:"protocol" validate:"required,min=0"`
			Algorithm int              `json:"algorithm" validate:"required,min=0"`
			PublicKey string           `json:"publicKey" validate:"required"`
			Events    []response.Event `json:"events,omitempty" validate:"dive,required"`
		} `json:"keyData,omitempty" validate:"dive"`
	} `json:"secureDNS,omitempty"`

	Entities []response.Entity `json:"entities,omitempty" validate:"dive,required"`
}
