/*
 * Npcf_UEPolicyControl
 *
 * UE Policy Control Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.525 V16.9.0; 5G System; UE Policy Control Service.
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.525/
 *
 * API version: 1.1.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type NgapIeType string

// List of NgapIeType
const (
	NgapIeType_PDU_RES_SETUP_REQ            NgapIeType = "PDU_RES_SETUP_REQ"
	NgapIeType_PDU_RES_REL_CMD              NgapIeType = "PDU_RES_REL_CMD"
	NgapIeType_PDU_RES_MOD_REQ              NgapIeType = "PDU_RES_MOD_REQ"
	NgapIeType_HANDOVER_CMD                 NgapIeType = "HANDOVER_CMD"
	NgapIeType_HANDOVER_REQUIRED            NgapIeType = "HANDOVER_REQUIRED"
	NgapIeType_HANDOVER_PREP_FAIL           NgapIeType = "HANDOVER_PREP_FAIL"
	NgapIeType_SRC_TO_TAR_CONTAINER         NgapIeType = "SRC_TO_TAR_CONTAINER"
	NgapIeType_TAR_TO_SRC_CONTAINER         NgapIeType = "TAR_TO_SRC_CONTAINER"
	NgapIeType_TAR_TO_SRC_FAIL_CONTAINER    NgapIeType = "TAR_TO_SRC_FAIL_CONTAINER"
	NgapIeType_RAN_STATUS_TRANS_CONTAINER   NgapIeType = "RAN_STATUS_TRANS_CONTAINER"
	NgapIeType_SON_CONFIG_TRANSFER          NgapIeType = "SON_CONFIG_TRANSFER"
	NgapIeType_NRPPA_PDU                    NgapIeType = "NRPPA_PDU"
	NgapIeType_UE_RADIO_CAPABILITY          NgapIeType = "UE_RADIO_CAPABILITY"
	NgapIeType_RIM_INFO_TRANSFER            NgapIeType = "RIM_INFO_TRANSFER"
	NgapIeType_SECONDARY_RAT_USAGE          NgapIeType = "SECONDARY_RAT_USAGE"
	NgapIeType_PC5_QOS_PARA                 NgapIeType = "PC5_QOS_PARA"
	NgapIeType_EARLY_STATUS_TRANS_CONTAINER NgapIeType = "EARLY_STATUS_TRANS_CONTAINER"
)
