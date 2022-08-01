// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Engine for Kubernetes API
//
// API for the Container Engine for Kubernetes service. Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Container Engine for Kubernetes (https://docs.cloud.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

//ContainerEngineClient a client for ContainerEngine
type ContainerEngineClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewContainerEngineClientWithConfigurationProvider Creates a new default ContainerEngine client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewContainerEngineClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ContainerEngineClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newContainerEngineClientFromBaseClient(baseClient, provider)
}

// NewContainerEngineClientWithOboToken Creates a new default ContainerEngine client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewContainerEngineClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ContainerEngineClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newContainerEngineClientFromBaseClient(baseClient, configProvider)
}

func newContainerEngineClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ContainerEngineClient, err error) {
	// ContainerEngine service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ContainerEngine"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ContainerEngineClient{BaseClient: baseClient}
	client.BasePath = "20180222"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ContainerEngineClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("containerengine", "https://containerengine.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ContainerEngineClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("Invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *ContainerEngineClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ClusterMigrateToNativeVcn Initiates cluster migration to use native VCN.
// A default retry strategy applies to this operation ClusterMigrateToNativeVcn()
func (client ContainerEngineClient) ClusterMigrateToNativeVcn(ctx context.Context, request ClusterMigrateToNativeVcnRequest) (response ClusterMigrateToNativeVcnResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.clusterMigrateToNativeVcn, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ClusterMigrateToNativeVcnResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ClusterMigrateToNativeVcnResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ClusterMigrateToNativeVcnResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ClusterMigrateToNativeVcnResponse")
	}
	return
}

// clusterMigrateToNativeVcn implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) clusterMigrateToNativeVcn(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusters/{clusterId}/actions/migrateToNativeVcn", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ClusterMigrateToNativeVcnResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/ClusterMigrateToNativeVcn"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ClusterMigrateToNativeVcn", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCluster Create a new cluster.
// A default retry strategy applies to this operation CreateCluster()
func (client ContainerEngineClient) CreateCluster(ctx context.Context, request CreateClusterRequest) (response CreateClusterResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateClusterResponse")
	}
	return
}

// createCluster implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) createCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/CreateCluster"
		err = common.PostProcessServiceError(err, "ContainerEngine", "CreateCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateKubeconfig Create the Kubeconfig YAML for a cluster.
// A default retry strategy applies to this operation CreateKubeconfig()
func (client ContainerEngineClient) CreateKubeconfig(ctx context.Context, request CreateKubeconfigRequest) (response CreateKubeconfigResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createKubeconfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateKubeconfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateKubeconfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateKubeconfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateKubeconfigResponse")
	}
	return
}

// createKubeconfig implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) createKubeconfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusters/{clusterId}/kubeconfig/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateKubeconfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/CreateKubeconfig"
		err = common.PostProcessServiceError(err, "ContainerEngine", "CreateKubeconfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateNodePool Create a new node pool.
// A default retry strategy applies to this operation CreateNodePool()
func (client ContainerEngineClient) CreateNodePool(ctx context.Context, request CreateNodePoolRequest) (response CreateNodePoolResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createNodePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateNodePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateNodePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateNodePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateNodePoolResponse")
	}
	return
}

// createNodePool implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) createNodePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/nodePools", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateNodePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/NodePool/CreateNodePool"
		err = common.PostProcessServiceError(err, "ContainerEngine", "CreateNodePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVirtualNodePool Create a new virtual node pool.
// A default retry strategy applies to this operation CreateVirtualNodePool()
func (client ContainerEngineClient) CreateVirtualNodePool(ctx context.Context, request CreateVirtualNodePoolRequest) (response CreateVirtualNodePoolResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createVirtualNodePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVirtualNodePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVirtualNodePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVirtualNodePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVirtualNodePoolResponse")
	}
	return
}

// createVirtualNodePool implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) createVirtualNodePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/virtualNodePools", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVirtualNodePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/VirtualNodePool/CreateVirtualNodePool"
		err = common.PostProcessServiceError(err, "ContainerEngine", "CreateVirtualNodePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCluster Delete a cluster.
// A default retry strategy applies to this operation DeleteCluster()
func (client ContainerEngineClient) DeleteCluster(ctx context.Context, request DeleteClusterRequest) (response DeleteClusterResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteClusterResponse")
	}
	return
}

// deleteCluster implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) deleteCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/clusters/{clusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/DeleteCluster"
		err = common.PostProcessServiceError(err, "ContainerEngine", "DeleteCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteNode Delete node.
// A default retry strategy applies to this operation DeleteNode()
func (client ContainerEngineClient) DeleteNode(ctx context.Context, request DeleteNodeRequest) (response DeleteNodeResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteNodeResponse")
	}
	return
}

// deleteNode implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) deleteNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/nodePools/{nodePoolId}/node/{nodeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/NodePool/DeleteNode"
		err = common.PostProcessServiceError(err, "ContainerEngine", "DeleteNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteNodePool Delete a node pool.
// A default retry strategy applies to this operation DeleteNodePool()
func (client ContainerEngineClient) DeleteNodePool(ctx context.Context, request DeleteNodePoolRequest) (response DeleteNodePoolResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteNodePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteNodePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteNodePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteNodePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteNodePoolResponse")
	}
	return
}

// deleteNodePool implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) deleteNodePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/nodePools/{nodePoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteNodePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/NodePool/DeleteNodePool"
		err = common.PostProcessServiceError(err, "ContainerEngine", "DeleteNodePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVirtualNode Delete virtual node.
// A default retry strategy applies to this operation DeleteVirtualNode()
func (client ContainerEngineClient) DeleteVirtualNode(ctx context.Context, request DeleteVirtualNodeRequest) (response DeleteVirtualNodeResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteVirtualNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVirtualNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVirtualNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVirtualNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVirtualNodeResponse")
	}
	return
}

// deleteVirtualNode implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) deleteVirtualNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/virtualNodePools/{virtualNodePoolId}/virtualNodes/{virtualNodeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVirtualNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/VirtualNodePool/DeleteVirtualNode"
		err = common.PostProcessServiceError(err, "ContainerEngine", "DeleteVirtualNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVirtualNodePool Delete a virtual node pool.
// A default retry strategy applies to this operation DeleteVirtualNodePool()
func (client ContainerEngineClient) DeleteVirtualNodePool(ctx context.Context, request DeleteVirtualNodePoolRequest) (response DeleteVirtualNodePoolResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteVirtualNodePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVirtualNodePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVirtualNodePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVirtualNodePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVirtualNodePoolResponse")
	}
	return
}

// deleteVirtualNodePool implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) deleteVirtualNodePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/virtualNodePools/{virtualNodePoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVirtualNodePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/VirtualNodePool/DeleteVirtualNodePool"
		err = common.PostProcessServiceError(err, "ContainerEngine", "DeleteVirtualNodePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteWorkRequest Cancel a work request that has not started.
// A default retry strategy applies to this operation DeleteWorkRequest()
func (client ContainerEngineClient) DeleteWorkRequest(ctx context.Context, request DeleteWorkRequestRequest) (response DeleteWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteWorkRequestResponse")
	}
	return
}

// deleteWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) deleteWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/WorkRequest/DeleteWorkRequest"
		err = common.PostProcessServiceError(err, "ContainerEngine", "DeleteWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCluster Get the details of a cluster.
// A default retry strategy applies to this operation GetCluster()
func (client ContainerEngineClient) GetCluster(ctx context.Context, request GetClusterRequest) (response GetClusterResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetClusterResponse")
	}
	return
}

// getCluster implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusters/{clusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/GetCluster"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetClusterMigrateToNativeVcnStatus Get details on a cluster's migration to native VCN.
// A default retry strategy applies to this operation GetClusterMigrateToNativeVcnStatus()
func (client ContainerEngineClient) GetClusterMigrateToNativeVcnStatus(ctx context.Context, request GetClusterMigrateToNativeVcnStatusRequest) (response GetClusterMigrateToNativeVcnStatusResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getClusterMigrateToNativeVcnStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetClusterMigrateToNativeVcnStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetClusterMigrateToNativeVcnStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetClusterMigrateToNativeVcnStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetClusterMigrateToNativeVcnStatusResponse")
	}
	return
}

// getClusterMigrateToNativeVcnStatus implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getClusterMigrateToNativeVcnStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusters/{clusterId}/migrateToNativeVcnStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetClusterMigrateToNativeVcnStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterMigrateToNativeVcnStatus/GetClusterMigrateToNativeVcnStatus"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetClusterMigrateToNativeVcnStatus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetClusterOptions Get options available for clusters.
// A default retry strategy applies to this operation GetClusterOptions()
func (client ContainerEngineClient) GetClusterOptions(ctx context.Context, request GetClusterOptionsRequest) (response GetClusterOptionsResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getClusterOptions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetClusterOptionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetClusterOptionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetClusterOptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetClusterOptionsResponse")
	}
	return
}

// getClusterOptions implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getClusterOptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusterOptions/{clusterOptionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetClusterOptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterOptions/GetClusterOptions"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetClusterOptions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetNodePool Get the details of a node pool.
// A default retry strategy applies to this operation GetNodePool()
func (client ContainerEngineClient) GetNodePool(ctx context.Context, request GetNodePoolRequest) (response GetNodePoolResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNodePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetNodePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetNodePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNodePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNodePoolResponse")
	}
	return
}

// getNodePool implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getNodePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/nodePools/{nodePoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetNodePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/NodePool/GetNodePool"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetNodePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetNodePoolOptions Get options available for node pools.
// A default retry strategy applies to this operation GetNodePoolOptions()
func (client ContainerEngineClient) GetNodePoolOptions(ctx context.Context, request GetNodePoolOptionsRequest) (response GetNodePoolOptionsResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNodePoolOptions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetNodePoolOptionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetNodePoolOptionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNodePoolOptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNodePoolOptionsResponse")
	}
	return
}

// getNodePoolOptions implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getNodePoolOptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/nodePoolOptions/{nodePoolOptionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetNodePoolOptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/NodePoolOptions/GetNodePoolOptions"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetNodePoolOptions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVirtualNode Get the details of a virtual node.
// A default retry strategy applies to this operation GetVirtualNode()
func (client ContainerEngineClient) GetVirtualNode(ctx context.Context, request GetVirtualNodeRequest) (response GetVirtualNodeResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVirtualNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVirtualNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVirtualNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVirtualNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVirtualNodeResponse")
	}
	return
}

// getVirtualNode implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getVirtualNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/virtualNodePools/{virtualNodePoolId}/virtualNodes/{virtualNodeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVirtualNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/VirtualNodePool/GetVirtualNode"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetVirtualNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVirtualNodePool Get the details of a virtual node pool.
// A default retry strategy applies to this operation GetVirtualNodePool()
func (client ContainerEngineClient) GetVirtualNodePool(ctx context.Context, request GetVirtualNodePoolRequest) (response GetVirtualNodePoolResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVirtualNodePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVirtualNodePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVirtualNodePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVirtualNodePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVirtualNodePoolResponse")
	}
	return
}

// getVirtualNodePool implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getVirtualNodePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/virtualNodePools/{virtualNodePoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVirtualNodePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/VirtualNodePool/GetVirtualNodePool"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetVirtualNodePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Get the details of a work request.
// A default retry strategy applies to this operation GetWorkRequest()
func (client ContainerEngineClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWorkRequestResponse")
	}
	return
}

// getWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListClusters List all the cluster objects in a compartment.
// A default retry strategy applies to this operation ListClusters()
func (client ContainerEngineClient) ListClusters(ctx context.Context, request ListClustersRequest) (response ListClustersResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listClusters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListClustersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListClustersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListClustersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListClustersResponse")
	}
	return
}

// listClusters implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) listClusters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListClustersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterSummary/ListClusters"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListClusters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListNodePools List all the node pools in a compartment, and optionally filter by cluster.
// A default retry strategy applies to this operation ListNodePools()
func (client ContainerEngineClient) ListNodePools(ctx context.Context, request ListNodePoolsRequest) (response ListNodePoolsResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listNodePools, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListNodePoolsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListNodePoolsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListNodePoolsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListNodePoolsResponse")
	}
	return
}

// listNodePools implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) listNodePools(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/nodePools", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListNodePoolsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/NodePoolSummary/ListNodePools"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListNodePools", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPodShapes List all the Pod Shapes in a compartment.
// A default retry strategy applies to this operation ListPodShapes()
func (client ContainerEngineClient) ListPodShapes(ctx context.Context, request ListPodShapesRequest) (response ListPodShapesResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPodShapes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPodShapesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPodShapesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPodShapesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPodShapesResponse")
	}
	return
}

// listPodShapes implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) listPodShapes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/podShapes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPodShapesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/PodShapeSummary/ListPodShapes"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListPodShapes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVirtualNodePools List all the virtual node pools in a compartment, and optionally filter by cluster.
// A default retry strategy applies to this operation ListVirtualNodePools()
func (client ContainerEngineClient) ListVirtualNodePools(ctx context.Context, request ListVirtualNodePoolsRequest) (response ListVirtualNodePoolsResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVirtualNodePools, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVirtualNodePoolsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVirtualNodePoolsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVirtualNodePoolsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVirtualNodePoolsResponse")
	}
	return
}

// listVirtualNodePools implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) listVirtualNodePools(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/virtualNodePools", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVirtualNodePoolsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/VirtualNodePoolSummary/ListVirtualNodePools"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListVirtualNodePools", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVirtualNodes List virtual nodes in a virtual node pool.
// A default retry strategy applies to this operation ListVirtualNodes()
func (client ContainerEngineClient) ListVirtualNodes(ctx context.Context, request ListVirtualNodesRequest) (response ListVirtualNodesResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVirtualNodes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVirtualNodesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVirtualNodesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVirtualNodesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVirtualNodesResponse")
	}
	return
}

// listVirtualNodes implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) listVirtualNodes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/virtualNodePools/{virtualNodePoolId}/virtualNodes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVirtualNodesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/VirtualNodePool/ListVirtualNodes"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListVirtualNodes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Get the errors of a work request.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client ContainerEngineClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestErrorsResponse")
	}
	return
}

// listWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Get the logs of a work request.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client ContainerEngineClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestLogsResponse")
	}
	return
}

// listWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests List all work requests in a compartment.
// A default retry strategy applies to this operation ListWorkRequests()
func (client ContainerEngineClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestsResponse")
	}
	return
}

// listWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/WorkRequestSummary/ListWorkRequests"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCluster Update the details of a cluster.
// A default retry strategy applies to this operation UpdateCluster()
func (client ContainerEngineClient) UpdateCluster(ctx context.Context, request UpdateClusterRequest) (response UpdateClusterResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateClusterResponse")
	}
	return
}

// updateCluster implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) updateCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/clusters/{clusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/UpdateCluster"
		err = common.PostProcessServiceError(err, "ContainerEngine", "UpdateCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateClusterEndpointConfig Update the details of the cluster endpoint configuration.
// A default retry strategy applies to this operation UpdateClusterEndpointConfig()
func (client ContainerEngineClient) UpdateClusterEndpointConfig(ctx context.Context, request UpdateClusterEndpointConfigRequest) (response UpdateClusterEndpointConfigResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateClusterEndpointConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateClusterEndpointConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateClusterEndpointConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateClusterEndpointConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateClusterEndpointConfigResponse")
	}
	return
}

// updateClusterEndpointConfig implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) updateClusterEndpointConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusters/{clusterId}/actions/updateEndpointConfig", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateClusterEndpointConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/UpdateClusterEndpointConfig"
		err = common.PostProcessServiceError(err, "ContainerEngine", "UpdateClusterEndpointConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateNodePool Update the details of a node pool.
// A default retry strategy applies to this operation UpdateNodePool()
func (client ContainerEngineClient) UpdateNodePool(ctx context.Context, request UpdateNodePoolRequest) (response UpdateNodePoolResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateNodePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateNodePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateNodePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateNodePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateNodePoolResponse")
	}
	return
}

// updateNodePool implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) updateNodePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/nodePools/{nodePoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateNodePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/NodePool/UpdateNodePool"
		err = common.PostProcessServiceError(err, "ContainerEngine", "UpdateNodePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateVirtualNodePool Update the details of a virtual node pool.
// A default retry strategy applies to this operation UpdateVirtualNodePool()
func (client ContainerEngineClient) UpdateVirtualNodePool(ctx context.Context, request UpdateVirtualNodePoolRequest) (response UpdateVirtualNodePoolResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.DefaultComplexRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateVirtualNodePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateVirtualNodePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateVirtualNodePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateVirtualNodePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateVirtualNodePoolResponse")
	}
	return
}

// updateVirtualNodePool implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) updateVirtualNodePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/virtualNodePools/{virtualNodePoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateVirtualNodePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/VirtualNodePool/UpdateVirtualNodePool"
		err = common.PostProcessServiceError(err, "ContainerEngine", "UpdateVirtualNodePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
