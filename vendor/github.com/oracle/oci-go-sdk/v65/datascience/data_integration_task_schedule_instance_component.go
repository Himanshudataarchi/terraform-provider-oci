// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataIntegrationTaskScheduleInstanceComponent Reference to Data Integration Task Schedule
type DataIntegrationTaskScheduleInstanceComponent struct {

	// Identifier of Task Schedule.
	Identifier *string `mandatory:"true" json:"identifier"`

	// OCID of Data Integration Workspace in which Task Schedule is located.
	WorkspaceId *string `mandatory:"true" json:"workspaceId"`

	// Key of Application where Task Schedule is located.
	ApplicationKey *string `mandatory:"true" json:"applicationKey"`

	// Key of Data Integration Task Schedule
	TaskScheduleKey *string `mandatory:"true" json:"taskScheduleKey"`

	// Key of Data Integration Task.
	TaskKey *string `mandatory:"true" json:"taskKey"`
}

func (m DataIntegrationTaskScheduleInstanceComponent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataIntegrationTaskScheduleInstanceComponent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataIntegrationTaskScheduleInstanceComponent) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataIntegrationTaskScheduleInstanceComponent DataIntegrationTaskScheduleInstanceComponent
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDataIntegrationTaskScheduleInstanceComponent
	}{
		"DATA_INTEGRATION_TASK_SCHEDULE",
		(MarshalTypeDataIntegrationTaskScheduleInstanceComponent)(m),
	}

	return json.Marshal(&s)
}
