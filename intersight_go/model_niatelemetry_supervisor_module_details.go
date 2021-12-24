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

// NiatelemetrySupervisorModuleDetails Supervisor module slot details in APIC.
type NiatelemetrySupervisorModuleDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Dn of the supervisor module in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Hardware version of supervisor module.
	HwVer *string `json:"HwVer,omitempty"`
	// Model of the supervisor module.
	Model *string `json:"Model,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Serial number of the supervisor module.
	Serial *string `json:"Serial,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                              `json:"SiteName,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetrySupervisorModuleDetails NiatelemetrySupervisorModuleDetails

// NewNiatelemetrySupervisorModuleDetails instantiates a new NiatelemetrySupervisorModuleDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetrySupervisorModuleDetails(classId string, objectType string) *NiatelemetrySupervisorModuleDetails {
	this := NiatelemetrySupervisorModuleDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetrySupervisorModuleDetailsWithDefaults instantiates a new NiatelemetrySupervisorModuleDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetrySupervisorModuleDetailsWithDefaults() *NiatelemetrySupervisorModuleDetails {
	this := NiatelemetrySupervisorModuleDetails{}
	var classId string = "niatelemetry.SupervisorModuleDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.SupervisorModuleDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetrySupervisorModuleDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySupervisorModuleDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetrySupervisorModuleDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetrySupervisorModuleDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySupervisorModuleDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetrySupervisorModuleDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetrySupervisorModuleDetails) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySupervisorModuleDetails) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetrySupervisorModuleDetails) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetrySupervisorModuleDetails) SetDn(v string) {
	o.Dn = &v
}

// GetHwVer returns the HwVer field value if set, zero value otherwise.
func (o *NiatelemetrySupervisorModuleDetails) GetHwVer() string {
	if o == nil || o.HwVer == nil {
		var ret string
		return ret
	}
	return *o.HwVer
}

// GetHwVerOk returns a tuple with the HwVer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySupervisorModuleDetails) GetHwVerOk() (*string, bool) {
	if o == nil || o.HwVer == nil {
		return nil, false
	}
	return o.HwVer, true
}

// HasHwVer returns a boolean if a field has been set.
func (o *NiatelemetrySupervisorModuleDetails) HasHwVer() bool {
	if o != nil && o.HwVer != nil {
		return true
	}

	return false
}

// SetHwVer gets a reference to the given string and assigns it to the HwVer field.
func (o *NiatelemetrySupervisorModuleDetails) SetHwVer(v string) {
	o.HwVer = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *NiatelemetrySupervisorModuleDetails) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySupervisorModuleDetails) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *NiatelemetrySupervisorModuleDetails) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *NiatelemetrySupervisorModuleDetails) SetModel(v string) {
	o.Model = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetrySupervisorModuleDetails) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySupervisorModuleDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetrySupervisorModuleDetails) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetrySupervisorModuleDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetrySupervisorModuleDetails) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySupervisorModuleDetails) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetrySupervisorModuleDetails) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetrySupervisorModuleDetails) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *NiatelemetrySupervisorModuleDetails) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySupervisorModuleDetails) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *NiatelemetrySupervisorModuleDetails) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *NiatelemetrySupervisorModuleDetails) SetSerial(v string) {
	o.Serial = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetrySupervisorModuleDetails) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySupervisorModuleDetails) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetrySupervisorModuleDetails) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetrySupervisorModuleDetails) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetrySupervisorModuleDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySupervisorModuleDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetrySupervisorModuleDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetrySupervisorModuleDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetrySupervisorModuleDetails) MarshalJSON() ([]byte, error) {
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
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.HwVer != nil {
		toSerialize["HwVer"] = o.HwVer
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetrySupervisorModuleDetails) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetrySupervisorModuleDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Dn of the supervisor module in APIC.
		Dn *string `json:"Dn,omitempty"`
		// Hardware version of supervisor module.
		HwVer *string `json:"HwVer,omitempty"`
		// Model of the supervisor module.
		Model *string `json:"Model,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Serial number of the supervisor module.
		Serial *string `json:"Serial,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName         *string                              `json:"SiteName,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetrySupervisorModuleDetailsWithoutEmbeddedStruct := NiatelemetrySupervisorModuleDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetrySupervisorModuleDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetrySupervisorModuleDetails := _NiatelemetrySupervisorModuleDetails{}
		varNiatelemetrySupervisorModuleDetails.ClassId = varNiatelemetrySupervisorModuleDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetrySupervisorModuleDetails.ObjectType = varNiatelemetrySupervisorModuleDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetrySupervisorModuleDetails.Dn = varNiatelemetrySupervisorModuleDetailsWithoutEmbeddedStruct.Dn
		varNiatelemetrySupervisorModuleDetails.HwVer = varNiatelemetrySupervisorModuleDetailsWithoutEmbeddedStruct.HwVer
		varNiatelemetrySupervisorModuleDetails.Model = varNiatelemetrySupervisorModuleDetailsWithoutEmbeddedStruct.Model
		varNiatelemetrySupervisorModuleDetails.RecordType = varNiatelemetrySupervisorModuleDetailsWithoutEmbeddedStruct.RecordType
		varNiatelemetrySupervisorModuleDetails.RecordVersion = varNiatelemetrySupervisorModuleDetailsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetrySupervisorModuleDetails.Serial = varNiatelemetrySupervisorModuleDetailsWithoutEmbeddedStruct.Serial
		varNiatelemetrySupervisorModuleDetails.SiteName = varNiatelemetrySupervisorModuleDetailsWithoutEmbeddedStruct.SiteName
		varNiatelemetrySupervisorModuleDetails.RegisteredDevice = varNiatelemetrySupervisorModuleDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetrySupervisorModuleDetails(varNiatelemetrySupervisorModuleDetails)
	} else {
		return err
	}

	varNiatelemetrySupervisorModuleDetails := _NiatelemetrySupervisorModuleDetails{}

	err = json.Unmarshal(bytes, &varNiatelemetrySupervisorModuleDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetrySupervisorModuleDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "HwVer")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableNiatelemetrySupervisorModuleDetails struct {
	value *NiatelemetrySupervisorModuleDetails
	isSet bool
}

func (v NullableNiatelemetrySupervisorModuleDetails) Get() *NiatelemetrySupervisorModuleDetails {
	return v.value
}

func (v *NullableNiatelemetrySupervisorModuleDetails) Set(val *NiatelemetrySupervisorModuleDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetrySupervisorModuleDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetrySupervisorModuleDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetrySupervisorModuleDetails(val *NiatelemetrySupervisorModuleDetails) *NullableNiatelemetrySupervisorModuleDetails {
	return &NullableNiatelemetrySupervisorModuleDetails{value: val, isSet: true}
}

func (v NullableNiatelemetrySupervisorModuleDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetrySupervisorModuleDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
