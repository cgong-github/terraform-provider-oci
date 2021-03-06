// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetDbNodeRequest wrapper for the GetDbNode operation
type GetDbNodeRequest struct {

	// The database node [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
	DbNodeId *string `mandatory:"true" contributesTo:"path" name:"dbNodeId"`
}

func (request GetDbNodeRequest) String() string {
	return common.PointerString(request)
}

// GetDbNodeResponse wrapper for the GetDbNode operation
type GetDbNodeResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DbNode instance
	DbNode `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDbNodeResponse) String() string {
	return common.PointerString(response)
}
