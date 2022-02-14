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
	"time"
)

// IamUserAllOf Definition of the list of properties defined in 'iam.User', excluding properties defined in parent classes.
type IamUserAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IP address from which the user last logged in to Intersight.
	ClientIpAddress *string `json:"ClientIpAddress,omitempty"`
	// Email of the user. Users are added to Intersight using the email configured in the IdP.
	Email *string `json:"Email,omitempty"`
	// First name of the user. This field is populated from the IdP attributes received after authentication.
	FirstName *string `json:"FirstName,omitempty"`
	// Last successful login time for user.
	LastLoginTime *time.Time `json:"LastLoginTime,omitempty"`
	// Last name of the user. This field is populated from the IdP attributes received after authentication.
	LastName *string `json:"LastName,omitempty"`
	// Last role modification time for user.
	LastRoleModifiedTime *time.Time `json:"LastRoleModifiedTime,omitempty"`
	// Name as configured in the IdP.
	Name *string `json:"Name,omitempty"`
	// UserID or email as configured in the IdP.
	UserIdOrEmail *string `json:"UserIdOrEmail,omitempty"`
	// Type of the User. If a user is added manually by specifying the email address, or has logged in using groups, based on the IdP attributes received during authentication. If added manually, the user type will be static, otherwise dynamic.
	UserType *string `json:"UserType,omitempty"`
	// Unique id of the user used by the identity provider to store the user.
	UserUniqueIdentifier *string `json:"UserUniqueIdentifier,omitempty"`
	// An array of relationships to iamApiKey resources.
	ApiKeys []IamApiKeyRelationship `json:"ApiKeys,omitempty"`
	// An array of relationships to iamAppRegistration resources.
	AppRegistrations  []IamAppRegistrationRelationship  `json:"AppRegistrations,omitempty"`
	Idp               *IamIdpRelationship               `json:"Idp,omitempty"`
	Idpreference      *IamIdpReferenceRelationship      `json:"Idpreference,omitempty"`
	LocalUserPassword *IamLocalUserPasswordRelationship `json:"LocalUserPassword,omitempty"`
	// An array of relationships to iamOAuthToken resources.
	OauthTokens []IamOAuthTokenRelationship `json:"OauthTokens,omitempty"`
	// An array of relationships to iamPermission resources.
	Permissions []IamPermissionRelationship `json:"Permissions,omitempty"`
	// An array of relationships to iamSession resources.
	Sessions             []IamSessionRelationship `json:"Sessions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamUserAllOf IamUserAllOf

// NewIamUserAllOf instantiates a new IamUserAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamUserAllOf(classId string, objectType string) *IamUserAllOf {
	this := IamUserAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamUserAllOfWithDefaults instantiates a new IamUserAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamUserAllOfWithDefaults() *IamUserAllOf {
	this := IamUserAllOf{}
	var classId string = "iam.User"
	this.ClassId = classId
	var objectType string = "iam.User"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamUserAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamUserAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamUserAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamUserAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamUserAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamUserAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClientIpAddress returns the ClientIpAddress field value if set, zero value otherwise.
func (o *IamUserAllOf) GetClientIpAddress() string {
	if o == nil || o.ClientIpAddress == nil {
		var ret string
		return ret
	}
	return *o.ClientIpAddress
}

// GetClientIpAddressOk returns a tuple with the ClientIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserAllOf) GetClientIpAddressOk() (*string, bool) {
	if o == nil || o.ClientIpAddress == nil {
		return nil, false
	}
	return o.ClientIpAddress, true
}

// HasClientIpAddress returns a boolean if a field has been set.
func (o *IamUserAllOf) HasClientIpAddress() bool {
	if o != nil && o.ClientIpAddress != nil {
		return true
	}

	return false
}

// SetClientIpAddress gets a reference to the given string and assigns it to the ClientIpAddress field.
func (o *IamUserAllOf) SetClientIpAddress(v string) {
	o.ClientIpAddress = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *IamUserAllOf) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserAllOf) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *IamUserAllOf) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *IamUserAllOf) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *IamUserAllOf) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserAllOf) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *IamUserAllOf) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *IamUserAllOf) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastLoginTime returns the LastLoginTime field value if set, zero value otherwise.
func (o *IamUserAllOf) GetLastLoginTime() time.Time {
	if o == nil || o.LastLoginTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastLoginTime
}

// GetLastLoginTimeOk returns a tuple with the LastLoginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserAllOf) GetLastLoginTimeOk() (*time.Time, bool) {
	if o == nil || o.LastLoginTime == nil {
		return nil, false
	}
	return o.LastLoginTime, true
}

// HasLastLoginTime returns a boolean if a field has been set.
func (o *IamUserAllOf) HasLastLoginTime() bool {
	if o != nil && o.LastLoginTime != nil {
		return true
	}

	return false
}

// SetLastLoginTime gets a reference to the given time.Time and assigns it to the LastLoginTime field.
func (o *IamUserAllOf) SetLastLoginTime(v time.Time) {
	o.LastLoginTime = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *IamUserAllOf) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserAllOf) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *IamUserAllOf) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *IamUserAllOf) SetLastName(v string) {
	o.LastName = &v
}

// GetLastRoleModifiedTime returns the LastRoleModifiedTime field value if set, zero value otherwise.
func (o *IamUserAllOf) GetLastRoleModifiedTime() time.Time {
	if o == nil || o.LastRoleModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastRoleModifiedTime
}

// GetLastRoleModifiedTimeOk returns a tuple with the LastRoleModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserAllOf) GetLastRoleModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastRoleModifiedTime == nil {
		return nil, false
	}
	return o.LastRoleModifiedTime, true
}

// HasLastRoleModifiedTime returns a boolean if a field has been set.
func (o *IamUserAllOf) HasLastRoleModifiedTime() bool {
	if o != nil && o.LastRoleModifiedTime != nil {
		return true
	}

	return false
}

// SetLastRoleModifiedTime gets a reference to the given time.Time and assigns it to the LastRoleModifiedTime field.
func (o *IamUserAllOf) SetLastRoleModifiedTime(v time.Time) {
	o.LastRoleModifiedTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamUserAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamUserAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamUserAllOf) SetName(v string) {
	o.Name = &v
}

// GetUserIdOrEmail returns the UserIdOrEmail field value if set, zero value otherwise.
func (o *IamUserAllOf) GetUserIdOrEmail() string {
	if o == nil || o.UserIdOrEmail == nil {
		var ret string
		return ret
	}
	return *o.UserIdOrEmail
}

// GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserAllOf) GetUserIdOrEmailOk() (*string, bool) {
	if o == nil || o.UserIdOrEmail == nil {
		return nil, false
	}
	return o.UserIdOrEmail, true
}

// HasUserIdOrEmail returns a boolean if a field has been set.
func (o *IamUserAllOf) HasUserIdOrEmail() bool {
	if o != nil && o.UserIdOrEmail != nil {
		return true
	}

	return false
}

// SetUserIdOrEmail gets a reference to the given string and assigns it to the UserIdOrEmail field.
func (o *IamUserAllOf) SetUserIdOrEmail(v string) {
	o.UserIdOrEmail = &v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *IamUserAllOf) GetUserType() string {
	if o == nil || o.UserType == nil {
		var ret string
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserAllOf) GetUserTypeOk() (*string, bool) {
	if o == nil || o.UserType == nil {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *IamUserAllOf) HasUserType() bool {
	if o != nil && o.UserType != nil {
		return true
	}

	return false
}

// SetUserType gets a reference to the given string and assigns it to the UserType field.
func (o *IamUserAllOf) SetUserType(v string) {
	o.UserType = &v
}

// GetUserUniqueIdentifier returns the UserUniqueIdentifier field value if set, zero value otherwise.
func (o *IamUserAllOf) GetUserUniqueIdentifier() string {
	if o == nil || o.UserUniqueIdentifier == nil {
		var ret string
		return ret
	}
	return *o.UserUniqueIdentifier
}

// GetUserUniqueIdentifierOk returns a tuple with the UserUniqueIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserAllOf) GetUserUniqueIdentifierOk() (*string, bool) {
	if o == nil || o.UserUniqueIdentifier == nil {
		return nil, false
	}
	return o.UserUniqueIdentifier, true
}

// HasUserUniqueIdentifier returns a boolean if a field has been set.
func (o *IamUserAllOf) HasUserUniqueIdentifier() bool {
	if o != nil && o.UserUniqueIdentifier != nil {
		return true
	}

	return false
}

// SetUserUniqueIdentifier gets a reference to the given string and assigns it to the UserUniqueIdentifier field.
func (o *IamUserAllOf) SetUserUniqueIdentifier(v string) {
	o.UserUniqueIdentifier = &v
}

// GetApiKeys returns the ApiKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUserAllOf) GetApiKeys() []IamApiKeyRelationship {
	if o == nil {
		var ret []IamApiKeyRelationship
		return ret
	}
	return o.ApiKeys
}

// GetApiKeysOk returns a tuple with the ApiKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUserAllOf) GetApiKeysOk() (*[]IamApiKeyRelationship, bool) {
	if o == nil || o.ApiKeys == nil {
		return nil, false
	}
	return &o.ApiKeys, true
}

// HasApiKeys returns a boolean if a field has been set.
func (o *IamUserAllOf) HasApiKeys() bool {
	if o != nil && o.ApiKeys != nil {
		return true
	}

	return false
}

// SetApiKeys gets a reference to the given []IamApiKeyRelationship and assigns it to the ApiKeys field.
func (o *IamUserAllOf) SetApiKeys(v []IamApiKeyRelationship) {
	o.ApiKeys = v
}

// GetAppRegistrations returns the AppRegistrations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUserAllOf) GetAppRegistrations() []IamAppRegistrationRelationship {
	if o == nil {
		var ret []IamAppRegistrationRelationship
		return ret
	}
	return o.AppRegistrations
}

// GetAppRegistrationsOk returns a tuple with the AppRegistrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUserAllOf) GetAppRegistrationsOk() (*[]IamAppRegistrationRelationship, bool) {
	if o == nil || o.AppRegistrations == nil {
		return nil, false
	}
	return &o.AppRegistrations, true
}

// HasAppRegistrations returns a boolean if a field has been set.
func (o *IamUserAllOf) HasAppRegistrations() bool {
	if o != nil && o.AppRegistrations != nil {
		return true
	}

	return false
}

// SetAppRegistrations gets a reference to the given []IamAppRegistrationRelationship and assigns it to the AppRegistrations field.
func (o *IamUserAllOf) SetAppRegistrations(v []IamAppRegistrationRelationship) {
	o.AppRegistrations = v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *IamUserAllOf) GetIdp() IamIdpRelationship {
	if o == nil || o.Idp == nil {
		var ret IamIdpRelationship
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserAllOf) GetIdpOk() (*IamIdpRelationship, bool) {
	if o == nil || o.Idp == nil {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *IamUserAllOf) HasIdp() bool {
	if o != nil && o.Idp != nil {
		return true
	}

	return false
}

// SetIdp gets a reference to the given IamIdpRelationship and assigns it to the Idp field.
func (o *IamUserAllOf) SetIdp(v IamIdpRelationship) {
	o.Idp = &v
}

// GetIdpreference returns the Idpreference field value if set, zero value otherwise.
func (o *IamUserAllOf) GetIdpreference() IamIdpReferenceRelationship {
	if o == nil || o.Idpreference == nil {
		var ret IamIdpReferenceRelationship
		return ret
	}
	return *o.Idpreference
}

// GetIdpreferenceOk returns a tuple with the Idpreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserAllOf) GetIdpreferenceOk() (*IamIdpReferenceRelationship, bool) {
	if o == nil || o.Idpreference == nil {
		return nil, false
	}
	return o.Idpreference, true
}

// HasIdpreference returns a boolean if a field has been set.
func (o *IamUserAllOf) HasIdpreference() bool {
	if o != nil && o.Idpreference != nil {
		return true
	}

	return false
}

// SetIdpreference gets a reference to the given IamIdpReferenceRelationship and assigns it to the Idpreference field.
func (o *IamUserAllOf) SetIdpreference(v IamIdpReferenceRelationship) {
	o.Idpreference = &v
}

// GetLocalUserPassword returns the LocalUserPassword field value if set, zero value otherwise.
func (o *IamUserAllOf) GetLocalUserPassword() IamLocalUserPasswordRelationship {
	if o == nil || o.LocalUserPassword == nil {
		var ret IamLocalUserPasswordRelationship
		return ret
	}
	return *o.LocalUserPassword
}

// GetLocalUserPasswordOk returns a tuple with the LocalUserPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserAllOf) GetLocalUserPasswordOk() (*IamLocalUserPasswordRelationship, bool) {
	if o == nil || o.LocalUserPassword == nil {
		return nil, false
	}
	return o.LocalUserPassword, true
}

// HasLocalUserPassword returns a boolean if a field has been set.
func (o *IamUserAllOf) HasLocalUserPassword() bool {
	if o != nil && o.LocalUserPassword != nil {
		return true
	}

	return false
}

// SetLocalUserPassword gets a reference to the given IamLocalUserPasswordRelationship and assigns it to the LocalUserPassword field.
func (o *IamUserAllOf) SetLocalUserPassword(v IamLocalUserPasswordRelationship) {
	o.LocalUserPassword = &v
}

// GetOauthTokens returns the OauthTokens field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUserAllOf) GetOauthTokens() []IamOAuthTokenRelationship {
	if o == nil {
		var ret []IamOAuthTokenRelationship
		return ret
	}
	return o.OauthTokens
}

// GetOauthTokensOk returns a tuple with the OauthTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUserAllOf) GetOauthTokensOk() (*[]IamOAuthTokenRelationship, bool) {
	if o == nil || o.OauthTokens == nil {
		return nil, false
	}
	return &o.OauthTokens, true
}

// HasOauthTokens returns a boolean if a field has been set.
func (o *IamUserAllOf) HasOauthTokens() bool {
	if o != nil && o.OauthTokens != nil {
		return true
	}

	return false
}

// SetOauthTokens gets a reference to the given []IamOAuthTokenRelationship and assigns it to the OauthTokens field.
func (o *IamUserAllOf) SetOauthTokens(v []IamOAuthTokenRelationship) {
	o.OauthTokens = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUserAllOf) GetPermissions() []IamPermissionRelationship {
	if o == nil {
		var ret []IamPermissionRelationship
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUserAllOf) GetPermissionsOk() (*[]IamPermissionRelationship, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *IamUserAllOf) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []IamPermissionRelationship and assigns it to the Permissions field.
func (o *IamUserAllOf) SetPermissions(v []IamPermissionRelationship) {
	o.Permissions = v
}

// GetSessions returns the Sessions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUserAllOf) GetSessions() []IamSessionRelationship {
	if o == nil {
		var ret []IamSessionRelationship
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUserAllOf) GetSessionsOk() (*[]IamSessionRelationship, bool) {
	if o == nil || o.Sessions == nil {
		return nil, false
	}
	return &o.Sessions, true
}

// HasSessions returns a boolean if a field has been set.
func (o *IamUserAllOf) HasSessions() bool {
	if o != nil && o.Sessions != nil {
		return true
	}

	return false
}

// SetSessions gets a reference to the given []IamSessionRelationship and assigns it to the Sessions field.
func (o *IamUserAllOf) SetSessions(v []IamSessionRelationship) {
	o.Sessions = v
}

func (o IamUserAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClientIpAddress != nil {
		toSerialize["ClientIpAddress"] = o.ClientIpAddress
	}
	if o.Email != nil {
		toSerialize["Email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["FirstName"] = o.FirstName
	}
	if o.LastLoginTime != nil {
		toSerialize["LastLoginTime"] = o.LastLoginTime
	}
	if o.LastName != nil {
		toSerialize["LastName"] = o.LastName
	}
	if o.LastRoleModifiedTime != nil {
		toSerialize["LastRoleModifiedTime"] = o.LastRoleModifiedTime
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.UserIdOrEmail != nil {
		toSerialize["UserIdOrEmail"] = o.UserIdOrEmail
	}
	if o.UserType != nil {
		toSerialize["UserType"] = o.UserType
	}
	if o.UserUniqueIdentifier != nil {
		toSerialize["UserUniqueIdentifier"] = o.UserUniqueIdentifier
	}
	if o.ApiKeys != nil {
		toSerialize["ApiKeys"] = o.ApiKeys
	}
	if o.AppRegistrations != nil {
		toSerialize["AppRegistrations"] = o.AppRegistrations
	}
	if o.Idp != nil {
		toSerialize["Idp"] = o.Idp
	}
	if o.Idpreference != nil {
		toSerialize["Idpreference"] = o.Idpreference
	}
	if o.LocalUserPassword != nil {
		toSerialize["LocalUserPassword"] = o.LocalUserPassword
	}
	if o.OauthTokens != nil {
		toSerialize["OauthTokens"] = o.OauthTokens
	}
	if o.Permissions != nil {
		toSerialize["Permissions"] = o.Permissions
	}
	if o.Sessions != nil {
		toSerialize["Sessions"] = o.Sessions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamUserAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamUserAllOf := _IamUserAllOf{}

	if err = json.Unmarshal(bytes, &varIamUserAllOf); err == nil {
		*o = IamUserAllOf(varIamUserAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClientIpAddress")
		delete(additionalProperties, "Email")
		delete(additionalProperties, "FirstName")
		delete(additionalProperties, "LastLoginTime")
		delete(additionalProperties, "LastName")
		delete(additionalProperties, "LastRoleModifiedTime")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "UserIdOrEmail")
		delete(additionalProperties, "UserType")
		delete(additionalProperties, "UserUniqueIdentifier")
		delete(additionalProperties, "ApiKeys")
		delete(additionalProperties, "AppRegistrations")
		delete(additionalProperties, "Idp")
		delete(additionalProperties, "Idpreference")
		delete(additionalProperties, "LocalUserPassword")
		delete(additionalProperties, "OauthTokens")
		delete(additionalProperties, "Permissions")
		delete(additionalProperties, "Sessions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamUserAllOf struct {
	value *IamUserAllOf
	isSet bool
}

func (v NullableIamUserAllOf) Get() *IamUserAllOf {
	return v.value
}

func (v *NullableIamUserAllOf) Set(val *IamUserAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamUserAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamUserAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamUserAllOf(val *IamUserAllOf) *NullableIamUserAllOf {
	return &NullableIamUserAllOf{value: val, isSet: true}
}

func (v NullableIamUserAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamUserAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
