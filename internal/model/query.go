package model

import (
	"log"
)

// RdapQuery represents the type of RDAP query to make to an RDAP server (domain, IPv4, or IPv6)
type RdapQuery int

const (
	// DomainQuery is a lookup on a DNS record - e.g., example.com
	DomainQuery RdapQuery = iota
	// IPv4Query is a lookup on an IPv4 address - e.g. 8.8.8.8
	IPv4Query
	// IPv6Query is a lookup on an IPv6 address - e.g. 2001:4860:4860::8888
	IPv6Query
)

func (q RdapQuery) String() string {
	switch q {
	case DomainQuery:
		return "domain"
	case IPv6Query:
		return "ip"
	case IPv4Query:
		return "ip"
	default:
		log.Panic("unknown RdapQuery type")
		return ""
	}
}
