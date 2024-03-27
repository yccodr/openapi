/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.9.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Describes an event to be subscribed
type AmfEvent struct {
	Type                  AmfEventType                        `json:"type" yaml:"type" bson:"type,omitempty"`
	ImmediateFlag         bool                                `json:"immediateFlag,omitempty" yaml:"immediateFlag" bson:"immediateFlag,omitempty"`
	AreaList              []AmfEventArea                      `json:"areaList,omitempty" yaml:"areaList" bson:"areaList,omitempty"`
	LocationFilterList    []LocationFilter                    `json:"locationFilterList,omitempty" yaml:"locationFilterList" bson:"locationFilterList,omitempty"`
	RefId                 int32                               `json:"refId,omitempty" yaml:"refId" bson:"refId,omitempty"`
	TrafficDescriptorList []AmfEventExposureTrafficDescriptor `json:"trafficDescriptorList,omitempty" yaml:"trafficDescriptorList" bson:"trafficDescriptorList,omitempty"`
	ReportUeReachable     bool                                `json:"reportUeReachable,omitempty" yaml:"reportUeReachable" bson:"reportUeReachable,omitempty"`
	ReachabilityFilter    ReachabilityFilter                  `json:"reachabilityFilter,omitempty" yaml:"reachabilityFilter" bson:"reachabilityFilter,omitempty"`
	UdmDetectInd          bool                                `json:"udmDetectInd,omitempty" yaml:"udmDetectInd" bson:"udmDetectInd,omitempty"`
	MaxReports            int32                               `json:"maxReports,omitempty" yaml:"maxReports" bson:"maxReports,omitempty"`
	// A map(list of key-value pairs) where praId serves as key.
	PresenceInfoList map[string]PresenceInfo `json:"presenceInfoList,omitempty" yaml:"presenceInfoList" bson:"presenceInfoList,omitempty"`
	// indicating a time in seconds.
	MaxResponseTime int32           `json:"maxResponseTime,omitempty" yaml:"maxResponseTime" bson:"maxResponseTime,omitempty"`
	TargetArea      *TargetArea     `json:"targetArea,omitempty" yaml:"targetArea" bson:"targetArea,omitempty"`
	SnssaiFilter    []ExtSnssai     `json:"snssaiFilter,omitempty" yaml:"snssaiFilter" bson:"snssaiFilter,omitempty"`
	UeInAreaFilter  *UeInAreaFilter `json:"ueInAreaFilter,omitempty" yaml:"ueInAreaFilter" bson:"ueInAreaFilter,omitempty"`
	// indicating a time in seconds.
	MinInterval int32 `json:"minInterval,omitempty" yaml:"minInterval" bson:"minInterval,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	NextReport     *time.Time      `json:"nextReport,omitempty" yaml:"nextReport" bson:"nextReport,omitempty"`
	IdleStatusInd  bool            `json:"idleStatusInd,omitempty" yaml:"idleStatusInd" bson:"idleStatusInd,omitempty"`
	DispersionArea *DispersionArea `json:"dispersionArea,omitempty" yaml:"dispersionArea" bson:"dispersionArea,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	NextPeriodicReportTime *time.Time `json:"nextPeriodicReportTime,omitempty" yaml:"nextPeriodicReportTime" bson:"nextPeriodicReportTime,omitempty"`
}
