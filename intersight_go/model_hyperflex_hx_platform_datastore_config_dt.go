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

// HyperflexHxPlatformDatastoreConfigDt Platform datastore configuration.
type HyperflexHxPlatformDatastoreConfigDt struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Size of datablock for this HyperFlex datastore.
	DataBlockSize *int64 `json:"DataBlockSize,omitempty"`
	// Unique name for the datastore.
	Name *string `json:"Name,omitempty"`
	// Number of copies of data maintained for data redundancy.
	NumMirrors *int64 `json:"NumMirrors,omitempty"`
	// Number of stripes to be used for large files in datastore.
	NumStripesForLargeFiles *int64 `json:"NumStripesForLargeFiles,omitempty"`
	// Provisioned capacity of datastore in bytes.
	ProvisionedCapacity *int64 `json:"ProvisionedCapacity,omitempty"`
	// Specifies if this datastore is a system datastore or not.
	SystemDatastore *bool `json:"SystemDatastore,omitempty"`
	// HyperFlex datastore usage type. * `NFS` - The HyperFlex datastore type is NFS. * `SMB` - The HyperFlex datastore type is SMB. * `ISCSI` - The HyperFlex datastore type is ISCSI. * `UNKNOWN` - The HyperFlex datastore type is unknown.
	UsageType            *string `json:"UsageType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxPlatformDatastoreConfigDt HyperflexHxPlatformDatastoreConfigDt

// NewHyperflexHxPlatformDatastoreConfigDt instantiates a new HyperflexHxPlatformDatastoreConfigDt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxPlatformDatastoreConfigDt(classId string, objectType string) *HyperflexHxPlatformDatastoreConfigDt {
	this := HyperflexHxPlatformDatastoreConfigDt{}
	this.ClassId = classId
	this.ObjectType = objectType
	var usageType string = "NFS"
	this.UsageType = &usageType
	return &this
}

// NewHyperflexHxPlatformDatastoreConfigDtWithDefaults instantiates a new HyperflexHxPlatformDatastoreConfigDt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxPlatformDatastoreConfigDtWithDefaults() *HyperflexHxPlatformDatastoreConfigDt {
	this := HyperflexHxPlatformDatastoreConfigDt{}
	var classId string = "hyperflex.HxPlatformDatastoreConfigDt"
	this.ClassId = classId
	var objectType string = "hyperflex.HxPlatformDatastoreConfigDt"
	this.ObjectType = objectType
	var usageType string = "NFS"
	this.UsageType = &usageType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxPlatformDatastoreConfigDt) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxPlatformDatastoreConfigDt) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxPlatformDatastoreConfigDt) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxPlatformDatastoreConfigDt) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxPlatformDatastoreConfigDt) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxPlatformDatastoreConfigDt) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDataBlockSize returns the DataBlockSize field value if set, zero value otherwise.
func (o *HyperflexHxPlatformDatastoreConfigDt) GetDataBlockSize() int64 {
	if o == nil || o.DataBlockSize == nil {
		var ret int64
		return ret
	}
	return *o.DataBlockSize
}

// GetDataBlockSizeOk returns a tuple with the DataBlockSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxPlatformDatastoreConfigDt) GetDataBlockSizeOk() (*int64, bool) {
	if o == nil || o.DataBlockSize == nil {
		return nil, false
	}
	return o.DataBlockSize, true
}

// HasDataBlockSize returns a boolean if a field has been set.
func (o *HyperflexHxPlatformDatastoreConfigDt) HasDataBlockSize() bool {
	if o != nil && o.DataBlockSize != nil {
		return true
	}

	return false
}

// SetDataBlockSize gets a reference to the given int64 and assigns it to the DataBlockSize field.
func (o *HyperflexHxPlatformDatastoreConfigDt) SetDataBlockSize(v int64) {
	o.DataBlockSize = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexHxPlatformDatastoreConfigDt) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxPlatformDatastoreConfigDt) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexHxPlatformDatastoreConfigDt) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexHxPlatformDatastoreConfigDt) SetName(v string) {
	o.Name = &v
}

// GetNumMirrors returns the NumMirrors field value if set, zero value otherwise.
func (o *HyperflexHxPlatformDatastoreConfigDt) GetNumMirrors() int64 {
	if o == nil || o.NumMirrors == nil {
		var ret int64
		return ret
	}
	return *o.NumMirrors
}

// GetNumMirrorsOk returns a tuple with the NumMirrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxPlatformDatastoreConfigDt) GetNumMirrorsOk() (*int64, bool) {
	if o == nil || o.NumMirrors == nil {
		return nil, false
	}
	return o.NumMirrors, true
}

// HasNumMirrors returns a boolean if a field has been set.
func (o *HyperflexHxPlatformDatastoreConfigDt) HasNumMirrors() bool {
	if o != nil && o.NumMirrors != nil {
		return true
	}

	return false
}

// SetNumMirrors gets a reference to the given int64 and assigns it to the NumMirrors field.
func (o *HyperflexHxPlatformDatastoreConfigDt) SetNumMirrors(v int64) {
	o.NumMirrors = &v
}

// GetNumStripesForLargeFiles returns the NumStripesForLargeFiles field value if set, zero value otherwise.
func (o *HyperflexHxPlatformDatastoreConfigDt) GetNumStripesForLargeFiles() int64 {
	if o == nil || o.NumStripesForLargeFiles == nil {
		var ret int64
		return ret
	}
	return *o.NumStripesForLargeFiles
}

// GetNumStripesForLargeFilesOk returns a tuple with the NumStripesForLargeFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxPlatformDatastoreConfigDt) GetNumStripesForLargeFilesOk() (*int64, bool) {
	if o == nil || o.NumStripesForLargeFiles == nil {
		return nil, false
	}
	return o.NumStripesForLargeFiles, true
}

// HasNumStripesForLargeFiles returns a boolean if a field has been set.
func (o *HyperflexHxPlatformDatastoreConfigDt) HasNumStripesForLargeFiles() bool {
	if o != nil && o.NumStripesForLargeFiles != nil {
		return true
	}

	return false
}

// SetNumStripesForLargeFiles gets a reference to the given int64 and assigns it to the NumStripesForLargeFiles field.
func (o *HyperflexHxPlatformDatastoreConfigDt) SetNumStripesForLargeFiles(v int64) {
	o.NumStripesForLargeFiles = &v
}

// GetProvisionedCapacity returns the ProvisionedCapacity field value if set, zero value otherwise.
func (o *HyperflexHxPlatformDatastoreConfigDt) GetProvisionedCapacity() int64 {
	if o == nil || o.ProvisionedCapacity == nil {
		var ret int64
		return ret
	}
	return *o.ProvisionedCapacity
}

// GetProvisionedCapacityOk returns a tuple with the ProvisionedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxPlatformDatastoreConfigDt) GetProvisionedCapacityOk() (*int64, bool) {
	if o == nil || o.ProvisionedCapacity == nil {
		return nil, false
	}
	return o.ProvisionedCapacity, true
}

// HasProvisionedCapacity returns a boolean if a field has been set.
func (o *HyperflexHxPlatformDatastoreConfigDt) HasProvisionedCapacity() bool {
	if o != nil && o.ProvisionedCapacity != nil {
		return true
	}

	return false
}

// SetProvisionedCapacity gets a reference to the given int64 and assigns it to the ProvisionedCapacity field.
func (o *HyperflexHxPlatformDatastoreConfigDt) SetProvisionedCapacity(v int64) {
	o.ProvisionedCapacity = &v
}

// GetSystemDatastore returns the SystemDatastore field value if set, zero value otherwise.
func (o *HyperflexHxPlatformDatastoreConfigDt) GetSystemDatastore() bool {
	if o == nil || o.SystemDatastore == nil {
		var ret bool
		return ret
	}
	return *o.SystemDatastore
}

// GetSystemDatastoreOk returns a tuple with the SystemDatastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxPlatformDatastoreConfigDt) GetSystemDatastoreOk() (*bool, bool) {
	if o == nil || o.SystemDatastore == nil {
		return nil, false
	}
	return o.SystemDatastore, true
}

// HasSystemDatastore returns a boolean if a field has been set.
func (o *HyperflexHxPlatformDatastoreConfigDt) HasSystemDatastore() bool {
	if o != nil && o.SystemDatastore != nil {
		return true
	}

	return false
}

// SetSystemDatastore gets a reference to the given bool and assigns it to the SystemDatastore field.
func (o *HyperflexHxPlatformDatastoreConfigDt) SetSystemDatastore(v bool) {
	o.SystemDatastore = &v
}

// GetUsageType returns the UsageType field value if set, zero value otherwise.
func (o *HyperflexHxPlatformDatastoreConfigDt) GetUsageType() string {
	if o == nil || o.UsageType == nil {
		var ret string
		return ret
	}
	return *o.UsageType
}

// GetUsageTypeOk returns a tuple with the UsageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxPlatformDatastoreConfigDt) GetUsageTypeOk() (*string, bool) {
	if o == nil || o.UsageType == nil {
		return nil, false
	}
	return o.UsageType, true
}

// HasUsageType returns a boolean if a field has been set.
func (o *HyperflexHxPlatformDatastoreConfigDt) HasUsageType() bool {
	if o != nil && o.UsageType != nil {
		return true
	}

	return false
}

// SetUsageType gets a reference to the given string and assigns it to the UsageType field.
func (o *HyperflexHxPlatformDatastoreConfigDt) SetUsageType(v string) {
	o.UsageType = &v
}

func (o HyperflexHxPlatformDatastoreConfigDt) MarshalJSON() ([]byte, error) {
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
	if o.DataBlockSize != nil {
		toSerialize["DataBlockSize"] = o.DataBlockSize
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.NumMirrors != nil {
		toSerialize["NumMirrors"] = o.NumMirrors
	}
	if o.NumStripesForLargeFiles != nil {
		toSerialize["NumStripesForLargeFiles"] = o.NumStripesForLargeFiles
	}
	if o.ProvisionedCapacity != nil {
		toSerialize["ProvisionedCapacity"] = o.ProvisionedCapacity
	}
	if o.SystemDatastore != nil {
		toSerialize["SystemDatastore"] = o.SystemDatastore
	}
	if o.UsageType != nil {
		toSerialize["UsageType"] = o.UsageType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxPlatformDatastoreConfigDt) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexHxPlatformDatastoreConfigDtWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Size of datablock for this HyperFlex datastore.
		DataBlockSize *int64 `json:"DataBlockSize,omitempty"`
		// Unique name for the datastore.
		Name *string `json:"Name,omitempty"`
		// Number of copies of data maintained for data redundancy.
		NumMirrors *int64 `json:"NumMirrors,omitempty"`
		// Number of stripes to be used for large files in datastore.
		NumStripesForLargeFiles *int64 `json:"NumStripesForLargeFiles,omitempty"`
		// Provisioned capacity of datastore in bytes.
		ProvisionedCapacity *int64 `json:"ProvisionedCapacity,omitempty"`
		// Specifies if this datastore is a system datastore or not.
		SystemDatastore *bool `json:"SystemDatastore,omitempty"`
		// HyperFlex datastore usage type. * `NFS` - The HyperFlex datastore type is NFS. * `SMB` - The HyperFlex datastore type is SMB. * `ISCSI` - The HyperFlex datastore type is ISCSI. * `UNKNOWN` - The HyperFlex datastore type is unknown.
		UsageType *string `json:"UsageType,omitempty"`
	}

	varHyperflexHxPlatformDatastoreConfigDtWithoutEmbeddedStruct := HyperflexHxPlatformDatastoreConfigDtWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexHxPlatformDatastoreConfigDtWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexHxPlatformDatastoreConfigDt := _HyperflexHxPlatformDatastoreConfigDt{}
		varHyperflexHxPlatformDatastoreConfigDt.ClassId = varHyperflexHxPlatformDatastoreConfigDtWithoutEmbeddedStruct.ClassId
		varHyperflexHxPlatformDatastoreConfigDt.ObjectType = varHyperflexHxPlatformDatastoreConfigDtWithoutEmbeddedStruct.ObjectType
		varHyperflexHxPlatformDatastoreConfigDt.DataBlockSize = varHyperflexHxPlatformDatastoreConfigDtWithoutEmbeddedStruct.DataBlockSize
		varHyperflexHxPlatformDatastoreConfigDt.Name = varHyperflexHxPlatformDatastoreConfigDtWithoutEmbeddedStruct.Name
		varHyperflexHxPlatformDatastoreConfigDt.NumMirrors = varHyperflexHxPlatformDatastoreConfigDtWithoutEmbeddedStruct.NumMirrors
		varHyperflexHxPlatformDatastoreConfigDt.NumStripesForLargeFiles = varHyperflexHxPlatformDatastoreConfigDtWithoutEmbeddedStruct.NumStripesForLargeFiles
		varHyperflexHxPlatformDatastoreConfigDt.ProvisionedCapacity = varHyperflexHxPlatformDatastoreConfigDtWithoutEmbeddedStruct.ProvisionedCapacity
		varHyperflexHxPlatformDatastoreConfigDt.SystemDatastore = varHyperflexHxPlatformDatastoreConfigDtWithoutEmbeddedStruct.SystemDatastore
		varHyperflexHxPlatformDatastoreConfigDt.UsageType = varHyperflexHxPlatformDatastoreConfigDtWithoutEmbeddedStruct.UsageType
		*o = HyperflexHxPlatformDatastoreConfigDt(varHyperflexHxPlatformDatastoreConfigDt)
	} else {
		return err
	}

	varHyperflexHxPlatformDatastoreConfigDt := _HyperflexHxPlatformDatastoreConfigDt{}

	err = json.Unmarshal(bytes, &varHyperflexHxPlatformDatastoreConfigDt)
	if err == nil {
		o.MoBaseComplexType = varHyperflexHxPlatformDatastoreConfigDt.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DataBlockSize")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NumMirrors")
		delete(additionalProperties, "NumStripesForLargeFiles")
		delete(additionalProperties, "ProvisionedCapacity")
		delete(additionalProperties, "SystemDatastore")
		delete(additionalProperties, "UsageType")

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

type NullableHyperflexHxPlatformDatastoreConfigDt struct {
	value *HyperflexHxPlatformDatastoreConfigDt
	isSet bool
}

func (v NullableHyperflexHxPlatformDatastoreConfigDt) Get() *HyperflexHxPlatformDatastoreConfigDt {
	return v.value
}

func (v *NullableHyperflexHxPlatformDatastoreConfigDt) Set(val *HyperflexHxPlatformDatastoreConfigDt) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxPlatformDatastoreConfigDt) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxPlatformDatastoreConfigDt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxPlatformDatastoreConfigDt(val *HyperflexHxPlatformDatastoreConfigDt) *NullableHyperflexHxPlatformDatastoreConfigDt {
	return &NullableHyperflexHxPlatformDatastoreConfigDt{value: val, isSet: true}
}

func (v NullableHyperflexHxPlatformDatastoreConfigDt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxPlatformDatastoreConfigDt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
