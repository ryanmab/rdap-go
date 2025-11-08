package dns

import (
	"fmt"

	"github.com/ryanmab/rdap-go/internal/model"
)

// Response represents the RDAP response structure for domain queries.
//
// See: https://datatracker.ietf.org/doc/rfc9083/
type Response struct {
	// An array of strings each providing a hint as to the
	// specifications used in the construction of the response.
	Conformance []string `json:"rdapConformance" validate:"dive,required"`

	ObjectType string `json:"objectClassName" validate:"required,eq=domain"`
	Handle     string `json:"handle" validate:"required"`

	// A string describing a domain name in LDH form
	LdhName string `json:"ldhName" validate:"required"`

	// A string describing a domain name in Unicode form
	UnicodeName *string `json:"unicodeName,omitempty"`

	Events []model.Event  `json:"events" validate:"dive,required"`
	Status []model.Status `json:"status" validate:"dive,required"`
	Links  []model.Link   `json:"links,omitempty" validate:"dive,required"`

	Nameservers []model.Nameserver `json:"nameservers" validate:"dive"`

	SecureDNS *struct {
		ZoneSigned       *bool `json:"zoneSigned,omitempty"`
		DelegationSigned *bool `json:"delegationSigned,omitempty"`
		MaxSignatureLife *int  `json:"maxSigLife,omitempty" validate:"omitempty,min=0"`

		DsData []struct {
			KeyTag     int           `json:"keyTag" validate:"required,min=0"`
			Algorithm  int           `json:"algorithm" validate:"required,min=0"`
			DigestType int           `json:"digestType" validate:"required,min=0"`
			Digest     string        `json:"digest" validate:"required"`
			Events     []model.Event `json:"events,omitempty" validate:"dive,required"`
		} `json:"dsData,omitempty" validate:"dive"`

		KeyData []struct {
			Flags     int           `json:"flags" validate:"required,min=0"`
			Protocol  int           `json:"protocol" validate:"required,min=0"`
			Algorithm int           `json:"algorithm" validate:"required,min=0"`
			PublicKey string        `json:"publicKey" validate:"required"`
			Events    []model.Event `json:"events,omitempty" validate:"dive,required"`
		} `json:"keyData,omitempty" validate:"dive"`
	} `json:"secureDNS,omitempty"`

	Entities []model.Entity `json:"entities,omitempty" validate:"dive,required"`
}

// GetServers returns the RDAP servers for a given TLD from the IANA bootstrap data.
//
// See: https://data.iana.org/rdap/
func GetServers(tld string) ([]string, error) {
	if servers, ok := Bootstrap[tld]; ok {
		return servers, nil
	}

	return nil, fmt.Errorf("no RDAP servers found for TLD: %s", tld)
}
