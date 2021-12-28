/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-12-27T12:26:28Z.
 *
 * API version: 0.0.1-37434
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// FirmwareEulaResponse - The response body of a HTTP GET request for the 'firmware.Eula' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'firmware.Eula' resources.
type FirmwareEulaResponse struct {
	FirmwareEulaList     *FirmwareEulaList
	MoAggregateTransform *MoAggregateTransform
	MoDocumentCount      *MoDocumentCount
	MoTagSummary         *MoTagSummary
}

// FirmwareEulaListAsFirmwareEulaResponse is a convenience function that returns FirmwareEulaList wrapped in FirmwareEulaResponse
func FirmwareEulaListAsFirmwareEulaResponse(v *FirmwareEulaList) FirmwareEulaResponse {
	return FirmwareEulaResponse{FirmwareEulaList: v}
}

// MoAggregateTransformAsFirmwareEulaResponse is a convenience function that returns MoAggregateTransform wrapped in FirmwareEulaResponse
func MoAggregateTransformAsFirmwareEulaResponse(v *MoAggregateTransform) FirmwareEulaResponse {
	return FirmwareEulaResponse{MoAggregateTransform: v}
}

// MoDocumentCountAsFirmwareEulaResponse is a convenience function that returns MoDocumentCount wrapped in FirmwareEulaResponse
func MoDocumentCountAsFirmwareEulaResponse(v *MoDocumentCount) FirmwareEulaResponse {
	return FirmwareEulaResponse{MoDocumentCount: v}
}

// MoTagSummaryAsFirmwareEulaResponse is a convenience function that returns MoTagSummary wrapped in FirmwareEulaResponse
func MoTagSummaryAsFirmwareEulaResponse(v *MoTagSummary) FirmwareEulaResponse {
	return FirmwareEulaResponse{MoTagSummary: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *FirmwareEulaResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'firmware.Eula.List'
	if jsonDict["ObjectType"] == "firmware.Eula.List" {
		// try to unmarshal JSON data into FirmwareEulaList
		err = json.Unmarshal(data, &dst.FirmwareEulaList)
		if err == nil {
			return nil // data stored in dst.FirmwareEulaList, return on the first match
		} else {
			dst.FirmwareEulaList = nil
			return fmt.Errorf("Failed to unmarshal FirmwareEulaResponse as FirmwareEulaList: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.AggregateTransform'
	if jsonDict["ObjectType"] == "mo.AggregateTransform" {
		// try to unmarshal JSON data into MoAggregateTransform
		err = json.Unmarshal(data, &dst.MoAggregateTransform)
		if err == nil {
			return nil // data stored in dst.MoAggregateTransform, return on the first match
		} else {
			dst.MoAggregateTransform = nil
			return fmt.Errorf("Failed to unmarshal FirmwareEulaResponse as MoAggregateTransform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.DocumentCount'
	if jsonDict["ObjectType"] == "mo.DocumentCount" {
		// try to unmarshal JSON data into MoDocumentCount
		err = json.Unmarshal(data, &dst.MoDocumentCount)
		if err == nil {
			return nil // data stored in dst.MoDocumentCount, return on the first match
		} else {
			dst.MoDocumentCount = nil
			return fmt.Errorf("Failed to unmarshal FirmwareEulaResponse as MoDocumentCount: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.TagSummary'
	if jsonDict["ObjectType"] == "mo.TagSummary" {
		// try to unmarshal JSON data into MoTagSummary
		err = json.Unmarshal(data, &dst.MoTagSummary)
		if err == nil {
			return nil // data stored in dst.MoTagSummary, return on the first match
		} else {
			dst.MoTagSummary = nil
			return fmt.Errorf("Failed to unmarshal FirmwareEulaResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FirmwareEulaResponse) MarshalJSON() ([]byte, error) {
	if src.FirmwareEulaList != nil {
		return json.Marshal(&src.FirmwareEulaList)
	}

	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FirmwareEulaResponse) GetActualInstance() interface{} {
	if obj.FirmwareEulaList != nil {
		return obj.FirmwareEulaList
	}

	if obj.MoAggregateTransform != nil {
		return obj.MoAggregateTransform
	}

	if obj.MoDocumentCount != nil {
		return obj.MoDocumentCount
	}

	if obj.MoTagSummary != nil {
		return obj.MoTagSummary
	}

	// all schemas are nil
	return nil
}

type NullableFirmwareEulaResponse struct {
	value *FirmwareEulaResponse
	isSet bool
}

func (v NullableFirmwareEulaResponse) Get() *FirmwareEulaResponse {
	return v.value
}

func (v *NullableFirmwareEulaResponse) Set(val *FirmwareEulaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareEulaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareEulaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareEulaResponse(val *FirmwareEulaResponse) *NullableFirmwareEulaResponse {
	return &NullableFirmwareEulaResponse{value: val, isSet: true}
}

func (v NullableFirmwareEulaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareEulaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
