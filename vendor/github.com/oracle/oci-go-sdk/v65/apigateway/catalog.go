// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Catalog Information about a catalog.
type Catalog struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which the
	// resource is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// List of APIs to be published.
	Apis []ApiPublication `mandatory:"false" json:"apis"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the catalog.
	// - The catalog will be in ACTIVE state if create and update are successful.
	// - The catalog will be in FAILED state if the create, update or delete fails.
	// - The catalog will be in DELETED state if the delete operation is successful.
	LifecycleState CatalogLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a
	// resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair
	// with no predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see
	// Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Catalog) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Catalog) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCatalogLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCatalogLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CatalogLifecycleStateEnum Enum with underlying type: string
type CatalogLifecycleStateEnum string

// Set of constants representing the allowable values for CatalogLifecycleStateEnum
const (
	CatalogLifecycleStateCreating CatalogLifecycleStateEnum = "CREATING"
	CatalogLifecycleStateActive   CatalogLifecycleStateEnum = "ACTIVE"
	CatalogLifecycleStateUpdating CatalogLifecycleStateEnum = "UPDATING"
	CatalogLifecycleStateDeleting CatalogLifecycleStateEnum = "DELETING"
	CatalogLifecycleStateDeleted  CatalogLifecycleStateEnum = "DELETED"
	CatalogLifecycleStateFailed   CatalogLifecycleStateEnum = "FAILED"
)

var mappingCatalogLifecycleStateEnum = map[string]CatalogLifecycleStateEnum{
	"CREATING": CatalogLifecycleStateCreating,
	"ACTIVE":   CatalogLifecycleStateActive,
	"UPDATING": CatalogLifecycleStateUpdating,
	"DELETING": CatalogLifecycleStateDeleting,
	"DELETED":  CatalogLifecycleStateDeleted,
	"FAILED":   CatalogLifecycleStateFailed,
}

var mappingCatalogLifecycleStateEnumLowerCase = map[string]CatalogLifecycleStateEnum{
	"creating": CatalogLifecycleStateCreating,
	"active":   CatalogLifecycleStateActive,
	"updating": CatalogLifecycleStateUpdating,
	"deleting": CatalogLifecycleStateDeleting,
	"deleted":  CatalogLifecycleStateDeleted,
	"failed":   CatalogLifecycleStateFailed,
}

// GetCatalogLifecycleStateEnumValues Enumerates the set of values for CatalogLifecycleStateEnum
func GetCatalogLifecycleStateEnumValues() []CatalogLifecycleStateEnum {
	values := make([]CatalogLifecycleStateEnum, 0)
	for _, v := range mappingCatalogLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCatalogLifecycleStateEnumStringValues Enumerates the set of values in String for CatalogLifecycleStateEnum
func GetCatalogLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingCatalogLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCatalogLifecycleStateEnum(val string) (CatalogLifecycleStateEnum, bool) {
	enum, ok := mappingCatalogLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
