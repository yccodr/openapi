/*
 * Npcf_MBSPolicyControl API
 *
 * MBS Policy Control Service   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.537 V17.3.0; 5G System; Multicast/Broadcast Policy Control Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.537/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents the reporting of MBS Policy decision level failure(s) and/or MBS PCC rule level failure(s).
type MbsErrorReport struct {
	MbsReports []MbsReport `json:"mbsReports,omitempty" yaml:"mbsReports" bson:"mbsReports,omitempty"`
}
