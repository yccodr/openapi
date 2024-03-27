/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V16.9.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.1.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type UpuInfo struct {
	UpuDataList      []UpuData1 `json:"upuDataList" yaml:"upuDataList" bson:"upuDataList,omitempty"`
	UpuRegInd        bool       `json:"upuRegInd" yaml:"upuRegInd" bson:"upuRegInd,omitempty"`
	UpuAckInd        bool       `json:"upuAckInd" yaml:"upuAckInd" bson:"upuAckInd,omitempty"`
	UpuMacIausf      string     `json:"upuMacIausf,omitempty" yaml:"upuMacIausf" bson:"upuMacIausf,omitempty"`
	CounterUpu       string     `json:"counterUpu,omitempty" yaml:"counterUpu" bson:"counterUpu,omitempty"`
	ProvisioningTime *time.Time `json:"provisioningTime" yaml:"provisioningTime" bson:"provisioningTime,omitempty"`
}
