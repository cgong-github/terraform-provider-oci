// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListBootVolumeAttachmentsRequest wrapper for the ListBootVolumeAttachments operation
type ListBootVolumeAttachmentsRequest struct {

	// The name of the Availability Domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" contributesTo:"query" name:"availabilityDomain"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return in a paginated "List" call.
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The OCID of the instance.
	InstanceId *string `mandatory:"false" contributesTo:"query" name:"instanceId"`

	// The OCID of the boot volume.
	BootVolumeId *string `mandatory:"false" contributesTo:"query" name:"bootVolumeId"`
}

func (request ListBootVolumeAttachmentsRequest) String() string {
	return common.PointerString(request)
}

// ListBootVolumeAttachmentsResponse wrapper for the ListBootVolumeAttachments operation
type ListBootVolumeAttachmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []BootVolumeAttachment instance
	Items []BootVolumeAttachment `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListBootVolumeAttachmentsResponse) String() string {
	return common.PointerString(response)
}
