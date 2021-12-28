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
	"time"
)

// ResourceReservation A Reservation is used to reserve a place for a new resource in the resource groups.
type ResourceReservation struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Expiration of the resource Reservation.
	Expiration *time.Time `json:"Expiration,omitempty"`
	// MarkFail is used to set the reservation status to Failed.
	MarkFail      *bool    `json:"MarkFail,omitempty"`
	ResourceMoids []string `json:"ResourceMoids,omitempty"`
	// Type of resources which will get filled into the resource groups.
	ResourceType *string `json:"ResourceType,omitempty"`
	// Status of the Reservation. * `Created` - By default, a reservation is in Created status. * `Processing` - A reservation is changed to Processing status for appliance mode resource claim requests. * `Failed` - A reservation is changed to Failed status if the validations on resources, resource groups fails. * `Finished` - A reservation is changed to Finished status if the validations on resources, resource groups are successful. The resource moids in reservation will be added to resource groups using OData filters.
	Status *string `json:"Status,omitempty"`
	// Moid of the user who created the reservation.
	UserMoid *string                 `json:"UserMoid,omitempty"`
	Account  *IamAccountRelationship `json:"Account,omitempty"`
	// An array of relationships to resourceGroup resources.
	Groups               []ResourceGroupRelationship `json:"Groups,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceReservation ResourceReservation

// NewResourceReservation instantiates a new ResourceReservation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceReservation(classId string, objectType string) *ResourceReservation {
	this := ResourceReservation{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "Created"
	this.Status = &status
	return &this
}

// NewResourceReservationWithDefaults instantiates a new ResourceReservation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceReservationWithDefaults() *ResourceReservation {
	this := ResourceReservation{}
	var classId string = "resource.Reservation"
	this.ClassId = classId
	var objectType string = "resource.Reservation"
	this.ObjectType = objectType
	var status string = "Created"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourceReservation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourceReservation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourceReservation) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ResourceReservation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourceReservation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourceReservation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *ResourceReservation) GetExpiration() time.Time {
	if o == nil || o.Expiration == nil {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceReservation) GetExpirationOk() (*time.Time, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *ResourceReservation) HasExpiration() bool {
	if o != nil && o.Expiration != nil {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *ResourceReservation) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetMarkFail returns the MarkFail field value if set, zero value otherwise.
func (o *ResourceReservation) GetMarkFail() bool {
	if o == nil || o.MarkFail == nil {
		var ret bool
		return ret
	}
	return *o.MarkFail
}

// GetMarkFailOk returns a tuple with the MarkFail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceReservation) GetMarkFailOk() (*bool, bool) {
	if o == nil || o.MarkFail == nil {
		return nil, false
	}
	return o.MarkFail, true
}

// HasMarkFail returns a boolean if a field has been set.
func (o *ResourceReservation) HasMarkFail() bool {
	if o != nil && o.MarkFail != nil {
		return true
	}

	return false
}

// SetMarkFail gets a reference to the given bool and assigns it to the MarkFail field.
func (o *ResourceReservation) SetMarkFail(v bool) {
	o.MarkFail = &v
}

// GetResourceMoids returns the ResourceMoids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceReservation) GetResourceMoids() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ResourceMoids
}

// GetResourceMoidsOk returns a tuple with the ResourceMoids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceReservation) GetResourceMoidsOk() (*[]string, bool) {
	if o == nil || o.ResourceMoids == nil {
		return nil, false
	}
	return &o.ResourceMoids, true
}

// HasResourceMoids returns a boolean if a field has been set.
func (o *ResourceReservation) HasResourceMoids() bool {
	if o != nil && o.ResourceMoids != nil {
		return true
	}

	return false
}

// SetResourceMoids gets a reference to the given []string and assigns it to the ResourceMoids field.
func (o *ResourceReservation) SetResourceMoids(v []string) {
	o.ResourceMoids = v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *ResourceReservation) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceReservation) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *ResourceReservation) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *ResourceReservation) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResourceReservation) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceReservation) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResourceReservation) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ResourceReservation) SetStatus(v string) {
	o.Status = &v
}

// GetUserMoid returns the UserMoid field value if set, zero value otherwise.
func (o *ResourceReservation) GetUserMoid() string {
	if o == nil || o.UserMoid == nil {
		var ret string
		return ret
	}
	return *o.UserMoid
}

// GetUserMoidOk returns a tuple with the UserMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceReservation) GetUserMoidOk() (*string, bool) {
	if o == nil || o.UserMoid == nil {
		return nil, false
	}
	return o.UserMoid, true
}

// HasUserMoid returns a boolean if a field has been set.
func (o *ResourceReservation) HasUserMoid() bool {
	if o != nil && o.UserMoid != nil {
		return true
	}

	return false
}

// SetUserMoid gets a reference to the given string and assigns it to the UserMoid field.
func (o *ResourceReservation) SetUserMoid(v string) {
	o.UserMoid = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ResourceReservation) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceReservation) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ResourceReservation) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ResourceReservation) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceReservation) GetGroups() []ResourceGroupRelationship {
	if o == nil {
		var ret []ResourceGroupRelationship
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceReservation) GetGroupsOk() (*[]ResourceGroupRelationship, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return &o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ResourceReservation) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []ResourceGroupRelationship and assigns it to the Groups field.
func (o *ResourceReservation) SetGroups(v []ResourceGroupRelationship) {
	o.Groups = v
}

func (o ResourceReservation) MarshalJSON() ([]byte, error) {
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
	if o.Expiration != nil {
		toSerialize["Expiration"] = o.Expiration
	}
	if o.MarkFail != nil {
		toSerialize["MarkFail"] = o.MarkFail
	}
	if o.ResourceMoids != nil {
		toSerialize["ResourceMoids"] = o.ResourceMoids
	}
	if o.ResourceType != nil {
		toSerialize["ResourceType"] = o.ResourceType
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.UserMoid != nil {
		toSerialize["UserMoid"] = o.UserMoid
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.Groups != nil {
		toSerialize["Groups"] = o.Groups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceReservation) UnmarshalJSON(bytes []byte) (err error) {
	type ResourceReservationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Expiration of the resource Reservation.
		Expiration *time.Time `json:"Expiration,omitempty"`
		// MarkFail is used to set the reservation status to Failed.
		MarkFail      *bool    `json:"MarkFail,omitempty"`
		ResourceMoids []string `json:"ResourceMoids,omitempty"`
		// Type of resources which will get filled into the resource groups.
		ResourceType *string `json:"ResourceType,omitempty"`
		// Status of the Reservation. * `Created` - By default, a reservation is in Created status. * `Processing` - A reservation is changed to Processing status for appliance mode resource claim requests. * `Failed` - A reservation is changed to Failed status if the validations on resources, resource groups fails. * `Finished` - A reservation is changed to Finished status if the validations on resources, resource groups are successful. The resource moids in reservation will be added to resource groups using OData filters.
		Status *string `json:"Status,omitempty"`
		// Moid of the user who created the reservation.
		UserMoid *string                 `json:"UserMoid,omitempty"`
		Account  *IamAccountRelationship `json:"Account,omitempty"`
		// An array of relationships to resourceGroup resources.
		Groups []ResourceGroupRelationship `json:"Groups,omitempty"`
	}

	varResourceReservationWithoutEmbeddedStruct := ResourceReservationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varResourceReservationWithoutEmbeddedStruct)
	if err == nil {
		varResourceReservation := _ResourceReservation{}
		varResourceReservation.ClassId = varResourceReservationWithoutEmbeddedStruct.ClassId
		varResourceReservation.ObjectType = varResourceReservationWithoutEmbeddedStruct.ObjectType
		varResourceReservation.Expiration = varResourceReservationWithoutEmbeddedStruct.Expiration
		varResourceReservation.MarkFail = varResourceReservationWithoutEmbeddedStruct.MarkFail
		varResourceReservation.ResourceMoids = varResourceReservationWithoutEmbeddedStruct.ResourceMoids
		varResourceReservation.ResourceType = varResourceReservationWithoutEmbeddedStruct.ResourceType
		varResourceReservation.Status = varResourceReservationWithoutEmbeddedStruct.Status
		varResourceReservation.UserMoid = varResourceReservationWithoutEmbeddedStruct.UserMoid
		varResourceReservation.Account = varResourceReservationWithoutEmbeddedStruct.Account
		varResourceReservation.Groups = varResourceReservationWithoutEmbeddedStruct.Groups
		*o = ResourceReservation(varResourceReservation)
	} else {
		return err
	}

	varResourceReservation := _ResourceReservation{}

	err = json.Unmarshal(bytes, &varResourceReservation)
	if err == nil {
		o.MoBaseMo = varResourceReservation.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Expiration")
		delete(additionalProperties, "MarkFail")
		delete(additionalProperties, "ResourceMoids")
		delete(additionalProperties, "ResourceType")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "UserMoid")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "Groups")

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

type NullableResourceReservation struct {
	value *ResourceReservation
	isSet bool
}

func (v NullableResourceReservation) Get() *ResourceReservation {
	return v.value
}

func (v *NullableResourceReservation) Set(val *ResourceReservation) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceReservation) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceReservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceReservation(val *ResourceReservation) *NullableResourceReservation {
	return &NullableResourceReservation{value: val, isSet: true}
}

func (v NullableResourceReservation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceReservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
