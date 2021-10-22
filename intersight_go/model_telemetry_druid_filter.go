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
	"fmt"
)

// TelemetryDruidFilter - struct for TelemetryDruidFilter
type TelemetryDruidFilter struct {
	TelemetryDruidAndFilter              *TelemetryDruidAndFilter
	TelemetryDruidColumnComparisonFilter *TelemetryDruidColumnComparisonFilter
	TelemetryDruidNotFilter              *TelemetryDruidNotFilter
	TelemetryDruidOrFilter               *TelemetryDruidOrFilter
	TelemetryDruidRegexFilter            *TelemetryDruidRegexFilter
	TelemetryDruidSelectorFilter         *TelemetryDruidSelectorFilter
}

// TelemetryDruidAndFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidAndFilter wrapped in TelemetryDruidFilter
func TelemetryDruidAndFilterAsTelemetryDruidFilter(v *TelemetryDruidAndFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{TelemetryDruidAndFilter: v}
}

// TelemetryDruidColumnComparisonFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidColumnComparisonFilter wrapped in TelemetryDruidFilter
func TelemetryDruidColumnComparisonFilterAsTelemetryDruidFilter(v *TelemetryDruidColumnComparisonFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{TelemetryDruidColumnComparisonFilter: v}
}

// TelemetryDruidNotFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidNotFilter wrapped in TelemetryDruidFilter
func TelemetryDruidNotFilterAsTelemetryDruidFilter(v *TelemetryDruidNotFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{TelemetryDruidNotFilter: v}
}

// TelemetryDruidOrFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidOrFilter wrapped in TelemetryDruidFilter
func TelemetryDruidOrFilterAsTelemetryDruidFilter(v *TelemetryDruidOrFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{TelemetryDruidOrFilter: v}
}

// TelemetryDruidRegexFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidRegexFilter wrapped in TelemetryDruidFilter
func TelemetryDruidRegexFilterAsTelemetryDruidFilter(v *TelemetryDruidRegexFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{TelemetryDruidRegexFilter: v}
}

// TelemetryDruidSelectorFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidSelectorFilter wrapped in TelemetryDruidFilter
func TelemetryDruidSelectorFilterAsTelemetryDruidFilter(v *TelemetryDruidSelectorFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{TelemetryDruidSelectorFilter: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TelemetryDruidFilter) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'and'
	if jsonDict["type"] == "and" {
		// try to unmarshal JSON data into TelemetryDruidAndFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidAndFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidAndFilter, return on the first match
		} else {
			dst.TelemetryDruidAndFilter = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidFilter as TelemetryDruidAndFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'columnComparison'
	if jsonDict["type"] == "columnComparison" {
		// try to unmarshal JSON data into TelemetryDruidColumnComparisonFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidColumnComparisonFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidColumnComparisonFilter, return on the first match
		} else {
			dst.TelemetryDruidColumnComparisonFilter = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidFilter as TelemetryDruidColumnComparisonFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'not'
	if jsonDict["type"] == "not" {
		// try to unmarshal JSON data into TelemetryDruidNotFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidNotFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidNotFilter, return on the first match
		} else {
			dst.TelemetryDruidNotFilter = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidFilter as TelemetryDruidNotFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'or'
	if jsonDict["type"] == "or" {
		// try to unmarshal JSON data into TelemetryDruidOrFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidOrFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidOrFilter, return on the first match
		} else {
			dst.TelemetryDruidOrFilter = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidFilter as TelemetryDruidOrFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'regex'
	if jsonDict["type"] == "regex" {
		// try to unmarshal JSON data into TelemetryDruidRegexFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidRegexFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidRegexFilter, return on the first match
		} else {
			dst.TelemetryDruidRegexFilter = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidFilter as TelemetryDruidRegexFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'selector'
	if jsonDict["type"] == "selector" {
		// try to unmarshal JSON data into TelemetryDruidSelectorFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidSelectorFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidSelectorFilter, return on the first match
		} else {
			dst.TelemetryDruidSelectorFilter = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidFilter as TelemetryDruidSelectorFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidAndFilter'
	if jsonDict["type"] == "telemetry.DruidAndFilter" {
		// try to unmarshal JSON data into TelemetryDruidAndFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidAndFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidAndFilter, return on the first match
		} else {
			dst.TelemetryDruidAndFilter = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidFilter as TelemetryDruidAndFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidColumnComparisonFilter'
	if jsonDict["type"] == "telemetry.DruidColumnComparisonFilter" {
		// try to unmarshal JSON data into TelemetryDruidColumnComparisonFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidColumnComparisonFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidColumnComparisonFilter, return on the first match
		} else {
			dst.TelemetryDruidColumnComparisonFilter = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidFilter as TelemetryDruidColumnComparisonFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidNotFilter'
	if jsonDict["type"] == "telemetry.DruidNotFilter" {
		// try to unmarshal JSON data into TelemetryDruidNotFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidNotFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidNotFilter, return on the first match
		} else {
			dst.TelemetryDruidNotFilter = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidFilter as TelemetryDruidNotFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidOrFilter'
	if jsonDict["type"] == "telemetry.DruidOrFilter" {
		// try to unmarshal JSON data into TelemetryDruidOrFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidOrFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidOrFilter, return on the first match
		} else {
			dst.TelemetryDruidOrFilter = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidFilter as TelemetryDruidOrFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidRegexFilter'
	if jsonDict["type"] == "telemetry.DruidRegexFilter" {
		// try to unmarshal JSON data into TelemetryDruidRegexFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidRegexFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidRegexFilter, return on the first match
		} else {
			dst.TelemetryDruidRegexFilter = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidFilter as TelemetryDruidRegexFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidSelectorFilter'
	if jsonDict["type"] == "telemetry.DruidSelectorFilter" {
		// try to unmarshal JSON data into TelemetryDruidSelectorFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidSelectorFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidSelectorFilter, return on the first match
		} else {
			dst.TelemetryDruidSelectorFilter = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidFilter as TelemetryDruidSelectorFilter: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TelemetryDruidFilter) MarshalJSON() ([]byte, error) {
	if src.TelemetryDruidAndFilter != nil {
		return json.Marshal(&src.TelemetryDruidAndFilter)
	}

	if src.TelemetryDruidColumnComparisonFilter != nil {
		return json.Marshal(&src.TelemetryDruidColumnComparisonFilter)
	}

	if src.TelemetryDruidNotFilter != nil {
		return json.Marshal(&src.TelemetryDruidNotFilter)
	}

	if src.TelemetryDruidOrFilter != nil {
		return json.Marshal(&src.TelemetryDruidOrFilter)
	}

	if src.TelemetryDruidRegexFilter != nil {
		return json.Marshal(&src.TelemetryDruidRegexFilter)
	}

	if src.TelemetryDruidSelectorFilter != nil {
		return json.Marshal(&src.TelemetryDruidSelectorFilter)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TelemetryDruidFilter) GetActualInstance() interface{} {
	if obj.TelemetryDruidAndFilter != nil {
		return obj.TelemetryDruidAndFilter
	}

	if obj.TelemetryDruidColumnComparisonFilter != nil {
		return obj.TelemetryDruidColumnComparisonFilter
	}

	if obj.TelemetryDruidNotFilter != nil {
		return obj.TelemetryDruidNotFilter
	}

	if obj.TelemetryDruidOrFilter != nil {
		return obj.TelemetryDruidOrFilter
	}

	if obj.TelemetryDruidRegexFilter != nil {
		return obj.TelemetryDruidRegexFilter
	}

	if obj.TelemetryDruidSelectorFilter != nil {
		return obj.TelemetryDruidSelectorFilter
	}

	// all schemas are nil
	return nil
}

type NullableTelemetryDruidFilter struct {
	value *TelemetryDruidFilter
	isSet bool
}

func (v NullableTelemetryDruidFilter) Get() *TelemetryDruidFilter {
	return v.value
}

func (v *NullableTelemetryDruidFilter) Set(val *TelemetryDruidFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidFilter(val *TelemetryDruidFilter) *NullableTelemetryDruidFilter {
	return &NullableTelemetryDruidFilter{value: val, isSet: true}
}

func (v NullableTelemetryDruidFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
