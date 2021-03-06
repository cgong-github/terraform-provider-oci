// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetCrossConnectRequest wrapper for the GetCrossConnect operation
type GetCrossConnectRequest struct {

	// The OCID of the cross-connect.
	CrossConnectId *string `mandatory:"true" contributesTo:"path" name:"crossConnectId"`
}

func (request GetCrossConnectRequest) String() string {
	return common.PointerString(request)
}

// GetCrossConnectResponse wrapper for the GetCrossConnect operation
type GetCrossConnectResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The CrossConnect instance
	CrossConnect `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetCrossConnectResponse) String() string {
	return common.PointerString(response)
}
