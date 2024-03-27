/*
 * Npcf_AMPolicyAuthorization Service API
 *
 * PCF Access and Mobility Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.534 V17.3.0; 5G System; Access and Mobility Policy Authorization Service; Stage 3.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.534/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Identifies the events the application subscribes to within an AM Policy Events Subscription subresource data. It may contain the notification of the already met events.
type AmEventsSubscRespData struct {
	// String providing an URI formatted according to RFC 3986.
	EventNotifUri string        `json:"eventNotifUri" yaml:"eventNotifUri" bson:"eventNotifUri,omitempty"`
	Events        []AmEventData `json:"events,omitempty" yaml:"events" bson:"events,omitempty"`
	// Contains the AM Policy Events Subscription resource identifier related to the event notification.
	AppAmContextId string                `json:"appAmContextId,omitempty" yaml:"appAmContextId" bson:"appAmContextId,omitempty"`
	RepEvents      []AmEventNotification `json:"repEvents" yaml:"repEvents" bson:"repEvents,omitempty"`
}
