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

// NiatelemetrySshVersionTwo Object to capture SSH V2 details in APIC.
type NiatelemetrySshVersionTwo struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin state of SSH V2 in APIC.
	AdminSt *string `json:"AdminSt,omitempty"`
	// Dn of SSH V2 attribute in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Password auth for SSH V2 in APIC.
	PasswordAuth *string `json:"PasswordAuth,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// SSH Ciphers for SSH V2 in APIC.
	SshCiphers *string `json:"SshCiphers,omitempty"`
	// SSH MACS for SSH V2 in APIC.
	SshMacs              *string                              `json:"SshMacs,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetrySshVersionTwo NiatelemetrySshVersionTwo

// NewNiatelemetrySshVersionTwo instantiates a new NiatelemetrySshVersionTwo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetrySshVersionTwo(classId string, objectType string) *NiatelemetrySshVersionTwo {
	this := NiatelemetrySshVersionTwo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetrySshVersionTwoWithDefaults instantiates a new NiatelemetrySshVersionTwo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetrySshVersionTwoWithDefaults() *NiatelemetrySshVersionTwo {
	this := NiatelemetrySshVersionTwo{}
	var classId string = "niatelemetry.SshVersionTwo"
	this.ClassId = classId
	var objectType string = "niatelemetry.SshVersionTwo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetrySshVersionTwo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetrySshVersionTwo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetrySshVersionTwo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetrySshVersionTwo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminSt returns the AdminSt field value if set, zero value otherwise.
func (o *NiatelemetrySshVersionTwo) GetAdminSt() string {
	if o == nil || o.AdminSt == nil {
		var ret string
		return ret
	}
	return *o.AdminSt
}

// GetAdminStOk returns a tuple with the AdminSt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwo) GetAdminStOk() (*string, bool) {
	if o == nil || o.AdminSt == nil {
		return nil, false
	}
	return o.AdminSt, true
}

// HasAdminSt returns a boolean if a field has been set.
func (o *NiatelemetrySshVersionTwo) HasAdminSt() bool {
	if o != nil && o.AdminSt != nil {
		return true
	}

	return false
}

// SetAdminSt gets a reference to the given string and assigns it to the AdminSt field.
func (o *NiatelemetrySshVersionTwo) SetAdminSt(v string) {
	o.AdminSt = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetrySshVersionTwo) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwo) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetrySshVersionTwo) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetrySshVersionTwo) SetDn(v string) {
	o.Dn = &v
}

// GetPasswordAuth returns the PasswordAuth field value if set, zero value otherwise.
func (o *NiatelemetrySshVersionTwo) GetPasswordAuth() string {
	if o == nil || o.PasswordAuth == nil {
		var ret string
		return ret
	}
	return *o.PasswordAuth
}

// GetPasswordAuthOk returns a tuple with the PasswordAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwo) GetPasswordAuthOk() (*string, bool) {
	if o == nil || o.PasswordAuth == nil {
		return nil, false
	}
	return o.PasswordAuth, true
}

// HasPasswordAuth returns a boolean if a field has been set.
func (o *NiatelemetrySshVersionTwo) HasPasswordAuth() bool {
	if o != nil && o.PasswordAuth != nil {
		return true
	}

	return false
}

// SetPasswordAuth gets a reference to the given string and assigns it to the PasswordAuth field.
func (o *NiatelemetrySshVersionTwo) SetPasswordAuth(v string) {
	o.PasswordAuth = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetrySshVersionTwo) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwo) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetrySshVersionTwo) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetrySshVersionTwo) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetrySshVersionTwo) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwo) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetrySshVersionTwo) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetrySshVersionTwo) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetrySshVersionTwo) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwo) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetrySshVersionTwo) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetrySshVersionTwo) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSshCiphers returns the SshCiphers field value if set, zero value otherwise.
func (o *NiatelemetrySshVersionTwo) GetSshCiphers() string {
	if o == nil || o.SshCiphers == nil {
		var ret string
		return ret
	}
	return *o.SshCiphers
}

// GetSshCiphersOk returns a tuple with the SshCiphers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwo) GetSshCiphersOk() (*string, bool) {
	if o == nil || o.SshCiphers == nil {
		return nil, false
	}
	return o.SshCiphers, true
}

// HasSshCiphers returns a boolean if a field has been set.
func (o *NiatelemetrySshVersionTwo) HasSshCiphers() bool {
	if o != nil && o.SshCiphers != nil {
		return true
	}

	return false
}

// SetSshCiphers gets a reference to the given string and assigns it to the SshCiphers field.
func (o *NiatelemetrySshVersionTwo) SetSshCiphers(v string) {
	o.SshCiphers = &v
}

// GetSshMacs returns the SshMacs field value if set, zero value otherwise.
func (o *NiatelemetrySshVersionTwo) GetSshMacs() string {
	if o == nil || o.SshMacs == nil {
		var ret string
		return ret
	}
	return *o.SshMacs
}

// GetSshMacsOk returns a tuple with the SshMacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwo) GetSshMacsOk() (*string, bool) {
	if o == nil || o.SshMacs == nil {
		return nil, false
	}
	return o.SshMacs, true
}

// HasSshMacs returns a boolean if a field has been set.
func (o *NiatelemetrySshVersionTwo) HasSshMacs() bool {
	if o != nil && o.SshMacs != nil {
		return true
	}

	return false
}

// SetSshMacs gets a reference to the given string and assigns it to the SshMacs field.
func (o *NiatelemetrySshVersionTwo) SetSshMacs(v string) {
	o.SshMacs = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetrySshVersionTwo) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwo) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetrySshVersionTwo) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetrySshVersionTwo) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetrySshVersionTwo) MarshalJSON() ([]byte, error) {
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
	if o.AdminSt != nil {
		toSerialize["AdminSt"] = o.AdminSt
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.PasswordAuth != nil {
		toSerialize["PasswordAuth"] = o.PasswordAuth
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.SshCiphers != nil {
		toSerialize["SshCiphers"] = o.SshCiphers
	}
	if o.SshMacs != nil {
		toSerialize["SshMacs"] = o.SshMacs
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetrySshVersionTwo) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetrySshVersionTwoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Admin state of SSH V2 in APIC.
		AdminSt *string `json:"AdminSt,omitempty"`
		// Dn of SSH V2 attribute in APIC.
		Dn *string `json:"Dn,omitempty"`
		// Password auth for SSH V2 in APIC.
		PasswordAuth *string `json:"PasswordAuth,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName *string `json:"SiteName,omitempty"`
		// SSH Ciphers for SSH V2 in APIC.
		SshCiphers *string `json:"SshCiphers,omitempty"`
		// SSH MACS for SSH V2 in APIC.
		SshMacs          *string                              `json:"SshMacs,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetrySshVersionTwoWithoutEmbeddedStruct := NiatelemetrySshVersionTwoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetrySshVersionTwoWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetrySshVersionTwo := _NiatelemetrySshVersionTwo{}
		varNiatelemetrySshVersionTwo.ClassId = varNiatelemetrySshVersionTwoWithoutEmbeddedStruct.ClassId
		varNiatelemetrySshVersionTwo.ObjectType = varNiatelemetrySshVersionTwoWithoutEmbeddedStruct.ObjectType
		varNiatelemetrySshVersionTwo.AdminSt = varNiatelemetrySshVersionTwoWithoutEmbeddedStruct.AdminSt
		varNiatelemetrySshVersionTwo.Dn = varNiatelemetrySshVersionTwoWithoutEmbeddedStruct.Dn
		varNiatelemetrySshVersionTwo.PasswordAuth = varNiatelemetrySshVersionTwoWithoutEmbeddedStruct.PasswordAuth
		varNiatelemetrySshVersionTwo.RecordType = varNiatelemetrySshVersionTwoWithoutEmbeddedStruct.RecordType
		varNiatelemetrySshVersionTwo.RecordVersion = varNiatelemetrySshVersionTwoWithoutEmbeddedStruct.RecordVersion
		varNiatelemetrySshVersionTwo.SiteName = varNiatelemetrySshVersionTwoWithoutEmbeddedStruct.SiteName
		varNiatelemetrySshVersionTwo.SshCiphers = varNiatelemetrySshVersionTwoWithoutEmbeddedStruct.SshCiphers
		varNiatelemetrySshVersionTwo.SshMacs = varNiatelemetrySshVersionTwoWithoutEmbeddedStruct.SshMacs
		varNiatelemetrySshVersionTwo.RegisteredDevice = varNiatelemetrySshVersionTwoWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetrySshVersionTwo(varNiatelemetrySshVersionTwo)
	} else {
		return err
	}

	varNiatelemetrySshVersionTwo := _NiatelemetrySshVersionTwo{}

	err = json.Unmarshal(bytes, &varNiatelemetrySshVersionTwo)
	if err == nil {
		o.MoBaseMo = varNiatelemetrySshVersionTwo.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminSt")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "PasswordAuth")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "SshCiphers")
		delete(additionalProperties, "SshMacs")
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

type NullableNiatelemetrySshVersionTwo struct {
	value *NiatelemetrySshVersionTwo
	isSet bool
}

func (v NullableNiatelemetrySshVersionTwo) Get() *NiatelemetrySshVersionTwo {
	return v.value
}

func (v *NullableNiatelemetrySshVersionTwo) Set(val *NiatelemetrySshVersionTwo) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetrySshVersionTwo) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetrySshVersionTwo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetrySshVersionTwo(val *NiatelemetrySshVersionTwo) *NullableNiatelemetrySshVersionTwo {
	return &NullableNiatelemetrySshVersionTwo{value: val, isSet: true}
}

func (v NullableNiatelemetrySshVersionTwo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetrySshVersionTwo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
