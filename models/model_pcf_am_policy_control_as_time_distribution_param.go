/*
 * Npcf_AMPolicyAuthorization Service API
 *
 * PCF Access and Mobility Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.534 V17.3.0; 5G System; Access and Mobility Policy Authorization Service; Stage 3.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.534/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains the 5G acess stratum time distribution parameters.
type PcfAmPolicyControlAsTimeDistributionParam struct {
	AsTimeDistInd bool `json:"asTimeDistInd,omitempty" yaml:"asTimeDistInd" bson:"asTimeDistInd,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property.
	UuErrorBudget int32 `json:"uuErrorBudget,omitempty" yaml:"uuErrorBudget" bson:"uuErrorBudget,omitempty"`
}
