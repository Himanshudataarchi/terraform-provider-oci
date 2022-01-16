// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package operationsinsights

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
	"strings"
)

// IngestSqlPlanLinesRequest wrapper for the IngestSqlPlanLines operation
type IngestSqlPlanLinesRequest struct {

	// Collection of SQL plan line objects for a particular database.
	IngestSqlPlanLinesDetails `contributesTo:"body"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Optional OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated DBaaS entity.
	DatabaseId *string `mandatory:"false" contributesTo:"query" name:"databaseId"`

	// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database insight resource.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Used for optimistic concurrency control. In the update or delete call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous get, create, or update response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// A token that uniquely identifies a request that can be retried in case of a timeout or
	// server error without risk of executing the same action again. Retry tokens expire after 24
	// hours.
	// *Note:* Retry tokens can be invalidated before the 24 hour time limit due to conflicting
	// operations, such as a resource being deleted or purged from the system.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request IngestSqlPlanLinesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request IngestSqlPlanLinesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request IngestSqlPlanLinesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request IngestSqlPlanLinesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request IngestSqlPlanLinesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IngestSqlPlanLinesResponse wrapper for the IngestSqlPlanLines operation
type IngestSqlPlanLinesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The IngestSqlPlanLinesResponseDetails instance
	IngestSqlPlanLinesResponseDetails `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response IngestSqlPlanLinesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response IngestSqlPlanLinesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}