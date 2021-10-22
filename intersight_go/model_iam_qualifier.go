/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-10-09T21:18:32Z.
 *
 * API version: 1.0.9-4809
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// IamQualifier The qualifier defines how a user qualifies to be part of a user group.
type IamQualifier struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the SAML attribute used to qualify a user to user group. By default this is memberOf attribute in SAML assertion.
	Name                 *string                   `json:"Name,omitempty"`
	Value                []string                  `json:"Value,omitempty"`
	Usergroup            *IamUserGroupRelationship `json:"Usergroup,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamQualifier IamQualifier

// NewIamQualifier instantiates a new IamQualifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamQualifier(classId string, objectType string) *IamQualifier {
	this := IamQualifier{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamQualifierWithDefaults instantiates a new IamQualifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamQualifierWithDefaults() *IamQualifier {
	this := IamQualifier{}
	var classId string = "iam.Qualifier"
	this.ClassId = classId
	var objectType string = "iam.Qualifier"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamQualifier) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamQualifier) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamQualifier) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamQualifier) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamQualifier) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamQualifier) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamQualifier) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamQualifier) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamQualifier) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamQualifier) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamQualifier) GetValue() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamQualifier) GetValueOk() (*[]string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IamQualifier) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *IamQualifier) SetValue(v []string) {
	o.Value = v
}

// GetUsergroup returns the Usergroup field value if set, zero value otherwise.
func (o *IamQualifier) GetUsergroup() IamUserGroupRelationship {
	if o == nil || o.Usergroup == nil {
		var ret IamUserGroupRelationship
		return ret
	}
	return *o.Usergroup
}

// GetUsergroupOk returns a tuple with the Usergroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamQualifier) GetUsergroupOk() (*IamUserGroupRelationship, bool) {
	if o == nil || o.Usergroup == nil {
		return nil, false
	}
	return o.Usergroup, true
}

// HasUsergroup returns a boolean if a field has been set.
func (o *IamQualifier) HasUsergroup() bool {
	if o != nil && o.Usergroup != nil {
		return true
	}

	return false
}

// SetUsergroup gets a reference to the given IamUserGroupRelationship and assigns it to the Usergroup field.
func (o *IamQualifier) SetUsergroup(v IamUserGroupRelationship) {
	o.Usergroup = &v
}

func (o IamQualifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	if o.Usergroup != nil {
		toSerialize["Usergroup"] = o.Usergroup
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamQualifier) UnmarshalJSON(bytes []byte) (err error) {
	type IamQualifierWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of the SAML attribute used to qualify a user to user group. By default this is memberOf attribute in SAML assertion.
		Name      *string                   `json:"Name,omitempty"`
		Value     []string                  `json:"Value,omitempty"`
		Usergroup *IamUserGroupRelationship `json:"Usergroup,omitempty"`
	}

	varIamQualifierWithoutEmbeddedStruct := IamQualifierWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamQualifierWithoutEmbeddedStruct)
	if err == nil {
		varIamQualifier := _IamQualifier{}
		varIamQualifier.ClassId = varIamQualifierWithoutEmbeddedStruct.ClassId
		varIamQualifier.ObjectType = varIamQualifierWithoutEmbeddedStruct.ObjectType
		varIamQualifier.Name = varIamQualifierWithoutEmbeddedStruct.Name
		varIamQualifier.Value = varIamQualifierWithoutEmbeddedStruct.Value
		varIamQualifier.Usergroup = varIamQualifierWithoutEmbeddedStruct.Usergroup
		*o = IamQualifier(varIamQualifier)
	} else {
		return err
	}

	varIamQualifier := _IamQualifier{}

	err = json.Unmarshal(bytes, &varIamQualifier)
	if err == nil {
		o.MoBaseMo = varIamQualifier.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Value")
		delete(additionalProperties, "Usergroup")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableIamQualifier struct {
	value *IamQualifier
	isSet bool
}

func (v NullableIamQualifier) Get() *IamQualifier {
	return v.value
}

func (v *NullableIamQualifier) Set(val *IamQualifier) {
	v.value = val
	v.isSet = true
}

func (v NullableIamQualifier) IsSet() bool {
	return v.isSet
}

func (v *NullableIamQualifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamQualifier(val *IamQualifier) *NullableIamQualifier {
	return &NullableIamQualifier{value: val, isSet: true}
}

func (v NullableIamQualifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamQualifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
