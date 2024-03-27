/*
 * NRF OAuth2
 *
 * NRF OAuth2 Authorization. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.510 V16.8.0; 5G System; Network Function Repository Services; Stage 3
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.510/
 *
 * API version: 1.1.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains information related to the access token request
type AccessTokenReq1 struct {
	GrantType           string   `json:"grant_type" yaml:"grant_type" bson:"grant_type,omitempty"`
	NfInstanceId        string   `json:"nfInstanceId" yaml:"nfInstanceId" bson:"nfInstanceId,omitempty"`
	NfType              NfType   `json:"nfType,omitempty" yaml:"nfType" bson:"nfType,omitempty"`
	TargetNfType        NfType   `json:"targetNfType,omitempty" yaml:"targetNfType" bson:"targetNfType,omitempty"`
	Scope               string   `json:"scope" yaml:"scope" bson:"scope,omitempty"`
	TargetNfInstanceId  string   `json:"targetNfInstanceId,omitempty" yaml:"targetNfInstanceId" bson:"targetNfInstanceId,omitempty"`
	RequesterPlmn       *PlmnId  `json:"requesterPlmn,omitempty" yaml:"requesterPlmn" bson:"requesterPlmn,omitempty"`
	RequesterPlmnList   []PlmnId `json:"requesterPlmnList,omitempty" yaml:"requesterPlmnList" bson:"requesterPlmnList,omitempty"`
	RequesterSnssaiList []Snssai `json:"requesterSnssaiList,omitempty" yaml:"requesterSnssaiList" bson:"requesterSnssaiList,omitempty"`
	// Fully Qualified Domain Name
	RequesterFqdn        string      `json:"requesterFqdn,omitempty" yaml:"requesterFqdn" bson:"requesterFqdn,omitempty"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty" yaml:"requesterSnpnList" bson:"requesterSnpnList,omitempty"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty" yaml:"targetPlmn" bson:"targetPlmn,omitempty"`
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty" yaml:"targetSnssaiList" bson:"targetSnssaiList,omitempty"`
	TargetNsiList        []string    `json:"targetNsiList,omitempty" yaml:"targetNsiList" bson:"targetNsiList,omitempty"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty" yaml:"targetNfSetId" bson:"targetNfSetId,omitempty"`
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty" yaml:"targetNfServiceSetId" bson:"targetNfServiceSetId,omitempty"`
}
