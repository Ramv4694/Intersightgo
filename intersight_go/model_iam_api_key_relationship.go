/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-12-24T09:42:08Z.
 *
 * API version: 0.0.1-37430
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// IamApiKeyRelationship - A relationship to the 'iam.ApiKey' resource, or the expanded 'iam.ApiKey' resource, or the 'null' value.
type IamApiKeyRelationship struct {
	IamApiKey *IamApiKey
	MoMoRef   *MoMoRef
}

// IamApiKeyAsIamApiKeyRelationship is a convenience function that returns IamApiKey wrapped in IamApiKeyRelationship
func IamApiKeyAsIamApiKeyRelationship(v *IamApiKey) IamApiKeyRelationship {
	return IamApiKeyRelationship{IamApiKey: v}
}

// MoMoRefAsIamApiKeyRelationship is a convenience function that returns MoMoRef wrapped in IamApiKeyRelationship
func MoMoRefAsIamApiKeyRelationship(v *MoMoRef) IamApiKeyRelationship {
	return IamApiKeyRelationship{MoMoRef: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IamApiKeyRelationship) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'iam.ApiKey'
	if jsonDict["ClassId"] == "iam.ApiKey" {
		// try to unmarshal JSON data into IamApiKey
		err = json.Unmarshal(data, &dst.IamApiKey)
		if err == nil {
			return nil // data stored in dst.IamApiKey, return on the first match
		} else {
			dst.IamApiKey = nil
			return fmt.Errorf("Failed to unmarshal IamApiKeyRelationship as IamApiKey: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			return nil // data stored in dst.MoMoRef, return on the first match
		} else {
			dst.MoMoRef = nil
			return fmt.Errorf("Failed to unmarshal IamApiKeyRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IamApiKeyRelationship) MarshalJSON() ([]byte, error) {
	if src.IamApiKey != nil {
		return json.Marshal(&src.IamApiKey)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IamApiKeyRelationship) GetActualInstance() interface{} {
	if obj.IamApiKey != nil {
		return obj.IamApiKey
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableIamApiKeyRelationship struct {
	value *IamApiKeyRelationship
	isSet bool
}

func (v NullableIamApiKeyRelationship) Get() *IamApiKeyRelationship {
	return v.value
}

func (v *NullableIamApiKeyRelationship) Set(val *IamApiKeyRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableIamApiKeyRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableIamApiKeyRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamApiKeyRelationship(val *IamApiKeyRelationship) *NullableIamApiKeyRelationship {
	return &NullableIamApiKeyRelationship{value: val, isSet: true}
}

func (v NullableIamApiKeyRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamApiKeyRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
