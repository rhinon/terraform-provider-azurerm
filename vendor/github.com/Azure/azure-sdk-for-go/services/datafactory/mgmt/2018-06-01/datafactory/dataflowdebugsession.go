package datafactory

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DataFlowDebugSessionClient is the the Azure Data Factory V2 management API provides a RESTful set of web services
// that interact with Azure Data Factory V2 services.
type DataFlowDebugSessionClient struct {
	BaseClient
}

// NewDataFlowDebugSessionClient creates an instance of the DataFlowDebugSessionClient client.
func NewDataFlowDebugSessionClient(subscriptionID string) DataFlowDebugSessionClient {
	return NewDataFlowDebugSessionClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDataFlowDebugSessionClientWithBaseURI creates an instance of the DataFlowDebugSessionClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewDataFlowDebugSessionClientWithBaseURI(baseURI string, subscriptionID string) DataFlowDebugSessionClient {
	return DataFlowDebugSessionClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// AddDataFlow add a data flow into debug session.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// request - data flow debug session definition with debug content.
func (client DataFlowDebugSessionClient) AddDataFlow(ctx context.Context, resourceGroupName string, factoryName string, request DataFlowDebugPackage) (result AddDataFlowToDebugSessionResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataFlowDebugSessionClient.AddDataFlow")
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
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: request,
			Constraints: []validation.Constraint{{Target: "request.Staging", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "request.Staging.LinkedService", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "request.Staging.LinkedService.Type", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "request.Staging.LinkedService.ReferenceName", Name: validation.Null, Rule: true, Chain: nil},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("datafactory.DataFlowDebugSessionClient", "AddDataFlow", err.Error())
	}

	req, err := client.AddDataFlowPreparer(ctx, resourceGroupName, factoryName, request)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionClient", "AddDataFlow", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddDataFlowSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionClient", "AddDataFlow", resp, "Failure sending request")
		return
	}

	result, err = client.AddDataFlowResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionClient", "AddDataFlow", resp, "Failure responding to request")
		return
	}

	return
}

// AddDataFlowPreparer prepares the AddDataFlow request.
func (client DataFlowDebugSessionClient) AddDataFlowPreparer(ctx context.Context, resourceGroupName string, factoryName string, request DataFlowDebugPackage) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/addDataFlowToDebugSession", pathParameters),
		autorest.WithJSON(request),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddDataFlowSender sends the AddDataFlow request. The method will close the
// http.Response Body if it receives an error.
func (client DataFlowDebugSessionClient) AddDataFlowSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// AddDataFlowResponder handles the response to the AddDataFlow request. The method always
// closes the http.Response Body.
func (client DataFlowDebugSessionClient) AddDataFlowResponder(resp *http.Response) (result AddDataFlowToDebugSessionResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create creates a data flow debug session.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// request - data flow debug session definition
func (client DataFlowDebugSessionClient) Create(ctx context.Context, resourceGroupName string, factoryName string, request CreateDataFlowDebugSessionRequest) (result DataFlowDebugSessionCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataFlowDebugSessionClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.DataFlowDebugSessionClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, factoryName, request)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client DataFlowDebugSessionClient) CreatePreparer(ctx context.Context, resourceGroupName string, factoryName string, request CreateDataFlowDebugSessionRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/createDataFlowDebugSession", pathParameters),
		autorest.WithJSON(request),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client DataFlowDebugSessionClient) CreateSender(req *http.Request) (future DataFlowDebugSessionCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DataFlowDebugSessionClient) (cdfdsr CreateDataFlowDebugSessionResponse, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionCreateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("datafactory.DataFlowDebugSessionCreateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		cdfdsr.Response.Response, err = future.GetResult(sender)
		if cdfdsr.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionCreateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && cdfdsr.Response.Response.StatusCode != http.StatusNoContent {
			cdfdsr, err = client.CreateResponder(cdfdsr.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionCreateFuture", "Result", cdfdsr.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client DataFlowDebugSessionClient) CreateResponder(resp *http.Response) (result CreateDataFlowDebugSessionResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a data flow debug session.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// request - data flow debug session definition for deletion
func (client DataFlowDebugSessionClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, request DeleteDataFlowDebugSessionRequest) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataFlowDebugSessionClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.DataFlowDebugSessionClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, factoryName, request)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DataFlowDebugSessionClient) DeletePreparer(ctx context.Context, resourceGroupName string, factoryName string, request DeleteDataFlowDebugSessionRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/deleteDataFlowDebugSession", pathParameters),
		autorest.WithJSON(request),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DataFlowDebugSessionClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DataFlowDebugSessionClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ExecuteCommand execute a data flow debug command.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// request - data flow debug command definition.
func (client DataFlowDebugSessionClient) ExecuteCommand(ctx context.Context, resourceGroupName string, factoryName string, request DataFlowDebugCommandRequest) (result DataFlowDebugSessionExecuteCommandFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataFlowDebugSessionClient.ExecuteCommand")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: request,
			Constraints: []validation.Constraint{{Target: "request.CommandPayload", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "request.CommandPayload.StreamName", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("datafactory.DataFlowDebugSessionClient", "ExecuteCommand", err.Error())
	}

	req, err := client.ExecuteCommandPreparer(ctx, resourceGroupName, factoryName, request)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionClient", "ExecuteCommand", nil, "Failure preparing request")
		return
	}

	result, err = client.ExecuteCommandSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionClient", "ExecuteCommand", nil, "Failure sending request")
		return
	}

	return
}

// ExecuteCommandPreparer prepares the ExecuteCommand request.
func (client DataFlowDebugSessionClient) ExecuteCommandPreparer(ctx context.Context, resourceGroupName string, factoryName string, request DataFlowDebugCommandRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/executeDataFlowDebugCommand", pathParameters),
		autorest.WithJSON(request),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ExecuteCommandSender sends the ExecuteCommand request. The method will close the
// http.Response Body if it receives an error.
func (client DataFlowDebugSessionClient) ExecuteCommandSender(req *http.Request) (future DataFlowDebugSessionExecuteCommandFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DataFlowDebugSessionClient) (dfdcr DataFlowDebugCommandResponse, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionExecuteCommandFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("datafactory.DataFlowDebugSessionExecuteCommandFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		dfdcr.Response.Response, err = future.GetResult(sender)
		if dfdcr.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionExecuteCommandFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && dfdcr.Response.Response.StatusCode != http.StatusNoContent {
			dfdcr, err = client.ExecuteCommandResponder(dfdcr.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionExecuteCommandFuture", "Result", dfdcr.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// ExecuteCommandResponder handles the response to the ExecuteCommand request. The method always
// closes the http.Response Body.
func (client DataFlowDebugSessionClient) ExecuteCommandResponder(resp *http.Response) (result DataFlowDebugCommandResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// QueryByFactory query all active data flow debug sessions.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
func (client DataFlowDebugSessionClient) QueryByFactory(ctx context.Context, resourceGroupName string, factoryName string) (result QueryDataFlowDebugSessionsResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataFlowDebugSessionClient.QueryByFactory")
		defer func() {
			sc := -1
			if result.qdfdsr.Response.Response != nil {
				sc = result.qdfdsr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.DataFlowDebugSessionClient", "QueryByFactory", err.Error())
	}

	result.fn = client.queryByFactoryNextResults
	req, err := client.QueryByFactoryPreparer(ctx, resourceGroupName, factoryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionClient", "QueryByFactory", nil, "Failure preparing request")
		return
	}

	resp, err := client.QueryByFactorySender(req)
	if err != nil {
		result.qdfdsr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionClient", "QueryByFactory", resp, "Failure sending request")
		return
	}

	result.qdfdsr, err = client.QueryByFactoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionClient", "QueryByFactory", resp, "Failure responding to request")
		return
	}
	if result.qdfdsr.hasNextLink() && result.qdfdsr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// QueryByFactoryPreparer prepares the QueryByFactory request.
func (client DataFlowDebugSessionClient) QueryByFactoryPreparer(ctx context.Context, resourceGroupName string, factoryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/queryDataFlowDebugSessions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// QueryByFactorySender sends the QueryByFactory request. The method will close the
// http.Response Body if it receives an error.
func (client DataFlowDebugSessionClient) QueryByFactorySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// QueryByFactoryResponder handles the response to the QueryByFactory request. The method always
// closes the http.Response Body.
func (client DataFlowDebugSessionClient) QueryByFactoryResponder(resp *http.Response) (result QueryDataFlowDebugSessionsResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// queryByFactoryNextResults retrieves the next set of results, if any.
func (client DataFlowDebugSessionClient) queryByFactoryNextResults(ctx context.Context, lastResults QueryDataFlowDebugSessionsResponse) (result QueryDataFlowDebugSessionsResponse, err error) {
	req, err := lastResults.queryDataFlowDebugSessionsResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionClient", "queryByFactoryNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.QueryByFactorySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionClient", "queryByFactoryNextResults", resp, "Failure sending next results request")
	}
	result, err = client.QueryByFactoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.DataFlowDebugSessionClient", "queryByFactoryNextResults", resp, "Failure responding to next results request")
	}
	return
}

// QueryByFactoryComplete enumerates all values, automatically crossing page boundaries as required.
func (client DataFlowDebugSessionClient) QueryByFactoryComplete(ctx context.Context, resourceGroupName string, factoryName string) (result QueryDataFlowDebugSessionsResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataFlowDebugSessionClient.QueryByFactory")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.QueryByFactory(ctx, resourceGroupName, factoryName)
	return
}
