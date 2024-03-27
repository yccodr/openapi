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

type ChfConvergedChargingPduSessionChargingInformation struct {
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	ChargingId    int32  `json:"chargingId,omitempty" yaml:"chargingId" bson:"chargingId,omitempty"`
	SMFchargingId string `json:"sMFchargingId,omitempty" yaml:"sMFchargingId" bson:"sMFchargingId,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	HomeProvidedChargingId       int32                                `json:"homeProvidedChargingId,omitempty" yaml:"homeProvidedChargingId" bson:"homeProvidedChargingId,omitempty"`
	SMFHomeProvidedChargingId    string                               `json:"sMFHomeProvidedChargingId,omitempty" yaml:"sMFHomeProvidedChargingId" bson:"sMFHomeProvidedChargingId,omitempty"`
	UserInformation              *ChfConvergedChargingUserInformation `json:"userInformation,omitempty" yaml:"userInformation" bson:"userInformation,omitempty"`
	UserLocationinfo             *UserLocation                        `json:"userLocationinfo,omitempty" yaml:"userLocationinfo" bson:"userLocationinfo,omitempty"`
	MAPDUNon3GPPUserLocationInfo *UserLocation                        `json:"mAPDUNon3GPPUserLocationInfo,omitempty" yaml:"mAPDUNon3GPPUserLocationInfo" bson:"mAPDUNon3GPPUserLocationInfo,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Non3GPPUserLocationTime *time.Time `json:"non3GPPUserLocationTime,omitempty" yaml:"non3GPPUserLocationTime" bson:"non3GPPUserLocationTime,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	MAPDUNon3GPPUserLocationTime     *time.Time              `json:"mAPDUNon3GPPUserLocationTime,omitempty" yaml:"mAPDUNon3GPPUserLocationTime" bson:"mAPDUNon3GPPUserLocationTime,omitempty"`
	PresenceReportingAreaInformation map[string]PresenceInfo `json:"presenceReportingAreaInformation,omitempty" yaml:"presenceReportingAreaInformation" bson:"presenceReportingAreaInformation,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.
	UetimeZone            string                                     `json:"uetimeZone,omitempty" yaml:"uetimeZone" bson:"uetimeZone,omitempty"`
	PduSessionInformation *ChfConvergedChargingPduSessionInformation `json:"pduSessionInformation,omitempty" yaml:"pduSessionInformation" bson:"pduSessionInformation,omitempty"`
	// indicating a time in seconds.
	UnitCountInactivityTimer   int32                       `json:"unitCountInactivityTimer,omitempty" yaml:"unitCountInactivityTimer" bson:"unitCountInactivityTimer,omitempty"`
	RANSecondaryRATUsageReport *RanSecondaryRatUsageReport `json:"rANSecondaryRATUsageReport,omitempty" yaml:"rANSecondaryRATUsageReport" bson:"rANSecondaryRATUsageReport,omitempty"`
}
