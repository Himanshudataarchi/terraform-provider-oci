// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package certificatesmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAssociationsRequest wrapper for the ListAssociations operation
type ListAssociationsRequest struct {

	// Unique Oracle-assigned identifier for the request. If provided, the returned request ID
	// will include this value. Otherwise, a random request ID will be
	// generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter that returns only resources that match the given compartment OCID.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter that returns only resources that match the given OCID of a certificate-related resource.
	CertificatesResourceId *string `mandatory:"false" contributesTo:"query" name:"certificatesResourceId"`

	// A filter that returns only resources that match the given OCID of an associated Oracle Cloud Infrastructure resource.
	AssociatedResourceId *string `mandatory:"false" contributesTo:"query" name:"associatedResourceId"`

	// The OCID of the association. If the parameter is set to null, the service lists all associations.
	AssociationId *string `mandatory:"false" contributesTo:"query" name:"associationId"`

	// A filter that returns only resources that match the specified name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort by. You can specify only one sort order. The default order for `TIMECREATED` is descending.
	// The default order for `NAME` is ascending.
	SortBy ListAssociationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAssociationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header
	// from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Type of associations to list. If the parameter is set to null, the service lists all types of associations.
	AssociationType ListAssociationsAssociationTypeEnum `mandatory:"false" contributesTo:"query" name:"associationType" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAssociationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAssociationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAssociationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAssociationsRequest) RetryPolicy() common.OCIRetry {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAssociationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAssociationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAssociationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssociationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAssociationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssociationsAssociationTypeEnum(string(request.AssociationType)); !ok && request.AssociationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssociationType: %s. Supported values are: %s.", request.AssociationType, strings.Join(GetListAssociationsAssociationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAssociationsResponse wrapper for the ListAssociations operation
type ListAssociationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AssociationCollection instances
	AssociationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#List_Pagination).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAssociationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAssociationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAssociationsSortByEnum Enum with underlying type: string
type ListAssociationsSortByEnum string

// Set of constants representing the allowable values for ListAssociationsSortByEnum
const (
	ListAssociationsSortByName        ListAssociationsSortByEnum = "NAME"
	ListAssociationsSortByTimecreated ListAssociationsSortByEnum = "TIMECREATED"
)

var mappingListAssociationsSortByEnum = map[string]ListAssociationsSortByEnum{
	"NAME":        ListAssociationsSortByName,
	"TIMECREATED": ListAssociationsSortByTimecreated,
}

var mappingListAssociationsSortByEnumLowerCase = map[string]ListAssociationsSortByEnum{
	"name":        ListAssociationsSortByName,
	"timecreated": ListAssociationsSortByTimecreated,
}

// GetListAssociationsSortByEnumValues Enumerates the set of values for ListAssociationsSortByEnum
func GetListAssociationsSortByEnumValues() []ListAssociationsSortByEnum {
	values := make([]ListAssociationsSortByEnum, 0)
	for _, v := range mappingListAssociationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssociationsSortByEnumStringValues Enumerates the set of values in String for ListAssociationsSortByEnum
func GetListAssociationsSortByEnumStringValues() []string {
	return []string{
		"NAME",
		"TIMECREATED",
	}
}

// GetMappingListAssociationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssociationsSortByEnum(val string) (ListAssociationsSortByEnum, bool) {
	enum, ok := mappingListAssociationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssociationsSortOrderEnum Enum with underlying type: string
type ListAssociationsSortOrderEnum string

// Set of constants representing the allowable values for ListAssociationsSortOrderEnum
const (
	ListAssociationsSortOrderAsc  ListAssociationsSortOrderEnum = "ASC"
	ListAssociationsSortOrderDesc ListAssociationsSortOrderEnum = "DESC"
)

var mappingListAssociationsSortOrderEnum = map[string]ListAssociationsSortOrderEnum{
	"ASC":  ListAssociationsSortOrderAsc,
	"DESC": ListAssociationsSortOrderDesc,
}

var mappingListAssociationsSortOrderEnumLowerCase = map[string]ListAssociationsSortOrderEnum{
	"asc":  ListAssociationsSortOrderAsc,
	"desc": ListAssociationsSortOrderDesc,
}

// GetListAssociationsSortOrderEnumValues Enumerates the set of values for ListAssociationsSortOrderEnum
func GetListAssociationsSortOrderEnumValues() []ListAssociationsSortOrderEnum {
	values := make([]ListAssociationsSortOrderEnum, 0)
	for _, v := range mappingListAssociationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssociationsSortOrderEnumStringValues Enumerates the set of values in String for ListAssociationsSortOrderEnum
func GetListAssociationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAssociationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssociationsSortOrderEnum(val string) (ListAssociationsSortOrderEnum, bool) {
	enum, ok := mappingListAssociationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssociationsAssociationTypeEnum Enum with underlying type: string
type ListAssociationsAssociationTypeEnum string

// Set of constants representing the allowable values for ListAssociationsAssociationTypeEnum
const (
	ListAssociationsAssociationTypeCertificate          ListAssociationsAssociationTypeEnum = "CERTIFICATE"
	ListAssociationsAssociationTypeCertificateAuthority ListAssociationsAssociationTypeEnum = "CERTIFICATE_AUTHORITY"
	ListAssociationsAssociationTypeCaBundle             ListAssociationsAssociationTypeEnum = "CA_BUNDLE"
)

var mappingListAssociationsAssociationTypeEnum = map[string]ListAssociationsAssociationTypeEnum{
	"CERTIFICATE":           ListAssociationsAssociationTypeCertificate,
	"CERTIFICATE_AUTHORITY": ListAssociationsAssociationTypeCertificateAuthority,
	"CA_BUNDLE":             ListAssociationsAssociationTypeCaBundle,
}

var mappingListAssociationsAssociationTypeEnumLowerCase = map[string]ListAssociationsAssociationTypeEnum{
	"certificate":           ListAssociationsAssociationTypeCertificate,
	"certificate_authority": ListAssociationsAssociationTypeCertificateAuthority,
	"ca_bundle":             ListAssociationsAssociationTypeCaBundle,
}

// GetListAssociationsAssociationTypeEnumValues Enumerates the set of values for ListAssociationsAssociationTypeEnum
func GetListAssociationsAssociationTypeEnumValues() []ListAssociationsAssociationTypeEnum {
	values := make([]ListAssociationsAssociationTypeEnum, 0)
	for _, v := range mappingListAssociationsAssociationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssociationsAssociationTypeEnumStringValues Enumerates the set of values in String for ListAssociationsAssociationTypeEnum
func GetListAssociationsAssociationTypeEnumStringValues() []string {
	return []string{
		"CERTIFICATE",
		"CERTIFICATE_AUTHORITY",
		"CA_BUNDLE",
	}
}

// GetMappingListAssociationsAssociationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssociationsAssociationTypeEnum(val string) (ListAssociationsAssociationTypeEnum, bool) {
	enum, ok := mappingListAssociationsAssociationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
