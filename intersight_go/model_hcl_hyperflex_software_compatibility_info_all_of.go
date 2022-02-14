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

// HclHyperflexSoftwareCompatibilityInfoAllOf Definition of the list of properties defined in 'hcl.HyperflexSoftwareCompatibilityInfo', excluding properties defined in parent classes.
type HclHyperflexSoftwareCompatibilityInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType  string          `json:"ObjectType"`
	Constraints []HclConstraint `json:"Constraints,omitempty"`
	// HXDP component software version.
	HxdpVersion *string `json:"HxdpVersion,omitempty"`
	// Type fo Hypervisor the HyperFlex components versions are compatible with. For example ESX, Hyperv or KVM. * `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * `HyperFlexAp` - The hypervisor of the virtualization platform is Cisco HyperFlex Application Platform. * `IWE` - The hypervisor of the virtualization platform is Cisco Intersight Workload Engine. * `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * `Unknown` - The hypervisor running on the HyperFlex cluster is not known.
	HypervisorType *string `json:"HypervisorType,omitempty"`
	// Hypervisor component software version.
	HypervisorVersion *string `json:"HypervisorVersion,omitempty"`
	// Type of the HXDP bundle mgmt or full.
	IsMgmtBuild *string `json:"IsMgmtBuild,omitempty"`
	// UCS Server Firmware component software version.
	ServerFwVersion      *string                          `json:"ServerFwVersion,omitempty"`
	AppCatalog           *HyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HclHyperflexSoftwareCompatibilityInfoAllOf HclHyperflexSoftwareCompatibilityInfoAllOf

// NewHclHyperflexSoftwareCompatibilityInfoAllOf instantiates a new HclHyperflexSoftwareCompatibilityInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclHyperflexSoftwareCompatibilityInfoAllOf(classId string, objectType string) *HclHyperflexSoftwareCompatibilityInfoAllOf {
	this := HclHyperflexSoftwareCompatibilityInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	return &this
}

// NewHclHyperflexSoftwareCompatibilityInfoAllOfWithDefaults instantiates a new HclHyperflexSoftwareCompatibilityInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclHyperflexSoftwareCompatibilityInfoAllOfWithDefaults() *HclHyperflexSoftwareCompatibilityInfoAllOf {
	this := HclHyperflexSoftwareCompatibilityInfoAllOf{}
	var classId string = "hcl.HyperflexSoftwareCompatibilityInfo"
	this.ClassId = classId
	var objectType string = "hcl.HyperflexSoftwareCompatibilityInfo"
	this.ObjectType = objectType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetConstraints() []HclConstraint {
	if o == nil {
		var ret []HclConstraint
		return ret
	}
	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetConstraintsOk() (*[]HclConstraint, bool) {
	if o == nil || o.Constraints == nil {
		return nil, false
	}
	return &o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) HasConstraints() bool {
	if o != nil && o.Constraints != nil {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given []HclConstraint and assigns it to the Constraints field.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) SetConstraints(v []HclConstraint) {
	o.Constraints = v
}

// GetHxdpVersion returns the HxdpVersion field value if set, zero value otherwise.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetHxdpVersion() string {
	if o == nil || o.HxdpVersion == nil {
		var ret string
		return ret
	}
	return *o.HxdpVersion
}

// GetHxdpVersionOk returns a tuple with the HxdpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetHxdpVersionOk() (*string, bool) {
	if o == nil || o.HxdpVersion == nil {
		return nil, false
	}
	return o.HxdpVersion, true
}

// HasHxdpVersion returns a boolean if a field has been set.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) HasHxdpVersion() bool {
	if o != nil && o.HxdpVersion != nil {
		return true
	}

	return false
}

// SetHxdpVersion gets a reference to the given string and assigns it to the HxdpVersion field.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) SetHxdpVersion(v string) {
	o.HxdpVersion = &v
}

// GetHypervisorType returns the HypervisorType field value if set, zero value otherwise.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetHypervisorType() string {
	if o == nil || o.HypervisorType == nil {
		var ret string
		return ret
	}
	return *o.HypervisorType
}

// GetHypervisorTypeOk returns a tuple with the HypervisorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetHypervisorTypeOk() (*string, bool) {
	if o == nil || o.HypervisorType == nil {
		return nil, false
	}
	return o.HypervisorType, true
}

// HasHypervisorType returns a boolean if a field has been set.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) HasHypervisorType() bool {
	if o != nil && o.HypervisorType != nil {
		return true
	}

	return false
}

// SetHypervisorType gets a reference to the given string and assigns it to the HypervisorType field.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) SetHypervisorType(v string) {
	o.HypervisorType = &v
}

// GetHypervisorVersion returns the HypervisorVersion field value if set, zero value otherwise.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetHypervisorVersion() string {
	if o == nil || o.HypervisorVersion == nil {
		var ret string
		return ret
	}
	return *o.HypervisorVersion
}

// GetHypervisorVersionOk returns a tuple with the HypervisorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetHypervisorVersionOk() (*string, bool) {
	if o == nil || o.HypervisorVersion == nil {
		return nil, false
	}
	return o.HypervisorVersion, true
}

// HasHypervisorVersion returns a boolean if a field has been set.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) HasHypervisorVersion() bool {
	if o != nil && o.HypervisorVersion != nil {
		return true
	}

	return false
}

// SetHypervisorVersion gets a reference to the given string and assigns it to the HypervisorVersion field.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) SetHypervisorVersion(v string) {
	o.HypervisorVersion = &v
}

// GetIsMgmtBuild returns the IsMgmtBuild field value if set, zero value otherwise.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetIsMgmtBuild() string {
	if o == nil || o.IsMgmtBuild == nil {
		var ret string
		return ret
	}
	return *o.IsMgmtBuild
}

// GetIsMgmtBuildOk returns a tuple with the IsMgmtBuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetIsMgmtBuildOk() (*string, bool) {
	if o == nil || o.IsMgmtBuild == nil {
		return nil, false
	}
	return o.IsMgmtBuild, true
}

// HasIsMgmtBuild returns a boolean if a field has been set.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) HasIsMgmtBuild() bool {
	if o != nil && o.IsMgmtBuild != nil {
		return true
	}

	return false
}

// SetIsMgmtBuild gets a reference to the given string and assigns it to the IsMgmtBuild field.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) SetIsMgmtBuild(v string) {
	o.IsMgmtBuild = &v
}

// GetServerFwVersion returns the ServerFwVersion field value if set, zero value otherwise.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetServerFwVersion() string {
	if o == nil || o.ServerFwVersion == nil {
		var ret string
		return ret
	}
	return *o.ServerFwVersion
}

// GetServerFwVersionOk returns a tuple with the ServerFwVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetServerFwVersionOk() (*string, bool) {
	if o == nil || o.ServerFwVersion == nil {
		return nil, false
	}
	return o.ServerFwVersion, true
}

// HasServerFwVersion returns a boolean if a field has been set.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) HasServerFwVersion() bool {
	if o != nil && o.ServerFwVersion != nil {
		return true
	}

	return false
}

// SetServerFwVersion gets a reference to the given string and assigns it to the ServerFwVersion field.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) SetServerFwVersion(v string) {
	o.ServerFwVersion = &v
}

// GetAppCatalog returns the AppCatalog field value if set, zero value otherwise.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetAppCatalog() HyperflexAppCatalogRelationship {
	if o == nil || o.AppCatalog == nil {
		var ret HyperflexAppCatalogRelationship
		return ret
	}
	return *o.AppCatalog
}

// GetAppCatalogOk returns a tuple with the AppCatalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool) {
	if o == nil || o.AppCatalog == nil {
		return nil, false
	}
	return o.AppCatalog, true
}

// HasAppCatalog returns a boolean if a field has been set.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) HasAppCatalog() bool {
	if o != nil && o.AppCatalog != nil {
		return true
	}

	return false
}

// SetAppCatalog gets a reference to the given HyperflexAppCatalogRelationship and assigns it to the AppCatalog field.
func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) SetAppCatalog(v HyperflexAppCatalogRelationship) {
	o.AppCatalog = &v
}

func (o HclHyperflexSoftwareCompatibilityInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Constraints != nil {
		toSerialize["Constraints"] = o.Constraints
	}
	if o.HxdpVersion != nil {
		toSerialize["HxdpVersion"] = o.HxdpVersion
	}
	if o.HypervisorType != nil {
		toSerialize["HypervisorType"] = o.HypervisorType
	}
	if o.HypervisorVersion != nil {
		toSerialize["HypervisorVersion"] = o.HypervisorVersion
	}
	if o.IsMgmtBuild != nil {
		toSerialize["IsMgmtBuild"] = o.IsMgmtBuild
	}
	if o.ServerFwVersion != nil {
		toSerialize["ServerFwVersion"] = o.ServerFwVersion
	}
	if o.AppCatalog != nil {
		toSerialize["AppCatalog"] = o.AppCatalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHclHyperflexSoftwareCompatibilityInfoAllOf := _HclHyperflexSoftwareCompatibilityInfoAllOf{}

	if err = json.Unmarshal(bytes, &varHclHyperflexSoftwareCompatibilityInfoAllOf); err == nil {
		*o = HclHyperflexSoftwareCompatibilityInfoAllOf(varHclHyperflexSoftwareCompatibilityInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Constraints")
		delete(additionalProperties, "HxdpVersion")
		delete(additionalProperties, "HypervisorType")
		delete(additionalProperties, "HypervisorVersion")
		delete(additionalProperties, "IsMgmtBuild")
		delete(additionalProperties, "ServerFwVersion")
		delete(additionalProperties, "AppCatalog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHclHyperflexSoftwareCompatibilityInfoAllOf struct {
	value *HclHyperflexSoftwareCompatibilityInfoAllOf
	isSet bool
}

func (v NullableHclHyperflexSoftwareCompatibilityInfoAllOf) Get() *HclHyperflexSoftwareCompatibilityInfoAllOf {
	return v.value
}

func (v *NullableHclHyperflexSoftwareCompatibilityInfoAllOf) Set(val *HclHyperflexSoftwareCompatibilityInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHclHyperflexSoftwareCompatibilityInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHclHyperflexSoftwareCompatibilityInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclHyperflexSoftwareCompatibilityInfoAllOf(val *HclHyperflexSoftwareCompatibilityInfoAllOf) *NullableHclHyperflexSoftwareCompatibilityInfoAllOf {
	return &NullableHclHyperflexSoftwareCompatibilityInfoAllOf{value: val, isSet: true}
}

func (v NullableHclHyperflexSoftwareCompatibilityInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclHyperflexSoftwareCompatibilityInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
