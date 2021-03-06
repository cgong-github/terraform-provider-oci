// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateLocalPeeringGatewayDetails The representation of UpdateLocalPeeringGatewayDetails
type UpdateLocalPeeringGatewayDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m UpdateLocalPeeringGatewayDetails) String() string {
	return common.PointerString(m)
}
