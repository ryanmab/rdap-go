package bootstrap

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/ryanmab/rdap-go/internal/query"
)

// Service represents a single service entry in the IANA RDAP bootstrap data.
type Service struct {
	Keys    []string
	Servers []string
}

// UnmarshalJSON custom unmarshals the Service struct from the IANA RDAP
// bootstrap JSON format - which is an array of two arrays.
func (dns *Service) UnmarshalJSON(data []byte) error {
	var raw [][]string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// The service data is an array of two arrays: the first array contains
	// the keys (i.e. tlds, IPv6's, IPv4's, etc.), the second array contains
	// server URIs (with a trailing slash).
	dns.Keys = raw[0]
	dns.Servers = raw[1]

	return nil
}

// Response represents the overall structure of the IANA RDAP bootstrap data.
type Response struct {
	Version     string    `json:"version"`
	Description string    `json:"description"`
	Publication time.Time `json:"publication,format:datetime"`

	Services []Service `json:"services"`
}

// FetchBootstrap wraps the IANA public bootstrap registries and allows for retrieving
// the latest registry files from IANA for a given type of RDAP query.
//
// For example, DNS bootstrap data is fetched for domain queries,
// IPv4 bootstrap data is fetched for IPv4 address queries, etc.
func FetchBootstrap(bootstrap query.RdapQuery) Response {
	var url string
	switch bootstrap {
	case query.DomainQuery:
		url = "https://data.iana.org/rdap/dns.json"
	case query.IPv4Query:
		url = "https://data.iana.org/rdap/ipv4.json"
	case query.IPv6Query:
		url = "https://data.iana.org/rdap/ipv6.json"
	default:
		panic("unknown bootstrap type")
	}

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var bootstrapResponse Response

	if err = json.NewDecoder(response.Body).Decode(&bootstrapResponse); err != nil {
		slog.Error("Failed to decode bootstrap response", url, err)
	}

	return bootstrapResponse
}
