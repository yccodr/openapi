/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.502 V17.11.0; 5G System; Session Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.502/
 *
 * API version: 1.2.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ModifyPduSessionResponse200 struct {
	JsonData                    *VsmfUpdatedData `json:"jsonData,omitempty" yaml:"jsonData" bson:"jsonData,omitempty"`
	BinaryDataN1SmInfoFromUe    []byte           `json:"binaryDataN1SmInfoFromUe,omitempty" yaml:"binaryDataN1SmInfoFromUe" bson:"binaryDataN1SmInfoFromUe,omitempty"`
	BinaryDataUnknownN1SmInfo   []byte           `json:"binaryDataUnknownN1SmInfo,omitempty" yaml:"binaryDataUnknownN1SmInfo" bson:"binaryDataUnknownN1SmInfo,omitempty"`
	BinaryDataN4Information     []byte           `json:"binaryDataN4Information,omitempty" yaml:"binaryDataN4Information" bson:"binaryDataN4Information,omitempty"`
	BinaryDataN4InformationExt1 []byte           `json:"binaryDataN4InformationExt1,omitempty" yaml:"binaryDataN4InformationExt1" bson:"binaryDataN4InformationExt1,omitempty"`
	BinaryDataN4InformationExt2 []byte           `json:"binaryDataN4InformationExt2,omitempty" yaml:"binaryDataN4InformationExt2" bson:"binaryDataN4InformationExt2,omitempty"`
	BinaryDataN4InformationExt3 []byte           `json:"binaryDataN4InformationExt3,omitempty" yaml:"binaryDataN4InformationExt3" bson:"binaryDataN4InformationExt3,omitempty"`
}
