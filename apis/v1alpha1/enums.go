// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type LanguageCodeString string

const (
	LanguageCodeString_de_DE  LanguageCodeString = "de-DE"
	LanguageCodeString_en_GB  LanguageCodeString = "en-GB"
	LanguageCodeString_en_US  LanguageCodeString = "en-US"
	LanguageCodeString_es_419 LanguageCodeString = "es-419"
	LanguageCodeString_es_ES  LanguageCodeString = "es-ES"
	LanguageCodeString_fr_CA  LanguageCodeString = "fr-CA"
	LanguageCodeString_fr_FR  LanguageCodeString = "fr-FR"
	LanguageCodeString_it_IT  LanguageCodeString = "it-IT"
	LanguageCodeString_ja_JP  LanguageCodeString = "ja-JP"
	LanguageCodeString_kr_KR  LanguageCodeString = "kr-KR"
	LanguageCodeString_pt_BR  LanguageCodeString = "pt-BR"
	LanguageCodeString_zh_CN  LanguageCodeString = "zh-CN"
	LanguageCodeString_zh_TW  LanguageCodeString = "zh-TW"
)

type NumberCapability string

const (
	NumberCapability_MMS   NumberCapability = "MMS"
	NumberCapability_SMS   NumberCapability = "SMS"
	NumberCapability_VOICE NumberCapability = "VOICE"
)

type RouteType string

const (
	RouteType_Premium       RouteType = "Premium"
	RouteType_Promotional   RouteType = "Promotional"
	RouteType_Transactional RouteType = "Transactional"
)

type SMSSandboxPhoneNumberVerificationStatus string

const (
	SMSSandboxPhoneNumberVerificationStatus_Pending  SMSSandboxPhoneNumberVerificationStatus = "Pending"
	SMSSandboxPhoneNumberVerificationStatus_Verified SMSSandboxPhoneNumberVerificationStatus = "Verified"
)
