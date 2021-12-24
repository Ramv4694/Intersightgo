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
	"reflect"
	"strings"
)

// CloudBaseBillingUnit A base billing unit definition that is extended by all concrete billingUnit objects.
type CloudBaseBillingUnit struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The ID of the paying account.
	Identity *string `json:"Identity,omitempty"`
	// The name of the paying account.
	Name *string `json:"Name,omitempty"`
	// Status of the account, for example Active, Suspended, etc.
	Status               *string `json:"Status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudBaseBillingUnit CloudBaseBillingUnit

// NewCloudBaseBillingUnit instantiates a new CloudBaseBillingUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBaseBillingUnit(classId string, objectType string) *CloudBaseBillingUnit {
	this := CloudBaseBillingUnit{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudBaseBillingUnitWithDefaults instantiates a new CloudBaseBillingUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBaseBillingUnitWithDefaults() *CloudBaseBillingUnit {
	this := CloudBaseBillingUnit{}
	var classId string = "cloud.AwsBillingUnit"
	this.ClassId = classId
	var objectType string = "cloud.AwsBillingUnit"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudBaseBillingUnit) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudBaseBillingUnit) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudBaseBillingUnit) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudBaseBillingUnit) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudBaseBillingUnit) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudBaseBillingUnit) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *CloudBaseBillingUnit) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseBillingUnit) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *CloudBaseBillingUnit) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *CloudBaseBillingUnit) SetIdentity(v string) {
	o.Identity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudBaseBillingUnit) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseBillingUnit) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudBaseBillingUnit) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudBaseBillingUnit) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CloudBaseBillingUnit) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseBillingUnit) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CloudBaseBillingUnit) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CloudBaseBillingUnit) SetStatus(v string) {
	o.Status = &v
}

func (o CloudBaseBillingUnit) MarshalJSON() ([]byte, error) {
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
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudBaseBillingUnit) UnmarshalJSON(bytes []byte) (err error) {
	type CloudBaseBillingUnitWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The ID of the paying account.
		Identity *string `json:"Identity,omitempty"`
		// The name of the paying account.
		Name *string `json:"Name,omitempty"`
		// Status of the account, for example Active, Suspended, etc.
		Status *string `json:"Status,omitempty"`
	}

	varCloudBaseBillingUnitWithoutEmbeddedStruct := CloudBaseBillingUnitWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudBaseBillingUnitWithoutEmbeddedStruct)
	if err == nil {
		varCloudBaseBillingUnit := _CloudBaseBillingUnit{}
		varCloudBaseBillingUnit.ClassId = varCloudBaseBillingUnitWithoutEmbeddedStruct.ClassId
		varCloudBaseBillingUnit.ObjectType = varCloudBaseBillingUnitWithoutEmbeddedStruct.ObjectType
		varCloudBaseBillingUnit.Identity = varCloudBaseBillingUnitWithoutEmbeddedStruct.Identity
		varCloudBaseBillingUnit.Name = varCloudBaseBillingUnitWithoutEmbeddedStruct.Name
		varCloudBaseBillingUnit.Status = varCloudBaseBillingUnitWithoutEmbeddedStruct.Status
		*o = CloudBaseBillingUnit(varCloudBaseBillingUnit)
	} else {
		return err
	}

	varCloudBaseBillingUnit := _CloudBaseBillingUnit{}

	err = json.Unmarshal(bytes, &varCloudBaseBillingUnit)
	if err == nil {
		o.MoBaseMo = varCloudBaseBillingUnit.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Status")

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

type NullableCloudBaseBillingUnit struct {
	value *CloudBaseBillingUnit
	isSet bool
}

func (v NullableCloudBaseBillingUnit) Get() *CloudBaseBillingUnit {
	return v.value
}

func (v *NullableCloudBaseBillingUnit) Set(val *CloudBaseBillingUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudBaseBillingUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudBaseBillingUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudBaseBillingUnit(val *CloudBaseBillingUnit) *NullableCloudBaseBillingUnit {
	return &NullableCloudBaseBillingUnit{value: val, isSet: true}
}

func (v NullableCloudBaseBillingUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudBaseBillingUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
