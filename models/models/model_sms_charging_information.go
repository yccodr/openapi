/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 32.291 V17.9.0: Telecommunication management; Charging management;  5G system, charging service; Stage 3.
 * Url: http://www.3gpp.org/ftp/Specs/archive/32_series/32.291/
 *
 * API version: 3.1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type SmsChargingInformation struct {
	OriginatorInfo *OriginatorInfo `json:"originatorInfo,omitempty" yaml:"originatorInfo" bson:"originatorInfo,omitempty"`
	RecipientInfo  []RecipientInfo `json:"recipientInfo,omitempty" yaml:"recipientInfo" bson:"recipientInfo,omitempty"`
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.
	UserEquipmentInfo string        `json:"userEquipmentInfo,omitempty" yaml:"userEquipmentInfo" bson:"userEquipmentInfo,omitempty"`
	RoamerInOut       RoamerInOut   `json:"roamerInOut,omitempty" yaml:"roamerInOut" bson:"roamerInOut,omitempty"`
	UserLocationinfo  *UserLocation `json:"userLocationinfo,omitempty" yaml:"userLocationinfo" bson:"userLocationinfo,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.
	UetimeZone           string             `json:"uetimeZone,omitempty" yaml:"uetimeZone" bson:"uetimeZone,omitempty"`
	RATType              RatType            `json:"rATType,omitempty" yaml:"rATType" bson:"rATType,omitempty"`
	SMSCAddress          string             `json:"sMSCAddress,omitempty" yaml:"sMSCAddress" bson:"sMSCAddress,omitempty"`
	SMDataCodingScheme   int32              `json:"sMDataCodingScheme,omitempty" yaml:"sMDataCodingScheme" bson:"sMDataCodingScheme,omitempty"`
	SMMessageType        SmMessageType      `json:"sMMessageType,omitempty" yaml:"sMMessageType" bson:"sMMessageType,omitempty"`
	SMReplyPathRequested ReplyPathRequested `json:"sMReplyPathRequested,omitempty" yaml:"sMReplyPathRequested" bson:"sMReplyPathRequested,omitempty"`
	SMUserDataHeader     string             `json:"sMUserDataHeader,omitempty" yaml:"sMUserDataHeader" bson:"sMUserDataHeader,omitempty"`
	SMStatus             string             `json:"sMStatus,omitempty" yaml:"sMStatus" bson:"sMStatus,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	SMDischargeTime *time.Time `json:"sMDischargeTime,omitempty" yaml:"sMDischargeTime" bson:"sMDischargeTime,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	NumberofMessagesSent int32         `json:"numberofMessagesSent,omitempty" yaml:"numberofMessagesSent" bson:"numberofMessagesSent,omitempty"`
	SMServiceType        SmServiceType `json:"sMServiceType,omitempty" yaml:"sMServiceType" bson:"sMServiceType,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	SMSequenceNumber int32 `json:"sMSequenceNumber,omitempty" yaml:"sMSequenceNumber" bson:"sMSequenceNumber,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	SMSresult int32 `json:"sMSresult,omitempty" yaml:"sMSresult" bson:"sMSresult,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	SubmissionTime   *time.Time `json:"submissionTime,omitempty" yaml:"submissionTime" bson:"submissionTime,omitempty"`
	SMPriority       SmPriority `json:"sMPriority,omitempty" yaml:"sMPriority" bson:"sMPriority,omitempty"`
	MessageReference string     `json:"messageReference,omitempty" yaml:"messageReference" bson:"messageReference,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	MessageSize             int32                   `json:"messageSize,omitempty" yaml:"messageSize" bson:"messageSize,omitempty"`
	MessageClass            *MessageClass           `json:"messageClass,omitempty" yaml:"messageClass" bson:"messageClass,omitempty"`
	DeliveryReportRequested DeliveryReportRequested `json:"deliveryReportRequested,omitempty" yaml:"deliveryReportRequested" bson:"deliveryReportRequested,omitempty"`
}
