/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2022-02-14T08:05:27Z.
 *
 * API version: 0.0.1-38392
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// NiaapiApicLatestMaintainedRelease This contains the latest maintained release information for each release on APIC.
type NiaapiApicLatestMaintainedRelease struct {
	NiaapiMaintainedRelease
	AdditionalProperties map[string]interface{}
}

type _NiaapiApicLatestMaintainedRelease NiaapiApicLatestMaintainedRelease

// NewNiaapiApicLatestMaintainedRelease instantiates a new NiaapiApicLatestMaintainedRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiApicLatestMaintainedRelease() *NiaapiApicLatestMaintainedRelease {
	this := NiaapiApicLatestMaintainedRelease{}
	return &this
}

// NewNiaapiApicLatestMaintainedReleaseWithDefaults instantiates a new NiaapiApicLatestMaintainedRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiApicLatestMaintainedReleaseWithDefaults() *NiaapiApicLatestMaintainedRelease {
	this := NiaapiApicLatestMaintainedRelease{}
	return &this
}

func (o NiaapiApicLatestMaintainedRelease) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedNiaapiMaintainedRelease, errNiaapiMaintainedRelease := json.Marshal(o.NiaapiMaintainedRelease)
	if errNiaapiMaintainedRelease != nil {
		return []byte{}, errNiaapiMaintainedRelease
	}
	errNiaapiMaintainedRelease = json.Unmarshal([]byte(serializedNiaapiMaintainedRelease), &toSerialize)
	if errNiaapiMaintainedRelease != nil {
		return []byte{}, errNiaapiMaintainedRelease
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiApicLatestMaintainedRelease) UnmarshalJSON(bytes []byte) (err error) {
	type NiaapiApicLatestMaintainedReleaseWithoutEmbeddedStruct struct {
	}

	varNiaapiApicLatestMaintainedReleaseWithoutEmbeddedStruct := NiaapiApicLatestMaintainedReleaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiaapiApicLatestMaintainedReleaseWithoutEmbeddedStruct)
	if err == nil {
		varNiaapiApicLatestMaintainedRelease := _NiaapiApicLatestMaintainedRelease{}
		*o = NiaapiApicLatestMaintainedRelease(varNiaapiApicLatestMaintainedRelease)
	} else {
		return err
	}

	varNiaapiApicLatestMaintainedRelease := _NiaapiApicLatestMaintainedRelease{}

	err = json.Unmarshal(bytes, &varNiaapiApicLatestMaintainedRelease)
	if err == nil {
		o.NiaapiMaintainedRelease = varNiaapiApicLatestMaintainedRelease.NiaapiMaintainedRelease
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectNiaapiMaintainedRelease := reflect.ValueOf(o.NiaapiMaintainedRelease)
		for i := 0; i < reflectNiaapiMaintainedRelease.Type().NumField(); i++ {
			t := reflectNiaapiMaintainedRelease.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiaapiApicLatestMaintainedRelease struct {
	value *NiaapiApicLatestMaintainedRelease
	isSet bool
}

func (v NullableNiaapiApicLatestMaintainedRelease) Get() *NiaapiApicLatestMaintainedRelease {
	return v.value
}

func (v *NullableNiaapiApicLatestMaintainedRelease) Set(val *NiaapiApicLatestMaintainedRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiApicLatestMaintainedRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiApicLatestMaintainedRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiApicLatestMaintainedRelease(val *NiaapiApicLatestMaintainedRelease) *NullableNiaapiApicLatestMaintainedRelease {
	return &NullableNiaapiApicLatestMaintainedRelease{value: val, isSet: true}
}

func (v NullableNiaapiApicLatestMaintainedRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiApicLatestMaintainedRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
