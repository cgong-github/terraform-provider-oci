// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetIdentityProviderRequest wrapper for the GetIdentityProvider operation
type GetIdentityProviderRequest struct {

	// The OCID of the identity provider.
	IdentityProviderId *string `mandatory:"true" contributesTo:"path" name:"identityProviderId"`
}

func (request GetIdentityProviderRequest) String() string {
	return common.PointerString(request)
}

// GetIdentityProviderResponse wrapper for the GetIdentityProvider operation
type GetIdentityProviderResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The IdentityProvider instance
	IdentityProvider `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response GetIdentityProviderResponse) String() string {
	return common.PointerString(response)
}
