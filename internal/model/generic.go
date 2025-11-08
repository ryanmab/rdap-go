package model

import "time"

// Status represents the RDAP specification's status of the Domain.
//
// See Section 10.2.2: https://datatracker.ietf.org/doc/rfc9083/
type Status string

const (
	// StatusValidated signifies that the data of the object instance has been
	// found to be accurate. This type of status is usually found on
	// entity object instances to note the validity of identifying
	// contact information.
	StatusValidated Status = "validated"

	// StatusRenewProhibited signifies that the renewal or reregistration of
	// the object instance is forbidden.
	StatusRenewProhibited Status = "renew prohibited"

	// StatusUpdateProhibited signifies that updates to the object instance are forbidden.
	StatusUpdateProhibited Status = "update prohibited"

	// StatusTransferProhibited signifies that transfers of the registration from
	// one registrar to another are forbidden. This type of status normally applies to
	// DNR domain names.
	StatusTransferProhibited Status = "transfer prohibited"

	// StatusDeleteProhibited signifies that deletion of the registration of the object
	// instance is forbidden. This type of status normally applies to DNR domain
	// names.
	StatusDeleteProhibited Status = "delete prohibited"

	// StatusProxy signifies that the registration of the object instance has been performed by
	// third party. This is most commonly applied to entities.
	StatusProxy Status = "proxy"

	// StatusPrivate signifies that the information of the object instance is not
	// designated for public consumption. This is most commonly applied
	// to entities.
	StatusPrivate Status = "private"

	// StatusRemoved signifies that some of the information of the object instance has not
	// been made available and has been removed. This is most commonly
	// applied to entities.
	StatusRemoved Status = "removed"

	// StatusObscured signifies that some of the information of the object instance has been
	// altered for the purposes of not readily revealing the actual
	// information of the object instance. This is most commonly applied
	// to entities.
	StatusObscured Status = "obscured"

	// StatusAssociated signifies that the object instance is associated with other object
	// instances in the registry. This is most commonly used to signify
	// that a nameserver is associated with a domain or that an entity is
	// associated with a network resource or domain.
	StatusAssociated Status = "associated"

	// StatusActive signifies that the object instance is in use. For domain names, it
	// signifies that the domain name is published in DNS. For network
	// and autnum registrations, it signifies that they are allocated or
	// assigned for use in operational networks. This maps to the "OK"
	// status of the Extensible Provisioning Protocol (EPP) [RFC5730].
	StatusActive Status = "active"

	// StatusInactive signifies that the object instance is not in use. See "active".
	StatusInactive Status = "inactive"

	// StatusLocked signifies that changes to the object instance cannot be made,
	// including the association of other object instances.
	StatusLocked Status = "locked"

	// StatusPendingCreate signifies that a request has been received for the creation of the
	// object instance, but this action is not yet complete.
	StatusPendingCreate Status = "pending create"

	// StatusPendingRenew signifies that a request has been received for the renewal of the
	// object instance, but this action is not yet complete.
	StatusPendingRenew Status = "pending renew"

	// StatusPendingTransfer signifies that a request has been received for the transfer of the
	// object instance, but this action is not yet complete.
	StatusPendingTransfer Status = "pending transfer"

	// StatusPendingUpdate signifies that a request has been received for the update or
	// modification of the object instance, but this action is not yet
	// complete.
	StatusPendingUpdate Status = "pending update"

	// StatusPendingDelete signifies that a request has been received for the deletion or removal
	// of the object instance, but this action is not yet complete. For
	// domains, this might mean that the name is no longer published in
	// DNS but has not yet been purged from the registry database.
	StatusPendingDelete Status = "pending delete"
)

// Event represents the RDAP specification's event object.
//
// See Section 4.5: https://datatracker.ietf.org/doc/rfc9083/
type Event struct {
	Action string    `json:"eventAction"`
	Actor  *string   `json:"eventActor,omitempty"`
	Date   time.Time `json:"eventDate,format:datetime"`
}

// Nameserver represents the RDAP specification's nameserver object.
//
// See Section 5.2: https://datatracker.ietf.org/doc/rfc9083/
type Nameserver struct {
	ObjectType  string   `json:"objectClassName" validate:"required,eq=nameserver"`
	Handle      *string  `json:"handle,omitempty"`
	LdhName     string   `json:"ldhName" validate:"required"`
	UnicodeName *string  `json:"unicodeName,omitempty"`
	Events      []Event  `json:"events,omitempty" validate:"dive,required"`
	Status      []Status `json:"status,omitempty" validate:"dive,required"`
	IPAddresses struct {
		V4 []string `json:"v4,omitempty" validate:"dive,ipv4"`
		V6 []string `json:"v6,omitempty" validate:"dive,ipv6"`
	} `json:"ipAddresses"`
}

// Entity represents the RDAP specification's entity object.
//
// See Section 5.1: https://datatracker.ietf.org/doc/rfc9083/
type Entity struct {
	ObjectType string   `json:"objectClassName" validate:"required,eq=entity"`
	Handle     string   `json:"handle"`
	VCardArray any      `json:"vcardArray" validate:"required"`
	Roles      []string `json:"roles,omitempty" validate:"dive,required"`
	PublicIds  []struct {
		Type       string `json:"type" validate:"required"`
		Identifier string `json:"identifier" validate:"required"`
	} `json:"publicIds,omitempty" validate:"dive,required"`
	Events       []Event  `json:"events,omitempty" validate:"dive,required"`
	Entities     []Entity `json:"entities,omitempty" validate:"dive,required"`
	AsEventActor []Event  `json:"asEventActor,omitempty" validate:"dive,required"`
	Status       []Status `json:"status,omitempty" validate:"dive,required"`
	WhoisURI     *string  `json:"port43,omitempty" validate:"omitempty"`
	Links        []Link   `json:"links,omitempty" validate:"dive,required"`
}

// Link represents the RDAP specification's link object.
//
// See Section 4.2: https://datatracker.ietf.org/doc/rfc9083/
type Link struct {
	Rel   string `json:"rel" validate:"required"`
	Href  string `json:"href" validate:"required,url"`
	Type  string `json:"type,omitempty" validate:"omitempty"`
	Value string `json:"value,omitempty" validate:"omitempty"`
}
