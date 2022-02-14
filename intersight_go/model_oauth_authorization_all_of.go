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

// OauthAuthorizationAllOf Definition of the list of properties defined in 'oauth.Authorization', excluding properties defined in parent classes.
type OauthAuthorizationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of OAuth Api. For example, Smart-licensing-API. * `Unknown` - Unknown is the default API type. * `SmartLicensing-API` - Smart licensing API type.
	ApiType *string `json:"ApiType,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Indicates whether the value of the 'userId' property has been set.
	IsUserIdSet *bool `json:"IsUserIdSet,omitempty"`
	// The password that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account's behalf.
	Password *string `json:"Password,omitempty"`
	// The username that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account's behalf.
	UserId               *string                       `json:"UserId,omitempty"`
	AccessToken          *OauthAccessTokenRelationship `json:"AccessToken,omitempty"`
	Account              *IamAccountRelationship       `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OauthAuthorizationAllOf OauthAuthorizationAllOf

// NewOauthAuthorizationAllOf instantiates a new OauthAuthorizationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOauthAuthorizationAllOf(classId string, objectType string) *OauthAuthorizationAllOf {
	this := OauthAuthorizationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var apiType string = "Unknown"
	this.ApiType = &apiType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	var isUserIdSet bool = false
	this.IsUserIdSet = &isUserIdSet
	return &this
}

// NewOauthAuthorizationAllOfWithDefaults instantiates a new OauthAuthorizationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOauthAuthorizationAllOfWithDefaults() *OauthAuthorizationAllOf {
	this := OauthAuthorizationAllOf{}
	var classId string = "oauth.Authorization"
	this.ClassId = classId
	var objectType string = "oauth.Authorization"
	this.ObjectType = objectType
	var apiType string = "Unknown"
	this.ApiType = &apiType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	var isUserIdSet bool = false
	this.IsUserIdSet = &isUserIdSet
	return &this
}

// GetClassId returns the ClassId field value
func (o *OauthAuthorizationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OauthAuthorizationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OauthAuthorizationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OauthAuthorizationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OauthAuthorizationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OauthAuthorizationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetApiType returns the ApiType field value if set, zero value otherwise.
func (o *OauthAuthorizationAllOf) GetApiType() string {
	if o == nil || o.ApiType == nil {
		var ret string
		return ret
	}
	return *o.ApiType
}

// GetApiTypeOk returns a tuple with the ApiType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthAuthorizationAllOf) GetApiTypeOk() (*string, bool) {
	if o == nil || o.ApiType == nil {
		return nil, false
	}
	return o.ApiType, true
}

// HasApiType returns a boolean if a field has been set.
func (o *OauthAuthorizationAllOf) HasApiType() bool {
	if o != nil && o.ApiType != nil {
		return true
	}

	return false
}

// SetApiType gets a reference to the given string and assigns it to the ApiType field.
func (o *OauthAuthorizationAllOf) SetApiType(v string) {
	o.ApiType = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *OauthAuthorizationAllOf) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthAuthorizationAllOf) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *OauthAuthorizationAllOf) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *OauthAuthorizationAllOf) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetIsUserIdSet returns the IsUserIdSet field value if set, zero value otherwise.
func (o *OauthAuthorizationAllOf) GetIsUserIdSet() bool {
	if o == nil || o.IsUserIdSet == nil {
		var ret bool
		return ret
	}
	return *o.IsUserIdSet
}

// GetIsUserIdSetOk returns a tuple with the IsUserIdSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthAuthorizationAllOf) GetIsUserIdSetOk() (*bool, bool) {
	if o == nil || o.IsUserIdSet == nil {
		return nil, false
	}
	return o.IsUserIdSet, true
}

// HasIsUserIdSet returns a boolean if a field has been set.
func (o *OauthAuthorizationAllOf) HasIsUserIdSet() bool {
	if o != nil && o.IsUserIdSet != nil {
		return true
	}

	return false
}

// SetIsUserIdSet gets a reference to the given bool and assigns it to the IsUserIdSet field.
func (o *OauthAuthorizationAllOf) SetIsUserIdSet(v bool) {
	o.IsUserIdSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *OauthAuthorizationAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthAuthorizationAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *OauthAuthorizationAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *OauthAuthorizationAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *OauthAuthorizationAllOf) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthAuthorizationAllOf) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *OauthAuthorizationAllOf) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *OauthAuthorizationAllOf) SetUserId(v string) {
	o.UserId = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *OauthAuthorizationAllOf) GetAccessToken() OauthAccessTokenRelationship {
	if o == nil || o.AccessToken == nil {
		var ret OauthAccessTokenRelationship
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthAuthorizationAllOf) GetAccessTokenOk() (*OauthAccessTokenRelationship, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *OauthAuthorizationAllOf) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given OauthAccessTokenRelationship and assigns it to the AccessToken field.
func (o *OauthAuthorizationAllOf) SetAccessToken(v OauthAccessTokenRelationship) {
	o.AccessToken = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *OauthAuthorizationAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthAuthorizationAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *OauthAuthorizationAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *OauthAuthorizationAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o OauthAuthorizationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ApiType != nil {
		toSerialize["ApiType"] = o.ApiType
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.IsUserIdSet != nil {
		toSerialize["IsUserIdSet"] = o.IsUserIdSet
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.UserId != nil {
		toSerialize["UserId"] = o.UserId
	}
	if o.AccessToken != nil {
		toSerialize["AccessToken"] = o.AccessToken
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OauthAuthorizationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOauthAuthorizationAllOf := _OauthAuthorizationAllOf{}

	if err = json.Unmarshal(bytes, &varOauthAuthorizationAllOf); err == nil {
		*o = OauthAuthorizationAllOf(varOauthAuthorizationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ApiType")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "IsUserIdSet")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "UserId")
		delete(additionalProperties, "AccessToken")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOauthAuthorizationAllOf struct {
	value *OauthAuthorizationAllOf
	isSet bool
}

func (v NullableOauthAuthorizationAllOf) Get() *OauthAuthorizationAllOf {
	return v.value
}

func (v *NullableOauthAuthorizationAllOf) Set(val *OauthAuthorizationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOauthAuthorizationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOauthAuthorizationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOauthAuthorizationAllOf(val *OauthAuthorizationAllOf) *NullableOauthAuthorizationAllOf {
	return &NullableOauthAuthorizationAllOf{value: val, isSet: true}
}

func (v NullableOauthAuthorizationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOauthAuthorizationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
