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

// TelemetryDruidConstantPostAggregatorAllOf struct for TelemetryDruidConstantPostAggregatorAllOf
type TelemetryDruidConstantPostAggregatorAllOf struct {
	// Output name for the post-aggregator.
	Name *string `json:"name,omitempty"`
	// The numerical value.
	Value                *float64 `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidConstantPostAggregatorAllOf TelemetryDruidConstantPostAggregatorAllOf

// NewTelemetryDruidConstantPostAggregatorAllOf instantiates a new TelemetryDruidConstantPostAggregatorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidConstantPostAggregatorAllOf() *TelemetryDruidConstantPostAggregatorAllOf {
	this := TelemetryDruidConstantPostAggregatorAllOf{}
	return &this
}

// NewTelemetryDruidConstantPostAggregatorAllOfWithDefaults instantiates a new TelemetryDruidConstantPostAggregatorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidConstantPostAggregatorAllOfWithDefaults() *TelemetryDruidConstantPostAggregatorAllOf {
	this := TelemetryDruidConstantPostAggregatorAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TelemetryDruidConstantPostAggregatorAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidConstantPostAggregatorAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TelemetryDruidConstantPostAggregatorAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TelemetryDruidConstantPostAggregatorAllOf) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TelemetryDruidConstantPostAggregatorAllOf) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidConstantPostAggregatorAllOf) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TelemetryDruidConstantPostAggregatorAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *TelemetryDruidConstantPostAggregatorAllOf) SetValue(v float64) {
	o.Value = &v
}

func (o TelemetryDruidConstantPostAggregatorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidConstantPostAggregatorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidConstantPostAggregatorAllOf := _TelemetryDruidConstantPostAggregatorAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidConstantPostAggregatorAllOf); err == nil {
		*o = TelemetryDruidConstantPostAggregatorAllOf(varTelemetryDruidConstantPostAggregatorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidConstantPostAggregatorAllOf struct {
	value *TelemetryDruidConstantPostAggregatorAllOf
	isSet bool
}

func (v NullableTelemetryDruidConstantPostAggregatorAllOf) Get() *TelemetryDruidConstantPostAggregatorAllOf {
	return v.value
}

func (v *NullableTelemetryDruidConstantPostAggregatorAllOf) Set(val *TelemetryDruidConstantPostAggregatorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidConstantPostAggregatorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidConstantPostAggregatorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidConstantPostAggregatorAllOf(val *TelemetryDruidConstantPostAggregatorAllOf) *NullableTelemetryDruidConstantPostAggregatorAllOf {
	return &NullableTelemetryDruidConstantPostAggregatorAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidConstantPostAggregatorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidConstantPostAggregatorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}