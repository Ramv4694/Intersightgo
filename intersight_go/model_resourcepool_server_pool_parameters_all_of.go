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
)

// ResourcepoolServerPoolParametersAllOf Definition of the list of properties defined in 'resourcepool.ServerPoolParameters', excluding properties defined in parent classes.
type ResourcepoolServerPoolParametersAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The platform for which the servers in resource pool are applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `IntersightStandalone` - Intersight Standalone mode of operation. * `UCSM` - Unified Computing System Manager mode of operation. * `Intersight` - Intersight managed mode of operation.
	ManagementMode       *string `json:"ManagementMode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourcepoolServerPoolParametersAllOf ResourcepoolServerPoolParametersAllOf

// NewResourcepoolServerPoolParametersAllOf instantiates a new ResourcepoolServerPoolParametersAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcepoolServerPoolParametersAllOf(classId string, objectType string) *ResourcepoolServerPoolParametersAllOf {
	this := ResourcepoolServerPoolParametersAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var managementMode string = "IntersightStandalone"
	this.ManagementMode = &managementMode
	return &this
}

// NewResourcepoolServerPoolParametersAllOfWithDefaults instantiates a new ResourcepoolServerPoolParametersAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcepoolServerPoolParametersAllOfWithDefaults() *ResourcepoolServerPoolParametersAllOf {
	this := ResourcepoolServerPoolParametersAllOf{}
	var classId string = "resourcepool.ServerPoolParameters"
	this.ClassId = classId
	var objectType string = "resourcepool.ServerPoolParameters"
	this.ObjectType = objectType
	var managementMode string = "IntersightStandalone"
	this.ManagementMode = &managementMode
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourcepoolServerPoolParametersAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourcepoolServerPoolParametersAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourcepoolServerPoolParametersAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ResourcepoolServerPoolParametersAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourcepoolServerPoolParametersAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourcepoolServerPoolParametersAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetManagementMode returns the ManagementMode field value if set, zero value otherwise.
func (o *ResourcepoolServerPoolParametersAllOf) GetManagementMode() string {
	if o == nil || o.ManagementMode == nil {
		var ret string
		return ret
	}
	return *o.ManagementMode
}

// GetManagementModeOk returns a tuple with the ManagementMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolServerPoolParametersAllOf) GetManagementModeOk() (*string, bool) {
	if o == nil || o.ManagementMode == nil {
		return nil, false
	}
	return o.ManagementMode, true
}

// HasManagementMode returns a boolean if a field has been set.
func (o *ResourcepoolServerPoolParametersAllOf) HasManagementMode() bool {
	if o != nil && o.ManagementMode != nil {
		return true
	}

	return false
}

// SetManagementMode gets a reference to the given string and assigns it to the ManagementMode field.
func (o *ResourcepoolServerPoolParametersAllOf) SetManagementMode(v string) {
	o.ManagementMode = &v
}

func (o ResourcepoolServerPoolParametersAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ManagementMode != nil {
		toSerialize["ManagementMode"] = o.ManagementMode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourcepoolServerPoolParametersAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varResourcepoolServerPoolParametersAllOf := _ResourcepoolServerPoolParametersAllOf{}

	if err = json.Unmarshal(bytes, &varResourcepoolServerPoolParametersAllOf); err == nil {
		*o = ResourcepoolServerPoolParametersAllOf(varResourcepoolServerPoolParametersAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ManagementMode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourcepoolServerPoolParametersAllOf struct {
	value *ResourcepoolServerPoolParametersAllOf
	isSet bool
}

func (v NullableResourcepoolServerPoolParametersAllOf) Get() *ResourcepoolServerPoolParametersAllOf {
	return v.value
}

func (v *NullableResourcepoolServerPoolParametersAllOf) Set(val *ResourcepoolServerPoolParametersAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcepoolServerPoolParametersAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcepoolServerPoolParametersAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcepoolServerPoolParametersAllOf(val *ResourcepoolServerPoolParametersAllOf) *NullableResourcepoolServerPoolParametersAllOf {
	return &NullableResourcepoolServerPoolParametersAllOf{value: val, isSet: true}
}

func (v NullableResourcepoolServerPoolParametersAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcepoolServerPoolParametersAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
