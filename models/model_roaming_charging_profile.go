/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.502 V16.9.0; 5G System; Session Management Services; Stage 3
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.502/
 *
 * API version: 1.1.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type RoamingChargingProfile struct {
	Triggers            []Trigger           `json:"triggers,omitempty" yaml:"triggers" bson:"triggers,omitempty"`
	PartialRecordMethod PartialRecordMethod `json:"partialRecordMethod,omitempty" yaml:"partialRecordMethod" bson:"partialRecordMethod,omitempty"`
}
