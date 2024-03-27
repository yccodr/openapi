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

type NwdafEventsSubscriptionCongestionType string

// List of NwdafEventsSubscriptionCongestionType
const (
	NwdafEventsSubscriptionCongestionType_USER_PLANE             NwdafEventsSubscriptionCongestionType = "USER_PLANE"
	NwdafEventsSubscriptionCongestionType_CONTROL_PLANE          NwdafEventsSubscriptionCongestionType = "CONTROL_PLANE"
	NwdafEventsSubscriptionCongestionType_USER_AND_CONTROL_PLANE NwdafEventsSubscriptionCongestionType = "USER_AND_CONTROL_PLANE"
)
