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

// FirmwareNetworkShare Firmware upgrade where the image is located in remote shared machine.
type FirmwareNetworkShare struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                     `json:"ObjectType"`
	CifsServer NullableFirmwareCifsServer `json:"CifsServer,omitempty"`
	HttpServer NullableFirmwareHttpServer `json:"HttpServer,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// File server protocols such as CIFS, NFS, WWW for HTTP (S) that hosts the image. * `nfs` - File server protocol used is NFS. * `cifs` - File server protocol used is CIFS. * `www` - File server protocol used is WWW.
	MapType   *string                   `json:"MapType,omitempty"`
	NfsServer NullableFirmwareNfsServer `json:"NfsServer,omitempty"`
	// Password as configured on the file server.
	Password *string `json:"Password,omitempty"`
	// Option to control the upgrade operation. Some examples, 1) nw_upgrade_mount_only - mount the image from a file server and run the upgrade on the next server boot and 2) nw_upgrade_full - mount the image and immediately run the upgrade. * `nw_upgrade_full` - Network upgrade option for full upgrade. * `nw_upgrade_mount_only` - Network upgrade mount only.
	Upgradeoption *string `json:"Upgradeoption,omitempty"`
	// Username as configured on the file server.
	Username             *string `json:"Username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareNetworkShare FirmwareNetworkShare

// NewFirmwareNetworkShare instantiates a new FirmwareNetworkShare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareNetworkShare(classId string, objectType string) *FirmwareNetworkShare {
	this := FirmwareNetworkShare{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	var mapType string = "nfs"
	this.MapType = &mapType
	var upgradeoption string = "nw_upgrade_full"
	this.Upgradeoption = &upgradeoption
	return &this
}

// NewFirmwareNetworkShareWithDefaults instantiates a new FirmwareNetworkShare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareNetworkShareWithDefaults() *FirmwareNetworkShare {
	this := FirmwareNetworkShare{}
	var classId string = "firmware.NetworkShare"
	this.ClassId = classId
	var objectType string = "firmware.NetworkShare"
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	var mapType string = "nfs"
	this.MapType = &mapType
	var upgradeoption string = "nw_upgrade_full"
	this.Upgradeoption = &upgradeoption
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareNetworkShare) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareNetworkShare) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareNetworkShare) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareNetworkShare) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareNetworkShare) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareNetworkShare) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCifsServer returns the CifsServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareNetworkShare) GetCifsServer() FirmwareCifsServer {
	if o == nil || o.CifsServer.Get() == nil {
		var ret FirmwareCifsServer
		return ret
	}
	return *o.CifsServer.Get()
}

// GetCifsServerOk returns a tuple with the CifsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareNetworkShare) GetCifsServerOk() (*FirmwareCifsServer, bool) {
	if o == nil {
		return nil, false
	}
	return o.CifsServer.Get(), o.CifsServer.IsSet()
}

// HasCifsServer returns a boolean if a field has been set.
func (o *FirmwareNetworkShare) HasCifsServer() bool {
	if o != nil && o.CifsServer.IsSet() {
		return true
	}

	return false
}

// SetCifsServer gets a reference to the given NullableFirmwareCifsServer and assigns it to the CifsServer field.
func (o *FirmwareNetworkShare) SetCifsServer(v FirmwareCifsServer) {
	o.CifsServer.Set(&v)
}

// SetCifsServerNil sets the value for CifsServer to be an explicit nil
func (o *FirmwareNetworkShare) SetCifsServerNil() {
	o.CifsServer.Set(nil)
}

// UnsetCifsServer ensures that no value is present for CifsServer, not even an explicit nil
func (o *FirmwareNetworkShare) UnsetCifsServer() {
	o.CifsServer.Unset()
}

// GetHttpServer returns the HttpServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareNetworkShare) GetHttpServer() FirmwareHttpServer {
	if o == nil || o.HttpServer.Get() == nil {
		var ret FirmwareHttpServer
		return ret
	}
	return *o.HttpServer.Get()
}

// GetHttpServerOk returns a tuple with the HttpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareNetworkShare) GetHttpServerOk() (*FirmwareHttpServer, bool) {
	if o == nil {
		return nil, false
	}
	return o.HttpServer.Get(), o.HttpServer.IsSet()
}

// HasHttpServer returns a boolean if a field has been set.
func (o *FirmwareNetworkShare) HasHttpServer() bool {
	if o != nil && o.HttpServer.IsSet() {
		return true
	}

	return false
}

// SetHttpServer gets a reference to the given NullableFirmwareHttpServer and assigns it to the HttpServer field.
func (o *FirmwareNetworkShare) SetHttpServer(v FirmwareHttpServer) {
	o.HttpServer.Set(&v)
}

// SetHttpServerNil sets the value for HttpServer to be an explicit nil
func (o *FirmwareNetworkShare) SetHttpServerNil() {
	o.HttpServer.Set(nil)
}

// UnsetHttpServer ensures that no value is present for HttpServer, not even an explicit nil
func (o *FirmwareNetworkShare) UnsetHttpServer() {
	o.HttpServer.Unset()
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *FirmwareNetworkShare) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareNetworkShare) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *FirmwareNetworkShare) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *FirmwareNetworkShare) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetMapType returns the MapType field value if set, zero value otherwise.
func (o *FirmwareNetworkShare) GetMapType() string {
	if o == nil || o.MapType == nil {
		var ret string
		return ret
	}
	return *o.MapType
}

// GetMapTypeOk returns a tuple with the MapType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareNetworkShare) GetMapTypeOk() (*string, bool) {
	if o == nil || o.MapType == nil {
		return nil, false
	}
	return o.MapType, true
}

// HasMapType returns a boolean if a field has been set.
func (o *FirmwareNetworkShare) HasMapType() bool {
	if o != nil && o.MapType != nil {
		return true
	}

	return false
}

// SetMapType gets a reference to the given string and assigns it to the MapType field.
func (o *FirmwareNetworkShare) SetMapType(v string) {
	o.MapType = &v
}

// GetNfsServer returns the NfsServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareNetworkShare) GetNfsServer() FirmwareNfsServer {
	if o == nil || o.NfsServer.Get() == nil {
		var ret FirmwareNfsServer
		return ret
	}
	return *o.NfsServer.Get()
}

// GetNfsServerOk returns a tuple with the NfsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareNetworkShare) GetNfsServerOk() (*FirmwareNfsServer, bool) {
	if o == nil {
		return nil, false
	}
	return o.NfsServer.Get(), o.NfsServer.IsSet()
}

// HasNfsServer returns a boolean if a field has been set.
func (o *FirmwareNetworkShare) HasNfsServer() bool {
	if o != nil && o.NfsServer.IsSet() {
		return true
	}

	return false
}

// SetNfsServer gets a reference to the given NullableFirmwareNfsServer and assigns it to the NfsServer field.
func (o *FirmwareNetworkShare) SetNfsServer(v FirmwareNfsServer) {
	o.NfsServer.Set(&v)
}

// SetNfsServerNil sets the value for NfsServer to be an explicit nil
func (o *FirmwareNetworkShare) SetNfsServerNil() {
	o.NfsServer.Set(nil)
}

// UnsetNfsServer ensures that no value is present for NfsServer, not even an explicit nil
func (o *FirmwareNetworkShare) UnsetNfsServer() {
	o.NfsServer.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *FirmwareNetworkShare) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareNetworkShare) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *FirmwareNetworkShare) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *FirmwareNetworkShare) SetPassword(v string) {
	o.Password = &v
}

// GetUpgradeoption returns the Upgradeoption field value if set, zero value otherwise.
func (o *FirmwareNetworkShare) GetUpgradeoption() string {
	if o == nil || o.Upgradeoption == nil {
		var ret string
		return ret
	}
	return *o.Upgradeoption
}

// GetUpgradeoptionOk returns a tuple with the Upgradeoption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareNetworkShare) GetUpgradeoptionOk() (*string, bool) {
	if o == nil || o.Upgradeoption == nil {
		return nil, false
	}
	return o.Upgradeoption, true
}

// HasUpgradeoption returns a boolean if a field has been set.
func (o *FirmwareNetworkShare) HasUpgradeoption() bool {
	if o != nil && o.Upgradeoption != nil {
		return true
	}

	return false
}

// SetUpgradeoption gets a reference to the given string and assigns it to the Upgradeoption field.
func (o *FirmwareNetworkShare) SetUpgradeoption(v string) {
	o.Upgradeoption = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *FirmwareNetworkShare) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareNetworkShare) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *FirmwareNetworkShare) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *FirmwareNetworkShare) SetUsername(v string) {
	o.Username = &v
}

func (o FirmwareNetworkShare) MarshalJSON() ([]byte, error) {
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
	if o.CifsServer.IsSet() {
		toSerialize["CifsServer"] = o.CifsServer.Get()
	}
	if o.HttpServer.IsSet() {
		toSerialize["HttpServer"] = o.HttpServer.Get()
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.MapType != nil {
		toSerialize["MapType"] = o.MapType
	}
	if o.NfsServer.IsSet() {
		toSerialize["NfsServer"] = o.NfsServer.Get()
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Upgradeoption != nil {
		toSerialize["Upgradeoption"] = o.Upgradeoption
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareNetworkShare) UnmarshalJSON(bytes []byte) (err error) {
	type FirmwareNetworkShareWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                     `json:"ObjectType"`
		CifsServer NullableFirmwareCifsServer `json:"CifsServer,omitempty"`
		HttpServer NullableFirmwareHttpServer `json:"HttpServer,omitempty"`
		// Indicates whether the value of the 'password' property has been set.
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
		// File server protocols such as CIFS, NFS, WWW for HTTP (S) that hosts the image. * `nfs` - File server protocol used is NFS. * `cifs` - File server protocol used is CIFS. * `www` - File server protocol used is WWW.
		MapType   *string                   `json:"MapType,omitempty"`
		NfsServer NullableFirmwareNfsServer `json:"NfsServer,omitempty"`
		// Password as configured on the file server.
		Password *string `json:"Password,omitempty"`
		// Option to control the upgrade operation. Some examples, 1) nw_upgrade_mount_only - mount the image from a file server and run the upgrade on the next server boot and 2) nw_upgrade_full - mount the image and immediately run the upgrade. * `nw_upgrade_full` - Network upgrade option for full upgrade. * `nw_upgrade_mount_only` - Network upgrade mount only.
		Upgradeoption *string `json:"Upgradeoption,omitempty"`
		// Username as configured on the file server.
		Username *string `json:"Username,omitempty"`
	}

	varFirmwareNetworkShareWithoutEmbeddedStruct := FirmwareNetworkShareWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFirmwareNetworkShareWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareNetworkShare := _FirmwareNetworkShare{}
		varFirmwareNetworkShare.ClassId = varFirmwareNetworkShareWithoutEmbeddedStruct.ClassId
		varFirmwareNetworkShare.ObjectType = varFirmwareNetworkShareWithoutEmbeddedStruct.ObjectType
		varFirmwareNetworkShare.CifsServer = varFirmwareNetworkShareWithoutEmbeddedStruct.CifsServer
		varFirmwareNetworkShare.HttpServer = varFirmwareNetworkShareWithoutEmbeddedStruct.HttpServer
		varFirmwareNetworkShare.IsPasswordSet = varFirmwareNetworkShareWithoutEmbeddedStruct.IsPasswordSet
		varFirmwareNetworkShare.MapType = varFirmwareNetworkShareWithoutEmbeddedStruct.MapType
		varFirmwareNetworkShare.NfsServer = varFirmwareNetworkShareWithoutEmbeddedStruct.NfsServer
		varFirmwareNetworkShare.Password = varFirmwareNetworkShareWithoutEmbeddedStruct.Password
		varFirmwareNetworkShare.Upgradeoption = varFirmwareNetworkShareWithoutEmbeddedStruct.Upgradeoption
		varFirmwareNetworkShare.Username = varFirmwareNetworkShareWithoutEmbeddedStruct.Username
		*o = FirmwareNetworkShare(varFirmwareNetworkShare)
	} else {
		return err
	}

	varFirmwareNetworkShare := _FirmwareNetworkShare{}

	err = json.Unmarshal(bytes, &varFirmwareNetworkShare)
	if err == nil {
		o.MoBaseComplexType = varFirmwareNetworkShare.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CifsServer")
		delete(additionalProperties, "HttpServer")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "MapType")
		delete(additionalProperties, "NfsServer")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Upgradeoption")
		delete(additionalProperties, "Username")

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

type NullableFirmwareNetworkShare struct {
	value *FirmwareNetworkShare
	isSet bool
}

func (v NullableFirmwareNetworkShare) Get() *FirmwareNetworkShare {
	return v.value
}

func (v *NullableFirmwareNetworkShare) Set(val *FirmwareNetworkShare) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareNetworkShare) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareNetworkShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareNetworkShare(val *FirmwareNetworkShare) *NullableFirmwareNetworkShare {
	return &NullableFirmwareNetworkShare{value: val, isSet: true}
}

func (v NullableFirmwareNetworkShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareNetworkShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
