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

// TunnelingTunnelAllOf Definition of the list of properties defined in 'tunneling.Tunnel', excluding properties defined in parent classes.
type TunnelingTunnelAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The multiplexer URL for the client to connect on.
	ClientUrl            *string `json:"ClientUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TunnelingTunnelAllOf TunnelingTunnelAllOf

// NewTunnelingTunnelAllOf instantiates a new TunnelingTunnelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTunnelingTunnelAllOf(classId string, objectType string) *TunnelingTunnelAllOf {
	this := TunnelingTunnelAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTunnelingTunnelAllOfWithDefaults instantiates a new TunnelingTunnelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTunnelingTunnelAllOfWithDefaults() *TunnelingTunnelAllOf {
	this := TunnelingTunnelAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *TunnelingTunnelAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TunnelingTunnelAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TunnelingTunnelAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TunnelingTunnelAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TunnelingTunnelAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TunnelingTunnelAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClientUrl returns the ClientUrl field value if set, zero value otherwise.
func (o *TunnelingTunnelAllOf) GetClientUrl() string {
	if o == nil || o.ClientUrl == nil {
		var ret string
		return ret
	}
	return *o.ClientUrl
}

// GetClientUrlOk returns a tuple with the ClientUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelingTunnelAllOf) GetClientUrlOk() (*string, bool) {
	if o == nil || o.ClientUrl == nil {
		return nil, false
	}
	return o.ClientUrl, true
}

// HasClientUrl returns a boolean if a field has been set.
func (o *TunnelingTunnelAllOf) HasClientUrl() bool {
	if o != nil && o.ClientUrl != nil {
		return true
	}

	return false
}

// SetClientUrl gets a reference to the given string and assigns it to the ClientUrl field.
func (o *TunnelingTunnelAllOf) SetClientUrl(v string) {
	o.ClientUrl = &v
}

func (o TunnelingTunnelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClientUrl != nil {
		toSerialize["ClientUrl"] = o.ClientUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TunnelingTunnelAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTunnelingTunnelAllOf := _TunnelingTunnelAllOf{}

	if err = json.Unmarshal(bytes, &varTunnelingTunnelAllOf); err == nil {
		*o = TunnelingTunnelAllOf(varTunnelingTunnelAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClientUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTunnelingTunnelAllOf struct {
	value *TunnelingTunnelAllOf
	isSet bool
}

func (v NullableTunnelingTunnelAllOf) Get() *TunnelingTunnelAllOf {
	return v.value
}

func (v *NullableTunnelingTunnelAllOf) Set(val *TunnelingTunnelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTunnelingTunnelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTunnelingTunnelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunnelingTunnelAllOf(val *TunnelingTunnelAllOf) *NullableTunnelingTunnelAllOf {
	return &NullableTunnelingTunnelAllOf{value: val, isSet: true}
}

func (v NullableTunnelingTunnelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunnelingTunnelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
