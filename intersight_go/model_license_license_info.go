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
	"time"
)

// LicenseLicenseInfo License state information for a specific license entitlement. Essentials license entitlement is supported currently. licenseState attribute is used for license enforcement. When license state is one of TrialPeriod, Compliance, or OutOfCompliance, the feature set defined for the license entitlement is granted to the customer.
type LicenseLicenseInfo struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The license administrative state. Set this property to 'true' to activate the license entitlements.
	ActiveAdmin *bool `json:"ActiveAdmin,omitempty"`
	// The number of days left for licenseState to stay in TrialPeriod or OutOfCompliance state.
	DaysLeft *int64 `json:"DaysLeft,omitempty"`
	// The date and time when the trial period expires. The value of the 'endTime' property is set when the account enters the TrialPeriod or OutOfCompliance state.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// The entitlement mode reported by Cisco Smart Software Manager.
	EnforceMode *string `json:"EnforceMode,omitempty"`
	// The detailed error message when there is any error related to this licensing entitlement.
	ErrorDesc *string `json:"ErrorDesc,omitempty"`
	// The default Trial or Grace period customer is entitled to.
	EvaluationPeriod *int64 `json:"EvaluationPeriod,omitempty"`
	// The date and time when the next expiration time of license subscription.
	ExpireTime *time.Time `json:"ExpireTime,omitempty"`
	// The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once.
	ExtraEvaluation *int64 `json:"ExtraEvaluation,omitempty"`
	// The total number of license consumed in the Intersight account.
	LicenseCount *int64 `json:"LicenseCount,omitempty"`
	// The total number of license purchased from cisco.
	LicenseCountPurchased *int64 `json:"LicenseCountPurchased,omitempty"`
	// The license state defined by Intersight. The value may be one of NotLicensed, TrialPeriod, OutOfCompliance, Compliance, GraceExpired, or TrialExpired. * `NotLicensed` - The license token is neither activated nor registered. * `GraceExpired` - The license grace period has expired. * `TrialPeriod` - The 90 days of trial period. * `OutOfCompliance` - The license is out of compliance. * `Compliance` - The license is in compliance. * `TrialExpired` - The trial period of 90 days has expired.
	LicenseState *string `json:"LicenseState,omitempty"`
	// The name of the Intersight license entitlement. For example, this property may be set to 'Essential'. * `Base` - Base as a License type. It is default license type. * `Essential` - Essential as a License type. * `Standard` - Standard as a License type. * `Advantage` - Advantage as a License type. * `Premier` - Premier as a License type. * `IWO-Essential` - IWO-Essential as a License type. * `IWO-Advantage` - IWO-Advantage as a License type. * `IWO-Premier` - IWO-Premier as a License type. * `IKS-Advantage` - IKS-Advantage as a License type.
	LicenseType *string `json:"LicenseType,omitempty"`
	// The date and time when the licenseState entered the TrialPeriod or OutOfCompliance state.
	StartTime *time.Time `json:"StartTime,omitempty"`
	// The administrative state of the trial license. When the LicenseState is set to 'NotLicensed', 'trialAdmin' can be set to true to start the trial period, i.e. licenseState is set to be TrialPeriod.
	TrialAdmin           *bool                                  `json:"TrialAdmin,omitempty"`
	AccountLicenseData   *LicenseAccountLicenseDataRelationship `json:"AccountLicenseData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LicenseLicenseInfo LicenseLicenseInfo

// NewLicenseLicenseInfo instantiates a new LicenseLicenseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseLicenseInfo(classId string, objectType string) *LicenseLicenseInfo {
	this := LicenseLicenseInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	var licenseState string = "NotLicensed"
	this.LicenseState = &licenseState
	var licenseType string = "Base"
	this.LicenseType = &licenseType
	return &this
}

// NewLicenseLicenseInfoWithDefaults instantiates a new LicenseLicenseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseLicenseInfoWithDefaults() *LicenseLicenseInfo {
	this := LicenseLicenseInfo{}
	var classId string = "license.LicenseInfo"
	this.ClassId = classId
	var objectType string = "license.LicenseInfo"
	this.ObjectType = objectType
	var licenseState string = "NotLicensed"
	this.LicenseState = &licenseState
	var licenseType string = "Base"
	this.LicenseType = &licenseType
	return &this
}

// GetClassId returns the ClassId field value
func (o *LicenseLicenseInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *LicenseLicenseInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *LicenseLicenseInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *LicenseLicenseInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *LicenseLicenseInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *LicenseLicenseInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActiveAdmin returns the ActiveAdmin field value if set, zero value otherwise.
func (o *LicenseLicenseInfo) GetActiveAdmin() bool {
	if o == nil || o.ActiveAdmin == nil {
		var ret bool
		return ret
	}
	return *o.ActiveAdmin
}

// GetActiveAdminOk returns a tuple with the ActiveAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseInfo) GetActiveAdminOk() (*bool, bool) {
	if o == nil || o.ActiveAdmin == nil {
		return nil, false
	}
	return o.ActiveAdmin, true
}

// HasActiveAdmin returns a boolean if a field has been set.
func (o *LicenseLicenseInfo) HasActiveAdmin() bool {
	if o != nil && o.ActiveAdmin != nil {
		return true
	}

	return false
}

// SetActiveAdmin gets a reference to the given bool and assigns it to the ActiveAdmin field.
func (o *LicenseLicenseInfo) SetActiveAdmin(v bool) {
	o.ActiveAdmin = &v
}

// GetDaysLeft returns the DaysLeft field value if set, zero value otherwise.
func (o *LicenseLicenseInfo) GetDaysLeft() int64 {
	if o == nil || o.DaysLeft == nil {
		var ret int64
		return ret
	}
	return *o.DaysLeft
}

// GetDaysLeftOk returns a tuple with the DaysLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseInfo) GetDaysLeftOk() (*int64, bool) {
	if o == nil || o.DaysLeft == nil {
		return nil, false
	}
	return o.DaysLeft, true
}

// HasDaysLeft returns a boolean if a field has been set.
func (o *LicenseLicenseInfo) HasDaysLeft() bool {
	if o != nil && o.DaysLeft != nil {
		return true
	}

	return false
}

// SetDaysLeft gets a reference to the given int64 and assigns it to the DaysLeft field.
func (o *LicenseLicenseInfo) SetDaysLeft(v int64) {
	o.DaysLeft = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *LicenseLicenseInfo) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseInfo) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *LicenseLicenseInfo) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *LicenseLicenseInfo) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetEnforceMode returns the EnforceMode field value if set, zero value otherwise.
func (o *LicenseLicenseInfo) GetEnforceMode() string {
	if o == nil || o.EnforceMode == nil {
		var ret string
		return ret
	}
	return *o.EnforceMode
}

// GetEnforceModeOk returns a tuple with the EnforceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseInfo) GetEnforceModeOk() (*string, bool) {
	if o == nil || o.EnforceMode == nil {
		return nil, false
	}
	return o.EnforceMode, true
}

// HasEnforceMode returns a boolean if a field has been set.
func (o *LicenseLicenseInfo) HasEnforceMode() bool {
	if o != nil && o.EnforceMode != nil {
		return true
	}

	return false
}

// SetEnforceMode gets a reference to the given string and assigns it to the EnforceMode field.
func (o *LicenseLicenseInfo) SetEnforceMode(v string) {
	o.EnforceMode = &v
}

// GetErrorDesc returns the ErrorDesc field value if set, zero value otherwise.
func (o *LicenseLicenseInfo) GetErrorDesc() string {
	if o == nil || o.ErrorDesc == nil {
		var ret string
		return ret
	}
	return *o.ErrorDesc
}

// GetErrorDescOk returns a tuple with the ErrorDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseInfo) GetErrorDescOk() (*string, bool) {
	if o == nil || o.ErrorDesc == nil {
		return nil, false
	}
	return o.ErrorDesc, true
}

// HasErrorDesc returns a boolean if a field has been set.
func (o *LicenseLicenseInfo) HasErrorDesc() bool {
	if o != nil && o.ErrorDesc != nil {
		return true
	}

	return false
}

// SetErrorDesc gets a reference to the given string and assigns it to the ErrorDesc field.
func (o *LicenseLicenseInfo) SetErrorDesc(v string) {
	o.ErrorDesc = &v
}

// GetEvaluationPeriod returns the EvaluationPeriod field value if set, zero value otherwise.
func (o *LicenseLicenseInfo) GetEvaluationPeriod() int64 {
	if o == nil || o.EvaluationPeriod == nil {
		var ret int64
		return ret
	}
	return *o.EvaluationPeriod
}

// GetEvaluationPeriodOk returns a tuple with the EvaluationPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseInfo) GetEvaluationPeriodOk() (*int64, bool) {
	if o == nil || o.EvaluationPeriod == nil {
		return nil, false
	}
	return o.EvaluationPeriod, true
}

// HasEvaluationPeriod returns a boolean if a field has been set.
func (o *LicenseLicenseInfo) HasEvaluationPeriod() bool {
	if o != nil && o.EvaluationPeriod != nil {
		return true
	}

	return false
}

// SetEvaluationPeriod gets a reference to the given int64 and assigns it to the EvaluationPeriod field.
func (o *LicenseLicenseInfo) SetEvaluationPeriod(v int64) {
	o.EvaluationPeriod = &v
}

// GetExpireTime returns the ExpireTime field value if set, zero value otherwise.
func (o *LicenseLicenseInfo) GetExpireTime() time.Time {
	if o == nil || o.ExpireTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpireTime
}

// GetExpireTimeOk returns a tuple with the ExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseInfo) GetExpireTimeOk() (*time.Time, bool) {
	if o == nil || o.ExpireTime == nil {
		return nil, false
	}
	return o.ExpireTime, true
}

// HasExpireTime returns a boolean if a field has been set.
func (o *LicenseLicenseInfo) HasExpireTime() bool {
	if o != nil && o.ExpireTime != nil {
		return true
	}

	return false
}

// SetExpireTime gets a reference to the given time.Time and assigns it to the ExpireTime field.
func (o *LicenseLicenseInfo) SetExpireTime(v time.Time) {
	o.ExpireTime = &v
}

// GetExtraEvaluation returns the ExtraEvaluation field value if set, zero value otherwise.
func (o *LicenseLicenseInfo) GetExtraEvaluation() int64 {
	if o == nil || o.ExtraEvaluation == nil {
		var ret int64
		return ret
	}
	return *o.ExtraEvaluation
}

// GetExtraEvaluationOk returns a tuple with the ExtraEvaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseInfo) GetExtraEvaluationOk() (*int64, bool) {
	if o == nil || o.ExtraEvaluation == nil {
		return nil, false
	}
	return o.ExtraEvaluation, true
}

// HasExtraEvaluation returns a boolean if a field has been set.
func (o *LicenseLicenseInfo) HasExtraEvaluation() bool {
	if o != nil && o.ExtraEvaluation != nil {
		return true
	}

	return false
}

// SetExtraEvaluation gets a reference to the given int64 and assigns it to the ExtraEvaluation field.
func (o *LicenseLicenseInfo) SetExtraEvaluation(v int64) {
	o.ExtraEvaluation = &v
}

// GetLicenseCount returns the LicenseCount field value if set, zero value otherwise.
func (o *LicenseLicenseInfo) GetLicenseCount() int64 {
	if o == nil || o.LicenseCount == nil {
		var ret int64
		return ret
	}
	return *o.LicenseCount
}

// GetLicenseCountOk returns a tuple with the LicenseCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseInfo) GetLicenseCountOk() (*int64, bool) {
	if o == nil || o.LicenseCount == nil {
		return nil, false
	}
	return o.LicenseCount, true
}

// HasLicenseCount returns a boolean if a field has been set.
func (o *LicenseLicenseInfo) HasLicenseCount() bool {
	if o != nil && o.LicenseCount != nil {
		return true
	}

	return false
}

// SetLicenseCount gets a reference to the given int64 and assigns it to the LicenseCount field.
func (o *LicenseLicenseInfo) SetLicenseCount(v int64) {
	o.LicenseCount = &v
}

// GetLicenseCountPurchased returns the LicenseCountPurchased field value if set, zero value otherwise.
func (o *LicenseLicenseInfo) GetLicenseCountPurchased() int64 {
	if o == nil || o.LicenseCountPurchased == nil {
		var ret int64
		return ret
	}
	return *o.LicenseCountPurchased
}

// GetLicenseCountPurchasedOk returns a tuple with the LicenseCountPurchased field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseInfo) GetLicenseCountPurchasedOk() (*int64, bool) {
	if o == nil || o.LicenseCountPurchased == nil {
		return nil, false
	}
	return o.LicenseCountPurchased, true
}

// HasLicenseCountPurchased returns a boolean if a field has been set.
func (o *LicenseLicenseInfo) HasLicenseCountPurchased() bool {
	if o != nil && o.LicenseCountPurchased != nil {
		return true
	}

	return false
}

// SetLicenseCountPurchased gets a reference to the given int64 and assigns it to the LicenseCountPurchased field.
func (o *LicenseLicenseInfo) SetLicenseCountPurchased(v int64) {
	o.LicenseCountPurchased = &v
}

// GetLicenseState returns the LicenseState field value if set, zero value otherwise.
func (o *LicenseLicenseInfo) GetLicenseState() string {
	if o == nil || o.LicenseState == nil {
		var ret string
		return ret
	}
	return *o.LicenseState
}

// GetLicenseStateOk returns a tuple with the LicenseState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseInfo) GetLicenseStateOk() (*string, bool) {
	if o == nil || o.LicenseState == nil {
		return nil, false
	}
	return o.LicenseState, true
}

// HasLicenseState returns a boolean if a field has been set.
func (o *LicenseLicenseInfo) HasLicenseState() bool {
	if o != nil && o.LicenseState != nil {
		return true
	}

	return false
}

// SetLicenseState gets a reference to the given string and assigns it to the LicenseState field.
func (o *LicenseLicenseInfo) SetLicenseState(v string) {
	o.LicenseState = &v
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *LicenseLicenseInfo) GetLicenseType() string {
	if o == nil || o.LicenseType == nil {
		var ret string
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseInfo) GetLicenseTypeOk() (*string, bool) {
	if o == nil || o.LicenseType == nil {
		return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *LicenseLicenseInfo) HasLicenseType() bool {
	if o != nil && o.LicenseType != nil {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given string and assigns it to the LicenseType field.
func (o *LicenseLicenseInfo) SetLicenseType(v string) {
	o.LicenseType = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *LicenseLicenseInfo) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseInfo) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *LicenseLicenseInfo) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *LicenseLicenseInfo) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetTrialAdmin returns the TrialAdmin field value if set, zero value otherwise.
func (o *LicenseLicenseInfo) GetTrialAdmin() bool {
	if o == nil || o.TrialAdmin == nil {
		var ret bool
		return ret
	}
	return *o.TrialAdmin
}

// GetTrialAdminOk returns a tuple with the TrialAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseInfo) GetTrialAdminOk() (*bool, bool) {
	if o == nil || o.TrialAdmin == nil {
		return nil, false
	}
	return o.TrialAdmin, true
}

// HasTrialAdmin returns a boolean if a field has been set.
func (o *LicenseLicenseInfo) HasTrialAdmin() bool {
	if o != nil && o.TrialAdmin != nil {
		return true
	}

	return false
}

// SetTrialAdmin gets a reference to the given bool and assigns it to the TrialAdmin field.
func (o *LicenseLicenseInfo) SetTrialAdmin(v bool) {
	o.TrialAdmin = &v
}

// GetAccountLicenseData returns the AccountLicenseData field value if set, zero value otherwise.
func (o *LicenseLicenseInfo) GetAccountLicenseData() LicenseAccountLicenseDataRelationship {
	if o == nil || o.AccountLicenseData == nil {
		var ret LicenseAccountLicenseDataRelationship
		return ret
	}
	return *o.AccountLicenseData
}

// GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseInfo) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool) {
	if o == nil || o.AccountLicenseData == nil {
		return nil, false
	}
	return o.AccountLicenseData, true
}

// HasAccountLicenseData returns a boolean if a field has been set.
func (o *LicenseLicenseInfo) HasAccountLicenseData() bool {
	if o != nil && o.AccountLicenseData != nil {
		return true
	}

	return false
}

// SetAccountLicenseData gets a reference to the given LicenseAccountLicenseDataRelationship and assigns it to the AccountLicenseData field.
func (o *LicenseLicenseInfo) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship) {
	o.AccountLicenseData = &v
}

func (o LicenseLicenseInfo) MarshalJSON() ([]byte, error) {
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
	if o.ActiveAdmin != nil {
		toSerialize["ActiveAdmin"] = o.ActiveAdmin
	}
	if o.DaysLeft != nil {
		toSerialize["DaysLeft"] = o.DaysLeft
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.EnforceMode != nil {
		toSerialize["EnforceMode"] = o.EnforceMode
	}
	if o.ErrorDesc != nil {
		toSerialize["ErrorDesc"] = o.ErrorDesc
	}
	if o.EvaluationPeriod != nil {
		toSerialize["EvaluationPeriod"] = o.EvaluationPeriod
	}
	if o.ExpireTime != nil {
		toSerialize["ExpireTime"] = o.ExpireTime
	}
	if o.ExtraEvaluation != nil {
		toSerialize["ExtraEvaluation"] = o.ExtraEvaluation
	}
	if o.LicenseCount != nil {
		toSerialize["LicenseCount"] = o.LicenseCount
	}
	if o.LicenseCountPurchased != nil {
		toSerialize["LicenseCountPurchased"] = o.LicenseCountPurchased
	}
	if o.LicenseState != nil {
		toSerialize["LicenseState"] = o.LicenseState
	}
	if o.LicenseType != nil {
		toSerialize["LicenseType"] = o.LicenseType
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.TrialAdmin != nil {
		toSerialize["TrialAdmin"] = o.TrialAdmin
	}
	if o.AccountLicenseData != nil {
		toSerialize["AccountLicenseData"] = o.AccountLicenseData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LicenseLicenseInfo) UnmarshalJSON(bytes []byte) (err error) {
	type LicenseLicenseInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The license administrative state. Set this property to 'true' to activate the license entitlements.
		ActiveAdmin *bool `json:"ActiveAdmin,omitempty"`
		// The number of days left for licenseState to stay in TrialPeriod or OutOfCompliance state.
		DaysLeft *int64 `json:"DaysLeft,omitempty"`
		// The date and time when the trial period expires. The value of the 'endTime' property is set when the account enters the TrialPeriod or OutOfCompliance state.
		EndTime *time.Time `json:"EndTime,omitempty"`
		// The entitlement mode reported by Cisco Smart Software Manager.
		EnforceMode *string `json:"EnforceMode,omitempty"`
		// The detailed error message when there is any error related to this licensing entitlement.
		ErrorDesc *string `json:"ErrorDesc,omitempty"`
		// The default Trial or Grace period customer is entitled to.
		EvaluationPeriod *int64 `json:"EvaluationPeriod,omitempty"`
		// The date and time when the next expiration time of license subscription.
		ExpireTime *time.Time `json:"ExpireTime,omitempty"`
		// The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once.
		ExtraEvaluation *int64 `json:"ExtraEvaluation,omitempty"`
		// The total number of license consumed in the Intersight account.
		LicenseCount *int64 `json:"LicenseCount,omitempty"`
		// The total number of license purchased from cisco.
		LicenseCountPurchased *int64 `json:"LicenseCountPurchased,omitempty"`
		// The license state defined by Intersight. The value may be one of NotLicensed, TrialPeriod, OutOfCompliance, Compliance, GraceExpired, or TrialExpired. * `NotLicensed` - The license token is neither activated nor registered. * `GraceExpired` - The license grace period has expired. * `TrialPeriod` - The 90 days of trial period. * `OutOfCompliance` - The license is out of compliance. * `Compliance` - The license is in compliance. * `TrialExpired` - The trial period of 90 days has expired.
		LicenseState *string `json:"LicenseState,omitempty"`
		// The name of the Intersight license entitlement. For example, this property may be set to 'Essential'. * `Base` - Base as a License type. It is default license type. * `Essential` - Essential as a License type. * `Standard` - Standard as a License type. * `Advantage` - Advantage as a License type. * `Premier` - Premier as a License type. * `IWO-Essential` - IWO-Essential as a License type. * `IWO-Advantage` - IWO-Advantage as a License type. * `IWO-Premier` - IWO-Premier as a License type. * `IKS-Advantage` - IKS-Advantage as a License type.
		LicenseType *string `json:"LicenseType,omitempty"`
		// The date and time when the licenseState entered the TrialPeriod or OutOfCompliance state.
		StartTime *time.Time `json:"StartTime,omitempty"`
		// The administrative state of the trial license. When the LicenseState is set to 'NotLicensed', 'trialAdmin' can be set to true to start the trial period, i.e. licenseState is set to be TrialPeriod.
		TrialAdmin         *bool                                  `json:"TrialAdmin,omitempty"`
		AccountLicenseData *LicenseAccountLicenseDataRelationship `json:"AccountLicenseData,omitempty"`
	}

	varLicenseLicenseInfoWithoutEmbeddedStruct := LicenseLicenseInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varLicenseLicenseInfoWithoutEmbeddedStruct)
	if err == nil {
		varLicenseLicenseInfo := _LicenseLicenseInfo{}
		varLicenseLicenseInfo.ClassId = varLicenseLicenseInfoWithoutEmbeddedStruct.ClassId
		varLicenseLicenseInfo.ObjectType = varLicenseLicenseInfoWithoutEmbeddedStruct.ObjectType
		varLicenseLicenseInfo.ActiveAdmin = varLicenseLicenseInfoWithoutEmbeddedStruct.ActiveAdmin
		varLicenseLicenseInfo.DaysLeft = varLicenseLicenseInfoWithoutEmbeddedStruct.DaysLeft
		varLicenseLicenseInfo.EndTime = varLicenseLicenseInfoWithoutEmbeddedStruct.EndTime
		varLicenseLicenseInfo.EnforceMode = varLicenseLicenseInfoWithoutEmbeddedStruct.EnforceMode
		varLicenseLicenseInfo.ErrorDesc = varLicenseLicenseInfoWithoutEmbeddedStruct.ErrorDesc
		varLicenseLicenseInfo.EvaluationPeriod = varLicenseLicenseInfoWithoutEmbeddedStruct.EvaluationPeriod
		varLicenseLicenseInfo.ExpireTime = varLicenseLicenseInfoWithoutEmbeddedStruct.ExpireTime
		varLicenseLicenseInfo.ExtraEvaluation = varLicenseLicenseInfoWithoutEmbeddedStruct.ExtraEvaluation
		varLicenseLicenseInfo.LicenseCount = varLicenseLicenseInfoWithoutEmbeddedStruct.LicenseCount
		varLicenseLicenseInfo.LicenseCountPurchased = varLicenseLicenseInfoWithoutEmbeddedStruct.LicenseCountPurchased
		varLicenseLicenseInfo.LicenseState = varLicenseLicenseInfoWithoutEmbeddedStruct.LicenseState
		varLicenseLicenseInfo.LicenseType = varLicenseLicenseInfoWithoutEmbeddedStruct.LicenseType
		varLicenseLicenseInfo.StartTime = varLicenseLicenseInfoWithoutEmbeddedStruct.StartTime
		varLicenseLicenseInfo.TrialAdmin = varLicenseLicenseInfoWithoutEmbeddedStruct.TrialAdmin
		varLicenseLicenseInfo.AccountLicenseData = varLicenseLicenseInfoWithoutEmbeddedStruct.AccountLicenseData
		*o = LicenseLicenseInfo(varLicenseLicenseInfo)
	} else {
		return err
	}

	varLicenseLicenseInfo := _LicenseLicenseInfo{}

	err = json.Unmarshal(bytes, &varLicenseLicenseInfo)
	if err == nil {
		o.MoBaseMo = varLicenseLicenseInfo.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActiveAdmin")
		delete(additionalProperties, "DaysLeft")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "EnforceMode")
		delete(additionalProperties, "ErrorDesc")
		delete(additionalProperties, "EvaluationPeriod")
		delete(additionalProperties, "ExpireTime")
		delete(additionalProperties, "ExtraEvaluation")
		delete(additionalProperties, "LicenseCount")
		delete(additionalProperties, "LicenseCountPurchased")
		delete(additionalProperties, "LicenseState")
		delete(additionalProperties, "LicenseType")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "TrialAdmin")
		delete(additionalProperties, "AccountLicenseData")

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

type NullableLicenseLicenseInfo struct {
	value *LicenseLicenseInfo
	isSet bool
}

func (v NullableLicenseLicenseInfo) Get() *LicenseLicenseInfo {
	return v.value
}

func (v *NullableLicenseLicenseInfo) Set(val *LicenseLicenseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseLicenseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseLicenseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseLicenseInfo(val *LicenseLicenseInfo) *NullableLicenseLicenseInfo {
	return &NullableLicenseLicenseInfo{value: val, isSet: true}
}

func (v NullableLicenseLicenseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseLicenseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}