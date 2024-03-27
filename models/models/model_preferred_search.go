/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.510 V17.12.0; 5G System; Network Function Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/
 *
 * API version: 1.2.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains information on whether the returned NFProfiles match the preferred query parameters
type PreferredSearch struct {
	PreferredTaiMatchInd               bool `json:"preferredTaiMatchInd,omitempty" yaml:"preferredTaiMatchInd" bson:"preferredTaiMatchInd,omitempty"`
	PreferredFullPlmnMatchInd          bool `json:"preferredFullPlmnMatchInd,omitempty" yaml:"preferredFullPlmnMatchInd" bson:"preferredFullPlmnMatchInd,omitempty"`
	PreferredApiVersionsMatchInd       bool `json:"preferredApiVersionsMatchInd,omitempty" yaml:"preferredApiVersionsMatchInd" bson:"preferredApiVersionsMatchInd,omitempty"`
	OtherApiVersionsInd                bool `json:"otherApiVersionsInd,omitempty" yaml:"otherApiVersionsInd" bson:"otherApiVersionsInd,omitempty"`
	PreferredLocalityMatchInd          bool `json:"preferredLocalityMatchInd,omitempty" yaml:"preferredLocalityMatchInd" bson:"preferredLocalityMatchInd,omitempty"`
	OtherLocalityInd                   bool `json:"otherLocalityInd,omitempty" yaml:"otherLocalityInd" bson:"otherLocalityInd,omitempty"`
	PreferredVendorSpecificFeaturesInd bool `json:"preferredVendorSpecificFeaturesInd,omitempty" yaml:"preferredVendorSpecificFeaturesInd" bson:"preferredVendorSpecificFeaturesInd,omitempty"`
	PreferredCollocatedNfTypeInd       bool `json:"preferredCollocatedNfTypeInd,omitempty" yaml:"preferredCollocatedNfTypeInd" bson:"preferredCollocatedNfTypeInd,omitempty"`
	PreferredPgwMatchInd               bool `json:"preferredPgwMatchInd,omitempty" yaml:"preferredPgwMatchInd" bson:"preferredPgwMatchInd,omitempty"`
	PreferredAnalyticsDelaysInd        bool `json:"preferredAnalyticsDelaysInd,omitempty" yaml:"preferredAnalyticsDelaysInd" bson:"preferredAnalyticsDelaysInd,omitempty"`
}
