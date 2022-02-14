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
	"reflect"
	"strings"
)

// IamIdp The SAML identity provider such as Cisco, that has been used to log in to Intersight.
type IamIdp struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Email domain name of the user for this IdP. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication.
	DomainName *string `json:"DomainName,omitempty"`
	// Setting that indicates whether 'Single Logout (SLO)' has been enabled for this IdP.
	EnableSingleLogout *bool `json:"EnableSingleLogout,omitempty"`
	// The Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP or Service Provider.
	IdpEntityId *string `json:"IdpEntityId,omitempty"`
	// SAML metadata of the IdP.
	Metadata *string `json:"Metadata,omitempty"`
	// The name of the Identity Provider, for example Cisco, Okta, or OneID.
	Name *string `json:"Name,omitempty"`
	// Authentication protocol used by the IdP. * `saml` - Use SAML as the authentication protocol for sign-on. * `oidc` - Open ID connect to be used as an authentication protocol for sign-on. * `local` - The local authentication method to be used for sign-on. Local type is set to default for the Intersight Appliance IdP.
	Type       *string                    `json:"Type,omitempty"`
	Account    *IamAccountRelationship    `json:"Account,omitempty"`
	LdapPolicy *IamLdapPolicyRelationship `json:"LdapPolicy,omitempty"`
	System     *IamSystemRelationship     `json:"System,omitempty"`
	// An array of relationships to iamUserPreference resources.
	UserPreferences []IamUserPreferenceRelationship `json:"UserPreferences,omitempty"`
	// An array of relationships to iamUserGroup resources.
	Usergroups []IamUserGroupRelationship `json:"Usergroups,omitempty"`
	// An array of relationships to iamUser resources.
	Users                []IamUserRelationship `json:"Users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamIdp IamIdp

// NewIamIdp instantiates a new IamIdp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamIdp(classId string, objectType string) *IamIdp {
	this := IamIdp{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "saml"
	this.Type = &type_
	return &this
}

// NewIamIdpWithDefaults instantiates a new IamIdp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamIdpWithDefaults() *IamIdp {
	this := IamIdp{}
	var classId string = "iam.Idp"
	this.ClassId = classId
	var objectType string = "iam.Idp"
	this.ObjectType = objectType
	var type_ string = "saml"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamIdp) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamIdp) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamIdp) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamIdp) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamIdp) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamIdp) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *IamIdp) GetDomainName() string {
	if o == nil || o.DomainName == nil {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdp) GetDomainNameOk() (*string, bool) {
	if o == nil || o.DomainName == nil {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *IamIdp) HasDomainName() bool {
	if o != nil && o.DomainName != nil {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *IamIdp) SetDomainName(v string) {
	o.DomainName = &v
}

// GetEnableSingleLogout returns the EnableSingleLogout field value if set, zero value otherwise.
func (o *IamIdp) GetEnableSingleLogout() bool {
	if o == nil || o.EnableSingleLogout == nil {
		var ret bool
		return ret
	}
	return *o.EnableSingleLogout
}

// GetEnableSingleLogoutOk returns a tuple with the EnableSingleLogout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdp) GetEnableSingleLogoutOk() (*bool, bool) {
	if o == nil || o.EnableSingleLogout == nil {
		return nil, false
	}
	return o.EnableSingleLogout, true
}

// HasEnableSingleLogout returns a boolean if a field has been set.
func (o *IamIdp) HasEnableSingleLogout() bool {
	if o != nil && o.EnableSingleLogout != nil {
		return true
	}

	return false
}

// SetEnableSingleLogout gets a reference to the given bool and assigns it to the EnableSingleLogout field.
func (o *IamIdp) SetEnableSingleLogout(v bool) {
	o.EnableSingleLogout = &v
}

// GetIdpEntityId returns the IdpEntityId field value if set, zero value otherwise.
func (o *IamIdp) GetIdpEntityId() string {
	if o == nil || o.IdpEntityId == nil {
		var ret string
		return ret
	}
	return *o.IdpEntityId
}

// GetIdpEntityIdOk returns a tuple with the IdpEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdp) GetIdpEntityIdOk() (*string, bool) {
	if o == nil || o.IdpEntityId == nil {
		return nil, false
	}
	return o.IdpEntityId, true
}

// HasIdpEntityId returns a boolean if a field has been set.
func (o *IamIdp) HasIdpEntityId() bool {
	if o != nil && o.IdpEntityId != nil {
		return true
	}

	return false
}

// SetIdpEntityId gets a reference to the given string and assigns it to the IdpEntityId field.
func (o *IamIdp) SetIdpEntityId(v string) {
	o.IdpEntityId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IamIdp) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdp) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IamIdp) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *IamIdp) SetMetadata(v string) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamIdp) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdp) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamIdp) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamIdp) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IamIdp) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdp) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IamIdp) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IamIdp) SetType(v string) {
	o.Type = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamIdp) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdp) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamIdp) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamIdp) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetLdapPolicy returns the LdapPolicy field value if set, zero value otherwise.
func (o *IamIdp) GetLdapPolicy() IamLdapPolicyRelationship {
	if o == nil || o.LdapPolicy == nil {
		var ret IamLdapPolicyRelationship
		return ret
	}
	return *o.LdapPolicy
}

// GetLdapPolicyOk returns a tuple with the LdapPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdp) GetLdapPolicyOk() (*IamLdapPolicyRelationship, bool) {
	if o == nil || o.LdapPolicy == nil {
		return nil, false
	}
	return o.LdapPolicy, true
}

// HasLdapPolicy returns a boolean if a field has been set.
func (o *IamIdp) HasLdapPolicy() bool {
	if o != nil && o.LdapPolicy != nil {
		return true
	}

	return false
}

// SetLdapPolicy gets a reference to the given IamLdapPolicyRelationship and assigns it to the LdapPolicy field.
func (o *IamIdp) SetLdapPolicy(v IamLdapPolicyRelationship) {
	o.LdapPolicy = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *IamIdp) GetSystem() IamSystemRelationship {
	if o == nil || o.System == nil {
		var ret IamSystemRelationship
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdp) GetSystemOk() (*IamSystemRelationship, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *IamIdp) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given IamSystemRelationship and assigns it to the System field.
func (o *IamIdp) SetSystem(v IamSystemRelationship) {
	o.System = &v
}

// GetUserPreferences returns the UserPreferences field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamIdp) GetUserPreferences() []IamUserPreferenceRelationship {
	if o == nil {
		var ret []IamUserPreferenceRelationship
		return ret
	}
	return o.UserPreferences
}

// GetUserPreferencesOk returns a tuple with the UserPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamIdp) GetUserPreferencesOk() (*[]IamUserPreferenceRelationship, bool) {
	if o == nil || o.UserPreferences == nil {
		return nil, false
	}
	return &o.UserPreferences, true
}

// HasUserPreferences returns a boolean if a field has been set.
func (o *IamIdp) HasUserPreferences() bool {
	if o != nil && o.UserPreferences != nil {
		return true
	}

	return false
}

// SetUserPreferences gets a reference to the given []IamUserPreferenceRelationship and assigns it to the UserPreferences field.
func (o *IamIdp) SetUserPreferences(v []IamUserPreferenceRelationship) {
	o.UserPreferences = v
}

// GetUsergroups returns the Usergroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamIdp) GetUsergroups() []IamUserGroupRelationship {
	if o == nil {
		var ret []IamUserGroupRelationship
		return ret
	}
	return o.Usergroups
}

// GetUsergroupsOk returns a tuple with the Usergroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamIdp) GetUsergroupsOk() (*[]IamUserGroupRelationship, bool) {
	if o == nil || o.Usergroups == nil {
		return nil, false
	}
	return &o.Usergroups, true
}

// HasUsergroups returns a boolean if a field has been set.
func (o *IamIdp) HasUsergroups() bool {
	if o != nil && o.Usergroups != nil {
		return true
	}

	return false
}

// SetUsergroups gets a reference to the given []IamUserGroupRelationship and assigns it to the Usergroups field.
func (o *IamIdp) SetUsergroups(v []IamUserGroupRelationship) {
	o.Usergroups = v
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamIdp) GetUsers() []IamUserRelationship {
	if o == nil {
		var ret []IamUserRelationship
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamIdp) GetUsersOk() (*[]IamUserRelationship, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return &o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *IamIdp) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []IamUserRelationship and assigns it to the Users field.
func (o *IamIdp) SetUsers(v []IamUserRelationship) {
	o.Users = v
}

func (o IamIdp) MarshalJSON() ([]byte, error) {
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
	if o.DomainName != nil {
		toSerialize["DomainName"] = o.DomainName
	}
	if o.EnableSingleLogout != nil {
		toSerialize["EnableSingleLogout"] = o.EnableSingleLogout
	}
	if o.IdpEntityId != nil {
		toSerialize["IdpEntityId"] = o.IdpEntityId
	}
	if o.Metadata != nil {
		toSerialize["Metadata"] = o.Metadata
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.LdapPolicy != nil {
		toSerialize["LdapPolicy"] = o.LdapPolicy
	}
	if o.System != nil {
		toSerialize["System"] = o.System
	}
	if o.UserPreferences != nil {
		toSerialize["UserPreferences"] = o.UserPreferences
	}
	if o.Usergroups != nil {
		toSerialize["Usergroups"] = o.Usergroups
	}
	if o.Users != nil {
		toSerialize["Users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamIdp) UnmarshalJSON(bytes []byte) (err error) {
	type IamIdpWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Email domain name of the user for this IdP. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication.
		DomainName *string `json:"DomainName,omitempty"`
		// Setting that indicates whether 'Single Logout (SLO)' has been enabled for this IdP.
		EnableSingleLogout *bool `json:"EnableSingleLogout,omitempty"`
		// The Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP or Service Provider.
		IdpEntityId *string `json:"IdpEntityId,omitempty"`
		// SAML metadata of the IdP.
		Metadata *string `json:"Metadata,omitempty"`
		// The name of the Identity Provider, for example Cisco, Okta, or OneID.
		Name *string `json:"Name,omitempty"`
		// Authentication protocol used by the IdP. * `saml` - Use SAML as the authentication protocol for sign-on. * `oidc` - Open ID connect to be used as an authentication protocol for sign-on. * `local` - The local authentication method to be used for sign-on. Local type is set to default for the Intersight Appliance IdP.
		Type       *string                    `json:"Type,omitempty"`
		Account    *IamAccountRelationship    `json:"Account,omitempty"`
		LdapPolicy *IamLdapPolicyRelationship `json:"LdapPolicy,omitempty"`
		System     *IamSystemRelationship     `json:"System,omitempty"`
		// An array of relationships to iamUserPreference resources.
		UserPreferences []IamUserPreferenceRelationship `json:"UserPreferences,omitempty"`
		// An array of relationships to iamUserGroup resources.
		Usergroups []IamUserGroupRelationship `json:"Usergroups,omitempty"`
		// An array of relationships to iamUser resources.
		Users []IamUserRelationship `json:"Users,omitempty"`
	}

	varIamIdpWithoutEmbeddedStruct := IamIdpWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamIdpWithoutEmbeddedStruct)
	if err == nil {
		varIamIdp := _IamIdp{}
		varIamIdp.ClassId = varIamIdpWithoutEmbeddedStruct.ClassId
		varIamIdp.ObjectType = varIamIdpWithoutEmbeddedStruct.ObjectType
		varIamIdp.DomainName = varIamIdpWithoutEmbeddedStruct.DomainName
		varIamIdp.EnableSingleLogout = varIamIdpWithoutEmbeddedStruct.EnableSingleLogout
		varIamIdp.IdpEntityId = varIamIdpWithoutEmbeddedStruct.IdpEntityId
		varIamIdp.Metadata = varIamIdpWithoutEmbeddedStruct.Metadata
		varIamIdp.Name = varIamIdpWithoutEmbeddedStruct.Name
		varIamIdp.Type = varIamIdpWithoutEmbeddedStruct.Type
		varIamIdp.Account = varIamIdpWithoutEmbeddedStruct.Account
		varIamIdp.LdapPolicy = varIamIdpWithoutEmbeddedStruct.LdapPolicy
		varIamIdp.System = varIamIdpWithoutEmbeddedStruct.System
		varIamIdp.UserPreferences = varIamIdpWithoutEmbeddedStruct.UserPreferences
		varIamIdp.Usergroups = varIamIdpWithoutEmbeddedStruct.Usergroups
		varIamIdp.Users = varIamIdpWithoutEmbeddedStruct.Users
		*o = IamIdp(varIamIdp)
	} else {
		return err
	}

	varIamIdp := _IamIdp{}

	err = json.Unmarshal(bytes, &varIamIdp)
	if err == nil {
		o.MoBaseMo = varIamIdp.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DomainName")
		delete(additionalProperties, "EnableSingleLogout")
		delete(additionalProperties, "IdpEntityId")
		delete(additionalProperties, "Metadata")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "LdapPolicy")
		delete(additionalProperties, "System")
		delete(additionalProperties, "UserPreferences")
		delete(additionalProperties, "Usergroups")
		delete(additionalProperties, "Users")

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

type NullableIamIdp struct {
	value *IamIdp
	isSet bool
}

func (v NullableIamIdp) Get() *IamIdp {
	return v.value
}

func (v *NullableIamIdp) Set(val *IamIdp) {
	v.value = val
	v.isSet = true
}

func (v NullableIamIdp) IsSet() bool {
	return v.isSet
}

func (v *NullableIamIdp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamIdp(val *IamIdp) *NullableIamIdp {
	return &NullableIamIdp{value: val, isSet: true}
}

func (v NullableIamIdp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamIdp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
