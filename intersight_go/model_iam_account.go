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

// IamAccount The Intersight Account used to access Intersight.
type IamAccount struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the Intersight account. By default, name is same as the MoID of the account.
	Name *string `json:"Name,omitempty"`
	// Status of the account. To activate the Intersight account, claim a device to the account.
	Status *string `json:"Status,omitempty"`
	// An array of relationships to iamAppRegistration resources.
	AppRegistrations []IamAppRegistrationRelationship `json:"AppRegistrations,omitempty"`
	// An array of relationships to iamDomainGroup resources.
	DomainGroups []IamDomainGroupRelationship `json:"DomainGroups,omitempty"`
	// An array of relationships to iamEndPointRole resources.
	EndPointRoles []IamEndPointRoleRelationship `json:"EndPointRoles,omitempty"`
	// An array of relationships to iamIdpReference resources.
	Idpreferences []IamIdpReferenceRelationship `json:"Idpreferences,omitempty"`
	// An array of relationships to iamIdp resources.
	Idps []IamIdpRelationship `json:"Idps,omitempty"`
	// An array of relationships to iamPermission resources.
	Permissions []IamPermissionRelationship `json:"Permissions,omitempty"`
	// An array of relationships to iamPrivilegeSet resources.
	PrivilegeSets []IamPrivilegeSetRelationship `json:"PrivilegeSets,omitempty"`
	// An array of relationships to iamPrivilege resources.
	Privileges     []IamPrivilegeRelationship     `json:"Privileges,omitempty"`
	ResourceLimits *IamResourceLimitsRelationship `json:"ResourceLimits,omitempty"`
	// An array of relationships to iamRole resources.
	Roles                []IamRoleRelationship          `json:"Roles,omitempty"`
	SecurityHolder       *IamSecurityHolderRelationship `json:"SecurityHolder,omitempty"`
	SessionLimits        *IamSessionLimitsRelationship  `json:"SessionLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamAccount IamAccount

// NewIamAccount instantiates a new IamAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamAccount(classId string, objectType string) *IamAccount {
	this := IamAccount{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamAccountWithDefaults instantiates a new IamAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamAccountWithDefaults() *IamAccount {
	this := IamAccount{}
	var classId string = "iam.Account"
	this.ClassId = classId
	var objectType string = "iam.Account"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamAccount) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamAccount) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamAccount) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamAccount) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamAccount) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamAccount) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamAccount) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAccount) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamAccount) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamAccount) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IamAccount) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAccount) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IamAccount) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IamAccount) SetStatus(v string) {
	o.Status = &v
}

// GetAppRegistrations returns the AppRegistrations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAccount) GetAppRegistrations() []IamAppRegistrationRelationship {
	if o == nil {
		var ret []IamAppRegistrationRelationship
		return ret
	}
	return o.AppRegistrations
}

// GetAppRegistrationsOk returns a tuple with the AppRegistrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAccount) GetAppRegistrationsOk() (*[]IamAppRegistrationRelationship, bool) {
	if o == nil || o.AppRegistrations == nil {
		return nil, false
	}
	return &o.AppRegistrations, true
}

// HasAppRegistrations returns a boolean if a field has been set.
func (o *IamAccount) HasAppRegistrations() bool {
	if o != nil && o.AppRegistrations != nil {
		return true
	}

	return false
}

// SetAppRegistrations gets a reference to the given []IamAppRegistrationRelationship and assigns it to the AppRegistrations field.
func (o *IamAccount) SetAppRegistrations(v []IamAppRegistrationRelationship) {
	o.AppRegistrations = v
}

// GetDomainGroups returns the DomainGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAccount) GetDomainGroups() []IamDomainGroupRelationship {
	if o == nil {
		var ret []IamDomainGroupRelationship
		return ret
	}
	return o.DomainGroups
}

// GetDomainGroupsOk returns a tuple with the DomainGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAccount) GetDomainGroupsOk() (*[]IamDomainGroupRelationship, bool) {
	if o == nil || o.DomainGroups == nil {
		return nil, false
	}
	return &o.DomainGroups, true
}

// HasDomainGroups returns a boolean if a field has been set.
func (o *IamAccount) HasDomainGroups() bool {
	if o != nil && o.DomainGroups != nil {
		return true
	}

	return false
}

// SetDomainGroups gets a reference to the given []IamDomainGroupRelationship and assigns it to the DomainGroups field.
func (o *IamAccount) SetDomainGroups(v []IamDomainGroupRelationship) {
	o.DomainGroups = v
}

// GetEndPointRoles returns the EndPointRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAccount) GetEndPointRoles() []IamEndPointRoleRelationship {
	if o == nil {
		var ret []IamEndPointRoleRelationship
		return ret
	}
	return o.EndPointRoles
}

// GetEndPointRolesOk returns a tuple with the EndPointRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAccount) GetEndPointRolesOk() (*[]IamEndPointRoleRelationship, bool) {
	if o == nil || o.EndPointRoles == nil {
		return nil, false
	}
	return &o.EndPointRoles, true
}

// HasEndPointRoles returns a boolean if a field has been set.
func (o *IamAccount) HasEndPointRoles() bool {
	if o != nil && o.EndPointRoles != nil {
		return true
	}

	return false
}

// SetEndPointRoles gets a reference to the given []IamEndPointRoleRelationship and assigns it to the EndPointRoles field.
func (o *IamAccount) SetEndPointRoles(v []IamEndPointRoleRelationship) {
	o.EndPointRoles = v
}

// GetIdpreferences returns the Idpreferences field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAccount) GetIdpreferences() []IamIdpReferenceRelationship {
	if o == nil {
		var ret []IamIdpReferenceRelationship
		return ret
	}
	return o.Idpreferences
}

// GetIdpreferencesOk returns a tuple with the Idpreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAccount) GetIdpreferencesOk() (*[]IamIdpReferenceRelationship, bool) {
	if o == nil || o.Idpreferences == nil {
		return nil, false
	}
	return &o.Idpreferences, true
}

// HasIdpreferences returns a boolean if a field has been set.
func (o *IamAccount) HasIdpreferences() bool {
	if o != nil && o.Idpreferences != nil {
		return true
	}

	return false
}

// SetIdpreferences gets a reference to the given []IamIdpReferenceRelationship and assigns it to the Idpreferences field.
func (o *IamAccount) SetIdpreferences(v []IamIdpReferenceRelationship) {
	o.Idpreferences = v
}

// GetIdps returns the Idps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAccount) GetIdps() []IamIdpRelationship {
	if o == nil {
		var ret []IamIdpRelationship
		return ret
	}
	return o.Idps
}

// GetIdpsOk returns a tuple with the Idps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAccount) GetIdpsOk() (*[]IamIdpRelationship, bool) {
	if o == nil || o.Idps == nil {
		return nil, false
	}
	return &o.Idps, true
}

// HasIdps returns a boolean if a field has been set.
func (o *IamAccount) HasIdps() bool {
	if o != nil && o.Idps != nil {
		return true
	}

	return false
}

// SetIdps gets a reference to the given []IamIdpRelationship and assigns it to the Idps field.
func (o *IamAccount) SetIdps(v []IamIdpRelationship) {
	o.Idps = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAccount) GetPermissions() []IamPermissionRelationship {
	if o == nil {
		var ret []IamPermissionRelationship
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAccount) GetPermissionsOk() (*[]IamPermissionRelationship, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *IamAccount) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []IamPermissionRelationship and assigns it to the Permissions field.
func (o *IamAccount) SetPermissions(v []IamPermissionRelationship) {
	o.Permissions = v
}

// GetPrivilegeSets returns the PrivilegeSets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAccount) GetPrivilegeSets() []IamPrivilegeSetRelationship {
	if o == nil {
		var ret []IamPrivilegeSetRelationship
		return ret
	}
	return o.PrivilegeSets
}

// GetPrivilegeSetsOk returns a tuple with the PrivilegeSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAccount) GetPrivilegeSetsOk() (*[]IamPrivilegeSetRelationship, bool) {
	if o == nil || o.PrivilegeSets == nil {
		return nil, false
	}
	return &o.PrivilegeSets, true
}

// HasPrivilegeSets returns a boolean if a field has been set.
func (o *IamAccount) HasPrivilegeSets() bool {
	if o != nil && o.PrivilegeSets != nil {
		return true
	}

	return false
}

// SetPrivilegeSets gets a reference to the given []IamPrivilegeSetRelationship and assigns it to the PrivilegeSets field.
func (o *IamAccount) SetPrivilegeSets(v []IamPrivilegeSetRelationship) {
	o.PrivilegeSets = v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAccount) GetPrivileges() []IamPrivilegeRelationship {
	if o == nil {
		var ret []IamPrivilegeRelationship
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAccount) GetPrivilegesOk() (*[]IamPrivilegeRelationship, bool) {
	if o == nil || o.Privileges == nil {
		return nil, false
	}
	return &o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *IamAccount) HasPrivileges() bool {
	if o != nil && o.Privileges != nil {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given []IamPrivilegeRelationship and assigns it to the Privileges field.
func (o *IamAccount) SetPrivileges(v []IamPrivilegeRelationship) {
	o.Privileges = v
}

// GetResourceLimits returns the ResourceLimits field value if set, zero value otherwise.
func (o *IamAccount) GetResourceLimits() IamResourceLimitsRelationship {
	if o == nil || o.ResourceLimits == nil {
		var ret IamResourceLimitsRelationship
		return ret
	}
	return *o.ResourceLimits
}

// GetResourceLimitsOk returns a tuple with the ResourceLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAccount) GetResourceLimitsOk() (*IamResourceLimitsRelationship, bool) {
	if o == nil || o.ResourceLimits == nil {
		return nil, false
	}
	return o.ResourceLimits, true
}

// HasResourceLimits returns a boolean if a field has been set.
func (o *IamAccount) HasResourceLimits() bool {
	if o != nil && o.ResourceLimits != nil {
		return true
	}

	return false
}

// SetResourceLimits gets a reference to the given IamResourceLimitsRelationship and assigns it to the ResourceLimits field.
func (o *IamAccount) SetResourceLimits(v IamResourceLimitsRelationship) {
	o.ResourceLimits = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAccount) GetRoles() []IamRoleRelationship {
	if o == nil {
		var ret []IamRoleRelationship
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAccount) GetRolesOk() (*[]IamRoleRelationship, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return &o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *IamAccount) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []IamRoleRelationship and assigns it to the Roles field.
func (o *IamAccount) SetRoles(v []IamRoleRelationship) {
	o.Roles = v
}

// GetSecurityHolder returns the SecurityHolder field value if set, zero value otherwise.
func (o *IamAccount) GetSecurityHolder() IamSecurityHolderRelationship {
	if o == nil || o.SecurityHolder == nil {
		var ret IamSecurityHolderRelationship
		return ret
	}
	return *o.SecurityHolder
}

// GetSecurityHolderOk returns a tuple with the SecurityHolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAccount) GetSecurityHolderOk() (*IamSecurityHolderRelationship, bool) {
	if o == nil || o.SecurityHolder == nil {
		return nil, false
	}
	return o.SecurityHolder, true
}

// HasSecurityHolder returns a boolean if a field has been set.
func (o *IamAccount) HasSecurityHolder() bool {
	if o != nil && o.SecurityHolder != nil {
		return true
	}

	return false
}

// SetSecurityHolder gets a reference to the given IamSecurityHolderRelationship and assigns it to the SecurityHolder field.
func (o *IamAccount) SetSecurityHolder(v IamSecurityHolderRelationship) {
	o.SecurityHolder = &v
}

// GetSessionLimits returns the SessionLimits field value if set, zero value otherwise.
func (o *IamAccount) GetSessionLimits() IamSessionLimitsRelationship {
	if o == nil || o.SessionLimits == nil {
		var ret IamSessionLimitsRelationship
		return ret
	}
	return *o.SessionLimits
}

// GetSessionLimitsOk returns a tuple with the SessionLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAccount) GetSessionLimitsOk() (*IamSessionLimitsRelationship, bool) {
	if o == nil || o.SessionLimits == nil {
		return nil, false
	}
	return o.SessionLimits, true
}

// HasSessionLimits returns a boolean if a field has been set.
func (o *IamAccount) HasSessionLimits() bool {
	if o != nil && o.SessionLimits != nil {
		return true
	}

	return false
}

// SetSessionLimits gets a reference to the given IamSessionLimitsRelationship and assigns it to the SessionLimits field.
func (o *IamAccount) SetSessionLimits(v IamSessionLimitsRelationship) {
	o.SessionLimits = &v
}

func (o IamAccount) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.AppRegistrations != nil {
		toSerialize["AppRegistrations"] = o.AppRegistrations
	}
	if o.DomainGroups != nil {
		toSerialize["DomainGroups"] = o.DomainGroups
	}
	if o.EndPointRoles != nil {
		toSerialize["EndPointRoles"] = o.EndPointRoles
	}
	if o.Idpreferences != nil {
		toSerialize["Idpreferences"] = o.Idpreferences
	}
	if o.Idps != nil {
		toSerialize["Idps"] = o.Idps
	}
	if o.Permissions != nil {
		toSerialize["Permissions"] = o.Permissions
	}
	if o.PrivilegeSets != nil {
		toSerialize["PrivilegeSets"] = o.PrivilegeSets
	}
	if o.Privileges != nil {
		toSerialize["Privileges"] = o.Privileges
	}
	if o.ResourceLimits != nil {
		toSerialize["ResourceLimits"] = o.ResourceLimits
	}
	if o.Roles != nil {
		toSerialize["Roles"] = o.Roles
	}
	if o.SecurityHolder != nil {
		toSerialize["SecurityHolder"] = o.SecurityHolder
	}
	if o.SessionLimits != nil {
		toSerialize["SessionLimits"] = o.SessionLimits
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamAccount) UnmarshalJSON(bytes []byte) (err error) {
	type IamAccountWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the Intersight account. By default, name is same as the MoID of the account.
		Name *string `json:"Name,omitempty"`
		// Status of the account. To activate the Intersight account, claim a device to the account.
		Status *string `json:"Status,omitempty"`
		// An array of relationships to iamAppRegistration resources.
		AppRegistrations []IamAppRegistrationRelationship `json:"AppRegistrations,omitempty"`
		// An array of relationships to iamDomainGroup resources.
		DomainGroups []IamDomainGroupRelationship `json:"DomainGroups,omitempty"`
		// An array of relationships to iamEndPointRole resources.
		EndPointRoles []IamEndPointRoleRelationship `json:"EndPointRoles,omitempty"`
		// An array of relationships to iamIdpReference resources.
		Idpreferences []IamIdpReferenceRelationship `json:"Idpreferences,omitempty"`
		// An array of relationships to iamIdp resources.
		Idps []IamIdpRelationship `json:"Idps,omitempty"`
		// An array of relationships to iamPermission resources.
		Permissions []IamPermissionRelationship `json:"Permissions,omitempty"`
		// An array of relationships to iamPrivilegeSet resources.
		PrivilegeSets []IamPrivilegeSetRelationship `json:"PrivilegeSets,omitempty"`
		// An array of relationships to iamPrivilege resources.
		Privileges     []IamPrivilegeRelationship     `json:"Privileges,omitempty"`
		ResourceLimits *IamResourceLimitsRelationship `json:"ResourceLimits,omitempty"`
		// An array of relationships to iamRole resources.
		Roles          []IamRoleRelationship          `json:"Roles,omitempty"`
		SecurityHolder *IamSecurityHolderRelationship `json:"SecurityHolder,omitempty"`
		SessionLimits  *IamSessionLimitsRelationship  `json:"SessionLimits,omitempty"`
	}

	varIamAccountWithoutEmbeddedStruct := IamAccountWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamAccountWithoutEmbeddedStruct)
	if err == nil {
		varIamAccount := _IamAccount{}
		varIamAccount.ClassId = varIamAccountWithoutEmbeddedStruct.ClassId
		varIamAccount.ObjectType = varIamAccountWithoutEmbeddedStruct.ObjectType
		varIamAccount.Name = varIamAccountWithoutEmbeddedStruct.Name
		varIamAccount.Status = varIamAccountWithoutEmbeddedStruct.Status
		varIamAccount.AppRegistrations = varIamAccountWithoutEmbeddedStruct.AppRegistrations
		varIamAccount.DomainGroups = varIamAccountWithoutEmbeddedStruct.DomainGroups
		varIamAccount.EndPointRoles = varIamAccountWithoutEmbeddedStruct.EndPointRoles
		varIamAccount.Idpreferences = varIamAccountWithoutEmbeddedStruct.Idpreferences
		varIamAccount.Idps = varIamAccountWithoutEmbeddedStruct.Idps
		varIamAccount.Permissions = varIamAccountWithoutEmbeddedStruct.Permissions
		varIamAccount.PrivilegeSets = varIamAccountWithoutEmbeddedStruct.PrivilegeSets
		varIamAccount.Privileges = varIamAccountWithoutEmbeddedStruct.Privileges
		varIamAccount.ResourceLimits = varIamAccountWithoutEmbeddedStruct.ResourceLimits
		varIamAccount.Roles = varIamAccountWithoutEmbeddedStruct.Roles
		varIamAccount.SecurityHolder = varIamAccountWithoutEmbeddedStruct.SecurityHolder
		varIamAccount.SessionLimits = varIamAccountWithoutEmbeddedStruct.SessionLimits
		*o = IamAccount(varIamAccount)
	} else {
		return err
	}

	varIamAccount := _IamAccount{}

	err = json.Unmarshal(bytes, &varIamAccount)
	if err == nil {
		o.MoBaseMo = varIamAccount.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "AppRegistrations")
		delete(additionalProperties, "DomainGroups")
		delete(additionalProperties, "EndPointRoles")
		delete(additionalProperties, "Idpreferences")
		delete(additionalProperties, "Idps")
		delete(additionalProperties, "Permissions")
		delete(additionalProperties, "PrivilegeSets")
		delete(additionalProperties, "Privileges")
		delete(additionalProperties, "ResourceLimits")
		delete(additionalProperties, "Roles")
		delete(additionalProperties, "SecurityHolder")
		delete(additionalProperties, "SessionLimits")

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

type NullableIamAccount struct {
	value *IamAccount
	isSet bool
}

func (v NullableIamAccount) Get() *IamAccount {
	return v.value
}

func (v *NullableIamAccount) Set(val *IamAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableIamAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableIamAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamAccount(val *IamAccount) *NullableIamAccount {
	return &NullableIamAccount{value: val, isSet: true}
}

func (v NullableIamAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
