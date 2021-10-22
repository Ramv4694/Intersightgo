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

// StorageVirtualDriveConfigurationAllOf Definition of the list of properties defined in 'storage.VirtualDriveConfiguration', excluding properties defined in parent classes.
type StorageVirtualDriveConfigurationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This flag enables this virtual drive to be used as a boot drive.
	BootDrive *bool `json:"BootDrive,omitempty"`
	// This flag enables the virtual drive to use all the space available in the disk group. When this flag is enabled, the size property is ignored.
	ExpandToAvailable *bool `json:"ExpandToAvailable,omitempty"`
	// The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed.
	Name *string `json:"Name,omitempty"`
	// Virtual drive size in MebiBytes. Size is mandatory field except when the Expand to Available option is enabled.
	Size                 *int64                            `json:"Size,omitempty"`
	VirtualDrivePolicy   NullableStorageVirtualDrivePolicy `json:"VirtualDrivePolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageVirtualDriveConfigurationAllOf StorageVirtualDriveConfigurationAllOf

// NewStorageVirtualDriveConfigurationAllOf instantiates a new StorageVirtualDriveConfigurationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageVirtualDriveConfigurationAllOf(classId string, objectType string) *StorageVirtualDriveConfigurationAllOf {
	this := StorageVirtualDriveConfigurationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageVirtualDriveConfigurationAllOfWithDefaults instantiates a new StorageVirtualDriveConfigurationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageVirtualDriveConfigurationAllOfWithDefaults() *StorageVirtualDriveConfigurationAllOf {
	this := StorageVirtualDriveConfigurationAllOf{}
	var classId string = "storage.VirtualDriveConfiguration"
	this.ClassId = classId
	var objectType string = "storage.VirtualDriveConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageVirtualDriveConfigurationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfigurationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageVirtualDriveConfigurationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageVirtualDriveConfigurationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfigurationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageVirtualDriveConfigurationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBootDrive returns the BootDrive field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfigurationAllOf) GetBootDrive() bool {
	if o == nil || o.BootDrive == nil {
		var ret bool
		return ret
	}
	return *o.BootDrive
}

// GetBootDriveOk returns a tuple with the BootDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfigurationAllOf) GetBootDriveOk() (*bool, bool) {
	if o == nil || o.BootDrive == nil {
		return nil, false
	}
	return o.BootDrive, true
}

// HasBootDrive returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfigurationAllOf) HasBootDrive() bool {
	if o != nil && o.BootDrive != nil {
		return true
	}

	return false
}

// SetBootDrive gets a reference to the given bool and assigns it to the BootDrive field.
func (o *StorageVirtualDriveConfigurationAllOf) SetBootDrive(v bool) {
	o.BootDrive = &v
}

// GetExpandToAvailable returns the ExpandToAvailable field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfigurationAllOf) GetExpandToAvailable() bool {
	if o == nil || o.ExpandToAvailable == nil {
		var ret bool
		return ret
	}
	return *o.ExpandToAvailable
}

// GetExpandToAvailableOk returns a tuple with the ExpandToAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfigurationAllOf) GetExpandToAvailableOk() (*bool, bool) {
	if o == nil || o.ExpandToAvailable == nil {
		return nil, false
	}
	return o.ExpandToAvailable, true
}

// HasExpandToAvailable returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfigurationAllOf) HasExpandToAvailable() bool {
	if o != nil && o.ExpandToAvailable != nil {
		return true
	}

	return false
}

// SetExpandToAvailable gets a reference to the given bool and assigns it to the ExpandToAvailable field.
func (o *StorageVirtualDriveConfigurationAllOf) SetExpandToAvailable(v bool) {
	o.ExpandToAvailable = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfigurationAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfigurationAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfigurationAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageVirtualDriveConfigurationAllOf) SetName(v string) {
	o.Name = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfigurationAllOf) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfigurationAllOf) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfigurationAllOf) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *StorageVirtualDriveConfigurationAllOf) SetSize(v int64) {
	o.Size = &v
}

// GetVirtualDrivePolicy returns the VirtualDrivePolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageVirtualDriveConfigurationAllOf) GetVirtualDrivePolicy() StorageVirtualDrivePolicy {
	if o == nil || o.VirtualDrivePolicy.Get() == nil {
		var ret StorageVirtualDrivePolicy
		return ret
	}
	return *o.VirtualDrivePolicy.Get()
}

// GetVirtualDrivePolicyOk returns a tuple with the VirtualDrivePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageVirtualDriveConfigurationAllOf) GetVirtualDrivePolicyOk() (*StorageVirtualDrivePolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualDrivePolicy.Get(), o.VirtualDrivePolicy.IsSet()
}

// HasVirtualDrivePolicy returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfigurationAllOf) HasVirtualDrivePolicy() bool {
	if o != nil && o.VirtualDrivePolicy.IsSet() {
		return true
	}

	return false
}

// SetVirtualDrivePolicy gets a reference to the given NullableStorageVirtualDrivePolicy and assigns it to the VirtualDrivePolicy field.
func (o *StorageVirtualDriveConfigurationAllOf) SetVirtualDrivePolicy(v StorageVirtualDrivePolicy) {
	o.VirtualDrivePolicy.Set(&v)
}

// SetVirtualDrivePolicyNil sets the value for VirtualDrivePolicy to be an explicit nil
func (o *StorageVirtualDriveConfigurationAllOf) SetVirtualDrivePolicyNil() {
	o.VirtualDrivePolicy.Set(nil)
}

// UnsetVirtualDrivePolicy ensures that no value is present for VirtualDrivePolicy, not even an explicit nil
func (o *StorageVirtualDriveConfigurationAllOf) UnsetVirtualDrivePolicy() {
	o.VirtualDrivePolicy.Unset()
}

func (o StorageVirtualDriveConfigurationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BootDrive != nil {
		toSerialize["BootDrive"] = o.BootDrive
	}
	if o.ExpandToAvailable != nil {
		toSerialize["ExpandToAvailable"] = o.ExpandToAvailable
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.VirtualDrivePolicy.IsSet() {
		toSerialize["VirtualDrivePolicy"] = o.VirtualDrivePolicy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageVirtualDriveConfigurationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageVirtualDriveConfigurationAllOf := _StorageVirtualDriveConfigurationAllOf{}

	if err = json.Unmarshal(bytes, &varStorageVirtualDriveConfigurationAllOf); err == nil {
		*o = StorageVirtualDriveConfigurationAllOf(varStorageVirtualDriveConfigurationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BootDrive")
		delete(additionalProperties, "ExpandToAvailable")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "VirtualDrivePolicy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageVirtualDriveConfigurationAllOf struct {
	value *StorageVirtualDriveConfigurationAllOf
	isSet bool
}

func (v NullableStorageVirtualDriveConfigurationAllOf) Get() *StorageVirtualDriveConfigurationAllOf {
	return v.value
}

func (v *NullableStorageVirtualDriveConfigurationAllOf) Set(val *StorageVirtualDriveConfigurationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageVirtualDriveConfigurationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageVirtualDriveConfigurationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageVirtualDriveConfigurationAllOf(val *StorageVirtualDriveConfigurationAllOf) *NullableStorageVirtualDriveConfigurationAllOf {
	return &NullableStorageVirtualDriveConfigurationAllOf{value: val, isSet: true}
}

func (v NullableStorageVirtualDriveConfigurationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageVirtualDriveConfigurationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
