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
)

// CloudAwsOrganizationalUnitAllOf Definition of the list of properties defined in 'cloud.AwsOrganizationalUnit', excluding properties defined in parent classes.
type CloudAwsOrganizationalUnitAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The identity of this organization. This entity is not manipulated by users. It aids in uniquely identifying the organization object.
	Identity *string `json:"Identity,omitempty"`
	// Name of the organizational unit.
	Name                 *string                                 `json:"Name,omitempty"`
	ParentOrg            *CloudAwsOrganizationalUnitRelationship `json:"ParentOrg,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudAwsOrganizationalUnitAllOf CloudAwsOrganizationalUnitAllOf

// NewCloudAwsOrganizationalUnitAllOf instantiates a new CloudAwsOrganizationalUnitAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudAwsOrganizationalUnitAllOf(classId string, objectType string) *CloudAwsOrganizationalUnitAllOf {
	this := CloudAwsOrganizationalUnitAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudAwsOrganizationalUnitAllOfWithDefaults instantiates a new CloudAwsOrganizationalUnitAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudAwsOrganizationalUnitAllOfWithDefaults() *CloudAwsOrganizationalUnitAllOf {
	this := CloudAwsOrganizationalUnitAllOf{}
	var classId string = "cloud.AwsOrganizationalUnit"
	this.ClassId = classId
	var objectType string = "cloud.AwsOrganizationalUnit"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudAwsOrganizationalUnitAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudAwsOrganizationalUnitAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudAwsOrganizationalUnitAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudAwsOrganizationalUnitAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudAwsOrganizationalUnitAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudAwsOrganizationalUnitAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *CloudAwsOrganizationalUnitAllOf) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsOrganizationalUnitAllOf) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *CloudAwsOrganizationalUnitAllOf) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *CloudAwsOrganizationalUnitAllOf) SetIdentity(v string) {
	o.Identity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudAwsOrganizationalUnitAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsOrganizationalUnitAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudAwsOrganizationalUnitAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudAwsOrganizationalUnitAllOf) SetName(v string) {
	o.Name = &v
}

// GetParentOrg returns the ParentOrg field value if set, zero value otherwise.
func (o *CloudAwsOrganizationalUnitAllOf) GetParentOrg() CloudAwsOrganizationalUnitRelationship {
	if o == nil || o.ParentOrg == nil {
		var ret CloudAwsOrganizationalUnitRelationship
		return ret
	}
	return *o.ParentOrg
}

// GetParentOrgOk returns a tuple with the ParentOrg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsOrganizationalUnitAllOf) GetParentOrgOk() (*CloudAwsOrganizationalUnitRelationship, bool) {
	if o == nil || o.ParentOrg == nil {
		return nil, false
	}
	return o.ParentOrg, true
}

// HasParentOrg returns a boolean if a field has been set.
func (o *CloudAwsOrganizationalUnitAllOf) HasParentOrg() bool {
	if o != nil && o.ParentOrg != nil {
		return true
	}

	return false
}

// SetParentOrg gets a reference to the given CloudAwsOrganizationalUnitRelationship and assigns it to the ParentOrg field.
func (o *CloudAwsOrganizationalUnitAllOf) SetParentOrg(v CloudAwsOrganizationalUnitRelationship) {
	o.ParentOrg = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *CloudAwsOrganizationalUnitAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsOrganizationalUnitAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *CloudAwsOrganizationalUnitAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *CloudAwsOrganizationalUnitAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o CloudAwsOrganizationalUnitAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ParentOrg != nil {
		toSerialize["ParentOrg"] = o.ParentOrg
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudAwsOrganizationalUnitAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudAwsOrganizationalUnitAllOf := _CloudAwsOrganizationalUnitAllOf{}

	if err = json.Unmarshal(bytes, &varCloudAwsOrganizationalUnitAllOf); err == nil {
		*o = CloudAwsOrganizationalUnitAllOf(varCloudAwsOrganizationalUnitAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ParentOrg")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudAwsOrganizationalUnitAllOf struct {
	value *CloudAwsOrganizationalUnitAllOf
	isSet bool
}

func (v NullableCloudAwsOrganizationalUnitAllOf) Get() *CloudAwsOrganizationalUnitAllOf {
	return v.value
}

func (v *NullableCloudAwsOrganizationalUnitAllOf) Set(val *CloudAwsOrganizationalUnitAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAwsOrganizationalUnitAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAwsOrganizationalUnitAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAwsOrganizationalUnitAllOf(val *CloudAwsOrganizationalUnitAllOf) *NullableCloudAwsOrganizationalUnitAllOf {
	return &NullableCloudAwsOrganizationalUnitAllOf{value: val, isSet: true}
}

func (v NullableCloudAwsOrganizationalUnitAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAwsOrganizationalUnitAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
