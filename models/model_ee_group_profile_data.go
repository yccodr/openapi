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

type EeGroupProfileData struct {
	RestrictedEventTypes []UdmEeEventType `json:"restrictedEventTypes,omitempty" yaml:"restrictedEventTypes" bson:"restrictedEventTypes,omitempty"`
	// A map (list of key-value pairs where EventType serves as key) of MTC provider lists. In addition to defined EventTypes, the key value \"ALL\" may be used to identify a map entry which contains a list of MtcProviders that are allowed monitoring all Event Types.
	AllowedMtcProvider map[string][]MtcProvider `json:"allowedMtcProvider,omitempty" yaml:"allowedMtcProvider" bson:"allowedMtcProvider,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
	IwkEpcRestricted  bool   `json:"iwkEpcRestricted,omitempty" yaml:"iwkEpcRestricted" bson:"iwkEpcRestricted,omitempty"`
	ExtGroupId        string `json:"extGroupId,omitempty" yaml:"extGroupId" bson:"extGroupId,omitempty"`
	// Identifier of a group of NFs.
	HssGroupId string `json:"hssGroupId,omitempty" yaml:"hssGroupId" bson:"hssGroupId,omitempty"`
}
