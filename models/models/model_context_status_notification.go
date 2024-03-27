/*
 * Namf_MBSBroadcast
 *
 * AMF MBSBroadcast Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.518 V17.12.0; 5G System; Access and Mobility Management Services
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.518/
 *
 * API version: 1.0.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Data within ContextStatusNotify Request
type ContextStatusNotification struct {
	MbsSessionId *MbsSessionId `json:"mbsSessionId" yaml:"mbsSessionId" bson:"mbsSessionId,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 16-bit integer.
	AreaSessionId   int32                        `json:"areaSessionId,omitempty" yaml:"areaSessionId" bson:"areaSessionId,omitempty"`
	N2MbsSmInfoList []AmfMbsBroadcastN2MbsSmInfo `json:"n2MbsSmInfoList,omitempty" yaml:"n2MbsSmInfoList" bson:"n2MbsSmInfoList,omitempty"`
	OperationEvents []OperationEvent             `json:"operationEvents,omitempty" yaml:"operationEvents" bson:"operationEvents,omitempty"`
	OperationStatus OperationStatus              `json:"operationStatus,omitempty" yaml:"operationStatus" bson:"operationStatus,omitempty"`
	ReleasedInd     bool                         `json:"releasedInd,omitempty" yaml:"releasedInd" bson:"releasedInd,omitempty"`
}
