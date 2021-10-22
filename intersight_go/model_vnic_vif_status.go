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

// VnicVifStatus The status of a vNIC or vHBA.
type VnicVifStatus struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the vNIC for which the status is reported.
	Name *string `json:"Name,omitempty"`
	// The reason for the status - it will be empty if status is ok or validating. If error, it will have the appropriate message indicating the reason for failure.
	Reason *string `json:"Reason,omitempty"`
	// Indicates if the vNIC / vHBA is ready for deploy or not. * `ok` - No issues with the LCP/SCP/VIF. * `error` - The LCP/SCP/VIF cannot be deployed due to error. * `validating` - Validation in progress for the LCP.
	Status               *string `json:"Status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicVifStatus VnicVifStatus

// NewVnicVifStatus instantiates a new VnicVifStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicVifStatus(classId string, objectType string) *VnicVifStatus {
	this := VnicVifStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "ok"
	this.Status = &status
	return &this
}

// NewVnicVifStatusWithDefaults instantiates a new VnicVifStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicVifStatusWithDefaults() *VnicVifStatus {
	this := VnicVifStatus{}
	var classId string = "vnic.VifStatus"
	this.ClassId = classId
	var objectType string = "vnic.VifStatus"
	this.ObjectType = objectType
	var status string = "ok"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicVifStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicVifStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicVifStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicVifStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicVifStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicVifStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VnicVifStatus) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVifStatus) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VnicVifStatus) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VnicVifStatus) SetName(v string) {
	o.Name = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *VnicVifStatus) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVifStatus) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *VnicVifStatus) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *VnicVifStatus) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VnicVifStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVifStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VnicVifStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VnicVifStatus) SetStatus(v string) {
	o.Status = &v
}

func (o VnicVifStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
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
	if o.Reason != nil {
		toSerialize["Reason"] = o.Reason
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicVifStatus) UnmarshalJSON(bytes []byte) (err error) {
	type VnicVifStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the vNIC for which the status is reported.
		Name *string `json:"Name,omitempty"`
		// The reason for the status - it will be empty if status is ok or validating. If error, it will have the appropriate message indicating the reason for failure.
		Reason *string `json:"Reason,omitempty"`
		// Indicates if the vNIC / vHBA is ready for deploy or not. * `ok` - No issues with the LCP/SCP/VIF. * `error` - The LCP/SCP/VIF cannot be deployed due to error. * `validating` - Validation in progress for the LCP.
		Status *string `json:"Status,omitempty"`
	}

	varVnicVifStatusWithoutEmbeddedStruct := VnicVifStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicVifStatusWithoutEmbeddedStruct)
	if err == nil {
		varVnicVifStatus := _VnicVifStatus{}
		varVnicVifStatus.ClassId = varVnicVifStatusWithoutEmbeddedStruct.ClassId
		varVnicVifStatus.ObjectType = varVnicVifStatusWithoutEmbeddedStruct.ObjectType
		varVnicVifStatus.Name = varVnicVifStatusWithoutEmbeddedStruct.Name
		varVnicVifStatus.Reason = varVnicVifStatusWithoutEmbeddedStruct.Reason
		varVnicVifStatus.Status = varVnicVifStatusWithoutEmbeddedStruct.Status
		*o = VnicVifStatus(varVnicVifStatus)
	} else {
		return err
	}

	varVnicVifStatus := _VnicVifStatus{}

	err = json.Unmarshal(bytes, &varVnicVifStatus)
	if err == nil {
		o.MoBaseComplexType = varVnicVifStatus.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Reason")
		delete(additionalProperties, "Status")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableVnicVifStatus struct {
	value *VnicVifStatus
	isSet bool
}

func (v NullableVnicVifStatus) Get() *VnicVifStatus {
	return v.value
}

func (v *NullableVnicVifStatus) Set(val *VnicVifStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicVifStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicVifStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicVifStatus(val *VnicVifStatus) *NullableVnicVifStatus {
	return &NullableVnicVifStatus{value: val, isSet: true}
}

func (v NullableVnicVifStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicVifStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
