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

// KubernetesDaemonSetStatusAllOf Definition of the list of properties defined in 'kubernetes.DaemonSetStatus', excluding properties defined in parent classes.
type KubernetesDaemonSetStatusAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod.
	CurrentNumberScheduled *int64 `json:"CurrentNumberScheduled,omitempty"`
	// The total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod).
	DesiredNumberScheduled *int64 `json:"DesiredNumberScheduled,omitempty"`
	// The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available (ready for at least spec.minReadySeconds).
	NumberAvailable *string `json:"NumberAvailable,omitempty"`
	// The number of nodes that are running the daemon pod, but are not supposed to run the daemon pod.
	NumberMisscheduled *int64 `json:"NumberMisscheduled,omitempty"`
	// The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready.
	NumberReady *int64 `json:"NumberReady,omitempty"`
	// The most recent generation observed by the daemon set controller.
	ObservedGeneration *int64 `json:"ObservedGeneration,omitempty"`
	// The total number of nodes that are running updated daemon pod.
	UpdatedNumberScheduled *string `json:"UpdatedNumberScheduled,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _KubernetesDaemonSetStatusAllOf KubernetesDaemonSetStatusAllOf

// NewKubernetesDaemonSetStatusAllOf instantiates a new KubernetesDaemonSetStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesDaemonSetStatusAllOf(classId string, objectType string) *KubernetesDaemonSetStatusAllOf {
	this := KubernetesDaemonSetStatusAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var currentNumberScheduled int64 = 0
	this.CurrentNumberScheduled = &currentNumberScheduled
	var desiredNumberScheduled int64 = 0
	this.DesiredNumberScheduled = &desiredNumberScheduled
	var numberMisscheduled int64 = 0
	this.NumberMisscheduled = &numberMisscheduled
	var numberReady int64 = 0
	this.NumberReady = &numberReady
	var observedGeneration int64 = 0
	this.ObservedGeneration = &observedGeneration
	return &this
}

// NewKubernetesDaemonSetStatusAllOfWithDefaults instantiates a new KubernetesDaemonSetStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesDaemonSetStatusAllOfWithDefaults() *KubernetesDaemonSetStatusAllOf {
	this := KubernetesDaemonSetStatusAllOf{}
	var classId string = "kubernetes.DaemonSetStatus"
	this.ClassId = classId
	var objectType string = "kubernetes.DaemonSetStatus"
	this.ObjectType = objectType
	var currentNumberScheduled int64 = 0
	this.CurrentNumberScheduled = &currentNumberScheduled
	var desiredNumberScheduled int64 = 0
	this.DesiredNumberScheduled = &desiredNumberScheduled
	var numberMisscheduled int64 = 0
	this.NumberMisscheduled = &numberMisscheduled
	var numberReady int64 = 0
	this.NumberReady = &numberReady
	var observedGeneration int64 = 0
	this.ObservedGeneration = &observedGeneration
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesDaemonSetStatusAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesDaemonSetStatusAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesDaemonSetStatusAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesDaemonSetStatusAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesDaemonSetStatusAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesDaemonSetStatusAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCurrentNumberScheduled returns the CurrentNumberScheduled field value if set, zero value otherwise.
func (o *KubernetesDaemonSetStatusAllOf) GetCurrentNumberScheduled() int64 {
	if o == nil || o.CurrentNumberScheduled == nil {
		var ret int64
		return ret
	}
	return *o.CurrentNumberScheduled
}

// GetCurrentNumberScheduledOk returns a tuple with the CurrentNumberScheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDaemonSetStatusAllOf) GetCurrentNumberScheduledOk() (*int64, bool) {
	if o == nil || o.CurrentNumberScheduled == nil {
		return nil, false
	}
	return o.CurrentNumberScheduled, true
}

// HasCurrentNumberScheduled returns a boolean if a field has been set.
func (o *KubernetesDaemonSetStatusAllOf) HasCurrentNumberScheduled() bool {
	if o != nil && o.CurrentNumberScheduled != nil {
		return true
	}

	return false
}

// SetCurrentNumberScheduled gets a reference to the given int64 and assigns it to the CurrentNumberScheduled field.
func (o *KubernetesDaemonSetStatusAllOf) SetCurrentNumberScheduled(v int64) {
	o.CurrentNumberScheduled = &v
}

// GetDesiredNumberScheduled returns the DesiredNumberScheduled field value if set, zero value otherwise.
func (o *KubernetesDaemonSetStatusAllOf) GetDesiredNumberScheduled() int64 {
	if o == nil || o.DesiredNumberScheduled == nil {
		var ret int64
		return ret
	}
	return *o.DesiredNumberScheduled
}

// GetDesiredNumberScheduledOk returns a tuple with the DesiredNumberScheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDaemonSetStatusAllOf) GetDesiredNumberScheduledOk() (*int64, bool) {
	if o == nil || o.DesiredNumberScheduled == nil {
		return nil, false
	}
	return o.DesiredNumberScheduled, true
}

// HasDesiredNumberScheduled returns a boolean if a field has been set.
func (o *KubernetesDaemonSetStatusAllOf) HasDesiredNumberScheduled() bool {
	if o != nil && o.DesiredNumberScheduled != nil {
		return true
	}

	return false
}

// SetDesiredNumberScheduled gets a reference to the given int64 and assigns it to the DesiredNumberScheduled field.
func (o *KubernetesDaemonSetStatusAllOf) SetDesiredNumberScheduled(v int64) {
	o.DesiredNumberScheduled = &v
}

// GetNumberAvailable returns the NumberAvailable field value if set, zero value otherwise.
func (o *KubernetesDaemonSetStatusAllOf) GetNumberAvailable() string {
	if o == nil || o.NumberAvailable == nil {
		var ret string
		return ret
	}
	return *o.NumberAvailable
}

// GetNumberAvailableOk returns a tuple with the NumberAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDaemonSetStatusAllOf) GetNumberAvailableOk() (*string, bool) {
	if o == nil || o.NumberAvailable == nil {
		return nil, false
	}
	return o.NumberAvailable, true
}

// HasNumberAvailable returns a boolean if a field has been set.
func (o *KubernetesDaemonSetStatusAllOf) HasNumberAvailable() bool {
	if o != nil && o.NumberAvailable != nil {
		return true
	}

	return false
}

// SetNumberAvailable gets a reference to the given string and assigns it to the NumberAvailable field.
func (o *KubernetesDaemonSetStatusAllOf) SetNumberAvailable(v string) {
	o.NumberAvailable = &v
}

// GetNumberMisscheduled returns the NumberMisscheduled field value if set, zero value otherwise.
func (o *KubernetesDaemonSetStatusAllOf) GetNumberMisscheduled() int64 {
	if o == nil || o.NumberMisscheduled == nil {
		var ret int64
		return ret
	}
	return *o.NumberMisscheduled
}

// GetNumberMisscheduledOk returns a tuple with the NumberMisscheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDaemonSetStatusAllOf) GetNumberMisscheduledOk() (*int64, bool) {
	if o == nil || o.NumberMisscheduled == nil {
		return nil, false
	}
	return o.NumberMisscheduled, true
}

// HasNumberMisscheduled returns a boolean if a field has been set.
func (o *KubernetesDaemonSetStatusAllOf) HasNumberMisscheduled() bool {
	if o != nil && o.NumberMisscheduled != nil {
		return true
	}

	return false
}

// SetNumberMisscheduled gets a reference to the given int64 and assigns it to the NumberMisscheduled field.
func (o *KubernetesDaemonSetStatusAllOf) SetNumberMisscheduled(v int64) {
	o.NumberMisscheduled = &v
}

// GetNumberReady returns the NumberReady field value if set, zero value otherwise.
func (o *KubernetesDaemonSetStatusAllOf) GetNumberReady() int64 {
	if o == nil || o.NumberReady == nil {
		var ret int64
		return ret
	}
	return *o.NumberReady
}

// GetNumberReadyOk returns a tuple with the NumberReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDaemonSetStatusAllOf) GetNumberReadyOk() (*int64, bool) {
	if o == nil || o.NumberReady == nil {
		return nil, false
	}
	return o.NumberReady, true
}

// HasNumberReady returns a boolean if a field has been set.
func (o *KubernetesDaemonSetStatusAllOf) HasNumberReady() bool {
	if o != nil && o.NumberReady != nil {
		return true
	}

	return false
}

// SetNumberReady gets a reference to the given int64 and assigns it to the NumberReady field.
func (o *KubernetesDaemonSetStatusAllOf) SetNumberReady(v int64) {
	o.NumberReady = &v
}

// GetObservedGeneration returns the ObservedGeneration field value if set, zero value otherwise.
func (o *KubernetesDaemonSetStatusAllOf) GetObservedGeneration() int64 {
	if o == nil || o.ObservedGeneration == nil {
		var ret int64
		return ret
	}
	return *o.ObservedGeneration
}

// GetObservedGenerationOk returns a tuple with the ObservedGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDaemonSetStatusAllOf) GetObservedGenerationOk() (*int64, bool) {
	if o == nil || o.ObservedGeneration == nil {
		return nil, false
	}
	return o.ObservedGeneration, true
}

// HasObservedGeneration returns a boolean if a field has been set.
func (o *KubernetesDaemonSetStatusAllOf) HasObservedGeneration() bool {
	if o != nil && o.ObservedGeneration != nil {
		return true
	}

	return false
}

// SetObservedGeneration gets a reference to the given int64 and assigns it to the ObservedGeneration field.
func (o *KubernetesDaemonSetStatusAllOf) SetObservedGeneration(v int64) {
	o.ObservedGeneration = &v
}

// GetUpdatedNumberScheduled returns the UpdatedNumberScheduled field value if set, zero value otherwise.
func (o *KubernetesDaemonSetStatusAllOf) GetUpdatedNumberScheduled() string {
	if o == nil || o.UpdatedNumberScheduled == nil {
		var ret string
		return ret
	}
	return *o.UpdatedNumberScheduled
}

// GetUpdatedNumberScheduledOk returns a tuple with the UpdatedNumberScheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDaemonSetStatusAllOf) GetUpdatedNumberScheduledOk() (*string, bool) {
	if o == nil || o.UpdatedNumberScheduled == nil {
		return nil, false
	}
	return o.UpdatedNumberScheduled, true
}

// HasUpdatedNumberScheduled returns a boolean if a field has been set.
func (o *KubernetesDaemonSetStatusAllOf) HasUpdatedNumberScheduled() bool {
	if o != nil && o.UpdatedNumberScheduled != nil {
		return true
	}

	return false
}

// SetUpdatedNumberScheduled gets a reference to the given string and assigns it to the UpdatedNumberScheduled field.
func (o *KubernetesDaemonSetStatusAllOf) SetUpdatedNumberScheduled(v string) {
	o.UpdatedNumberScheduled = &v
}

func (o KubernetesDaemonSetStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CurrentNumberScheduled != nil {
		toSerialize["CurrentNumberScheduled"] = o.CurrentNumberScheduled
	}
	if o.DesiredNumberScheduled != nil {
		toSerialize["DesiredNumberScheduled"] = o.DesiredNumberScheduled
	}
	if o.NumberAvailable != nil {
		toSerialize["NumberAvailable"] = o.NumberAvailable
	}
	if o.NumberMisscheduled != nil {
		toSerialize["NumberMisscheduled"] = o.NumberMisscheduled
	}
	if o.NumberReady != nil {
		toSerialize["NumberReady"] = o.NumberReady
	}
	if o.ObservedGeneration != nil {
		toSerialize["ObservedGeneration"] = o.ObservedGeneration
	}
	if o.UpdatedNumberScheduled != nil {
		toSerialize["UpdatedNumberScheduled"] = o.UpdatedNumberScheduled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesDaemonSetStatusAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesDaemonSetStatusAllOf := _KubernetesDaemonSetStatusAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesDaemonSetStatusAllOf); err == nil {
		*o = KubernetesDaemonSetStatusAllOf(varKubernetesDaemonSetStatusAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CurrentNumberScheduled")
		delete(additionalProperties, "DesiredNumberScheduled")
		delete(additionalProperties, "NumberAvailable")
		delete(additionalProperties, "NumberMisscheduled")
		delete(additionalProperties, "NumberReady")
		delete(additionalProperties, "ObservedGeneration")
		delete(additionalProperties, "UpdatedNumberScheduled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesDaemonSetStatusAllOf struct {
	value *KubernetesDaemonSetStatusAllOf
	isSet bool
}

func (v NullableKubernetesDaemonSetStatusAllOf) Get() *KubernetesDaemonSetStatusAllOf {
	return v.value
}

func (v *NullableKubernetesDaemonSetStatusAllOf) Set(val *KubernetesDaemonSetStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesDaemonSetStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesDaemonSetStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesDaemonSetStatusAllOf(val *KubernetesDaemonSetStatusAllOf) *NullableKubernetesDaemonSetStatusAllOf {
	return &NullableKubernetesDaemonSetStatusAllOf{value: val, isSet: true}
}

func (v NullableKubernetesDaemonSetStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesDaemonSetStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
