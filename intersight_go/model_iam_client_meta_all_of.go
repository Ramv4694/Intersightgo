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

// IamClientMetaAllOf Definition of the list of properties defined in 'iam.ClientMeta', excluding properties defined in parent classes.
type IamClientMetaAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Parsed device model from raw User-Agent.
	DeviceModel *string `json:"DeviceModel,omitempty"`
	// The value of the \"User-Agent\" HTTP header, as sent by the HTTP client when it initiate a session to Intersight. This can be used to identify the client operating system, browser type and browser version. Example - Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36 It is set when User successfully passed OAuth login flow and receives Access Token.
	UserAgent            *string `json:"UserAgent,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamClientMetaAllOf IamClientMetaAllOf

// NewIamClientMetaAllOf instantiates a new IamClientMetaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamClientMetaAllOf(classId string, objectType string) *IamClientMetaAllOf {
	this := IamClientMetaAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamClientMetaAllOfWithDefaults instantiates a new IamClientMetaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamClientMetaAllOfWithDefaults() *IamClientMetaAllOf {
	this := IamClientMetaAllOf{}
	var classId string = "iam.ClientMeta"
	this.ClassId = classId
	var objectType string = "iam.ClientMeta"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamClientMetaAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamClientMetaAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamClientMetaAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamClientMetaAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamClientMetaAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamClientMetaAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise.
func (o *IamClientMetaAllOf) GetDeviceModel() string {
	if o == nil || o.DeviceModel == nil {
		var ret string
		return ret
	}
	return *o.DeviceModel
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamClientMetaAllOf) GetDeviceModelOk() (*string, bool) {
	if o == nil || o.DeviceModel == nil {
		return nil, false
	}
	return o.DeviceModel, true
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *IamClientMetaAllOf) HasDeviceModel() bool {
	if o != nil && o.DeviceModel != nil {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.
func (o *IamClientMetaAllOf) SetDeviceModel(v string) {
	o.DeviceModel = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *IamClientMetaAllOf) GetUserAgent() string {
	if o == nil || o.UserAgent == nil {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamClientMetaAllOf) GetUserAgentOk() (*string, bool) {
	if o == nil || o.UserAgent == nil {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *IamClientMetaAllOf) HasUserAgent() bool {
	if o != nil && o.UserAgent != nil {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *IamClientMetaAllOf) SetUserAgent(v string) {
	o.UserAgent = &v
}

func (o IamClientMetaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DeviceModel != nil {
		toSerialize["DeviceModel"] = o.DeviceModel
	}
	if o.UserAgent != nil {
		toSerialize["UserAgent"] = o.UserAgent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamClientMetaAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamClientMetaAllOf := _IamClientMetaAllOf{}

	if err = json.Unmarshal(bytes, &varIamClientMetaAllOf); err == nil {
		*o = IamClientMetaAllOf(varIamClientMetaAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceModel")
		delete(additionalProperties, "UserAgent")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamClientMetaAllOf struct {
	value *IamClientMetaAllOf
	isSet bool
}

func (v NullableIamClientMetaAllOf) Get() *IamClientMetaAllOf {
	return v.value
}

func (v *NullableIamClientMetaAllOf) Set(val *IamClientMetaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamClientMetaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamClientMetaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamClientMetaAllOf(val *IamClientMetaAllOf) *NullableIamClientMetaAllOf {
	return &NullableIamClientMetaAllOf{value: val, isSet: true}
}

func (v NullableIamClientMetaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamClientMetaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}