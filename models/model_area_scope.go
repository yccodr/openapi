/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.518 V17.12.0; 5G System; Access and Mobility Management Services
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.518/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contain the area based on Cells or Tracking Areas.
type AreaScope struct {
	EutraCellIdList []string `json:"eutraCellIdList,omitempty" yaml:"eutraCellIdList" bson:"eutraCellIdList,omitempty"`
	NrCellIdList    []string `json:"nrCellIdList,omitempty" yaml:"nrCellIdList" bson:"nrCellIdList,omitempty"`
	TacList         []string `json:"tacList,omitempty" yaml:"tacList" bson:"tacList,omitempty"`
	// A map (list of key-value pairs) where PlmnId converted to a string serves as key
	TacInfoPerPlmn map[string]TacInfo `json:"tacInfoPerPlmn,omitempty" yaml:"tacInfoPerPlmn" bson:"tacInfoPerPlmn,omitempty"`
}
