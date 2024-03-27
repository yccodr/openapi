/*
 * Nudm_UEAU
 *
 * UDM UE Authentication Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.10.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 1.2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type Av5GHeAka struct {
	AvType   AvType `json:"avType" yaml:"avType" bson:"avType,omitempty"`
	Rand     string `json:"rand" yaml:"rand" bson:"rand,omitempty"`
	XresStar string `json:"xresStar" yaml:"xresStar" bson:"xresStar,omitempty"`
	Autn     string `json:"autn" yaml:"autn" bson:"autn,omitempty"`
	Kausf    string `json:"kausf" yaml:"kausf" bson:"kausf,omitempty"`
}
