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

// N2 information container
type N2InfoContainer struct {
	N2InformationClass N2InformationClass `json:"n2InformationClass" yaml:"n2InformationClass" bson:"n2InformationClass,omitempty"`
	SmInfo             *N2SmInformation   `json:"smInfo,omitempty" yaml:"smInfo" bson:"smInfo,omitempty"`
	RanInfo            *N2RanInformation  `json:"ranInfo,omitempty" yaml:"ranInfo" bson:"ranInfo,omitempty"`
	NrppaInfo          *NrppaInformation  `json:"nrppaInfo,omitempty" yaml:"nrppaInfo" bson:"nrppaInfo,omitempty"`
	PwsInfo            *PwsInformation    `json:"pwsInfo,omitempty" yaml:"pwsInfo" bson:"pwsInfo,omitempty"`
	V2xInfo            *V2xInformation    `json:"v2xInfo,omitempty" yaml:"v2xInfo" bson:"v2xInfo,omitempty"`
	ProseInfo          *ProSeInformation  `json:"proseInfo,omitempty" yaml:"proseInfo" bson:"proseInfo,omitempty"`
}
