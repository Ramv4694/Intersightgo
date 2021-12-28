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
	"reflect"
	"strings"
)

// NtpAuthNtpServer All the properties of a single authenticated NTP server.
type NtpAuthNtpServer struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of symmetric key to use for this server. * `SHA1` - Key type used by the authentication is SHA1.
	KeyType *string `json:"KeyType,omitempty"`
	// Server hostname or IP address.
	ServerName *string `json:"ServerName,omitempty"`
	// The key ID is a positive integer that identifies a cryptographic key used to authenticate NTP messages.
	SymKeyId *int64 `json:"SymKeyId,omitempty"`
	// The value of the symmetric key.
	SymKeyValue          *string `json:"SymKeyValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NtpAuthNtpServer NtpAuthNtpServer

// NewNtpAuthNtpServer instantiates a new NtpAuthNtpServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNtpAuthNtpServer(classId string, objectType string) *NtpAuthNtpServer {
	this := NtpAuthNtpServer{}
	this.ClassId = classId
	this.ObjectType = objectType
	var keyType string = "SHA1"
	this.KeyType = &keyType
	return &this
}

// NewNtpAuthNtpServerWithDefaults instantiates a new NtpAuthNtpServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNtpAuthNtpServerWithDefaults() *NtpAuthNtpServer {
	this := NtpAuthNtpServer{}
	var classId string = "ntp.AuthNtpServer"
	this.ClassId = classId
	var objectType string = "ntp.AuthNtpServer"
	this.ObjectType = objectType
	var keyType string = "SHA1"
	this.KeyType = &keyType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NtpAuthNtpServer) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NtpAuthNtpServer) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NtpAuthNtpServer) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NtpAuthNtpServer) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NtpAuthNtpServer) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NtpAuthNtpServer) SetObjectType(v string) {
	o.ObjectType = v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *NtpAuthNtpServer) GetKeyType() string {
	if o == nil || o.KeyType == nil {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpAuthNtpServer) GetKeyTypeOk() (*string, bool) {
	if o == nil || o.KeyType == nil {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *NtpAuthNtpServer) HasKeyType() bool {
	if o != nil && o.KeyType != nil {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *NtpAuthNtpServer) SetKeyType(v string) {
	o.KeyType = &v
}

// GetServerName returns the ServerName field value if set, zero value otherwise.
func (o *NtpAuthNtpServer) GetServerName() string {
	if o == nil || o.ServerName == nil {
		var ret string
		return ret
	}
	return *o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpAuthNtpServer) GetServerNameOk() (*string, bool) {
	if o == nil || o.ServerName == nil {
		return nil, false
	}
	return o.ServerName, true
}

// HasServerName returns a boolean if a field has been set.
func (o *NtpAuthNtpServer) HasServerName() bool {
	if o != nil && o.ServerName != nil {
		return true
	}

	return false
}

// SetServerName gets a reference to the given string and assigns it to the ServerName field.
func (o *NtpAuthNtpServer) SetServerName(v string) {
	o.ServerName = &v
}

// GetSymKeyId returns the SymKeyId field value if set, zero value otherwise.
func (o *NtpAuthNtpServer) GetSymKeyId() int64 {
	if o == nil || o.SymKeyId == nil {
		var ret int64
		return ret
	}
	return *o.SymKeyId
}

// GetSymKeyIdOk returns a tuple with the SymKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpAuthNtpServer) GetSymKeyIdOk() (*int64, bool) {
	if o == nil || o.SymKeyId == nil {
		return nil, false
	}
	return o.SymKeyId, true
}

// HasSymKeyId returns a boolean if a field has been set.
func (o *NtpAuthNtpServer) HasSymKeyId() bool {
	if o != nil && o.SymKeyId != nil {
		return true
	}

	return false
}

// SetSymKeyId gets a reference to the given int64 and assigns it to the SymKeyId field.
func (o *NtpAuthNtpServer) SetSymKeyId(v int64) {
	o.SymKeyId = &v
}

// GetSymKeyValue returns the SymKeyValue field value if set, zero value otherwise.
func (o *NtpAuthNtpServer) GetSymKeyValue() string {
	if o == nil || o.SymKeyValue == nil {
		var ret string
		return ret
	}
	return *o.SymKeyValue
}

// GetSymKeyValueOk returns a tuple with the SymKeyValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpAuthNtpServer) GetSymKeyValueOk() (*string, bool) {
	if o == nil || o.SymKeyValue == nil {
		return nil, false
	}
	return o.SymKeyValue, true
}

// HasSymKeyValue returns a boolean if a field has been set.
func (o *NtpAuthNtpServer) HasSymKeyValue() bool {
	if o != nil && o.SymKeyValue != nil {
		return true
	}

	return false
}

// SetSymKeyValue gets a reference to the given string and assigns it to the SymKeyValue field.
func (o *NtpAuthNtpServer) SetSymKeyValue(v string) {
	o.SymKeyValue = &v
}

func (o NtpAuthNtpServer) MarshalJSON() ([]byte, error) {
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
	if o.KeyType != nil {
		toSerialize["KeyType"] = o.KeyType
	}
	if o.ServerName != nil {
		toSerialize["ServerName"] = o.ServerName
	}
	if o.SymKeyId != nil {
		toSerialize["SymKeyId"] = o.SymKeyId
	}
	if o.SymKeyValue != nil {
		toSerialize["SymKeyValue"] = o.SymKeyValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NtpAuthNtpServer) UnmarshalJSON(bytes []byte) (err error) {
	type NtpAuthNtpServerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Type of symmetric key to use for this server. * `SHA1` - Key type used by the authentication is SHA1.
		KeyType *string `json:"KeyType,omitempty"`
		// Server hostname or IP address.
		ServerName *string `json:"ServerName,omitempty"`
		// The key ID is a positive integer that identifies a cryptographic key used to authenticate NTP messages.
		SymKeyId *int64 `json:"SymKeyId,omitempty"`
		// The value of the symmetric key.
		SymKeyValue *string `json:"SymKeyValue,omitempty"`
	}

	varNtpAuthNtpServerWithoutEmbeddedStruct := NtpAuthNtpServerWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNtpAuthNtpServerWithoutEmbeddedStruct)
	if err == nil {
		varNtpAuthNtpServer := _NtpAuthNtpServer{}
		varNtpAuthNtpServer.ClassId = varNtpAuthNtpServerWithoutEmbeddedStruct.ClassId
		varNtpAuthNtpServer.ObjectType = varNtpAuthNtpServerWithoutEmbeddedStruct.ObjectType
		varNtpAuthNtpServer.KeyType = varNtpAuthNtpServerWithoutEmbeddedStruct.KeyType
		varNtpAuthNtpServer.ServerName = varNtpAuthNtpServerWithoutEmbeddedStruct.ServerName
		varNtpAuthNtpServer.SymKeyId = varNtpAuthNtpServerWithoutEmbeddedStruct.SymKeyId
		varNtpAuthNtpServer.SymKeyValue = varNtpAuthNtpServerWithoutEmbeddedStruct.SymKeyValue
		*o = NtpAuthNtpServer(varNtpAuthNtpServer)
	} else {
		return err
	}

	varNtpAuthNtpServer := _NtpAuthNtpServer{}

	err = json.Unmarshal(bytes, &varNtpAuthNtpServer)
	if err == nil {
		o.MoBaseComplexType = varNtpAuthNtpServer.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "KeyType")
		delete(additionalProperties, "ServerName")
		delete(additionalProperties, "SymKeyId")
		delete(additionalProperties, "SymKeyValue")

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

type NullableNtpAuthNtpServer struct {
	value *NtpAuthNtpServer
	isSet bool
}

func (v NullableNtpAuthNtpServer) Get() *NtpAuthNtpServer {
	return v.value
}

func (v *NullableNtpAuthNtpServer) Set(val *NtpAuthNtpServer) {
	v.value = val
	v.isSet = true
}

func (v NullableNtpAuthNtpServer) IsSet() bool {
	return v.isSet
}

func (v *NullableNtpAuthNtpServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNtpAuthNtpServer(val *NtpAuthNtpServer) *NullableNtpAuthNtpServer {
	return &NullableNtpAuthNtpServer{value: val, isSet: true}
}

func (v NullableNtpAuthNtpServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNtpAuthNtpServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
