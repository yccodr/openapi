/*
 * Npcf_BDTPolicyControl Service API
 *
 * PCF BDT Policy Control Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.554 V16.7.0; 5G System; Background Data Transfer Policy Control Service.
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.554/
 *
 * API version: 1.1.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Describes the updates in authorization data of an Individual BDT Policy created by the PCF.
type BdtPolicyPatchBdtPolicy struct {
	BdtPolData *BdtPolicyBdtPolicyDataPatch `json:"bdtPolData,omitempty" yaml:"bdtPolData" bson:"bdtPolData,omitempty"`
	BdtReqData *BdtPolicyBdtReqDataPatch    `json:"bdtReqData,omitempty" yaml:"bdtReqData" bson:"bdtReqData,omitempty"`
}
