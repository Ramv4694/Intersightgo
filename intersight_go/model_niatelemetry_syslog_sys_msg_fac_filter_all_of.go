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
)

// NiatelemetrySyslogSysMsgFacFilterAllOf Definition of the list of properties defined in 'niatelemetry.SyslogSysMsgFacFilter', excluding properties defined in parent classes.
type NiatelemetrySyslogSysMsgFacFilterAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Parent common policy for syslog system msg in APIC.
	CommonPolicy *string `json:"CommonPolicy,omitempty"`
	// Dn of the Syslog System msg facility filter in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Facility of Syslog System msg facility filter in APIC.
	Facility *string `json:"Facility,omitempty"`
	// Minimum severity of Syslog System msg facility filter in APIC.
	MinSev *string `json:"MinSev,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// Parent syslog msg for syslog sys msg facility filter in APIC.
	SyslogSysMsg         *string                              `json:"SyslogSysMsg,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetrySyslogSysMsgFacFilterAllOf NiatelemetrySyslogSysMsgFacFilterAllOf

// NewNiatelemetrySyslogSysMsgFacFilterAllOf instantiates a new NiatelemetrySyslogSysMsgFacFilterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetrySyslogSysMsgFacFilterAllOf(classId string, objectType string) *NiatelemetrySyslogSysMsgFacFilterAllOf {
	this := NiatelemetrySyslogSysMsgFacFilterAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetrySyslogSysMsgFacFilterAllOfWithDefaults instantiates a new NiatelemetrySyslogSysMsgFacFilterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetrySyslogSysMsgFacFilterAllOfWithDefaults() *NiatelemetrySyslogSysMsgFacFilterAllOf {
	this := NiatelemetrySyslogSysMsgFacFilterAllOf{}
	var classId string = "niatelemetry.SyslogSysMsgFacFilter"
	this.ClassId = classId
	var objectType string = "niatelemetry.SyslogSysMsgFacFilter"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCommonPolicy returns the CommonPolicy field value if set, zero value otherwise.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetCommonPolicy() string {
	if o == nil || o.CommonPolicy == nil {
		var ret string
		return ret
	}
	return *o.CommonPolicy
}

// GetCommonPolicyOk returns a tuple with the CommonPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetCommonPolicyOk() (*string, bool) {
	if o == nil || o.CommonPolicy == nil {
		return nil, false
	}
	return o.CommonPolicy, true
}

// HasCommonPolicy returns a boolean if a field has been set.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) HasCommonPolicy() bool {
	if o != nil && o.CommonPolicy != nil {
		return true
	}

	return false
}

// SetCommonPolicy gets a reference to the given string and assigns it to the CommonPolicy field.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetCommonPolicy(v string) {
	o.CommonPolicy = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetFacility returns the Facility field value if set, zero value otherwise.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetFacility() string {
	if o == nil || o.Facility == nil {
		var ret string
		return ret
	}
	return *o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetFacilityOk() (*string, bool) {
	if o == nil || o.Facility == nil {
		return nil, false
	}
	return o.Facility, true
}

// HasFacility returns a boolean if a field has been set.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) HasFacility() bool {
	if o != nil && o.Facility != nil {
		return true
	}

	return false
}

// SetFacility gets a reference to the given string and assigns it to the Facility field.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetFacility(v string) {
	o.Facility = &v
}

// GetMinSev returns the MinSev field value if set, zero value otherwise.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetMinSev() string {
	if o == nil || o.MinSev == nil {
		var ret string
		return ret
	}
	return *o.MinSev
}

// GetMinSevOk returns a tuple with the MinSev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetMinSevOk() (*string, bool) {
	if o == nil || o.MinSev == nil {
		return nil, false
	}
	return o.MinSev, true
}

// HasMinSev returns a boolean if a field has been set.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) HasMinSev() bool {
	if o != nil && o.MinSev != nil {
		return true
	}

	return false
}

// SetMinSev gets a reference to the given string and assigns it to the MinSev field.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetMinSev(v string) {
	o.MinSev = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSyslogSysMsg returns the SyslogSysMsg field value if set, zero value otherwise.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetSyslogSysMsg() string {
	if o == nil || o.SyslogSysMsg == nil {
		var ret string
		return ret
	}
	return *o.SyslogSysMsg
}

// GetSyslogSysMsgOk returns a tuple with the SyslogSysMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetSyslogSysMsgOk() (*string, bool) {
	if o == nil || o.SyslogSysMsg == nil {
		return nil, false
	}
	return o.SyslogSysMsg, true
}

// HasSyslogSysMsg returns a boolean if a field has been set.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) HasSyslogSysMsg() bool {
	if o != nil && o.SyslogSysMsg != nil {
		return true
	}

	return false
}

// SetSyslogSysMsg gets a reference to the given string and assigns it to the SyslogSysMsg field.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetSyslogSysMsg(v string) {
	o.SyslogSysMsg = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetrySyslogSysMsgFacFilterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CommonPolicy != nil {
		toSerialize["CommonPolicy"] = o.CommonPolicy
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.Facility != nil {
		toSerialize["Facility"] = o.Facility
	}
	if o.MinSev != nil {
		toSerialize["MinSev"] = o.MinSev
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

func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetrySyslogSysMsgFacFilterAllOf := _NiatelemetrySyslogSysMsgFacFilterAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetrySyslogSysMsgFacFilterAllOf); err == nil {
		*o = NiatelemetrySyslogSysMsgFacFilterAllOf(varNiatelemetrySyslogSysMsgFacFilterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CommonPolicy")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "Facility")
		delete(additionalProperties, "MinSev")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "SyslogSysMsg")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetrySyslogSysMsgFacFilterAllOf struct {
	value *NiatelemetrySyslogSysMsgFacFilterAllOf
	isSet bool
}

func (v NullableNiatelemetrySyslogSysMsgFacFilterAllOf) Get() *NiatelemetrySyslogSysMsgFacFilterAllOf {
	return v.value
}

func (v *NullableNiatelemetrySyslogSysMsgFacFilterAllOf) Set(val *NiatelemetrySyslogSysMsgFacFilterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetrySyslogSysMsgFacFilterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetrySyslogSysMsgFacFilterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetrySyslogSysMsgFacFilterAllOf(val *NiatelemetrySyslogSysMsgFacFilterAllOf) *NullableNiatelemetrySyslogSysMsgFacFilterAllOf {
	return &NullableNiatelemetrySyslogSysMsgFacFilterAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetrySyslogSysMsgFacFilterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetrySyslogSysMsgFacFilterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}