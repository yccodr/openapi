/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.10.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents an Individual NWDAF Event Subscription resource.
type NnwdafEventsSubscription struct {
	// Subscribed events
	EventSubscriptions []NwdafEventsSubscriptionEventSubscription `json:"eventSubscriptions" yaml:"eventSubscriptions" bson:"eventSubscriptions,omitempty"`
	EvtReq             *ReportingInformation                      `json:"evtReq,omitempty" yaml:"evtReq" bson:"evtReq,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	NotificationURI string `json:"notificationURI,omitempty" yaml:"notificationURI" bson:"notificationURI,omitempty"`
	// Notification correlation identifier.
	NotifCorrId string `json:"notifCorrId,omitempty" yaml:"notifCorrId" bson:"notifCorrId,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures  string                                     `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
	EventNotifications []NwdafEventsSubscriptionEventNotification `json:"eventNotifications,omitempty" yaml:"eventNotifications" bson:"eventNotifications,omitempty"`
	FailEventReports   []FailureEventInfo                         `json:"failEventReports,omitempty" yaml:"failEventReports" bson:"failEventReports,omitempty"`
	PrevSub            *PrevSubInfo                               `json:"prevSub,omitempty" yaml:"prevSub" bson:"prevSub,omitempty"`
	ConsNfInfo         *ConsumerNfInformation                     `json:"consNfInfo,omitempty" yaml:"consNfInfo" bson:"consNfInfo,omitempty"`
}
