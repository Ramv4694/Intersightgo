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
	"reflect"
	"strings"
)

// VirtualizationBaseDistributedNetwork Common attributes of any distributed virtual network managed by a hypervisor.
type VirtualizationBaseDistributedNetwork struct {
	VirtualizationBaseNetwork
	AdditionalProperties map[string]interface{}
}

type _VirtualizationBaseDistributedNetwork VirtualizationBaseDistributedNetwork

// NewVirtualizationBaseDistributedNetwork instantiates a new VirtualizationBaseDistributedNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBaseDistributedNetwork() *VirtualizationBaseDistributedNetwork {
	this := VirtualizationBaseDistributedNetwork{}
	return &this
}

// NewVirtualizationBaseDistributedNetworkWithDefaults instantiates a new VirtualizationBaseDistributedNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBaseDistributedNetworkWithDefaults() *VirtualizationBaseDistributedNetwork {
	this := VirtualizationBaseDistributedNetwork{}
	return &this
}

func (o VirtualizationBaseDistributedNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseNetwork, errVirtualizationBaseNetwork := json.Marshal(o.VirtualizationBaseNetwork)
	if errVirtualizationBaseNetwork != nil {
		return []byte{}, errVirtualizationBaseNetwork
	}
	errVirtualizationBaseNetwork = json.Unmarshal([]byte(serializedVirtualizationBaseNetwork), &toSerialize)
	if errVirtualizationBaseNetwork != nil {
		return []byte{}, errVirtualizationBaseNetwork
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationBaseDistributedNetwork) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationBaseDistributedNetworkWithoutEmbeddedStruct struct {
	}

	varVirtualizationBaseDistributedNetworkWithoutEmbeddedStruct := VirtualizationBaseDistributedNetworkWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationBaseDistributedNetworkWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationBaseDistributedNetwork := _VirtualizationBaseDistributedNetwork{}
		*o = VirtualizationBaseDistributedNetwork(varVirtualizationBaseDistributedNetwork)
	} else {
		return err
	}

	varVirtualizationBaseDistributedNetwork := _VirtualizationBaseDistributedNetwork{}

	err = json.Unmarshal(bytes, &varVirtualizationBaseDistributedNetwork)
	if err == nil {
		o.VirtualizationBaseNetwork = varVirtualizationBaseDistributedNetwork.VirtualizationBaseNetwork
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectVirtualizationBaseNetwork := reflect.ValueOf(o.VirtualizationBaseNetwork)
		for i := 0; i < reflectVirtualizationBaseNetwork.Type().NumField(); i++ {
			t := reflectVirtualizationBaseNetwork.Type().Field(i)

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

type NullableVirtualizationBaseDistributedNetwork struct {
	value *VirtualizationBaseDistributedNetwork
	isSet bool
}

func (v NullableVirtualizationBaseDistributedNetwork) Get() *VirtualizationBaseDistributedNetwork {
	return v.value
}

func (v *NullableVirtualizationBaseDistributedNetwork) Set(val *VirtualizationBaseDistributedNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBaseDistributedNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBaseDistributedNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBaseDistributedNetwork(val *VirtualizationBaseDistributedNetwork) *NullableVirtualizationBaseDistributedNetwork {
	return &NullableVirtualizationBaseDistributedNetwork{value: val, isSet: true}
}

func (v NullableVirtualizationBaseDistributedNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBaseDistributedNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
