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

// Information of an SMS-IWMSC NF Instance
type IwmscInfo struct {
	MsisdnRanges []IdentityRange `json:"msisdnRanges,omitempty" yaml:"msisdnRanges" bson:"msisdnRanges,omitempty"`
	SupiRanges   []SupiRange     `json:"supiRanges,omitempty" yaml:"supiRanges" bson:"supiRanges,omitempty"`
	TaiRangeList []TaiRange      `json:"taiRangeList,omitempty" yaml:"taiRangeList" bson:"taiRangeList,omitempty"`
	ScNumber     string          `json:"scNumber,omitempty" yaml:"scNumber" bson:"scNumber,omitempty"`
}
