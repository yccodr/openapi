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

// MBS Session Identifier
type MbsSessionId struct {
	Tmgi *Tmgi `json:"tmgi,omitempty" yaml:"tmgi" bson:"tmgi,omitempty"`
	Ssm  *Ssm  `json:"ssm,omitempty" yaml:"ssm" bson:"ssm,omitempty"`
	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).
	Nid string `json:"nid,omitempty" yaml:"nid" bson:"nid,omitempty"`
}
