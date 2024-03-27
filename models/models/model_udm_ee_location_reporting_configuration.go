/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.9.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type UdmEeLocationReportingConfiguration struct {
	CurrentLocation bool                  `json:"currentLocation" yaml:"currentLocation" bson:"currentLocation,omitempty"`
	OneTime         bool                  `json:"oneTime,omitempty" yaml:"oneTime" bson:"oneTime,omitempty"`
	Accuracy        UdmEeLocationAccuracy `json:"accuracy,omitempty" yaml:"accuracy" bson:"accuracy,omitempty"`
	N3gppAccuracy   UdmEeLocationAccuracy `json:"n3gppAccuracy,omitempty" yaml:"n3gppAccuracy" bson:"n3gppAccuracy,omitempty"`
}
