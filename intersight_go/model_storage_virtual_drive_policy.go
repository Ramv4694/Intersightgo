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

// StorageVirtualDrivePolicy This models the manual drive selection configuration.
type StorageVirtualDrivePolicy struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Access policy that host has on this virtual drive. * `Default` - Use platform default access mode. * `ReadWrite` - Enables host to perform read-write on the VD. * `ReadOnly` - Host can only read from the VD. * `Blocked` - Host can neither read nor write to the VD.
	AccessPolicy *string `json:"AccessPolicy,omitempty"`
	// Disk cache policy for the virtual drive. * `Default` - Use platform default drive cache mode. * `NoChange` - Drive cache policy is unchanged. * `Enable` - Enables IO caching on the drive. * `Disable` - Disables IO caching on the drive.
	DriveCache *string `json:"DriveCache,omitempty"`
	// Read ahead mode to be used to read data from this virtual drive. * `Default` - Use platform default read ahead mode. * `ReadAhead` - Use read ahead mode for the policy. * `NoReadAhead` - Do not use read ahead mode for the policy.
	ReadPolicy *string `json:"ReadPolicy,omitempty"`
	// Desired strip size - Allowed values are 64KiB, 128KiB, 256KiB, 512KiB, 1024KiB. * `64` - Number of bytes in a strip is 64 Kibibytes. * `128` - Number of bytes in a strip is 128 Kibibytes. * `256` - Number of bytes in a strip is 256 Kibibytes. * `512` - Number of bytes in a strip is 512 Kibibytes. * `1024` - Number of bytes in a strip is 1024 Kibibytes or 1 Mebibyte.
	StripSize *int32 `json:"StripSize,omitempty"`
	// Write mode to be used to write data to this virtual drive. * `Default` - Use platform default write mode. * `WriteThrough` - Data is written through the cache and to the physical drives. Performance is improved, because subsequent reads of that data can be satisfied from the cache. * `WriteBackGoodBbu` - Data is stored in the cache, and is only written to the physical drives when space in the cache is needed. Virtual drives requesting this policy fall back to Write Through caching when the battery backup unit (BBU) cannot guarantee the safety of the cache in the event of a power failure. * `AlwaysWriteBack` - With this policy, write caching remains Write Back even if the battery backup unit is defective or discharged.
	WritePolicy          *string `json:"WritePolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageVirtualDrivePolicy StorageVirtualDrivePolicy

// NewStorageVirtualDrivePolicy instantiates a new StorageVirtualDrivePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageVirtualDrivePolicy(classId string, objectType string) *StorageVirtualDrivePolicy {
	this := StorageVirtualDrivePolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var accessPolicy string = "Default"
	this.AccessPolicy = &accessPolicy
	var driveCache string = "Default"
	this.DriveCache = &driveCache
	var readPolicy string = "Default"
	this.ReadPolicy = &readPolicy
	var stripSize int32 = 64
	this.StripSize = &stripSize
	var writePolicy string = "Default"
	this.WritePolicy = &writePolicy
	return &this
}

// NewStorageVirtualDrivePolicyWithDefaults instantiates a new StorageVirtualDrivePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageVirtualDrivePolicyWithDefaults() *StorageVirtualDrivePolicy {
	this := StorageVirtualDrivePolicy{}
	var classId string = "storage.VirtualDrivePolicy"
	this.ClassId = classId
	var objectType string = "storage.VirtualDrivePolicy"
	this.ObjectType = objectType
	var accessPolicy string = "Default"
	this.AccessPolicy = &accessPolicy
	var driveCache string = "Default"
	this.DriveCache = &driveCache
	var readPolicy string = "Default"
	this.ReadPolicy = &readPolicy
	var stripSize int32 = 64
	this.StripSize = &stripSize
	var writePolicy string = "Default"
	this.WritePolicy = &writePolicy
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageVirtualDrivePolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDrivePolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageVirtualDrivePolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageVirtualDrivePolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDrivePolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageVirtualDrivePolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccessPolicy returns the AccessPolicy field value if set, zero value otherwise.
func (o *StorageVirtualDrivePolicy) GetAccessPolicy() string {
	if o == nil || o.AccessPolicy == nil {
		var ret string
		return ret
	}
	return *o.AccessPolicy
}

// GetAccessPolicyOk returns a tuple with the AccessPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDrivePolicy) GetAccessPolicyOk() (*string, bool) {
	if o == nil || o.AccessPolicy == nil {
		return nil, false
	}
	return o.AccessPolicy, true
}

// HasAccessPolicy returns a boolean if a field has been set.
func (o *StorageVirtualDrivePolicy) HasAccessPolicy() bool {
	if o != nil && o.AccessPolicy != nil {
		return true
	}

	return false
}

// SetAccessPolicy gets a reference to the given string and assigns it to the AccessPolicy field.
func (o *StorageVirtualDrivePolicy) SetAccessPolicy(v string) {
	o.AccessPolicy = &v
}

// GetDriveCache returns the DriveCache field value if set, zero value otherwise.
func (o *StorageVirtualDrivePolicy) GetDriveCache() string {
	if o == nil || o.DriveCache == nil {
		var ret string
		return ret
	}
	return *o.DriveCache
}

// GetDriveCacheOk returns a tuple with the DriveCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDrivePolicy) GetDriveCacheOk() (*string, bool) {
	if o == nil || o.DriveCache == nil {
		return nil, false
	}
	return o.DriveCache, true
}

// HasDriveCache returns a boolean if a field has been set.
func (o *StorageVirtualDrivePolicy) HasDriveCache() bool {
	if o != nil && o.DriveCache != nil {
		return true
	}

	return false
}

// SetDriveCache gets a reference to the given string and assigns it to the DriveCache field.
func (o *StorageVirtualDrivePolicy) SetDriveCache(v string) {
	o.DriveCache = &v
}

// GetReadPolicy returns the ReadPolicy field value if set, zero value otherwise.
func (o *StorageVirtualDrivePolicy) GetReadPolicy() string {
	if o == nil || o.ReadPolicy == nil {
		var ret string
		return ret
	}
	return *o.ReadPolicy
}

// GetReadPolicyOk returns a tuple with the ReadPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDrivePolicy) GetReadPolicyOk() (*string, bool) {
	if o == nil || o.ReadPolicy == nil {
		return nil, false
	}
	return o.ReadPolicy, true
}

// HasReadPolicy returns a boolean if a field has been set.
func (o *StorageVirtualDrivePolicy) HasReadPolicy() bool {
	if o != nil && o.ReadPolicy != nil {
		return true
	}

	return false
}

// SetReadPolicy gets a reference to the given string and assigns it to the ReadPolicy field.
func (o *StorageVirtualDrivePolicy) SetReadPolicy(v string) {
	o.ReadPolicy = &v
}

// GetStripSize returns the StripSize field value if set, zero value otherwise.
func (o *StorageVirtualDrivePolicy) GetStripSize() int32 {
	if o == nil || o.StripSize == nil {
		var ret int32
		return ret
	}
	return *o.StripSize
}

// GetStripSizeOk returns a tuple with the StripSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDrivePolicy) GetStripSizeOk() (*int32, bool) {
	if o == nil || o.StripSize == nil {
		return nil, false
	}
	return o.StripSize, true
}

// HasStripSize returns a boolean if a field has been set.
func (o *StorageVirtualDrivePolicy) HasStripSize() bool {
	if o != nil && o.StripSize != nil {
		return true
	}

	return false
}

// SetStripSize gets a reference to the given int32 and assigns it to the StripSize field.
func (o *StorageVirtualDrivePolicy) SetStripSize(v int32) {
	o.StripSize = &v
}

// GetWritePolicy returns the WritePolicy field value if set, zero value otherwise.
func (o *StorageVirtualDrivePolicy) GetWritePolicy() string {
	if o == nil || o.WritePolicy == nil {
		var ret string
		return ret
	}
	return *o.WritePolicy
}

// GetWritePolicyOk returns a tuple with the WritePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDrivePolicy) GetWritePolicyOk() (*string, bool) {
	if o == nil || o.WritePolicy == nil {
		return nil, false
	}
	return o.WritePolicy, true
}

// HasWritePolicy returns a boolean if a field has been set.
func (o *StorageVirtualDrivePolicy) HasWritePolicy() bool {
	if o != nil && o.WritePolicy != nil {
		return true
	}

	return false
}

// SetWritePolicy gets a reference to the given string and assigns it to the WritePolicy field.
func (o *StorageVirtualDrivePolicy) SetWritePolicy(v string) {
	o.WritePolicy = &v
}

func (o StorageVirtualDrivePolicy) MarshalJSON() ([]byte, error) {
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
	if o.AccessPolicy != nil {
		toSerialize["AccessPolicy"] = o.AccessPolicy
	}
	if o.DriveCache != nil {
		toSerialize["DriveCache"] = o.DriveCache
	}
	if o.ReadPolicy != nil {
		toSerialize["ReadPolicy"] = o.ReadPolicy
	}
	if o.StripSize != nil {
		toSerialize["StripSize"] = o.StripSize
	}
	if o.WritePolicy != nil {
		toSerialize["WritePolicy"] = o.WritePolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageVirtualDrivePolicy) UnmarshalJSON(bytes []byte) (err error) {
	type StorageVirtualDrivePolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Access policy that host has on this virtual drive. * `Default` - Use platform default access mode. * `ReadWrite` - Enables host to perform read-write on the VD. * `ReadOnly` - Host can only read from the VD. * `Blocked` - Host can neither read nor write to the VD.
		AccessPolicy *string `json:"AccessPolicy,omitempty"`
		// Disk cache policy for the virtual drive. * `Default` - Use platform default drive cache mode. * `NoChange` - Drive cache policy is unchanged. * `Enable` - Enables IO caching on the drive. * `Disable` - Disables IO caching on the drive.
		DriveCache *string `json:"DriveCache,omitempty"`
		// Read ahead mode to be used to read data from this virtual drive. * `Default` - Use platform default read ahead mode. * `ReadAhead` - Use read ahead mode for the policy. * `NoReadAhead` - Do not use read ahead mode for the policy.
		ReadPolicy *string `json:"ReadPolicy,omitempty"`
		// Desired strip size - Allowed values are 64KiB, 128KiB, 256KiB, 512KiB, 1024KiB. * `64` - Number of bytes in a strip is 64 Kibibytes. * `128` - Number of bytes in a strip is 128 Kibibytes. * `256` - Number of bytes in a strip is 256 Kibibytes. * `512` - Number of bytes in a strip is 512 Kibibytes. * `1024` - Number of bytes in a strip is 1024 Kibibytes or 1 Mebibyte.
		StripSize *int32 `json:"StripSize,omitempty"`
		// Write mode to be used to write data to this virtual drive. * `Default` - Use platform default write mode. * `WriteThrough` - Data is written through the cache and to the physical drives. Performance is improved, because subsequent reads of that data can be satisfied from the cache. * `WriteBackGoodBbu` - Data is stored in the cache, and is only written to the physical drives when space in the cache is needed. Virtual drives requesting this policy fall back to Write Through caching when the battery backup unit (BBU) cannot guarantee the safety of the cache in the event of a power failure. * `AlwaysWriteBack` - With this policy, write caching remains Write Back even if the battery backup unit is defective or discharged.
		WritePolicy *string `json:"WritePolicy,omitempty"`
	}

	varStorageVirtualDrivePolicyWithoutEmbeddedStruct := StorageVirtualDrivePolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageVirtualDrivePolicyWithoutEmbeddedStruct)
	if err == nil {
		varStorageVirtualDrivePolicy := _StorageVirtualDrivePolicy{}
		varStorageVirtualDrivePolicy.ClassId = varStorageVirtualDrivePolicyWithoutEmbeddedStruct.ClassId
		varStorageVirtualDrivePolicy.ObjectType = varStorageVirtualDrivePolicyWithoutEmbeddedStruct.ObjectType
		varStorageVirtualDrivePolicy.AccessPolicy = varStorageVirtualDrivePolicyWithoutEmbeddedStruct.AccessPolicy
		varStorageVirtualDrivePolicy.DriveCache = varStorageVirtualDrivePolicyWithoutEmbeddedStruct.DriveCache
		varStorageVirtualDrivePolicy.ReadPolicy = varStorageVirtualDrivePolicyWithoutEmbeddedStruct.ReadPolicy
		varStorageVirtualDrivePolicy.StripSize = varStorageVirtualDrivePolicyWithoutEmbeddedStruct.StripSize
		varStorageVirtualDrivePolicy.WritePolicy = varStorageVirtualDrivePolicyWithoutEmbeddedStruct.WritePolicy
		*o = StorageVirtualDrivePolicy(varStorageVirtualDrivePolicy)
	} else {
		return err
	}

	varStorageVirtualDrivePolicy := _StorageVirtualDrivePolicy{}

	err = json.Unmarshal(bytes, &varStorageVirtualDrivePolicy)
	if err == nil {
		o.MoBaseComplexType = varStorageVirtualDrivePolicy.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccessPolicy")
		delete(additionalProperties, "DriveCache")
		delete(additionalProperties, "ReadPolicy")
		delete(additionalProperties, "StripSize")
		delete(additionalProperties, "WritePolicy")

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

type NullableStorageVirtualDrivePolicy struct {
	value *StorageVirtualDrivePolicy
	isSet bool
}

func (v NullableStorageVirtualDrivePolicy) Get() *StorageVirtualDrivePolicy {
	return v.value
}

func (v *NullableStorageVirtualDrivePolicy) Set(val *StorageVirtualDrivePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageVirtualDrivePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageVirtualDrivePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageVirtualDrivePolicy(val *StorageVirtualDrivePolicy) *NullableStorageVirtualDrivePolicy {
	return &NullableStorageVirtualDrivePolicy{value: val, isSet: true}
}

func (v NullableStorageVirtualDrivePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageVirtualDrivePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
