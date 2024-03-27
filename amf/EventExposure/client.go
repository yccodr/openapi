/*
 * Namf_EventExposure
 *
 * AMF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.518 V17.10.0; 5G System; Access and Mobility Management Services
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.518/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package EventExposure

// APIClient manages communication with the Namf_EventExposure API v1.2.3
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	IndividualSubscriptionDocumentApi    *IndividualSubscriptionDocumentApiService
	SubscriptionsCollectionCollectionApi *SubscriptionsCollectionCollectionApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.IndividualSubscriptionDocumentApi = (*IndividualSubscriptionDocumentApiService)(&c.common)
	c.SubscriptionsCollectionCollectionApi = (*SubscriptionsCollectionCollectionApiService)(&c.common)

	return c
}
