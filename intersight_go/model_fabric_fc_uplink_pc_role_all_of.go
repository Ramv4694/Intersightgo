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
)

// FabricFcUplinkPcRoleAllOf Definition of the list of properties defined in 'fabric.FcUplinkPcRole', excluding properties defined in parent classes.
type FabricFcUplinkPcRoleAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin configured speed for the port. * `Auto` - Admin configurable speed AUTO ( default ). * `8Gbps` - Admin configurable speed 8Gbps. * `16Gbps` - Admin configurable speed 16Gbps. * `32Gbps` - Admin configurable speed 32Gbps.
	AdminSpeed *string `json:"AdminSpeed,omitempty"`
	// Fill pattern to differentiate the configs in NPIV. * `Idle` - Fc Fill Pattern type Idle. * `Arbff` - Fc Fill Pattern type Arbff.
	FillPattern *string `json:"FillPattern,omitempty"`
	// Virtual San Identifier associated to the FC port.
	VsanId               *int64 `json:"VsanId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricFcUplinkPcRoleAllOf FabricFcUplinkPcRoleAllOf

// NewFabricFcUplinkPcRoleAllOf instantiates a new FabricFcUplinkPcRoleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricFcUplinkPcRoleAllOf(classId string, objectType string) *FabricFcUplinkPcRoleAllOf {
	this := FabricFcUplinkPcRoleAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	var fillPattern string = "Idle"
	this.FillPattern = &fillPattern
	return &this
}

// NewFabricFcUplinkPcRoleAllOfWithDefaults instantiates a new FabricFcUplinkPcRoleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricFcUplinkPcRoleAllOfWithDefaults() *FabricFcUplinkPcRoleAllOf {
	this := FabricFcUplinkPcRoleAllOf{}
	var classId string = "fabric.FcUplinkPcRole"
	this.ClassId = classId
	var objectType string = "fabric.FcUplinkPcRole"
	this.ObjectType = objectType
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	var fillPattern string = "Idle"
	this.FillPattern = &fillPattern
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricFcUplinkPcRoleAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricFcUplinkPcRoleAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricFcUplinkPcRoleAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricFcUplinkPcRoleAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricFcUplinkPcRoleAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricFcUplinkPcRoleAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminSpeed returns the AdminSpeed field value if set, zero value otherwise.
func (o *FabricFcUplinkPcRoleAllOf) GetAdminSpeed() string {
	if o == nil || o.AdminSpeed == nil {
		var ret string
		return ret
	}
	return *o.AdminSpeed
}

// GetAdminSpeedOk returns a tuple with the AdminSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcUplinkPcRoleAllOf) GetAdminSpeedOk() (*string, bool) {
	if o == nil || o.AdminSpeed == nil {
		return nil, false
	}
	return o.AdminSpeed, true
}

// HasAdminSpeed returns a boolean if a field has been set.
func (o *FabricFcUplinkPcRoleAllOf) HasAdminSpeed() bool {
	if o != nil && o.AdminSpeed != nil {
		return true
	}

	return false
}

// SetAdminSpeed gets a reference to the given string and assigns it to the AdminSpeed field.
func (o *FabricFcUplinkPcRoleAllOf) SetAdminSpeed(v string) {
	o.AdminSpeed = &v
}

// GetFillPattern returns the FillPattern field value if set, zero value otherwise.
func (o *FabricFcUplinkPcRoleAllOf) GetFillPattern() string {
	if o == nil || o.FillPattern == nil {
		var ret string
		return ret
	}
	return *o.FillPattern
}

// GetFillPatternOk returns a tuple with the FillPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcUplinkPcRoleAllOf) GetFillPatternOk() (*string, bool) {
	if o == nil || o.FillPattern == nil {
		return nil, false
	}
	return o.FillPattern, true
}

// HasFillPattern returns a boolean if a field has been set.
func (o *FabricFcUplinkPcRoleAllOf) HasFillPattern() bool {
	if o != nil && o.FillPattern != nil {
		return true
	}

	return false
}

// SetFillPattern gets a reference to the given string and assigns it to the FillPattern field.
func (o *FabricFcUplinkPcRoleAllOf) SetFillPattern(v string) {
	o.FillPattern = &v
}

// GetVsanId returns the VsanId field value if set, zero value otherwise.
func (o *FabricFcUplinkPcRoleAllOf) GetVsanId() int64 {
	if o == nil || o.VsanId == nil {
		var ret int64
		return ret
	}
	return *o.VsanId
}

// GetVsanIdOk returns a tuple with the VsanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcUplinkPcRoleAllOf) GetVsanIdOk() (*int64, bool) {
	if o == nil || o.VsanId == nil {
		return nil, false
	}
	return o.VsanId, true
}

// HasVsanId returns a boolean if a field has been set.
func (o *FabricFcUplinkPcRoleAllOf) HasVsanId() bool {
	if o != nil && o.VsanId != nil {
		return true
	}

	return false
}

// SetVsanId gets a reference to the given int64 and assigns it to the VsanId field.
func (o *FabricFcUplinkPcRoleAllOf) SetVsanId(v int64) {
	o.VsanId = &v
}

func (o FabricFcUplinkPcRoleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminSpeed != nil {
		toSerialize["AdminSpeed"] = o.AdminSpeed
	}
	if o.FillPattern != nil {
		toSerialize["FillPattern"] = o.FillPattern
	}
	if o.VsanId != nil {
		toSerialize["VsanId"] = o.VsanId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricFcUplinkPcRoleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricFcUplinkPcRoleAllOf := _FabricFcUplinkPcRoleAllOf{}

	if err = json.Unmarshal(bytes, &varFabricFcUplinkPcRoleAllOf); err == nil {
		*o = FabricFcUplinkPcRoleAllOf(varFabricFcUplinkPcRoleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminSpeed")
		delete(additionalProperties, "FillPattern")
		delete(additionalProperties, "VsanId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricFcUplinkPcRoleAllOf struct {
	value *FabricFcUplinkPcRoleAllOf
	isSet bool
}

func (v NullableFabricFcUplinkPcRoleAllOf) Get() *FabricFcUplinkPcRoleAllOf {
	return v.value
}

func (v *NullableFabricFcUplinkPcRoleAllOf) Set(val *FabricFcUplinkPcRoleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricFcUplinkPcRoleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricFcUplinkPcRoleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricFcUplinkPcRoleAllOf(val *FabricFcUplinkPcRoleAllOf) *NullableFabricFcUplinkPcRoleAllOf {
	return &NullableFabricFcUplinkPcRoleAllOf{value: val, isSet: true}
}

func (v NullableFabricFcUplinkPcRoleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricFcUplinkPcRoleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
