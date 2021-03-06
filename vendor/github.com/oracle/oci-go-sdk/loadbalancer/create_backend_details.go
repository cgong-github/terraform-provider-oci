// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateBackendDetails The configuration details for creating a backend server in a backend set.
// For more information on backend server configuration, see
// [Managing Backend Servers]({{DOC_SERVER_URL}}/Content/Balance/tasks/managingbackendservers.htm).
type CreateBackendDetails struct {

	// The IP address of the backend server.
	// Example: `10.10.10.4`
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// The communication port for the backend server.
	// Example: `8080`
	Port *int `mandatory:"true" json:"port"`

	// Whether the load balancer should treat this server as a backup unit. If `true`, the load balancer forwards no ingress
	// traffic to this backend server unless all other backend servers not marked as "backup" fail the health check policy.
	// Example: `true`
	Backup *bool `mandatory:"false" json:"backup"`

	// Whether the load balancer should drain this server. Servers marked "drain" receive no new
	// incoming traffic.
	// Example: `true`
	Drain *bool `mandatory:"false" json:"drain"`

	// Whether the load balancer should treat this server as offline. Offline servers receive no incoming
	// traffic.
	// Example: `true`
	Offline *bool `mandatory:"false" json:"offline"`

	// The load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger
	// proportion of incoming traffic. For example, a server weighted '3' receives 3 times the number of new connections
	// as a server weighted '1'.
	// For more information on load balancing policies, see
	// [How Load Balancing Policies Work]({{DOC_SERVER_URL}}/Content/Balance/Reference/lbpolicies.htm).
	// Example: `3`
	Weight *int `mandatory:"false" json:"weight"`
}

func (m CreateBackendDetails) String() string {
	return common.PointerString(m)
}
