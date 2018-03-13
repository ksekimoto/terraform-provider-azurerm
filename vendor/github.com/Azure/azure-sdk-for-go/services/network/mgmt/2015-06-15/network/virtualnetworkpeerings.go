package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// VirtualNetworkPeeringsClient is the network Client
type VirtualNetworkPeeringsClient struct {
	BaseClient
}

// NewVirtualNetworkPeeringsClient creates an instance of the VirtualNetworkPeeringsClient client.
func NewVirtualNetworkPeeringsClient(subscriptionID string) VirtualNetworkPeeringsClient {
	return NewVirtualNetworkPeeringsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualNetworkPeeringsClientWithBaseURI creates an instance of the VirtualNetworkPeeringsClient client.
func NewVirtualNetworkPeeringsClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworkPeeringsClient {
	return VirtualNetworkPeeringsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a peering in the specified virtual network.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is the name of the virtual network.
// virtualNetworkPeeringName is the name of the peering. virtualNetworkPeeringParameters is parameters supplied to the
// create or update virtual network peering operation.
func (client VirtualNetworkPeeringsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, virtualNetworkPeeringParameters VirtualNetworkPeering) (result VirtualNetworkPeeringsCreateOrUpdateFuture, err error) {
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, virtualNetworkName, virtualNetworkPeeringName, virtualNetworkPeeringParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkPeeringsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkPeeringsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualNetworkPeeringsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, virtualNetworkPeeringParameters VirtualNetworkPeering) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName":        autorest.Encode("path", virtualNetworkName),
		"virtualNetworkPeeringName": autorest.Encode("path", virtualNetworkPeeringName),
	}

	const APIVersion = "2015-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}", pathParameters),
		autorest.WithJSON(virtualNetworkPeeringParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkPeeringsClient) CreateOrUpdateSender(req *http.Request) (future VirtualNetworkPeeringsCreateOrUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated))
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualNetworkPeeringsClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualNetworkPeering, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified virtual network peering.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is the name of the virtual network.
// virtualNetworkPeeringName is the name of the virtual network peering.
func (client VirtualNetworkPeeringsClient) Delete(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string) (result VirtualNetworkPeeringsDeleteFuture, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, virtualNetworkName, virtualNetworkPeeringName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkPeeringsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkPeeringsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VirtualNetworkPeeringsClient) DeletePreparer(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName":        autorest.Encode("path", virtualNetworkName),
		"virtualNetworkPeeringName": autorest.Encode("path", virtualNetworkPeeringName),
	}

	const APIVersion = "2015-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkPeeringsClient) DeleteSender(req *http.Request) (future VirtualNetworkPeeringsDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualNetworkPeeringsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified virtual network peering.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is the name of the virtual network.
// virtualNetworkPeeringName is the name of the virtual network peering.
func (client VirtualNetworkPeeringsClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string) (result VirtualNetworkPeering, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, virtualNetworkName, virtualNetworkPeeringName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkPeeringsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkPeeringsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkPeeringsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualNetworkPeeringsClient) GetPreparer(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName":        autorest.Encode("path", virtualNetworkName),
		"virtualNetworkPeeringName": autorest.Encode("path", virtualNetworkPeeringName),
	}

	const APIVersion = "2015-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkPeeringsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualNetworkPeeringsClient) GetResponder(resp *http.Response) (result VirtualNetworkPeering, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all virtual network peerings in a virtual network.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is the name of the virtual network.
func (client VirtualNetworkPeeringsClient) List(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result VirtualNetworkPeeringListResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, virtualNetworkName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkPeeringsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.vnplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkPeeringsClient", "List", resp, "Failure sending request")
		return
	}

	result.vnplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkPeeringsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualNetworkPeeringsClient) ListPreparer(ctx context.Context, resourceGroupName string, virtualNetworkName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
	}

	const APIVersion = "2015-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkPeeringsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualNetworkPeeringsClient) ListResponder(resp *http.Response) (result VirtualNetworkPeeringListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client VirtualNetworkPeeringsClient) listNextResults(lastResults VirtualNetworkPeeringListResult) (result VirtualNetworkPeeringListResult, err error) {
	req, err := lastResults.virtualNetworkPeeringListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworkPeeringsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworkPeeringsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkPeeringsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualNetworkPeeringsClient) ListComplete(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result VirtualNetworkPeeringListResultIterator, err error) {
	result.page, err = client.List(ctx, resourceGroupName, virtualNetworkName)
	return
}