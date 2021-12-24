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

// CertificatemanagementCertificateBase Lists properties that are common to all certificate types.
type CertificatemanagementCertificateBase struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType  string                  `json:"ObjectType"`
	Certificate NullableX509Certificate `json:"Certificate,omitempty"`
	// Enable/Disable the certificate in Certificate Management policy.
	Enabled *bool `json:"Enabled,omitempty"`
	// Indicates whether the value of the 'privatekey' property has been set.
	IsPrivatekeySet *bool `json:"IsPrivatekeySet,omitempty"`
	// Private Key which is used to validate the certificate.
	Privatekey           *string `json:"Privatekey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificatemanagementCertificateBase CertificatemanagementCertificateBase

// NewCertificatemanagementCertificateBase instantiates a new CertificatemanagementCertificateBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificatemanagementCertificateBase(classId string, objectType string) *CertificatemanagementCertificateBase {
	this := CertificatemanagementCertificateBase{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var isPrivatekeySet bool = false
	this.IsPrivatekeySet = &isPrivatekeySet
	return &this
}

// NewCertificatemanagementCertificateBaseWithDefaults instantiates a new CertificatemanagementCertificateBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificatemanagementCertificateBaseWithDefaults() *CertificatemanagementCertificateBase {
	this := CertificatemanagementCertificateBase{}
	var classId string = "certificatemanagement.Imc"
	this.ClassId = classId
	var objectType string = "certificatemanagement.Imc"
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var isPrivatekeySet bool = false
	this.IsPrivatekeySet = &isPrivatekeySet
	return &this
}

// GetClassId returns the ClassId field value
func (o *CertificatemanagementCertificateBase) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CertificatemanagementCertificateBase) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CertificatemanagementCertificateBase) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CertificatemanagementCertificateBase) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CertificatemanagementCertificateBase) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CertificatemanagementCertificateBase) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificatemanagementCertificateBase) GetCertificate() X509Certificate {
	if o == nil || o.Certificate.Get() == nil {
		var ret X509Certificate
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificatemanagementCertificateBase) GetCertificateOk() (*X509Certificate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *CertificatemanagementCertificateBase) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableX509Certificate and assigns it to the Certificate field.
func (o *CertificatemanagementCertificateBase) SetCertificate(v X509Certificate) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *CertificatemanagementCertificateBase) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *CertificatemanagementCertificateBase) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CertificatemanagementCertificateBase) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificatemanagementCertificateBase) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CertificatemanagementCertificateBase) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CertificatemanagementCertificateBase) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIsPrivatekeySet returns the IsPrivatekeySet field value if set, zero value otherwise.
func (o *CertificatemanagementCertificateBase) GetIsPrivatekeySet() bool {
	if o == nil || o.IsPrivatekeySet == nil {
		var ret bool
		return ret
	}
	return *o.IsPrivatekeySet
}

// GetIsPrivatekeySetOk returns a tuple with the IsPrivatekeySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificatemanagementCertificateBase) GetIsPrivatekeySetOk() (*bool, bool) {
	if o == nil || o.IsPrivatekeySet == nil {
		return nil, false
	}
	return o.IsPrivatekeySet, true
}

// HasIsPrivatekeySet returns a boolean if a field has been set.
func (o *CertificatemanagementCertificateBase) HasIsPrivatekeySet() bool {
	if o != nil && o.IsPrivatekeySet != nil {
		return true
	}

	return false
}

// SetIsPrivatekeySet gets a reference to the given bool and assigns it to the IsPrivatekeySet field.
func (o *CertificatemanagementCertificateBase) SetIsPrivatekeySet(v bool) {
	o.IsPrivatekeySet = &v
}

// GetPrivatekey returns the Privatekey field value if set, zero value otherwise.
func (o *CertificatemanagementCertificateBase) GetPrivatekey() string {
	if o == nil || o.Privatekey == nil {
		var ret string
		return ret
	}
	return *o.Privatekey
}

// GetPrivatekeyOk returns a tuple with the Privatekey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificatemanagementCertificateBase) GetPrivatekeyOk() (*string, bool) {
	if o == nil || o.Privatekey == nil {
		return nil, false
	}
	return o.Privatekey, true
}

// HasPrivatekey returns a boolean if a field has been set.
func (o *CertificatemanagementCertificateBase) HasPrivatekey() bool {
	if o != nil && o.Privatekey != nil {
		return true
	}

	return false
}

// SetPrivatekey gets a reference to the given string and assigns it to the Privatekey field.
func (o *CertificatemanagementCertificateBase) SetPrivatekey(v string) {
	o.Privatekey = &v
}

func (o CertificatemanagementCertificateBase) MarshalJSON() ([]byte, error) {
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
	if o.Certificate.IsSet() {
		toSerialize["Certificate"] = o.Certificate.Get()
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.IsPrivatekeySet != nil {
		toSerialize["IsPrivatekeySet"] = o.IsPrivatekeySet
	}
	if o.Privatekey != nil {
		toSerialize["Privatekey"] = o.Privatekey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CertificatemanagementCertificateBase) UnmarshalJSON(bytes []byte) (err error) {
	type CertificatemanagementCertificateBaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType  string                  `json:"ObjectType"`
		Certificate NullableX509Certificate `json:"Certificate,omitempty"`
		// Enable/Disable the certificate in Certificate Management policy.
		Enabled *bool `json:"Enabled,omitempty"`
		// Indicates whether the value of the 'privatekey' property has been set.
		IsPrivatekeySet *bool `json:"IsPrivatekeySet,omitempty"`
		// Private Key which is used to validate the certificate.
		Privatekey *string `json:"Privatekey,omitempty"`
	}

	varCertificatemanagementCertificateBaseWithoutEmbeddedStruct := CertificatemanagementCertificateBaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCertificatemanagementCertificateBaseWithoutEmbeddedStruct)
	if err == nil {
		varCertificatemanagementCertificateBase := _CertificatemanagementCertificateBase{}
		varCertificatemanagementCertificateBase.ClassId = varCertificatemanagementCertificateBaseWithoutEmbeddedStruct.ClassId
		varCertificatemanagementCertificateBase.ObjectType = varCertificatemanagementCertificateBaseWithoutEmbeddedStruct.ObjectType
		varCertificatemanagementCertificateBase.Certificate = varCertificatemanagementCertificateBaseWithoutEmbeddedStruct.Certificate
		varCertificatemanagementCertificateBase.Enabled = varCertificatemanagementCertificateBaseWithoutEmbeddedStruct.Enabled
		varCertificatemanagementCertificateBase.IsPrivatekeySet = varCertificatemanagementCertificateBaseWithoutEmbeddedStruct.IsPrivatekeySet
		varCertificatemanagementCertificateBase.Privatekey = varCertificatemanagementCertificateBaseWithoutEmbeddedStruct.Privatekey
		*o = CertificatemanagementCertificateBase(varCertificatemanagementCertificateBase)
	} else {
		return err
	}

	varCertificatemanagementCertificateBase := _CertificatemanagementCertificateBase{}

	err = json.Unmarshal(bytes, &varCertificatemanagementCertificateBase)
	if err == nil {
		o.MoBaseComplexType = varCertificatemanagementCertificateBase.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Certificate")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "IsPrivatekeySet")
		delete(additionalProperties, "Privatekey")

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

type NullableCertificatemanagementCertificateBase struct {
	value *CertificatemanagementCertificateBase
	isSet bool
}

func (v NullableCertificatemanagementCertificateBase) Get() *CertificatemanagementCertificateBase {
	return v.value
}

func (v *NullableCertificatemanagementCertificateBase) Set(val *CertificatemanagementCertificateBase) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificatemanagementCertificateBase) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificatemanagementCertificateBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificatemanagementCertificateBase(val *CertificatemanagementCertificateBase) *NullableCertificatemanagementCertificateBase {
	return &NullableCertificatemanagementCertificateBase{value: val, isSet: true}
}

func (v NullableCertificatemanagementCertificateBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificatemanagementCertificateBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
