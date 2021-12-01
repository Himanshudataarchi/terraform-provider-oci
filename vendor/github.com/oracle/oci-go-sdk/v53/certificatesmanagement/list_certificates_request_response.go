// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package certificatesmanagement

import (
	"github.com/oracle/oci-go-sdk/v53/common"
	"net/http"
)

// ListCertificatesRequest wrapper for the ListCertificates operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/ListCertificates.go.html to see an example of how to use ListCertificatesRequest.
type ListCertificatesRequest struct {

	// Unique Oracle-assigned identifier for the request. If provided, the returned request ID
	// will include this value. Otherwise, a random request ID will be
	// generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter that returns only resources that match the given compartment OCID.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter that returns only resources that match the given lifecycle state. The state value is case-insensitive.
	LifecycleState ListCertificatesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter that returns only resources that match the specified name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort by. You can specify only one sort order. The default
	// order for `EXPIRATIONDATE` and 'TIMECREATED' is descending. The default order for `NAME`
	// is ascending.
	SortBy ListCertificatesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListCertificatesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header
	// from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The OCID of the certificate authority (CA). If the parameter is set to null, the service lists all CAs.
	IssuerCertificateAuthorityId *string `mandatory:"false" contributesTo:"query" name:"issuerCertificateAuthorityId"`

	// The OCID of the certificate. If the parameter is set to null, the service lists all certificates.
	CertificateId *string `mandatory:"false" contributesTo:"query" name:"certificateId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCertificatesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCertificatesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCertificatesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCertificatesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListCertificatesResponse wrapper for the ListCertificates operation
type ListCertificatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CertificateCollection instances
	CertificateCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#List_Pagination).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCertificatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCertificatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCertificatesLifecycleStateEnum Enum with underlying type: string
type ListCertificatesLifecycleStateEnum string

// Set of constants representing the allowable values for ListCertificatesLifecycleStateEnum
const (
	ListCertificatesLifecycleStateCreating           ListCertificatesLifecycleStateEnum = "CREATING"
	ListCertificatesLifecycleStateActive             ListCertificatesLifecycleStateEnum = "ACTIVE"
	ListCertificatesLifecycleStateUpdating           ListCertificatesLifecycleStateEnum = "UPDATING"
	ListCertificatesLifecycleStateDeleting           ListCertificatesLifecycleStateEnum = "DELETING"
	ListCertificatesLifecycleStateDeleted            ListCertificatesLifecycleStateEnum = "DELETED"
	ListCertificatesLifecycleStateSchedulingDeletion ListCertificatesLifecycleStateEnum = "SCHEDULING_DELETION"
	ListCertificatesLifecycleStatePendingDeletion    ListCertificatesLifecycleStateEnum = "PENDING_DELETION"
	ListCertificatesLifecycleStateCancellingDeletion ListCertificatesLifecycleStateEnum = "CANCELLING_DELETION"
	ListCertificatesLifecycleStateFailed             ListCertificatesLifecycleStateEnum = "FAILED"
)

var mappingListCertificatesLifecycleState = map[string]ListCertificatesLifecycleStateEnum{
	"CREATING":            ListCertificatesLifecycleStateCreating,
	"ACTIVE":              ListCertificatesLifecycleStateActive,
	"UPDATING":            ListCertificatesLifecycleStateUpdating,
	"DELETING":            ListCertificatesLifecycleStateDeleting,
	"DELETED":             ListCertificatesLifecycleStateDeleted,
	"SCHEDULING_DELETION": ListCertificatesLifecycleStateSchedulingDeletion,
	"PENDING_DELETION":    ListCertificatesLifecycleStatePendingDeletion,
	"CANCELLING_DELETION": ListCertificatesLifecycleStateCancellingDeletion,
	"FAILED":              ListCertificatesLifecycleStateFailed,
}

// GetListCertificatesLifecycleStateEnumValues Enumerates the set of values for ListCertificatesLifecycleStateEnum
func GetListCertificatesLifecycleStateEnumValues() []ListCertificatesLifecycleStateEnum {
	values := make([]ListCertificatesLifecycleStateEnum, 0)
	for _, v := range mappingListCertificatesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListCertificatesSortByEnum Enum with underlying type: string
type ListCertificatesSortByEnum string

// Set of constants representing the allowable values for ListCertificatesSortByEnum
const (
	ListCertificatesSortByName           ListCertificatesSortByEnum = "NAME"
	ListCertificatesSortByExpirationdate ListCertificatesSortByEnum = "EXPIRATIONDATE"
	ListCertificatesSortByTimecreated    ListCertificatesSortByEnum = "TIMECREATED"
)

var mappingListCertificatesSortBy = map[string]ListCertificatesSortByEnum{
	"NAME":           ListCertificatesSortByName,
	"EXPIRATIONDATE": ListCertificatesSortByExpirationdate,
	"TIMECREATED":    ListCertificatesSortByTimecreated,
}

// GetListCertificatesSortByEnumValues Enumerates the set of values for ListCertificatesSortByEnum
func GetListCertificatesSortByEnumValues() []ListCertificatesSortByEnum {
	values := make([]ListCertificatesSortByEnum, 0)
	for _, v := range mappingListCertificatesSortBy {
		values = append(values, v)
	}
	return values
}

// ListCertificatesSortOrderEnum Enum with underlying type: string
type ListCertificatesSortOrderEnum string

// Set of constants representing the allowable values for ListCertificatesSortOrderEnum
const (
	ListCertificatesSortOrderAsc  ListCertificatesSortOrderEnum = "ASC"
	ListCertificatesSortOrderDesc ListCertificatesSortOrderEnum = "DESC"
)

var mappingListCertificatesSortOrder = map[string]ListCertificatesSortOrderEnum{
	"ASC":  ListCertificatesSortOrderAsc,
	"DESC": ListCertificatesSortOrderDesc,
}

// GetListCertificatesSortOrderEnumValues Enumerates the set of values for ListCertificatesSortOrderEnum
func GetListCertificatesSortOrderEnumValues() []ListCertificatesSortOrderEnum {
	values := make([]ListCertificatesSortOrderEnum, 0)
	for _, v := range mappingListCertificatesSortOrder {
		values = append(values, v)
	}
	return values
}
