/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.512 V17.11.0; 5G System; Session Management Policy Control Service.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Contains the parameters used to create an Individual SM policy resource.
type SmPolicyContextData struct {
	AccNetChId      *AccNetChId            `json:"accNetChId,omitempty" yaml:"accNetChId" bson:"accNetChId,omitempty"`
	ChargEntityAddr *AccNetChargingAddress `json:"chargEntityAddr,omitempty" yaml:"chargEntityAddr" bson:"chargEntityAddr,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi string `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi string `json:"supi" yaml:"supi" bson:"supi,omitempty"`
	// When this attribute is included and set to true, it indicates that the supi attribute contains an invalid value.This attribute shall be present if the SUPI is not available in the SMF or the SUPI is unauthenticated. When present it shall be set to true for an invalid SUPI and false (default) for a valid SUPI.
	InvalidSupi bool     `json:"invalidSupi,omitempty" yaml:"invalidSupi" bson:"invalidSupi,omitempty"`
	InterGrpIds []string `json:"interGrpIds,omitempty" yaml:"interGrpIds" bson:"interGrpIds,omitempty"`
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.
	PduSessionId            int32          `json:"pduSessionId" yaml:"pduSessionId" bson:"pduSessionId,omitempty"`
	PduSessionType          PduSessionType `json:"pduSessionType" yaml:"pduSessionType" bson:"pduSessionType,omitempty"`
	Chargingcharacteristics string         `json:"chargingcharacteristics,omitempty" yaml:"chargingcharacteristics" bson:"chargingcharacteristics,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn        string           `json:"dnn" yaml:"dnn" bson:"dnn,omitempty"`
	DnnSelMode DnnSelectionMode `json:"dnnSelMode,omitempty" yaml:"dnnSelMode" bson:"dnnSelMode,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	NotificationUri  string                `json:"notificationUri" yaml:"notificationUri" bson:"notificationUri,omitempty"`
	AccessType       AccessType            `json:"accessType,omitempty" yaml:"accessType" bson:"accessType,omitempty"`
	RatType          RatType               `json:"ratType,omitempty" yaml:"ratType" bson:"ratType,omitempty"`
	AddAccessInfo    *AdditionalAccessInfo `json:"addAccessInfo,omitempty" yaml:"addAccessInfo" bson:"addAccessInfo,omitempty"`
	ServingNetwork   *PlmnIdNid            `json:"servingNetwork,omitempty" yaml:"servingNetwork" bson:"servingNetwork,omitempty"`
	UserLocationInfo *UserLocation         `json:"userLocationInfo,omitempty" yaml:"userLocationInfo" bson:"userLocationInfo,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.
	UeTimeZone string `json:"ueTimeZone,omitempty" yaml:"ueTimeZone" bson:"ueTimeZone,omitempty"`
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.
	Pei string `json:"pei,omitempty" yaml:"pei" bson:"pei,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	Ipv4Address       string `json:"ipv4Address,omitempty" yaml:"ipv4Address" bson:"ipv4Address,omitempty"`
	Ipv6AddressPrefix string `json:"ipv6AddressPrefix,omitempty" yaml:"ipv6AddressPrefix" bson:"ipv6AddressPrefix,omitempty"`
	// Indicates the IPv4 address domain
	IpDomain     string `json:"ipDomain,omitempty" yaml:"ipDomain" bson:"ipDomain,omitempty"`
	SubsSessAmbr *Ambr  `json:"subsSessAmbr,omitempty" yaml:"subsSessAmbr" bson:"subsSessAmbr,omitempty"`
	// Indicates the DN-AAA authorization profile index
	AuthProfIndex string                `json:"authProfIndex,omitempty" yaml:"authProfIndex" bson:"authProfIndex,omitempty"`
	SubsDefQos    *SubscribedDefaultQos `json:"subsDefQos,omitempty" yaml:"subsDefQos" bson:"subsDefQos,omitempty"`
	VplmnQos      *VplmnQos             `json:"vplmnQos,omitempty" yaml:"vplmnQos" bson:"vplmnQos,omitempty"`
	// Contains the number of supported packet filter for signalled QoS rules.
	NumOfPackFilter int32 `json:"numOfPackFilter,omitempty" yaml:"numOfPackFilter" bson:"numOfPackFilter,omitempty"`
	// If it is included and set to true, the online charging is applied to the PDU session.
	Online bool `json:"online,omitempty" yaml:"online" bson:"online,omitempty"`
	// If it is included and set to true, the offline charging is applied to the PDU session.
	Offline bool `json:"offline,omitempty" yaml:"offline" bson:"offline,omitempty"`
	// If it is included and set to true, the 3GPP PS Data Off is activated by the UE.
	Var3gppPsDataOffStatus bool `json:"3gppPsDataOffStatus,omitempty" yaml:"3gppPsDataOffStatus" bson:"3gppPsDataOffStatus,omitempty"`
	// If it is included and set to true, the reflective QoS is supported by the UE.
	RefQosIndication bool               `json:"refQosIndication,omitempty" yaml:"refQosIndication" bson:"refQosIndication,omitempty"`
	TraceReq         *TraceData         `json:"traceReq,omitempty" yaml:"traceReq" bson:"traceReq,omitempty"`
	SliceInfo        *Snssai            `json:"sliceInfo" yaml:"sliceInfo" bson:"sliceInfo,omitempty"`
	QosFlowUsage     QosFlowUsage       `json:"qosFlowUsage,omitempty" yaml:"qosFlowUsage" bson:"qosFlowUsage,omitempty"`
	ServNfId         *ServingNfIdentity `json:"servNfId,omitempty" yaml:"servNfId" bson:"servNfId,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat string `json:"suppFeat,omitempty" yaml:"suppFeat" bson:"suppFeat,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	SmfId string `json:"smfId,omitempty" yaml:"smfId" bson:"smfId,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RecoveryTime        *time.Time                        `json:"recoveryTime,omitempty" yaml:"recoveryTime" bson:"recoveryTime,omitempty"`
	MaPduInd            MaPduIndication                   `json:"maPduInd,omitempty" yaml:"maPduInd" bson:"maPduInd,omitempty"`
	AtsssCapab          PcfSmPolicyControlAtsssCapability `json:"atsssCapab,omitempty" yaml:"atsssCapab" bson:"atsssCapab,omitempty"`
	Ipv4FrameRouteList  []string                          `json:"ipv4FrameRouteList,omitempty" yaml:"ipv4FrameRouteList" bson:"ipv4FrameRouteList,omitempty"`
	Ipv6FrameRouteList  []string                          `json:"ipv6FrameRouteList,omitempty" yaml:"ipv6FrameRouteList" bson:"ipv6FrameRouteList,omitempty"`
	SatBackhaulCategory SatelliteBackhaulCategory         `json:"satBackhaulCategory,omitempty" yaml:"satBackhaulCategory" bson:"satBackhaulCategory,omitempty"`
	PcfUeInfo           *PcfUeCallbackInfo                `json:"pcfUeInfo,omitempty" yaml:"pcfUeInfo" bson:"pcfUeInfo,omitempty"`
	PvsInfo             []ServerAddressingInfo            `json:"pvsInfo,omitempty" yaml:"pvsInfo" bson:"pvsInfo,omitempty"`
	// If it is included and set to true, it indicates that the PDU session is used for UE Onboarding.
	OnboardInd bool        `json:"onboardInd,omitempty" yaml:"onboardInd" bson:"onboardInd,omitempty"`
	NwdafDatas []NwdafData `json:"nwdafDatas,omitempty" yaml:"nwdafDatas" bson:"nwdafDatas,omitempty"`
}
