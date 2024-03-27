/*
 * 3gpp-traffic-influence
 *
 * API for AF traffic influence   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.522 V17.7.0; 5G System; Network Exposure Function Northbound APIs.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.522/
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents the description of invalid parameters, for a request rejected due to invalid parameters.
type InvalidParam struct {
	// Attribute's name encoded as a JSON Pointer, or header's name.
	Param string `json:"param" yaml:"param" bson:"param,omitempty"`
	// A human-readable reason, e.g. \"must be a positive integer\".
	Reason string `json:"reason,omitempty" yaml:"reason" bson:"reason,omitempty"`
}
