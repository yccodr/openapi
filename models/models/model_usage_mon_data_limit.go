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

// Contains usage monitoring control data for a subscriber.
type UsageMonDataLimit struct {
	LimitId string `json:"limitId" yaml:"limitId" bson:"limitId,omitempty"`
	// Identifies the SNSSAI and DNN combinations to which the usage monitoring data limit applies. The S-NSSAI is the key of the map.
	Scopes  map[string]UsageMonDataScope `json:"scopes,omitempty" yaml:"scopes" bson:"scopes,omitempty"`
	UmLevel UsageMonLevel                `json:"umLevel,omitempty" yaml:"umLevel" bson:"umLevel,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	StartDate *time.Time `json:"startDate,omitempty" yaml:"startDate" bson:"startDate,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	EndDate     *time.Time      `json:"endDate,omitempty" yaml:"endDate" bson:"endDate,omitempty"`
	UsageLimit  *UsageThreshold `json:"usageLimit,omitempty" yaml:"usageLimit" bson:"usageLimit,omitempty"`
	ResetPeriod *TimePeriod     `json:"resetPeriod,omitempty" yaml:"resetPeriod" bson:"resetPeriod,omitempty"`
}
