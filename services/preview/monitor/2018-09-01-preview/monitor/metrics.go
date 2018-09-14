package monitor

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

// MetricsClient is the monitor Management Client
type MetricsClient struct {
	BaseClient
}

// NewMetricsClient creates an instance of the MetricsClient client.
func NewMetricsClient() MetricsClient {
	return NewMetricsClientWithBaseURI(DefaultBaseURI)
}

// NewMetricsClientWithBaseURI creates an instance of the MetricsClient client.
func NewMetricsClientWithBaseURI(baseURI string) MetricsClient {
	return MetricsClient{NewWithBaseURI(baseURI)}
}

// Create **Post the metric values for a resource**.
// Parameters:
// subscriptionID - the azure subscription id
// resourceGroupName - the ARM resource group name
// resourceProvider - the ARM resource provider name
// resourceTypeName - the ARM resource type name
// resourceName - the ARM resource name
// body - the Azure metrics document json payload
// contentType - supports application/json and application/x-ndjson
// contentLength - content length of the payload
func (client MetricsClient) Create(ctx context.Context, subscriptionID string, resourceGroupName string, resourceProvider string, resourceTypeName string, resourceName string, body AzureMetricsDocument, contentType string, contentLength *int32) (result AzureMetricsResult, err error) {
	req, err := client.CreatePreparer(ctx, subscriptionID, resourceGroupName, resourceProvider, resourceTypeName, resourceName, body, contentType, contentLength)
	if err != nil {
		err = autorest.NewErrorWithError(err, "monitor.MetricsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "monitor.MetricsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "monitor.MetricsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client MetricsClient) CreatePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, resourceProvider string, resourceTypeName string, resourceName string, body AzureMetricsDocument, contentType string, contentLength *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"resourceProvider":  autorest.Encode("path", resourceProvider),
		"resourceTypeName":  autorest.Encode("path", resourceTypeName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProvider}/{resourceTypeName}/{resourceName}/metrics", pathParameters),
		autorest.WithJSON(body))
	if len(contentType) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Content-Type", autorest.String(contentType)))
	}
	if contentLength != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Content-Length", autorest.String(contentLength)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client MetricsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client MetricsClient) CreateResponder(resp *http.Response) (result AzureMetricsResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}