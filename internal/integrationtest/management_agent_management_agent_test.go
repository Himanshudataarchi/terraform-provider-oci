// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v57/common"
	oci_management_agent "github.com/oracle/oci-go-sdk/v57/managementagent"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	ManagementAgentResourceConfig = ManagementAgentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent", "test_management_agent", acctest.Optional, acctest.Update, managementAgentRepresentation)

	managementAgentSingularDataSourceRepresentation = map[string]interface{}{
		"management_agent_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_management_agent_management_agent.test_management_agent.id}`},
	}

	managementAgentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_status":  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"host_id":              acctest.Representation{RepType: acctest.Optional, Create: ``},
		"install_type":         acctest.Representation{RepType: acctest.Optional, Create: `AGENT`},
		"is_customer_deployed": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"platform_type":        acctest.Representation{RepType: acctest.Optional, Create: []string{`LINUX`}},
		"state":                acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":               acctest.RepresentationGroup{RepType: acctest.Required, Group: managementAgentDataSourceFilterRepresentation}}
	managementAgentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_management_agent_management_agent.test_management_agent.id}`}},
	}

	managementAgentRepresentation = map[string]interface{}{
		"managed_agent_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.managed_agent_id}`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"deploy_plugins_id": acctest.Representation{RepType: acctest.Optional, Create: []string{`${data.oci_management_agent_management_agent_plugins.test_management_agent_plugins.management_agent_plugins.0.id}`}},
	}

	ManagementAgentResourceDependencies = ""
)

// issue-routing-tag: management_agent/default
func TestManagementAgentManagementAgentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagementAgentManagementAgentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managementAgentId := utils.GetEnvSettingWithBlankDefault("managed_agent_id")
	if managementAgentId == "" {
		t.Skip("Manual install agent and set managed_agent_id to run this test")
	}
	managementAgentIdVariableStr := fmt.Sprintf("variable \"managed_agent_id\" { default = \"%s\" }\n", managementAgentId)

	resourceName := "oci_management_agent_management_agent.test_management_agent"
	datasourceName := "data.oci_management_agent_management_agents.test_management_agents"
	singularDatasourceName := "data.oci_management_agent_management_agent.test_management_agent"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ManagementAgentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent", "test_management_agent", acctest.Optional, acctest.Create, managementAgentRepresentation), "managementagent", "managementAgent", t)

	acctest.ResourceTest(t, testAccCheckManagementAgentManagementAgentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + ManagementAgentResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_plugins", "test_management_agent_plugins", acctest.Required, acctest.Create, managementAgentPluginDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent", "test_management_agent", acctest.Required, acctest.Create, managementAgentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + ManagementAgentResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_plugins", "test_management_agent_plugins", acctest.Required, acctest.Create, managementAgentPluginDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent", "test_management_agent", acctest.Optional, acctest.Update, managementAgentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agents", "test_management_agents", acctest.Optional, acctest.Update, managementAgentDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_plugins", "test_management_agent_plugins", acctest.Required, acctest.Create, managementAgentPluginDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentIdVariableStr + ManagementAgentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent", "test_management_agent", acctest.Optional, acctest.Update, managementAgentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "availability_status", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "install_type", "AGENT"),
				resource.TestCheckResourceAttr(datasourceName, "is_customer_deployed", "true"),
				resource.TestCheckResourceAttr(datasourceName, "platform_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "management_agents.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.availability_status"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.host"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.install_key_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.install_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.is_customer_deployed"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.is_agent_auto_upgradable"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.platform_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.platform_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.platform_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.time_last_heartbeat"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.version"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent", "test_management_agent", acctest.Required, acctest.Create, managementAgentSingularDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_plugins", "test_management_agent_plugins", acctest.Required, acctest.Create, managementAgentPluginDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentIdVariableStr + ManagementAgentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "platform_type", "LINUX"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "management_agent_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "install_key_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "install_path"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "install_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_agent_auto_upgradable"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_customer_deployed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "platform_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "platform_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_heartbeat"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + ManagementAgentResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_plugins", "test_management_agent_plugins", acctest.Required, acctest.Create, managementAgentPluginDataSourceRepresentation),
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       false,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckManagementAgentManagementAgentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ManagementAgentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_management_agent_management_agent" {
			noResourceFound = false
			request := oci_management_agent.GetManagementAgentRequest{}

			tmp := rs.Primary.ID
			request.ManagementAgentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "management_agent")

			response, err := client.GetManagementAgent(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_management_agent.LifecycleStatesTerminated): true, string(oci_management_agent.LifecycleStatesDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("ManagementAgentManagementAgent") {
		resource.AddTestSweepers("ManagementAgentManagementAgent", &resource.Sweeper{
			Name:         "ManagementAgentManagementAgent",
			Dependencies: acctest.DependencyGraph["managementAgent"],
			F:            sweepManagementAgentManagementAgentResource,
		})
	}
}

func sweepManagementAgentManagementAgentResource(compartment string) error {
	managementAgentClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementAgentClient()
	managementAgentIds, err := getManagementAgentIds(compartment)
	if err != nil {
		return err
	}
	for _, managementAgentId := range managementAgentIds {
		if ok := acctest.SweeperDefaultResourceId[managementAgentId]; !ok {
			deleteManagementAgentRequest := oci_management_agent.DeleteManagementAgentRequest{}

			deleteManagementAgentRequest.ManagementAgentId = &managementAgentId

			deleteManagementAgentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "management_agent")
			_, error := managementAgentClient.DeleteManagementAgent(context.Background(), deleteManagementAgentRequest)
			if error != nil {
				fmt.Printf("Error deleting ManagementAgent %s %s, It is possible that the resource is already deleted. Please verify manually \n", managementAgentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &managementAgentId, managementAgentSweepWaitCondition, time.Duration(3*time.Minute),
				managementAgentSweepResponseFetchOperation, "management_agent", true)
		}
	}
	return nil
}

func getManagementAgentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ManagementAgentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	managementAgentClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementAgentClient()

	listManagementAgentsRequest := oci_management_agent.ListManagementAgentsRequest{}
	listManagementAgentsRequest.CompartmentId = &compartmentId
	listManagementAgentsRequest.LifecycleState = oci_management_agent.ListManagementAgentsLifecycleStateActive
	listManagementAgentsResponse, err := managementAgentClient.ListManagementAgents(context.Background(), listManagementAgentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ManagementAgent list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, managementAgent := range listManagementAgentsResponse.Items {
		id := *managementAgent.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ManagementAgentId", id)
	}
	return resourceIds, nil
}

func managementAgentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if managementAgentResponse, ok := response.Response.(oci_management_agent.GetManagementAgentResponse); ok {
		return managementAgentResponse.LifecycleState != oci_management_agent.LifecycleStatesTerminated
	}
	return false
}

func managementAgentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ManagementAgentClient().GetManagementAgent(context.Background(), oci_management_agent.GetManagementAgentRequest{
		ManagementAgentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
