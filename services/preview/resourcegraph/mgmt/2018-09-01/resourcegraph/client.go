// Package resourcegraph implements the Azure ARM Resourcegraph service API version 2018-09-01-preview.
//
//
package resourcegraph

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

const (
	// DefaultBaseURI is the default URI used for the service Resourcegraph
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Resourcegraph.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client.
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// ResourceChangeDetails get resource change details.
// Parameters:
// parameters - the parameters for this request for resource change details.
func (client BaseClient) ResourceChangeDetails(ctx context.Context, parameters ResourceChangeDetailsRequestParameters) (result ResourceChangeData, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.ResourceChangeDetails")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ResourceID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.ChangeID", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("resourcegraph.BaseClient", "ResourceChangeDetails", err.Error())
	}

	req, err := client.ResourceChangeDetailsPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "ResourceChangeDetails", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResourceChangeDetailsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "ResourceChangeDetails", resp, "Failure sending request")
		return
	}

	result, err = client.ResourceChangeDetailsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "ResourceChangeDetails", resp, "Failure responding to request")
	}

	return
}

// ResourceChangeDetailsPreparer prepares the ResourceChangeDetails request.
func (client BaseClient) ResourceChangeDetailsPreparer(ctx context.Context, parameters ResourceChangeDetailsRequestParameters) (*http.Request, error) {
	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.ResourceGraph/resourceChangeDetails"),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResourceChangeDetailsSender sends the ResourceChangeDetails request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) ResourceChangeDetailsSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ResourceChangeDetailsResponder handles the response to the ResourceChangeDetails request. The method always
// closes the http.Response Body.
func (client BaseClient) ResourceChangeDetailsResponder(resp *http.Response) (result ResourceChangeData, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ResourceChanges list changes to a resource for a given time interval.
// Parameters:
// parameters - the parameters for this request for changes.
func (client BaseClient) ResourceChanges(ctx context.Context, parameters ResourceChangesRequestParameters) (result ResourceChangeList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.ResourceChanges")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ResourceID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Interval", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Top", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.Top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
						{Target: "parameters.Top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewError("resourcegraph.BaseClient", "ResourceChanges", err.Error())
	}

	req, err := client.ResourceChangesPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "ResourceChanges", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResourceChangesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "ResourceChanges", resp, "Failure sending request")
		return
	}

	result, err = client.ResourceChangesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "ResourceChanges", resp, "Failure responding to request")
	}

	return
}

// ResourceChangesPreparer prepares the ResourceChanges request.
func (client BaseClient) ResourceChangesPreparer(ctx context.Context, parameters ResourceChangesRequestParameters) (*http.Request, error) {
	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.ResourceGraph/resourceChanges"),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResourceChangesSender sends the ResourceChanges request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) ResourceChangesSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ResourceChangesResponder handles the response to the ResourceChanges request. The method always
// closes the http.Response Body.
func (client BaseClient) ResourceChangesResponder(resp *http.Response) (result ResourceChangeList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Resources queries the resources managed by Azure Resource Manager for all subscriptions specified in the request.
// Parameters:
// query - request specifying query and its options.
func (client BaseClient) Resources(ctx context.Context, query QueryRequest) (result QueryResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.Resources")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: query,
			Constraints: []validation.Constraint{{Target: "query.Subscriptions", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "query.Query", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "query.Options", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "query.Options.Top", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "query.Options.Top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
							{Target: "query.Options.Top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
						}},
						{Target: "query.Options.Skip", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "query.Options.Skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}},
					}}}}}); err != nil {
		return result, validation.NewError("resourcegraph.BaseClient", "Resources", err.Error())
	}

	req, err := client.ResourcesPreparer(ctx, query)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "Resources", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResourcesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "Resources", resp, "Failure sending request")
		return
	}

	result, err = client.ResourcesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegraph.BaseClient", "Resources", resp, "Failure responding to request")
	}

	return
}

// ResourcesPreparer prepares the Resources request.
func (client BaseClient) ResourcesPreparer(ctx context.Context, query QueryRequest) (*http.Request, error) {
	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.ResourceGraph/resources"),
		autorest.WithJSON(query),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResourcesSender sends the Resources request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) ResourcesSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ResourcesResponder handles the response to the Resources request. The method always
// closes the http.Response Body.
func (client BaseClient) ResourcesResponder(resp *http.Response) (result QueryResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
