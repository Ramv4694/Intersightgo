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

// KubernetesContainerRuntimePolicyAllOf Definition of the list of properties defined in 'kubernetes.ContainerRuntimePolicy', excluding properties defined in parent classes.
type KubernetesContainerRuntimePolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Bridge IP (--bip) including Prefix (e.g., 172.17.0.5/24) that Docker will use for the default bridge network (docker0). Containers will connect to this if no other network is configured, not used by kubernetes pods because their network is managed by CNI. However this address space must not collide with other CIDRs on your networks, including the cluster's Service CIDR, Pod Network CIDR and IP Pools.
	DockerBridgeNetworkCidr *string                       `json:"DockerBridgeNetworkCidr,omitempty"`
	DockerHttpProxy         NullableKubernetesProxyConfig `json:"DockerHttpProxy,omitempty"`
	DockerHttpsProxy        NullableKubernetesProxyConfig `json:"DockerHttpsProxy,omitempty"`
	DockerNoProxy           []string                      `json:"DockerNoProxy,omitempty"`
	// An array of relationships to kubernetesClusterProfile resources.
	ClusterProfiles      []KubernetesClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship  `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesContainerRuntimePolicyAllOf KubernetesContainerRuntimePolicyAllOf

// NewKubernetesContainerRuntimePolicyAllOf instantiates a new KubernetesContainerRuntimePolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesContainerRuntimePolicyAllOf(classId string, objectType string) *KubernetesContainerRuntimePolicyAllOf {
	this := KubernetesContainerRuntimePolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesContainerRuntimePolicyAllOfWithDefaults instantiates a new KubernetesContainerRuntimePolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesContainerRuntimePolicyAllOfWithDefaults() *KubernetesContainerRuntimePolicyAllOf {
	this := KubernetesContainerRuntimePolicyAllOf{}
	var classId string = "kubernetes.ContainerRuntimePolicy"
	this.ClassId = classId
	var objectType string = "kubernetes.ContainerRuntimePolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesContainerRuntimePolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesContainerRuntimePolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesContainerRuntimePolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesContainerRuntimePolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesContainerRuntimePolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesContainerRuntimePolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDockerBridgeNetworkCidr returns the DockerBridgeNetworkCidr field value if set, zero value otherwise.
func (o *KubernetesContainerRuntimePolicyAllOf) GetDockerBridgeNetworkCidr() string {
	if o == nil || o.DockerBridgeNetworkCidr == nil {
		var ret string
		return ret
	}
	return *o.DockerBridgeNetworkCidr
}

// GetDockerBridgeNetworkCidrOk returns a tuple with the DockerBridgeNetworkCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesContainerRuntimePolicyAllOf) GetDockerBridgeNetworkCidrOk() (*string, bool) {
	if o == nil || o.DockerBridgeNetworkCidr == nil {
		return nil, false
	}
	return o.DockerBridgeNetworkCidr, true
}

// HasDockerBridgeNetworkCidr returns a boolean if a field has been set.
func (o *KubernetesContainerRuntimePolicyAllOf) HasDockerBridgeNetworkCidr() bool {
	if o != nil && o.DockerBridgeNetworkCidr != nil {
		return true
	}

	return false
}

// SetDockerBridgeNetworkCidr gets a reference to the given string and assigns it to the DockerBridgeNetworkCidr field.
func (o *KubernetesContainerRuntimePolicyAllOf) SetDockerBridgeNetworkCidr(v string) {
	o.DockerBridgeNetworkCidr = &v
}

// GetDockerHttpProxy returns the DockerHttpProxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesContainerRuntimePolicyAllOf) GetDockerHttpProxy() KubernetesProxyConfig {
	if o == nil || o.DockerHttpProxy.Get() == nil {
		var ret KubernetesProxyConfig
		return ret
	}
	return *o.DockerHttpProxy.Get()
}

// GetDockerHttpProxyOk returns a tuple with the DockerHttpProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesContainerRuntimePolicyAllOf) GetDockerHttpProxyOk() (*KubernetesProxyConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.DockerHttpProxy.Get(), o.DockerHttpProxy.IsSet()
}

// HasDockerHttpProxy returns a boolean if a field has been set.
func (o *KubernetesContainerRuntimePolicyAllOf) HasDockerHttpProxy() bool {
	if o != nil && o.DockerHttpProxy.IsSet() {
		return true
	}

	return false
}

// SetDockerHttpProxy gets a reference to the given NullableKubernetesProxyConfig and assigns it to the DockerHttpProxy field.
func (o *KubernetesContainerRuntimePolicyAllOf) SetDockerHttpProxy(v KubernetesProxyConfig) {
	o.DockerHttpProxy.Set(&v)
}

// SetDockerHttpProxyNil sets the value for DockerHttpProxy to be an explicit nil
func (o *KubernetesContainerRuntimePolicyAllOf) SetDockerHttpProxyNil() {
	o.DockerHttpProxy.Set(nil)
}

// UnsetDockerHttpProxy ensures that no value is present for DockerHttpProxy, not even an explicit nil
func (o *KubernetesContainerRuntimePolicyAllOf) UnsetDockerHttpProxy() {
	o.DockerHttpProxy.Unset()
}

// GetDockerHttpsProxy returns the DockerHttpsProxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesContainerRuntimePolicyAllOf) GetDockerHttpsProxy() KubernetesProxyConfig {
	if o == nil || o.DockerHttpsProxy.Get() == nil {
		var ret KubernetesProxyConfig
		return ret
	}
	return *o.DockerHttpsProxy.Get()
}

// GetDockerHttpsProxyOk returns a tuple with the DockerHttpsProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesContainerRuntimePolicyAllOf) GetDockerHttpsProxyOk() (*KubernetesProxyConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.DockerHttpsProxy.Get(), o.DockerHttpsProxy.IsSet()
}

// HasDockerHttpsProxy returns a boolean if a field has been set.
func (o *KubernetesContainerRuntimePolicyAllOf) HasDockerHttpsProxy() bool {
	if o != nil && o.DockerHttpsProxy.IsSet() {
		return true
	}

	return false
}

// SetDockerHttpsProxy gets a reference to the given NullableKubernetesProxyConfig and assigns it to the DockerHttpsProxy field.
func (o *KubernetesContainerRuntimePolicyAllOf) SetDockerHttpsProxy(v KubernetesProxyConfig) {
	o.DockerHttpsProxy.Set(&v)
}

// SetDockerHttpsProxyNil sets the value for DockerHttpsProxy to be an explicit nil
func (o *KubernetesContainerRuntimePolicyAllOf) SetDockerHttpsProxyNil() {
	o.DockerHttpsProxy.Set(nil)
}

// UnsetDockerHttpsProxy ensures that no value is present for DockerHttpsProxy, not even an explicit nil
func (o *KubernetesContainerRuntimePolicyAllOf) UnsetDockerHttpsProxy() {
	o.DockerHttpsProxy.Unset()
}

// GetDockerNoProxy returns the DockerNoProxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesContainerRuntimePolicyAllOf) GetDockerNoProxy() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DockerNoProxy
}

// GetDockerNoProxyOk returns a tuple with the DockerNoProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesContainerRuntimePolicyAllOf) GetDockerNoProxyOk() (*[]string, bool) {
	if o == nil || o.DockerNoProxy == nil {
		return nil, false
	}
	return &o.DockerNoProxy, true
}

// HasDockerNoProxy returns a boolean if a field has been set.
func (o *KubernetesContainerRuntimePolicyAllOf) HasDockerNoProxy() bool {
	if o != nil && o.DockerNoProxy != nil {
		return true
	}

	return false
}

// SetDockerNoProxy gets a reference to the given []string and assigns it to the DockerNoProxy field.
func (o *KubernetesContainerRuntimePolicyAllOf) SetDockerNoProxy(v []string) {
	o.DockerNoProxy = v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesContainerRuntimePolicyAllOf) GetClusterProfiles() []KubernetesClusterProfileRelationship {
	if o == nil {
		var ret []KubernetesClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesContainerRuntimePolicyAllOf) GetClusterProfilesOk() (*[]KubernetesClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *KubernetesContainerRuntimePolicyAllOf) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []KubernetesClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *KubernetesContainerRuntimePolicyAllOf) SetClusterProfiles(v []KubernetesClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesContainerRuntimePolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesContainerRuntimePolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesContainerRuntimePolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesContainerRuntimePolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o KubernetesContainerRuntimePolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DockerBridgeNetworkCidr != nil {
		toSerialize["DockerBridgeNetworkCidr"] = o.DockerBridgeNetworkCidr
	}
	if o.DockerHttpProxy.IsSet() {
		toSerialize["DockerHttpProxy"] = o.DockerHttpProxy.Get()
	}
	if o.DockerHttpsProxy.IsSet() {
		toSerialize["DockerHttpsProxy"] = o.DockerHttpsProxy.Get()
	}
	if o.DockerNoProxy != nil {
		toSerialize["DockerNoProxy"] = o.DockerNoProxy
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesContainerRuntimePolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesContainerRuntimePolicyAllOf := _KubernetesContainerRuntimePolicyAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesContainerRuntimePolicyAllOf); err == nil {
		*o = KubernetesContainerRuntimePolicyAllOf(varKubernetesContainerRuntimePolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DockerBridgeNetworkCidr")
		delete(additionalProperties, "DockerHttpProxy")
		delete(additionalProperties, "DockerHttpsProxy")
		delete(additionalProperties, "DockerNoProxy")
		delete(additionalProperties, "ClusterProfiles")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesContainerRuntimePolicyAllOf struct {
	value *KubernetesContainerRuntimePolicyAllOf
	isSet bool
}

func (v NullableKubernetesContainerRuntimePolicyAllOf) Get() *KubernetesContainerRuntimePolicyAllOf {
	return v.value
}

func (v *NullableKubernetesContainerRuntimePolicyAllOf) Set(val *KubernetesContainerRuntimePolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesContainerRuntimePolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesContainerRuntimePolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesContainerRuntimePolicyAllOf(val *KubernetesContainerRuntimePolicyAllOf) *NullableKubernetesContainerRuntimePolicyAllOf {
	return &NullableKubernetesContainerRuntimePolicyAllOf{value: val, isSet: true}
}

func (v NullableKubernetesContainerRuntimePolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesContainerRuntimePolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}