/*
 * Nudm_EE
 *
 * Nudm Event Exposure Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 16.9.0
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 1.1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type MonitoringReport struct {
	ReferenceId              int32                     `json:"referenceId" yaml:"referenceId" bson:"referenceId,omitempty"`
	EventType                EventType                 `json:"eventType" yaml:"eventType" bson:"eventType,omitempty"`
	Report                   *Report                   `json:"report,omitempty" yaml:"report" bson:"report,omitempty"`
	ReachabilityForSmsReport *ReachabilityForSmsReport `json:"reachabilityForSmsReport,omitempty" yaml:"reachabilityForSmsReport" bson:"reachabilityForSmsReport,omitempty"`
	Gpsi                     string                    `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi,omitempty"`
	TimeStamp                *time.Time                `json:"timeStamp" yaml:"timeStamp" bson:"timeStamp,omitempty"`
}
