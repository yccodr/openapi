/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.502 V17.11.0; 5G System; Session Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.502/
 *
 * API version: 1.2.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Complete SM Context
type SmContext struct {
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.
	PduSessionId int32 `json:"pduSessionId" yaml:"pduSessionId" bson:"pduSessionId,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn" yaml:"dnn" bson:"dnn,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	SelectedDnn    string         `json:"selectedDnn,omitempty" yaml:"selectedDnn" bson:"selectedDnn,omitempty"`
	SNssai         *Snssai        `json:"sNssai" yaml:"sNssai" bson:"sNssai,omitempty"`
	HplmnSnssai    *Snssai        `json:"hplmnSnssai,omitempty" yaml:"hplmnSnssai" bson:"hplmnSnssai,omitempty"`
	PduSessionType PduSessionType `json:"pduSessionType" yaml:"pduSessionType" bson:"pduSessionType,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi string `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	HSmfUri string `json:"hSmfUri,omitempty" yaml:"hSmfUri" bson:"hSmfUri,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	SmfUri string `json:"smfUri,omitempty" yaml:"smfUri" bson:"smfUri,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	PduSessionRef string `json:"pduSessionRef,omitempty" yaml:"pduSessionRef" bson:"pduSessionRef,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	InterPlmnApiRoot string `json:"interPlmnApiRoot,omitempty" yaml:"interPlmnApiRoot" bson:"interPlmnApiRoot,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	IntraPlmnApiRoot string `json:"intraPlmnApiRoot,omitempty" yaml:"intraPlmnApiRoot" bson:"intraPlmnApiRoot,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	PcfId string `json:"pcfId,omitempty" yaml:"pcfId" bson:"pcfId,omitempty"`
	// Identifier of a group of NFs.
	PcfGroupId string `json:"pcfGroupId,omitempty" yaml:"pcfGroupId" bson:"pcfGroupId,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.
	PcfSetId string           `json:"pcfSetId,omitempty" yaml:"pcfSetId" bson:"pcfSetId,omitempty"`
	SelMode  DnnSelectionMode `json:"selMode,omitempty" yaml:"selMode" bson:"selMode,omitempty"`
	// Identifier of a group of NFs.
	UdmGroupId       string             `json:"udmGroupId,omitempty" yaml:"udmGroupId" bson:"udmGroupId,omitempty"`
	RoutingIndicator string             `json:"routingIndicator,omitempty" yaml:"routingIndicator" bson:"routingIndicator,omitempty"`
	HNwPubKeyId      int32              `json:"hNwPubKeyId,omitempty" yaml:"hNwPubKeyId" bson:"hNwPubKeyId,omitempty"`
	SessionAmbr      *Ambr              `json:"sessionAmbr" yaml:"sessionAmbr" bson:"sessionAmbr,omitempty"`
	QosFlowsList     []QosFlowSetupItem `json:"qosFlowsList" yaml:"qosFlowsList" bson:"qosFlowsList,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	HSmfInstanceId string `json:"hSmfInstanceId,omitempty" yaml:"hSmfInstanceId" bson:"hSmfInstanceId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	SmfInstanceId string `json:"smfInstanceId,omitempty" yaml:"smfInstanceId" bson:"smfInstanceId,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.
	PduSessionSmfSetId string `json:"pduSessionSmfSetId,omitempty" yaml:"pduSessionSmfSetId" bson:"pduSessionSmfSetId,omitempty"`
	// NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \"set<Set ID>.sn<Service Name>.nfi<NF Instance ID>.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.sn<ServiceName>.nfi<NFInstanceID>.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)   <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NID> encoded as defined in clause 5.4.2 (\"Nid\" data type definition)  <NFInstanceId> encoded as defined in clause 5.3.2  <ServiceName> encoded as defined in 3GPP TS 29.510  <Set ID> encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit.
	PduSessionSmfServiceSetId string          `json:"pduSessionSmfServiceSetId,omitempty" yaml:"pduSessionSmfServiceSetId" bson:"pduSessionSmfServiceSetId,omitempty"`
	PduSessionSmfBinding      SbiBindingLevel `json:"pduSessionSmfBinding,omitempty" yaml:"pduSessionSmfBinding" bson:"pduSessionSmfBinding,omitempty"`
	EnablePauseCharging       bool            `json:"enablePauseCharging,omitempty" yaml:"enablePauseCharging" bson:"enablePauseCharging,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	UeIpv4Address                   string                        `json:"ueIpv4Address,omitempty" yaml:"ueIpv4Address" bson:"ueIpv4Address,omitempty"`
	UeIpv6Prefix                    string                        `json:"ueIpv6Prefix,omitempty" yaml:"ueIpv6Prefix" bson:"ueIpv6Prefix,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                `json:"epsPdnCnxInfo,omitempty" yaml:"epsPdnCnxInfo" bson:"epsPdnCnxInfo,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo               `json:"epsBearerInfo,omitempty" yaml:"epsBearerInfo" bson:"epsBearerInfo,omitempty"`
	MaxIntegrityProtectedDataRate   MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRate,omitempty" yaml:"maxIntegrityProtectedDataRate" bson:"maxIntegrityProtectedDataRate,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateDl,omitempty" yaml:"maxIntegrityProtectedDataRateDl" bson:"maxIntegrityProtectedDataRateDl,omitempty"`
	AlwaysOnGranted                 bool                          `json:"alwaysOnGranted,omitempty" yaml:"alwaysOnGranted" bson:"alwaysOnGranted,omitempty"`
	UpSecurity                      *UpSecurity                   `json:"upSecurity,omitempty" yaml:"upSecurity" bson:"upSecurity,omitempty"`
	HSmfServiceInstanceId           string                        `json:"hSmfServiceInstanceId,omitempty" yaml:"hSmfServiceInstanceId" bson:"hSmfServiceInstanceId,omitempty"`
	SmfServiceInstanceId            string                        `json:"smfServiceInstanceId,omitempty" yaml:"smfServiceInstanceId" bson:"smfServiceInstanceId,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RecoveryTime           *time.Time                                  `json:"recoveryTime,omitempty" yaml:"recoveryTime" bson:"recoveryTime,omitempty"`
	ForwardingInd          bool                                        `json:"forwardingInd,omitempty" yaml:"forwardingInd" bson:"forwardingInd,omitempty"`
	PsaTunnelInfo          *TunnelInfo                                 `json:"psaTunnelInfo,omitempty" yaml:"psaTunnelInfo" bson:"psaTunnelInfo,omitempty"`
	ChargingId             string                                      `json:"chargingId,omitempty" yaml:"chargingId" bson:"chargingId,omitempty"`
	ChargingInfo           *ChargingInformation                        `json:"chargingInfo,omitempty" yaml:"chargingInfo" bson:"chargingInfo,omitempty"`
	RoamingChargingProfile *ChfConvergedChargingRoamingChargingProfile `json:"roamingChargingProfile,omitempty" yaml:"roamingChargingProfile" bson:"roamingChargingProfile,omitempty"`
	NefExtBufSupportInd    bool                                        `json:"nefExtBufSupportInd,omitempty" yaml:"nefExtBufSupportInd" bson:"nefExtBufSupportInd,omitempty"`
	// Represents information that identifies which IP pool or external server is used to allocate the IP address.
	Ipv6Index               int32                           `json:"ipv6Index,omitempty" yaml:"ipv6Index" bson:"ipv6Index,omitempty"`
	DnAaaAddress            *SmfPduSessionIpAddress         `json:"dnAaaAddress,omitempty" yaml:"dnAaaAddress" bson:"dnAaaAddress,omitempty"`
	RedundantPduSessionInfo *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty" yaml:"redundantPduSessionInfo" bson:"redundantPduSessionInfo,omitempty"`
	RanTunnelInfo           *QosFlowTunnel                  `json:"ranTunnelInfo,omitempty" yaml:"ranTunnelInfo" bson:"ranTunnelInfo,omitempty"`
	AddRanTunnelInfo        []QosFlowTunnel                 `json:"addRanTunnelInfo,omitempty" yaml:"addRanTunnelInfo" bson:"addRanTunnelInfo,omitempty"`
	RedRanTunnelInfo        *QosFlowTunnel                  `json:"redRanTunnelInfo,omitempty" yaml:"redRanTunnelInfo" bson:"redRanTunnelInfo,omitempty"`
	AddRedRanTunnelInfo     []QosFlowTunnel                 `json:"addRedRanTunnelInfo,omitempty" yaml:"addRedRanTunnelInfo" bson:"addRedRanTunnelInfo,omitempty"`
	NspuSupportInd          bool                            `json:"nspuSupportInd,omitempty" yaml:"nspuSupportInd" bson:"nspuSupportInd,omitempty"`
	SmfBindingInfo          string                          `json:"smfBindingInfo,omitempty" yaml:"smfBindingInfo" bson:"smfBindingInfo,omitempty"`
	SatelliteBackhaulCat    SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty" yaml:"satelliteBackhaulCat" bson:"satelliteBackhaulCat,omitempty"`
	SscMode                 string                          `json:"sscMode,omitempty" yaml:"sscMode" bson:"sscMode,omitempty"`
	DlsetSupportInd         bool                            `json:"dlsetSupportInd,omitempty" yaml:"dlsetSupportInd" bson:"dlsetSupportInd,omitempty"`
	N9fscSupportInd         bool                            `json:"n9fscSupportInd,omitempty" yaml:"n9fscSupportInd" bson:"n9fscSupportInd,omitempty"`
	AnchorSmfOauth2Required bool                            `json:"anchorSmfOauth2Required,omitempty" yaml:"anchorSmfOauth2Required" bson:"anchorSmfOauth2Required,omitempty"`
}
