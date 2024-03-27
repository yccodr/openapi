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

import (
	"time"
)

// Used to store the status of the latest UPU data update.
type UpuData struct {
	// string with format 'date-time' as defined in OpenAPI.
	ProvisioningTime *time.Time     `json:"provisioningTime" yaml:"provisioningTime" bson:"provisioningTime,omitempty"`
	UeUpdateStatus   UeUpdateStatus `json:"ueUpdateStatus" yaml:"ueUpdateStatus" bson:"ueUpdateStatus,omitempty"`
	// MAC value for protecting UPU procedure (UPU-MAC-IAUSF and UPU-MAC-IUE).
	UpuXmacIue string `json:"upuXmacIue,omitempty" yaml:"upuXmacIue" bson:"upuXmacIue,omitempty"`
	// MAC value for protecting UPU procedure (UPU-MAC-IAUSF and UPU-MAC-IUE).
	UpuMacIue string `json:"upuMacIue,omitempty" yaml:"upuMacIue" bson:"upuMacIue,omitempty"`
}
