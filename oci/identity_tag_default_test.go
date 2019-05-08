// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	TagDefaultResourceConfig = TagDefaultResourceDependencies +
		generateResourceFromRepresentationMap("oci_identity_tag_default", "test_tag_default", Optional, Update, tagDefaultRepresentation)

	tagDefaultSingularDataSourceRepresentation = map[string]interface{}{
		"tag_default_id": Representation{repType: Required, create: `${oci_identity_tag_default.test_tag_default.id}`},
	}

	tagDefaultDataSourceRepresentationWithCompartmentIdFilter = map[string]interface{}{
		"compartment_id": Representation{repType: Optional, create: `${var.compartment_id}`},
		"filter":         RepresentationGroup{Required, tagDefaultDataSourceFilterRepresentation}}
	tagDefaultDataSourceRepresentationWithIdFilter = map[string]interface{}{
		"id":     Representation{repType: Optional, create: `${oci_identity_tag_default.test_tag_default.id}`},
		"filter": RepresentationGroup{Required, tagDefaultDataSourceFilterRepresentation}}
	tagDefaultDataSourceRepresentationWithStateFilter = map[string]interface{}{
		"compartment_id": Representation{repType: Optional, create: `${var.compartment_id}`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"filter":         RepresentationGroup{Required, tagDefaultDataSourceFilterRepresentation}}
	tagDefaultDataSourceRepresentationWithTagDefinitionIdFilter = map[string]interface{}{
		"compartment_id":    Representation{repType: Optional, create: `${var.compartment_id}`},
		"tag_definition_id": Representation{repType: Optional, create: `${oci_identity_tag.test_tag.id}`},
		"filter":            RepresentationGroup{Required, tagDefaultDataSourceFilterRepresentation}}
	tagDefaultDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_identity_tag_default.test_tag_default.id}`}},
	}

	tagDefaultRepresentation = map[string]interface{}{
		"compartment_id":    Representation{repType: Required, create: `${var.compartment_id}`},
		"tag_definition_id": Representation{repType: Required, create: `${oci_identity_tag.test_tag.id}`},
		"value":             Representation{repType: Required, create: `W123`, update: `value2`},
	}

	TagDefaultResourceDependencies = TagRequiredOnlyResource
)

func TestIdentityTagDefaultResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityTagDefaultResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentIdCreate := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentId := getEnvSettingWithDefault("compartment_id_for_update", compartmentIdCreate)
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_tag_default.test_tag_default"
	datasourceName := "data.oci_identity_tag_defaults.test_tag_defaults"
	singularDatasourceName := "data.oci_identity_tag_default.test_tag_default"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckIdentityTagDefaultDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + TagDefaultResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_tag_default", "test_tag_default", Required, Create, tagDefaultRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "tag_definition_id"),
					resource.TestCheckResourceAttr(resourceName, "value", "W123"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + TagDefaultResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_tag_default", "test_tag_default", Optional, Update, tagDefaultRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "tag_definition_id"),
					resource.TestCheckResourceAttrSet(resourceName, "tag_definition_name"),
					resource.TestCheckResourceAttrSet(resourceName, "tag_namespace_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "value", "value2"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_identity_tag_defaults", "test_tag_defaults_with_compartment_id_filter", Optional, Update, tagDefaultDataSourceRepresentationWithCompartmentIdFilter) +
					generateDataSourceFromRepresentationMap("oci_identity_tag_defaults", "test_tag_defaults_with_id_filter", Optional, Update, tagDefaultDataSourceRepresentationWithIdFilter) +
					generateDataSourceFromRepresentationMap("oci_identity_tag_defaults", "test_tag_defaults_with_state_filter", Optional, Update, tagDefaultDataSourceRepresentationWithStateFilter) +
					generateDataSourceFromRepresentationMap("oci_identity_tag_defaults", "test_tag_defaults_with_tag_definition_id_filter", Optional, Update, tagDefaultDataSourceRepresentationWithTagDefinitionIdFilter) +
					compartmentIdVariableStr + TagDefaultResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_tag_default", "test_tag_default", Optional, Update, tagDefaultRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName+"_with_compartment_id_filter", "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName+"_with_compartment_id_filter", "tag_defaults.#"),
					resource.TestCheckResourceAttr(datasourceName+"_with_compartment_id_filter", "tag_defaults.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_compartment_id_filter", "tag_defaults.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_compartment_id_filter", "tag_defaults.0.tag_definition_id"),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_compartment_id_filter", "tag_defaults.0.tag_definition_name"),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_compartment_id_filter", "tag_defaults.0.tag_namespace_id"),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_compartment_id_filter", "tag_defaults.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName+"_with_compartment_id_filter", "tag_defaults.0.value", "value2"),

					resource.TestCheckResourceAttrSet(datasourceName+"_with_id_filter", "id"),

					resource.TestCheckResourceAttrSet(datasourceName+"_with_id_filter", "tag_defaults.#"),
					resource.TestCheckResourceAttr(datasourceName+"_with_id_filter", "tag_defaults.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_id_filter", "tag_defaults.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_id_filter", "tag_defaults.0.tag_definition_id"),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_id_filter", "tag_defaults.0.tag_definition_name"),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_id_filter", "tag_defaults.0.tag_namespace_id"),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_id_filter", "tag_defaults.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName+"_with_id_filter", "tag_defaults.0.value", "value2"),

					resource.TestCheckResourceAttr(datasourceName+"_with_state_filter", "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName+"_with_state_filter", "state", "AVAILABLE"),

					resource.TestCheckResourceAttrSet(datasourceName+"_with_state_filter", "tag_defaults.#"),
					resource.TestCheckResourceAttr(datasourceName+"_with_state_filter", "tag_defaults.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_state_filter", "tag_defaults.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_state_filter", "tag_defaults.0.tag_definition_id"),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_state_filter", "tag_defaults.0.tag_definition_name"),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_state_filter", "tag_defaults.0.tag_namespace_id"),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_state_filter", "tag_defaults.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName+"_with_state_filter", "tag_defaults.0.value", "value2"),

					resource.TestCheckResourceAttr(datasourceName+"_with_tag_definition_id_filter", "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_tag_definition_id_filter", "tag_definition_id"),

					resource.TestCheckResourceAttrSet(datasourceName+"_with_tag_definition_id_filter", "tag_defaults.#"),
					resource.TestCheckResourceAttr(datasourceName+"_with_tag_definition_id_filter", "tag_defaults.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_tag_definition_id_filter", "tag_defaults.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_tag_definition_id_filter", "tag_defaults.0.tag_definition_id"),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_tag_definition_id_filter", "tag_defaults.0.tag_definition_name"),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_tag_definition_id_filter", "tag_defaults.0.tag_namespace_id"),
					resource.TestCheckResourceAttrSet(datasourceName+"_with_tag_definition_id_filter", "tag_defaults.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName+"_with_tag_definition_id_filter", "tag_defaults.0.value", "value2"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_identity_tag_default", "test_tag_default", Required, Create, tagDefaultSingularDataSourceRepresentation) +
					compartmentIdVariableStr + TagDefaultResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "tag_default_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "tag_definition_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "tag_definition_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "tag_namespace_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttr(singularDatasourceName, "value", "value2"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + TagDefaultResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckIdentityTagDefaultDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_tag_default" {
			noResourceFound = false
			request := oci_identity.GetTagDefaultRequest{}

			tmp := rs.Primary.ID
			request.TagDefaultId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "identity")

			_, err := client.GetTagDefault(context.Background(), request)

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
	resource.AddTestSweepers("IdentityTagDefault", &resource.Sweeper{
		Name:         "IdentityTagDefault",
		Dependencies: DependencyGraph["tagDefault"],
		F:            sweepIdentityTagDefaultResource,
	})
}

func sweepIdentityTagDefaultResource(compartment string) error {
	identityClient := GetTestClients(&schema.ResourceData{}).identityClient
	tagDefaultIds, err := getTagDefaultIds(compartment)
	if err != nil {
		return err
	}
	for _, tagDefaultId := range tagDefaultIds {
		if ok := SweeperDefaultResourceId[tagDefaultId]; !ok {
			deleteTagDefaultRequest := oci_identity.DeleteTagDefaultRequest{}

			deleteTagDefaultRequest.TagDefaultId = &tagDefaultId

			deleteTagDefaultRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "identity")
			_, error := identityClient.DeleteTagDefault(context.Background(), deleteTagDefaultRequest)
			if error != nil {
				fmt.Printf("Error deleting TagDefault %s %s, It is possible that the resource is already deleted. Please verify manually \n", tagDefaultId, error)
				continue
			}
		}
	}
	return nil
}

func getTagDefaultIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "TagDefaultId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityClient := GetTestClients(&schema.ResourceData{}).identityClient

	listTagDefaultsRequest := oci_identity.ListTagDefaultsRequest{}
	listTagDefaultsRequest.CompartmentId = &compartmentId
	listTagDefaultsResponse, err := identityClient.ListTagDefaults(context.Background(), listTagDefaultsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting TagDefault list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, tagDefault := range listTagDefaultsResponse.Items {
		id := *tagDefault.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "TagDefaultId", id)
	}
	return resourceIds, nil
}
