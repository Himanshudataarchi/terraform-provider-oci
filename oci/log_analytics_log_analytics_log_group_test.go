// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v53/common"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v53/loganalytics"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	LogAnalyticsLogGroupRequiredOnlyResource = LogAnalyticsLogGroupResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_log_group", "test_log_analytics_log_group", Required, Create, logAnalyticsLogGroupRepresentation)

	LogAnalyticsLogGroupResourceConfig = LogAnalyticsLogGroupResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_log_group", "test_log_analytics_log_group", Optional, Update, logAnalyticsLogGroupRepresentation)

	logAnalyticsLogGroupSingularDataSourceRepresentation = map[string]interface{}{
		"log_analytics_log_group_id": Representation{RepType: Required, Create: `${oci_log_analytics_log_analytics_log_group.test_log_analytics_log_group.id}`},
		"namespace":                  Representation{RepType: Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	logAnalyticsLogGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"namespace":      Representation{RepType: Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"filter":         RepresentationGroup{Required, logAnalyticsLogGroupDataSourceFilterRepresentation}}
	logAnalyticsLogGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_log_analytics_log_analytics_log_group.test_log_analytics_log_group.id}`}},
	}

	logAnalyticsLogGroupRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Required, Create: `displayName`, Update: `displayName2`},
		"namespace":      Representation{RepType: Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	LogAnalyticsLogGroupResourceDependencies = DefinedTagsDependencies +
		GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsLogGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsLogGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_log_analytics_log_analytics_log_group.test_log_analytics_log_group"
	datasourceName := "data.oci_log_analytics_log_analytics_log_groups.test_log_analytics_log_groups"
	singularDatasourceName := "data.oci_log_analytics_log_analytics_log_group.test_log_analytics_log_group"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+LogAnalyticsLogGroupResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_log_group", "test_log_analytics_log_group", Optional, Create, logAnalyticsLogGroupRepresentation), "loganalytics", "logAnalyticsLogGroup", t)

	ResourceTest(t, testAccCheckLogAnalyticsLogAnalyticsLogGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogGroupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_log_group", "test_log_analytics_log_group", Required, Create, logAnalyticsLogGroupRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogGroupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogGroupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_log_group", "test_log_analytics_log_group", Optional, Create, logAnalyticsLogGroupRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + LogAnalyticsLogGroupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_log_group", "test_log_analytics_log_group", Optional, Create,
					RepresentationCopyWithNewProperties(logAnalyticsLogGroupRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogGroupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_log_group", "test_log_analytics_log_group", Optional, Update, logAnalyticsLogGroupRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_log_groups", "test_log_analytics_log_groups", Optional, Update, logAnalyticsLogGroupDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsLogGroupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_log_group", "test_log_analytics_log_group", Optional, Update, logAnalyticsLogGroupRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "namespace"),

				resource.TestCheckResourceAttr(datasourceName, "log_analytics_log_group_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "log_analytics_log_group_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_log_group", "test_log_analytics_log_group", Required, Create, logAnalyticsLogGroupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsLogGroupResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "log_analytics_log_group_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogGroupResourceConfig,
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + LogAnalyticsLogGroupResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateIdFunc:       getLogAnalyticsLogGroupsEndpointImportId(resourceName),
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckLogAnalyticsLogAnalyticsLogGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).logAnalyticsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_log_analytics_log_analytics_log_group" {
			noResourceFound = false
			request := oci_log_analytics.GetLogAnalyticsLogGroupRequest{}

			tmp := rs.Primary.ID
			request.LogAnalyticsLogGroupId = &tmp

			if value, ok := rs.Primary.Attributes["namespace"]; ok {
				request.NamespaceName = &value
			}

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "log_analytics")

			_, err := client.GetLogAnalyticsLogGroup(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("LogAnalyticsLogAnalyticsLogGroup") {
		resource.AddTestSweepers("LogAnalyticsLogAnalyticsLogGroup", &resource.Sweeper{
			Name:         "LogAnalyticsLogAnalyticsLogGroup",
			Dependencies: DependencyGraph["logAnalyticsLogGroup"],
			F:            sweepLogAnalyticsLogAnalyticsLogGroupResource,
		})
	}
}

func sweepLogAnalyticsLogAnalyticsLogGroupResource(compartment string) error {
	logAnalyticsClient := GetTestClients(&schema.ResourceData{}).logAnalyticsClient()
	logAnalyticsLogGroupIds, err := getLogAnalyticsLogGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, logAnalyticsLogGroupId := range logAnalyticsLogGroupIds {
		if ok := SweeperDefaultResourceId[logAnalyticsLogGroupId]; !ok {
			deleteLogAnalyticsLogGroupRequest := oci_log_analytics.DeleteLogAnalyticsLogGroupRequest{}

			deleteLogAnalyticsLogGroupRequest.LogAnalyticsLogGroupId = &logAnalyticsLogGroupId

			deleteLogAnalyticsLogGroupRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "log_analytics")
			_, error := logAnalyticsClient.DeleteLogAnalyticsLogGroup(context.Background(), deleteLogAnalyticsLogGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting LogAnalyticsLogGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", logAnalyticsLogGroupId, error)
				continue
			}
		}
	}
	return nil
}

func getLogAnalyticsLogGroupIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "LogAnalyticsLogGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	logAnalyticsClient := GetTestClients(&schema.ResourceData{}).logAnalyticsClient()

	listLogAnalyticsLogGroupsRequest := oci_log_analytics.ListLogAnalyticsLogGroupsRequest{}
	listLogAnalyticsLogGroupsRequest.CompartmentId = &compartmentId

	namespaces, error := getNamespaces(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting namespace required for LogAnalyticsLogGroup resource requests \n")
	}
	for _, namespace := range namespaces {
		listLogAnalyticsLogGroupsRequest.NamespaceName = &namespace

		listLogAnalyticsLogGroupsResponse, err := logAnalyticsClient.ListLogAnalyticsLogGroups(context.Background(), listLogAnalyticsLogGroupsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting LogAnalyticsLogGroup list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, logAnalyticsLogGroup := range listLogAnalyticsLogGroupsResponse.Items {
			id := *logAnalyticsLogGroup.Id
			resourceIds = append(resourceIds, id)
			AddResourceIdToSweeperResourceIdMap(compartmentId, "LogAnalyticsLogGroupId", id)
		}

	}
	return resourceIds, nil
}

func getLogAnalyticsLogGroupsEndpointImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("namespaces/" + rs.Primary.Attributes["namespace"] + "/logAnalyticsLogGroups/" + rs.Primary.Attributes["id"]), nil
	}
}
