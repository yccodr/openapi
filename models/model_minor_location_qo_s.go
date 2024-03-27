/*
 * LMF Location
 *
 * LMF Location Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.572 V17.9.0; 5G System; Location Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.572/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contain Minor Location QoS.
type MinorLocationQoS struct {
	// Indicates value of accuracy.
	HAccuracy float32 `json:"hAccuracy,omitempty" yaml:"hAccuracy" bson:"hAccuracy,omitempty"`
	// Indicates value of accuracy.
	VAccuracy float32 `json:"vAccuracy,omitempty" yaml:"vAccuracy" bson:"vAccuracy,omitempty"`
}
