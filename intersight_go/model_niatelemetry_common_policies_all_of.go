/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-12-27T12:26:28Z.
 *
 * API version: 0.0.1-37434
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// NiatelemetryCommonPoliciesAllOf Definition of the list of properties defined in 'niatelemetry.CommonPolicies', excluding properties defined in parent classes.
type NiatelemetryCommonPoliciesAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Dn of the Common Policy in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// List of Dn of SNMP Src for the above common pol.
	SnmpSrc *string `json:"SnmpSrc,omitempty"`
	// List of Dn of Syslog Src for the above common pol.
	SyslogSrc *string `json:"SyslogSrc,omitempty"`
	// List of Dn of Syslog Sys Msg the above common pol.
	SyslogSysMsg         *string                              `json:"SyslogSysMsg,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryCommonPoliciesAllOf NiatelemetryCommonPoliciesAllOf

// NewNiatelemetryCommonPoliciesAllOf instantiates a new NiatelemetryCommonPoliciesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryCommonPoliciesAllOf(classId string, objectType string) *NiatelemetryCommonPoliciesAllOf {
	this := NiatelemetryCommonPoliciesAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryCommonPoliciesAllOfWithDefaults instantiates a new NiatelemetryCommonPoliciesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryCommonPoliciesAllOfWithDefaults() *NiatelemetryCommonPoliciesAllOf {
	this := NiatelemetryCommonPoliciesAllOf{}
	var classId string = "niatelemetry.CommonPolicies"
	this.ClassId = classId
	var objectType string = "niatelemetry.CommonPolicies"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryCommonPoliciesAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryCommonPoliciesAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryCommonPoliciesAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryCommonPoliciesAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryCommonPoliciesAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryCommonPoliciesAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryCommonPoliciesAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryCommonPoliciesAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryCommonPoliciesAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryCommonPoliciesAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryCommonPoliciesAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryCommonPoliciesAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryCommonPoliciesAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryCommonPoliciesAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryCommonPoliciesAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryCommonPoliciesAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryCommonPoliciesAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryCommonPoliciesAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryCommonPoliciesAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryCommonPoliciesAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryCommonPoliciesAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryCommonPoliciesAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSnmpSrc returns the SnmpSrc field value if set, zero value otherwise.
func (o *NiatelemetryCommonPoliciesAllOf) GetSnmpSrc() string {
	if o == nil || o.SnmpSrc == nil {
		var ret string
		return ret
	}
	return *o.SnmpSrc
}

// GetSnmpSrcOk returns a tuple with the SnmpSrc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryCommonPoliciesAllOf) GetSnmpSrcOk() (*string, bool) {
	if o == nil || o.SnmpSrc == nil {
		return nil, false
	}
	return o.SnmpSrc, true
}

// HasSnmpSrc returns a boolean if a field has been set.
func (o *NiatelemetryCommonPoliciesAllOf) HasSnmpSrc() bool {
	if o != nil && o.SnmpSrc != nil {
		return true
	}

	return false
}

// SetSnmpSrc gets a reference to the given string and assigns it to the SnmpSrc field.
func (o *NiatelemetryCommonPoliciesAllOf) SetSnmpSrc(v string) {
	o.SnmpSrc = &v
}

// GetSyslogSrc returns the SyslogSrc field value if set, zero value otherwise.
func (o *NiatelemetryCommonPoliciesAllOf) GetSyslogSrc() string {
	if o == nil || o.SyslogSrc == nil {
		var ret string
		return ret
	}
	return *o.SyslogSrc
}

// GetSyslogSrcOk returns a tuple with the SyslogSrc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryCommonPoliciesAllOf) GetSyslogSrcOk() (*string, bool) {
	if o == nil || o.SyslogSrc == nil {
		return nil, false
	}
	return o.SyslogSrc, true
}

// HasSyslogSrc returns a boolean if a field has been set.
func (o *NiatelemetryCommonPoliciesAllOf) HasSyslogSrc() bool {
	if o != nil && o.SyslogSrc != nil {
		return true
	}

	return false
}

// SetSyslogSrc gets a reference to the given string and assigns it to the SyslogSrc field.
func (o *NiatelemetryCommonPoliciesAllOf) SetSyslogSrc(v string) {
	o.SyslogSrc = &v
}

// GetSyslogSysMsg returns the SyslogSysMsg field value if set, zero value otherwise.
func (o *NiatelemetryCommonPoliciesAllOf) GetSyslogSysMsg() string {
	if o == nil || o.SyslogSysMsg == nil {
		var ret string
		return ret
	}
	return *o.SyslogSysMsg
}

// GetSyslogSysMsgOk returns a tuple with the SyslogSysMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryCommonPoliciesAllOf) GetSyslogSysMsgOk() (*string, bool) {
	if o == nil || o.SyslogSysMsg == nil {
		return nil, false
	}
	return o.SyslogSysMsg, true
}

// HasSyslogSysMsg returns a boolean if a field has been set.
func (o *NiatelemetryCommonPoliciesAllOf) HasSyslogSysMsg() bool {
	if o != nil && o.SyslogSysMsg != nil {
		return true
	}

	return false
}

// SetSyslogSysMsg gets a reference to the given string and assigns it to the SyslogSysMsg field.
func (o *NiatelemetryCommonPoliciesAllOf) SetSyslogSysMsg(v string) {
	o.SyslogSysMsg = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryCommonPoliciesAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryCommonPoliciesAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryCommonPoliciesAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryCommonPoliciesAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryCommonPoliciesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
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
	if o.SnmpSrc != nil {
		toSerialize["SnmpSrc"] = o.SnmpSrc
	}
	if o.SyslogSrc != nil {
		toSerialize["SyslogSrc"] = o.SyslogSrc
	}
	if o.SyslogSysMsg != nil {
		toSerialize["SyslogSysMsg"] = o.SyslogSysMsg
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryCommonPoliciesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryCommonPoliciesAllOf := _NiatelemetryCommonPoliciesAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryCommonPoliciesAllOf); err == nil {
		*o = NiatelemetryCommonPoliciesAllOf(varNiatelemetryCommonPoliciesAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "SnmpSrc")
		delete(additionalProperties, "SyslogSrc")
		delete(additionalProperties, "SyslogSysMsg")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryCommonPoliciesAllOf struct {
	value *NiatelemetryCommonPoliciesAllOf
	isSet bool
}

func (v NullableNiatelemetryCommonPoliciesAllOf) Get() *NiatelemetryCommonPoliciesAllOf {
	return v.value
}

func (v *NullableNiatelemetryCommonPoliciesAllOf) Set(val *NiatelemetryCommonPoliciesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryCommonPoliciesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryCommonPoliciesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryCommonPoliciesAllOf(val *NiatelemetryCommonPoliciesAllOf) *NullableNiatelemetryCommonPoliciesAllOf {
	return &NullableNiatelemetryCommonPoliciesAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryCommonPoliciesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryCommonPoliciesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
