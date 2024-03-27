/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V17.12.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type RegistrationLocationInfo struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	AmfInstanceId  string        `json:"amfInstanceId" yaml:"amfInstanceId" bson:"amfInstanceId,omitempty"`
	Guami          *Guami        `json:"guami,omitempty" yaml:"guami" bson:"guami,omitempty"`
	PlmnId         *PlmnId       `json:"plmnId,omitempty" yaml:"plmnId" bson:"plmnId,omitempty"`
	VgmlcAddress   *VgmlcAddress `json:"vgmlcAddress,omitempty" yaml:"vgmlcAddress" bson:"vgmlcAddress,omitempty"`
	AccessTypeList []AccessType  `json:"accessTypeList" yaml:"accessTypeList" bson:"accessTypeList,omitempty"`
}
