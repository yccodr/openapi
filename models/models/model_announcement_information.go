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

type AnnouncementInformation struct {
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	AnnouncementIdentifier int32 `json:"announcementIdentifier,omitempty" yaml:"announcementIdentifier" bson:"announcementIdentifier,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	AnnouncementReference string         `json:"announcementReference,omitempty" yaml:"announcementReference" bson:"announcementReference,omitempty"`
	VariableParts         []VariablePart `json:"variableParts,omitempty" yaml:"variableParts" bson:"variableParts,omitempty"`
	// indicating a time in seconds.
	TimeToPlay                int32                     `json:"timeToPlay,omitempty" yaml:"timeToPlay" bson:"timeToPlay,omitempty"`
	QuotaConsumptionIndicator QuotaConsumptionIndicator `json:"quotaConsumptionIndicator,omitempty" yaml:"quotaConsumptionIndicator" bson:"quotaConsumptionIndicator,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	AnnouncementPriority         int32                        `json:"announcementPriority,omitempty" yaml:"announcementPriority" bson:"announcementPriority,omitempty"`
	PlayToParty                  PlayToParty                  `json:"playToParty,omitempty" yaml:"playToParty" bson:"playToParty,omitempty"`
	AnnouncementPrivacyIndicator AnnouncementPrivacyIndicator `json:"announcementPrivacyIndicator,omitempty" yaml:"announcementPrivacyIndicator" bson:"announcementPrivacyIndicator,omitempty"`
	Language                     string                       `json:"Language,omitempty" yaml:"Language" bson:"Language,omitempty"`
}
