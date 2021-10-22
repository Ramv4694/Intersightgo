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

// AssetOauthBearerTokenCredential An credential which performs authentication based on a bearer token.
type AssetOauthBearerTokenCredential struct {
	AssetCredential
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the value of the 'token' property has been set.
	IsTokenSet *bool `json:"IsTokenSet,omitempty"`
	// Scope type for the crendetial i.e. User, Organization, Team.
	ScopeType *string `json:"ScopeType,omitempty"`
	// Scope value for the credential i.e. username, organization name or team name.
	ScopeValue *string `json:"ScopeValue,omitempty"`
	// The token used to authenticate with a managed target.
	Token                *string `json:"Token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetOauthBearerTokenCredential AssetOauthBearerTokenCredential

// NewAssetOauthBearerTokenCredential instantiates a new AssetOauthBearerTokenCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetOauthBearerTokenCredential(classId string, objectType string) *AssetOauthBearerTokenCredential {
	this := AssetOauthBearerTokenCredential{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isTokenSet bool = false
	this.IsTokenSet = &isTokenSet
	return &this
}

// NewAssetOauthBearerTokenCredentialWithDefaults instantiates a new AssetOauthBearerTokenCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetOauthBearerTokenCredentialWithDefaults() *AssetOauthBearerTokenCredential {
	this := AssetOauthBearerTokenCredential{}
	var classId string = "asset.OauthBearerTokenCredential"
	this.ClassId = classId
	var objectType string = "asset.OauthBearerTokenCredential"
	this.ObjectType = objectType
	var isTokenSet bool = false
	this.IsTokenSet = &isTokenSet
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetOauthBearerTokenCredential) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetOauthBearerTokenCredential) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetOauthBearerTokenCredential) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetOauthBearerTokenCredential) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetOauthBearerTokenCredential) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetOauthBearerTokenCredential) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsTokenSet returns the IsTokenSet field value if set, zero value otherwise.
func (o *AssetOauthBearerTokenCredential) GetIsTokenSet() bool {
	if o == nil || o.IsTokenSet == nil {
		var ret bool
		return ret
	}
	return *o.IsTokenSet
}

// GetIsTokenSetOk returns a tuple with the IsTokenSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOauthBearerTokenCredential) GetIsTokenSetOk() (*bool, bool) {
	if o == nil || o.IsTokenSet == nil {
		return nil, false
	}
	return o.IsTokenSet, true
}

// HasIsTokenSet returns a boolean if a field has been set.
func (o *AssetOauthBearerTokenCredential) HasIsTokenSet() bool {
	if o != nil && o.IsTokenSet != nil {
		return true
	}

	return false
}

// SetIsTokenSet gets a reference to the given bool and assigns it to the IsTokenSet field.
func (o *AssetOauthBearerTokenCredential) SetIsTokenSet(v bool) {
	o.IsTokenSet = &v
}

// GetScopeType returns the ScopeType field value if set, zero value otherwise.
func (o *AssetOauthBearerTokenCredential) GetScopeType() string {
	if o == nil || o.ScopeType == nil {
		var ret string
		return ret
	}
	return *o.ScopeType
}

// GetScopeTypeOk returns a tuple with the ScopeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOauthBearerTokenCredential) GetScopeTypeOk() (*string, bool) {
	if o == nil || o.ScopeType == nil {
		return nil, false
	}
	return o.ScopeType, true
}

// HasScopeType returns a boolean if a field has been set.
func (o *AssetOauthBearerTokenCredential) HasScopeType() bool {
	if o != nil && o.ScopeType != nil {
		return true
	}

	return false
}

// SetScopeType gets a reference to the given string and assigns it to the ScopeType field.
func (o *AssetOauthBearerTokenCredential) SetScopeType(v string) {
	o.ScopeType = &v
}

// GetScopeValue returns the ScopeValue field value if set, zero value otherwise.
func (o *AssetOauthBearerTokenCredential) GetScopeValue() string {
	if o == nil || o.ScopeValue == nil {
		var ret string
		return ret
	}
	return *o.ScopeValue
}

// GetScopeValueOk returns a tuple with the ScopeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOauthBearerTokenCredential) GetScopeValueOk() (*string, bool) {
	if o == nil || o.ScopeValue == nil {
		return nil, false
	}
	return o.ScopeValue, true
}

// HasScopeValue returns a boolean if a field has been set.
func (o *AssetOauthBearerTokenCredential) HasScopeValue() bool {
	if o != nil && o.ScopeValue != nil {
		return true
	}

	return false
}

// SetScopeValue gets a reference to the given string and assigns it to the ScopeValue field.
func (o *AssetOauthBearerTokenCredential) SetScopeValue(v string) {
	o.ScopeValue = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AssetOauthBearerTokenCredential) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOauthBearerTokenCredential) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AssetOauthBearerTokenCredential) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AssetOauthBearerTokenCredential) SetToken(v string) {
	o.Token = &v
}

func (o AssetOauthBearerTokenCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetCredential, errAssetCredential := json.Marshal(o.AssetCredential)
	if errAssetCredential != nil {
		return []byte{}, errAssetCredential
	}
	errAssetCredential = json.Unmarshal([]byte(serializedAssetCredential), &toSerialize)
	if errAssetCredential != nil {
		return []byte{}, errAssetCredential
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IsTokenSet != nil {
		toSerialize["IsTokenSet"] = o.IsTokenSet
	}
	if o.ScopeType != nil {
		toSerialize["ScopeType"] = o.ScopeType
	}
	if o.ScopeValue != nil {
		toSerialize["ScopeValue"] = o.ScopeValue
	}
	if o.Token != nil {
		toSerialize["Token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetOauthBearerTokenCredential) UnmarshalJSON(bytes []byte) (err error) {
	type AssetOauthBearerTokenCredentialWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Indicates whether the value of the 'token' property has been set.
		IsTokenSet *bool `json:"IsTokenSet,omitempty"`
		// Scope type for the crendetial i.e. User, Organization, Team.
		ScopeType *string `json:"ScopeType,omitempty"`
		// Scope value for the credential i.e. username, organization name or team name.
		ScopeValue *string `json:"ScopeValue,omitempty"`
		// The token used to authenticate with a managed target.
		Token *string `json:"Token,omitempty"`
	}

	varAssetOauthBearerTokenCredentialWithoutEmbeddedStruct := AssetOauthBearerTokenCredentialWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetOauthBearerTokenCredentialWithoutEmbeddedStruct)
	if err == nil {
		varAssetOauthBearerTokenCredential := _AssetOauthBearerTokenCredential{}
		varAssetOauthBearerTokenCredential.ClassId = varAssetOauthBearerTokenCredentialWithoutEmbeddedStruct.ClassId
		varAssetOauthBearerTokenCredential.ObjectType = varAssetOauthBearerTokenCredentialWithoutEmbeddedStruct.ObjectType
		varAssetOauthBearerTokenCredential.IsTokenSet = varAssetOauthBearerTokenCredentialWithoutEmbeddedStruct.IsTokenSet
		varAssetOauthBearerTokenCredential.ScopeType = varAssetOauthBearerTokenCredentialWithoutEmbeddedStruct.ScopeType
		varAssetOauthBearerTokenCredential.ScopeValue = varAssetOauthBearerTokenCredentialWithoutEmbeddedStruct.ScopeValue
		varAssetOauthBearerTokenCredential.Token = varAssetOauthBearerTokenCredentialWithoutEmbeddedStruct.Token
		*o = AssetOauthBearerTokenCredential(varAssetOauthBearerTokenCredential)
	} else {
		return err
	}

	varAssetOauthBearerTokenCredential := _AssetOauthBearerTokenCredential{}

	err = json.Unmarshal(bytes, &varAssetOauthBearerTokenCredential)
	if err == nil {
		o.AssetCredential = varAssetOauthBearerTokenCredential.AssetCredential
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsTokenSet")
		delete(additionalProperties, "ScopeType")
		delete(additionalProperties, "ScopeValue")
		delete(additionalProperties, "Token")

		// remove fields from embedded structs
		reflectAssetCredential := reflect.ValueOf(o.AssetCredential)
		for i := 0; i < reflectAssetCredential.Type().NumField(); i++ {
			t := reflectAssetCredential.Type().Field(i)

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

type NullableAssetOauthBearerTokenCredential struct {
	value *AssetOauthBearerTokenCredential
	isSet bool
}

func (v NullableAssetOauthBearerTokenCredential) Get() *AssetOauthBearerTokenCredential {
	return v.value
}

func (v *NullableAssetOauthBearerTokenCredential) Set(val *AssetOauthBearerTokenCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetOauthBearerTokenCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetOauthBearerTokenCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetOauthBearerTokenCredential(val *AssetOauthBearerTokenCredential) *NullableAssetOauthBearerTokenCredential {
	return &NullableAssetOauthBearerTokenCredential{value: val, isSet: true}
}

func (v NullableAssetOauthBearerTokenCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetOauthBearerTokenCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
