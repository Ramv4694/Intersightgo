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

// SoftwarerepositoryCatalog A container MO that holds references to the files in an account's image repository. It is internally created for each account and is used to hold information about all user uploaded files.
type SoftwarerepositoryCatalog struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The status of the image catalog synchronization operation.
	IsImagePullFailure *bool `json:"IsImagePullFailure,omitempty"`
	// The name of the catalog. The names are populated and predefined during MO creation.
	Name                 *string                               `json:"Name,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	System               *IamSystemRelationship                `json:"System,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryCatalog SoftwarerepositoryCatalog

// NewSoftwarerepositoryCatalog instantiates a new SoftwarerepositoryCatalog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryCatalog(classId string, objectType string) *SoftwarerepositoryCatalog {
	this := SoftwarerepositoryCatalog{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwarerepositoryCatalogWithDefaults instantiates a new SoftwarerepositoryCatalog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryCatalogWithDefaults() *SoftwarerepositoryCatalog {
	this := SoftwarerepositoryCatalog{}
	var classId string = "softwarerepository.Catalog"
	this.ClassId = classId
	var objectType string = "softwarerepository.Catalog"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryCatalog) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCatalog) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryCatalog) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryCatalog) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCatalog) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryCatalog) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsImagePullFailure returns the IsImagePullFailure field value if set, zero value otherwise.
func (o *SoftwarerepositoryCatalog) GetIsImagePullFailure() bool {
	if o == nil || o.IsImagePullFailure == nil {
		var ret bool
		return ret
	}
	return *o.IsImagePullFailure
}

// GetIsImagePullFailureOk returns a tuple with the IsImagePullFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCatalog) GetIsImagePullFailureOk() (*bool, bool) {
	if o == nil || o.IsImagePullFailure == nil {
		return nil, false
	}
	return o.IsImagePullFailure, true
}

// HasIsImagePullFailure returns a boolean if a field has been set.
func (o *SoftwarerepositoryCatalog) HasIsImagePullFailure() bool {
	if o != nil && o.IsImagePullFailure != nil {
		return true
	}

	return false
}

// SetIsImagePullFailure gets a reference to the given bool and assigns it to the IsImagePullFailure field.
func (o *SoftwarerepositoryCatalog) SetIsImagePullFailure(v bool) {
	o.IsImagePullFailure = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SoftwarerepositoryCatalog) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCatalog) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SoftwarerepositoryCatalog) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SoftwarerepositoryCatalog) SetName(v string) {
	o.Name = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *SoftwarerepositoryCatalog) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCatalog) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *SoftwarerepositoryCatalog) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *SoftwarerepositoryCatalog) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *SoftwarerepositoryCatalog) GetSystem() IamSystemRelationship {
	if o == nil || o.System == nil {
		var ret IamSystemRelationship
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCatalog) GetSystemOk() (*IamSystemRelationship, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *SoftwarerepositoryCatalog) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given IamSystemRelationship and assigns it to the System field.
func (o *SoftwarerepositoryCatalog) SetSystem(v IamSystemRelationship) {
	o.System = &v
}

func (o SoftwarerepositoryCatalog) MarshalJSON() ([]byte, error) {
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
	if o.IsImagePullFailure != nil {
		toSerialize["IsImagePullFailure"] = o.IsImagePullFailure
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.System != nil {
		toSerialize["System"] = o.System
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryCatalog) UnmarshalJSON(bytes []byte) (err error) {
	type SoftwarerepositoryCatalogWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The status of the image catalog synchronization operation.
		IsImagePullFailure *bool `json:"IsImagePullFailure,omitempty"`
		// The name of the catalog. The names are populated and predefined during MO creation.
		Name         *string                               `json:"Name,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		System       *IamSystemRelationship                `json:"System,omitempty"`
	}

	varSoftwarerepositoryCatalogWithoutEmbeddedStruct := SoftwarerepositoryCatalogWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSoftwarerepositoryCatalogWithoutEmbeddedStruct)
	if err == nil {
		varSoftwarerepositoryCatalog := _SoftwarerepositoryCatalog{}
		varSoftwarerepositoryCatalog.ClassId = varSoftwarerepositoryCatalogWithoutEmbeddedStruct.ClassId
		varSoftwarerepositoryCatalog.ObjectType = varSoftwarerepositoryCatalogWithoutEmbeddedStruct.ObjectType
		varSoftwarerepositoryCatalog.IsImagePullFailure = varSoftwarerepositoryCatalogWithoutEmbeddedStruct.IsImagePullFailure
		varSoftwarerepositoryCatalog.Name = varSoftwarerepositoryCatalogWithoutEmbeddedStruct.Name
		varSoftwarerepositoryCatalog.Organization = varSoftwarerepositoryCatalogWithoutEmbeddedStruct.Organization
		varSoftwarerepositoryCatalog.System = varSoftwarerepositoryCatalogWithoutEmbeddedStruct.System
		*o = SoftwarerepositoryCatalog(varSoftwarerepositoryCatalog)
	} else {
		return err
	}

	varSoftwarerepositoryCatalog := _SoftwarerepositoryCatalog{}

	err = json.Unmarshal(bytes, &varSoftwarerepositoryCatalog)
	if err == nil {
		o.MoBaseMo = varSoftwarerepositoryCatalog.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsImagePullFailure")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "System")

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

type NullableSoftwarerepositoryCatalog struct {
	value *SoftwarerepositoryCatalog
	isSet bool
}

func (v NullableSoftwarerepositoryCatalog) Get() *SoftwarerepositoryCatalog {
	return v.value
}

func (v *NullableSoftwarerepositoryCatalog) Set(val *SoftwarerepositoryCatalog) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryCatalog) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryCatalog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryCatalog(val *SoftwarerepositoryCatalog) *NullableSoftwarerepositoryCatalog {
	return &NullableSoftwarerepositoryCatalog{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryCatalog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryCatalog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
