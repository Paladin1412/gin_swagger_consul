package commerce

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RateCardClient is the client for the RateCard methods of the Commerce service.
type RateCardClient struct {
	BaseClient
}

// NewRateCardClient creates an instance of the RateCardClient client.
func NewRateCardClient(subscriptionID string) RateCardClient {
	return NewRateCardClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRateCardClientWithBaseURI creates an instance of the RateCardClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRateCardClientWithBaseURI(baseURI string, subscriptionID string) RateCardClient {
	return RateCardClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get enables you to query for the resource/meter metadata and related prices used in a given subscription by Offer
// ID, Currency, Locale and Region. The metadata associated with the billing meters, including but not limited to
// service names, types, resources, units of measure, and regions, is subject to change at any time and without notice.
// If you intend to use this billing data in an automated fashion, please use the billing meter GUID to uniquely
// identify each billable item. If the billing meter GUID is scheduled to change due to a new billing model, you will
// be notified in advance of the change.
// Parameters:
// filter - the filter to apply on the operation. It ONLY supports the 'eq' and 'and' logical operators at this
// time. All the 4 query parameters 'OfferDurableId',  'Currency', 'Locale', 'Region' are required to be a part
// of the $filter.
func (client RateCardClient) Get(ctx context.Context, filter string) (result ResourceRateCardInfo, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RateCardClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commerce.RateCardClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "commerce.RateCardClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commerce.RateCardClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RateCardClient) GetPreparer(ctx context.Context, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"$filter":     autorest.Encode("query", filter),
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Commerce/RateCard", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RateCardClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RateCardClient) GetResponder(resp *http.Response) (result ResourceRateCardInfo, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
