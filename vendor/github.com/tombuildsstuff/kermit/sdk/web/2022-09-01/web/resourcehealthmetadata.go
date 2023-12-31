package web

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
)

// ResourceHealthMetadataClient is the webSite Management Client
type ResourceHealthMetadataClient struct {
	BaseClient
}

// NewResourceHealthMetadataClient creates an instance of the ResourceHealthMetadataClient client.
func NewResourceHealthMetadataClient(subscriptionID string) ResourceHealthMetadataClient {
	return NewResourceHealthMetadataClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewResourceHealthMetadataClientWithBaseURI creates an instance of the ResourceHealthMetadataClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewResourceHealthMetadataClientWithBaseURI(baseURI string, subscriptionID string) ResourceHealthMetadataClient {
	return ResourceHealthMetadataClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetBySite description for Gets the category of ResourceHealthMetadata to use for the given site
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// name - name of web app
func (client ResourceHealthMetadataClient) GetBySite(ctx context.Context, resourceGroupName string, name string) (result ResourceHealthMetadata, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceHealthMetadataClient.GetBySite")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.ResourceHealthMetadataClient", "GetBySite", err.Error())
	}

	req, err := client.GetBySitePreparer(ctx, resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "GetBySite", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetBySiteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "GetBySite", resp, "Failure sending request")
		return
	}

	result, err = client.GetBySiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "GetBySite", resp, "Failure responding to request")
		return
	}

	return
}

// GetBySitePreparer prepares the GetBySite request.
func (client ResourceHealthMetadataClient) GetBySitePreparer(ctx context.Context, resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/resourceHealthMetadata/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetBySiteSender sends the GetBySite request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceHealthMetadataClient) GetBySiteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetBySiteResponder handles the response to the GetBySite request. The method always
// closes the http.Response Body.
func (client ResourceHealthMetadataClient) GetBySiteResponder(resp *http.Response) (result ResourceHealthMetadata, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetBySiteSlot description for Gets the category of ResourceHealthMetadata to use for the given site
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// name - name of web app
// slot - name of web app slot. If not specified then will default to production slot.
func (client ResourceHealthMetadataClient) GetBySiteSlot(ctx context.Context, resourceGroupName string, name string, slot string) (result ResourceHealthMetadata, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceHealthMetadataClient.GetBySiteSlot")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.ResourceHealthMetadataClient", "GetBySiteSlot", err.Error())
	}

	req, err := client.GetBySiteSlotPreparer(ctx, resourceGroupName, name, slot)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "GetBySiteSlot", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetBySiteSlotSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "GetBySiteSlot", resp, "Failure sending request")
		return
	}

	result, err = client.GetBySiteSlotResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "GetBySiteSlot", resp, "Failure responding to request")
		return
	}

	return
}

// GetBySiteSlotPreparer prepares the GetBySiteSlot request.
func (client ResourceHealthMetadataClient) GetBySiteSlotPreparer(ctx context.Context, resourceGroupName string, name string, slot string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"slot":              autorest.Encode("path", slot),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/resourceHealthMetadata/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetBySiteSlotSender sends the GetBySiteSlot request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceHealthMetadataClient) GetBySiteSlotSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetBySiteSlotResponder handles the response to the GetBySiteSlot request. The method always
// closes the http.Response Body.
func (client ResourceHealthMetadataClient) GetBySiteSlotResponder(resp *http.Response) (result ResourceHealthMetadata, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List description for List all ResourceHealthMetadata for all sites in the subscription.
func (client ResourceHealthMetadataClient) List(ctx context.Context) (result ResourceHealthMetadataCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceHealthMetadataClient.List")
		defer func() {
			sc := -1
			if result.rhmc.Response.Response != nil {
				sc = result.rhmc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rhmc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "List", resp, "Failure sending request")
		return
	}

	result.rhmc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "List", resp, "Failure responding to request")
		return
	}
	if result.rhmc.hasNextLink() && result.rhmc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ResourceHealthMetadataClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Web/resourceHealthMetadata", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceHealthMetadataClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ResourceHealthMetadataClient) ListResponder(resp *http.Response) (result ResourceHealthMetadataCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ResourceHealthMetadataClient) listNextResults(ctx context.Context, lastResults ResourceHealthMetadataCollection) (result ResourceHealthMetadataCollection, err error) {
	req, err := lastResults.resourceHealthMetadataCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ResourceHealthMetadataClient) ListComplete(ctx context.Context) (result ResourceHealthMetadataCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceHealthMetadataClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup description for List all ResourceHealthMetadata for all sites in the resource group in the
// subscription.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
func (client ResourceHealthMetadataClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ResourceHealthMetadataCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceHealthMetadataClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.rhmc.Response.Response != nil {
				sc = result.rhmc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.ResourceHealthMetadataClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.rhmc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.rhmc, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.rhmc.hasNextLink() && result.rhmc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ResourceHealthMetadataClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/resourceHealthMetadata", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceHealthMetadataClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ResourceHealthMetadataClient) ListByResourceGroupResponder(resp *http.Response) (result ResourceHealthMetadataCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ResourceHealthMetadataClient) listByResourceGroupNextResults(ctx context.Context, lastResults ResourceHealthMetadataCollection) (result ResourceHealthMetadataCollection, err error) {
	req, err := lastResults.resourceHealthMetadataCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ResourceHealthMetadataClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ResourceHealthMetadataCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceHealthMetadataClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// ListBySite description for Gets the category of ResourceHealthMetadata to use for the given site as a collection
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// name - name of web app.
func (client ResourceHealthMetadataClient) ListBySite(ctx context.Context, resourceGroupName string, name string) (result ResourceHealthMetadataCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceHealthMetadataClient.ListBySite")
		defer func() {
			sc := -1
			if result.rhmc.Response.Response != nil {
				sc = result.rhmc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.ResourceHealthMetadataClient", "ListBySite", err.Error())
	}

	result.fn = client.listBySiteNextResults
	req, err := client.ListBySitePreparer(ctx, resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "ListBySite", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySiteSender(req)
	if err != nil {
		result.rhmc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "ListBySite", resp, "Failure sending request")
		return
	}

	result.rhmc, err = client.ListBySiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "ListBySite", resp, "Failure responding to request")
		return
	}
	if result.rhmc.hasNextLink() && result.rhmc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySitePreparer prepares the ListBySite request.
func (client ResourceHealthMetadataClient) ListBySitePreparer(ctx context.Context, resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/resourceHealthMetadata", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySiteSender sends the ListBySite request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceHealthMetadataClient) ListBySiteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySiteResponder handles the response to the ListBySite request. The method always
// closes the http.Response Body.
func (client ResourceHealthMetadataClient) ListBySiteResponder(resp *http.Response) (result ResourceHealthMetadataCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySiteNextResults retrieves the next set of results, if any.
func (client ResourceHealthMetadataClient) listBySiteNextResults(ctx context.Context, lastResults ResourceHealthMetadataCollection) (result ResourceHealthMetadataCollection, err error) {
	req, err := lastResults.resourceHealthMetadataCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listBySiteNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySiteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listBySiteNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listBySiteNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySiteComplete enumerates all values, automatically crossing page boundaries as required.
func (client ResourceHealthMetadataClient) ListBySiteComplete(ctx context.Context, resourceGroupName string, name string) (result ResourceHealthMetadataCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceHealthMetadataClient.ListBySite")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySite(ctx, resourceGroupName, name)
	return
}

// ListBySiteSlot description for Gets the category of ResourceHealthMetadata to use for the given site as a collection
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// name - name of web app.
// slot - name of web app slot. If not specified then will default to production slot.
func (client ResourceHealthMetadataClient) ListBySiteSlot(ctx context.Context, resourceGroupName string, name string, slot string) (result ResourceHealthMetadataCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceHealthMetadataClient.ListBySiteSlot")
		defer func() {
			sc := -1
			if result.rhmc.Response.Response != nil {
				sc = result.rhmc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.ResourceHealthMetadataClient", "ListBySiteSlot", err.Error())
	}

	result.fn = client.listBySiteSlotNextResults
	req, err := client.ListBySiteSlotPreparer(ctx, resourceGroupName, name, slot)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "ListBySiteSlot", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySiteSlotSender(req)
	if err != nil {
		result.rhmc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "ListBySiteSlot", resp, "Failure sending request")
		return
	}

	result.rhmc, err = client.ListBySiteSlotResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "ListBySiteSlot", resp, "Failure responding to request")
		return
	}
	if result.rhmc.hasNextLink() && result.rhmc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySiteSlotPreparer prepares the ListBySiteSlot request.
func (client ResourceHealthMetadataClient) ListBySiteSlotPreparer(ctx context.Context, resourceGroupName string, name string, slot string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"slot":              autorest.Encode("path", slot),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/resourceHealthMetadata", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySiteSlotSender sends the ListBySiteSlot request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceHealthMetadataClient) ListBySiteSlotSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySiteSlotResponder handles the response to the ListBySiteSlot request. The method always
// closes the http.Response Body.
func (client ResourceHealthMetadataClient) ListBySiteSlotResponder(resp *http.Response) (result ResourceHealthMetadataCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySiteSlotNextResults retrieves the next set of results, if any.
func (client ResourceHealthMetadataClient) listBySiteSlotNextResults(ctx context.Context, lastResults ResourceHealthMetadataCollection) (result ResourceHealthMetadataCollection, err error) {
	req, err := lastResults.resourceHealthMetadataCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listBySiteSlotNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySiteSlotSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listBySiteSlotNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySiteSlotResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ResourceHealthMetadataClient", "listBySiteSlotNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySiteSlotComplete enumerates all values, automatically crossing page boundaries as required.
func (client ResourceHealthMetadataClient) ListBySiteSlotComplete(ctx context.Context, resourceGroupName string, name string, slot string) (result ResourceHealthMetadataCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceHealthMetadataClient.ListBySiteSlot")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySiteSlot(ctx, resourceGroupName, name, slot)
	return
}
