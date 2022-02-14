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

// BulkRequestAllOf Definition of the list of properties defined in 'bulk.Request', excluding properties defined in parent classes.
type BulkRequestAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The action to be taken when an error occurs during processing of the request. * `Stop` - Stop the processing of the request after the first error. * `Proceed` - Proceed with the processing of the request even when an error occurs.
	ActionOnError *string  `json:"ActionOnError,omitempty"`
	Actions       []string `json:"Actions,omitempty"`
	// The timestamp when the request processing completed.
	CompletionTime *string          `json:"CompletionTime,omitempty"`
	Headers        []BulkHttpHeader `json:"Headers,omitempty"`
	// The number of sub requests received in this request.
	NumSubRequests *int64 `json:"NumSubRequests,omitempty"`
	// The moid of the organization under which this request was issued.
	OrgMoid *string `json:"OrgMoid,omitempty"`
	// The timestamp when the request was received.
	RequestReceivedTime *string          `json:"RequestReceivedTime,omitempty"`
	Requests            []BulkSubRequest `json:"Requests,omitempty"`
	Results             []BulkApiResult  `json:"Results,omitempty"`
	// Skip the already present objects.
	SkipDuplicates *bool `json:"SkipDuplicates,omitempty"`
	// The processing status of the Request. * `NotStarted` - Indicates that the request processing has not begun yet. * `ObjPresenceCheckInProgress` - Indicates that the object presence check is in progress for this request. * `ObjPresenceCheckComplete` - Indicates that the object presence check is complete. * `ExecutionInProgress` - Indicates that the request processing is in progress. * `Completed` - Indicates that the request processing has been completed successfully. * `Failed` - Indicates that the processing of this request failed.
	Status *string `json:"Status,omitempty"`
	// The status message corresponding to the status.
	StatusMessage *string `json:"StatusMessage,omitempty"`
	// The URI on which this bulk action is to be performed. The value will be used when there is no override in the SubRequest.
	Uri *string `json:"Uri,omitempty"`
	// The type of operation to be performed. One of - Post (Create), Patch (Update) or Delete (Remove). The value will be used when there is no override in the SubRequest. * `POST` - Used to create a REST resource. * `PATCH` - Used to update a REST resource. * `DELETE` - Used to delete a REST resource.
	Verb *string `json:"Verb,omitempty"`
	// An array of relationships to bulkSubRequestObj resources.
	AsyncResults []BulkSubRequestObjRelationship `json:"AsyncResults,omitempty"`
	// An array of relationships to bulkSubRequestObj resources.
	AsyncResultsFailed   []BulkSubRequestObjRelationship       `json:"AsyncResultsFailed,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	WorkflowInfo         *WorkflowWorkflowInfoRelationship     `json:"WorkflowInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkRequestAllOf BulkRequestAllOf

// NewBulkRequestAllOf instantiates a new BulkRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkRequestAllOf(classId string, objectType string) *BulkRequestAllOf {
	this := BulkRequestAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var actionOnError string = "Stop"
	this.ActionOnError = &actionOnError
	var status string = "NotStarted"
	this.Status = &status
	var verb string = "POST"
	this.Verb = &verb
	return &this
}

// NewBulkRequestAllOfWithDefaults instantiates a new BulkRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkRequestAllOfWithDefaults() *BulkRequestAllOf {
	this := BulkRequestAllOf{}
	var classId string = "bulk.Request"
	this.ClassId = classId
	var objectType string = "bulk.Request"
	this.ObjectType = objectType
	var actionOnError string = "Stop"
	this.ActionOnError = &actionOnError
	var status string = "NotStarted"
	this.Status = &status
	var verb string = "POST"
	this.Verb = &verb
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkRequestAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkRequestAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkRequestAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BulkRequestAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkRequestAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkRequestAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActionOnError returns the ActionOnError field value if set, zero value otherwise.
func (o *BulkRequestAllOf) GetActionOnError() string {
	if o == nil || o.ActionOnError == nil {
		var ret string
		return ret
	}
	return *o.ActionOnError
}

// GetActionOnErrorOk returns a tuple with the ActionOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequestAllOf) GetActionOnErrorOk() (*string, bool) {
	if o == nil || o.ActionOnError == nil {
		return nil, false
	}
	return o.ActionOnError, true
}

// HasActionOnError returns a boolean if a field has been set.
func (o *BulkRequestAllOf) HasActionOnError() bool {
	if o != nil && o.ActionOnError != nil {
		return true
	}

	return false
}

// SetActionOnError gets a reference to the given string and assigns it to the ActionOnError field.
func (o *BulkRequestAllOf) SetActionOnError(v string) {
	o.ActionOnError = &v
}

// GetActions returns the Actions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkRequestAllOf) GetActions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkRequestAllOf) GetActionsOk() (*[]string, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return &o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *BulkRequestAllOf) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []string and assigns it to the Actions field.
func (o *BulkRequestAllOf) SetActions(v []string) {
	o.Actions = v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *BulkRequestAllOf) GetCompletionTime() string {
	if o == nil || o.CompletionTime == nil {
		var ret string
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequestAllOf) GetCompletionTimeOk() (*string, bool) {
	if o == nil || o.CompletionTime == nil {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *BulkRequestAllOf) HasCompletionTime() bool {
	if o != nil && o.CompletionTime != nil {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given string and assigns it to the CompletionTime field.
func (o *BulkRequestAllOf) SetCompletionTime(v string) {
	o.CompletionTime = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkRequestAllOf) GetHeaders() []BulkHttpHeader {
	if o == nil {
		var ret []BulkHttpHeader
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkRequestAllOf) GetHeadersOk() (*[]BulkHttpHeader, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *BulkRequestAllOf) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []BulkHttpHeader and assigns it to the Headers field.
func (o *BulkRequestAllOf) SetHeaders(v []BulkHttpHeader) {
	o.Headers = v
}

// GetNumSubRequests returns the NumSubRequests field value if set, zero value otherwise.
func (o *BulkRequestAllOf) GetNumSubRequests() int64 {
	if o == nil || o.NumSubRequests == nil {
		var ret int64
		return ret
	}
	return *o.NumSubRequests
}

// GetNumSubRequestsOk returns a tuple with the NumSubRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequestAllOf) GetNumSubRequestsOk() (*int64, bool) {
	if o == nil || o.NumSubRequests == nil {
		return nil, false
	}
	return o.NumSubRequests, true
}

// HasNumSubRequests returns a boolean if a field has been set.
func (o *BulkRequestAllOf) HasNumSubRequests() bool {
	if o != nil && o.NumSubRequests != nil {
		return true
	}

	return false
}

// SetNumSubRequests gets a reference to the given int64 and assigns it to the NumSubRequests field.
func (o *BulkRequestAllOf) SetNumSubRequests(v int64) {
	o.NumSubRequests = &v
}

// GetOrgMoid returns the OrgMoid field value if set, zero value otherwise.
func (o *BulkRequestAllOf) GetOrgMoid() string {
	if o == nil || o.OrgMoid == nil {
		var ret string
		return ret
	}
	return *o.OrgMoid
}

// GetOrgMoidOk returns a tuple with the OrgMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequestAllOf) GetOrgMoidOk() (*string, bool) {
	if o == nil || o.OrgMoid == nil {
		return nil, false
	}
	return o.OrgMoid, true
}

// HasOrgMoid returns a boolean if a field has been set.
func (o *BulkRequestAllOf) HasOrgMoid() bool {
	if o != nil && o.OrgMoid != nil {
		return true
	}

	return false
}

// SetOrgMoid gets a reference to the given string and assigns it to the OrgMoid field.
func (o *BulkRequestAllOf) SetOrgMoid(v string) {
	o.OrgMoid = &v
}

// GetRequestReceivedTime returns the RequestReceivedTime field value if set, zero value otherwise.
func (o *BulkRequestAllOf) GetRequestReceivedTime() string {
	if o == nil || o.RequestReceivedTime == nil {
		var ret string
		return ret
	}
	return *o.RequestReceivedTime
}

// GetRequestReceivedTimeOk returns a tuple with the RequestReceivedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequestAllOf) GetRequestReceivedTimeOk() (*string, bool) {
	if o == nil || o.RequestReceivedTime == nil {
		return nil, false
	}
	return o.RequestReceivedTime, true
}

// HasRequestReceivedTime returns a boolean if a field has been set.
func (o *BulkRequestAllOf) HasRequestReceivedTime() bool {
	if o != nil && o.RequestReceivedTime != nil {
		return true
	}

	return false
}

// SetRequestReceivedTime gets a reference to the given string and assigns it to the RequestReceivedTime field.
func (o *BulkRequestAllOf) SetRequestReceivedTime(v string) {
	o.RequestReceivedTime = &v
}

// GetRequests returns the Requests field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkRequestAllOf) GetRequests() []BulkSubRequest {
	if o == nil {
		var ret []BulkSubRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkRequestAllOf) GetRequestsOk() (*[]BulkSubRequest, bool) {
	if o == nil || o.Requests == nil {
		return nil, false
	}
	return &o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *BulkRequestAllOf) HasRequests() bool {
	if o != nil && o.Requests != nil {
		return true
	}

	return false
}

// SetRequests gets a reference to the given []BulkSubRequest and assigns it to the Requests field.
func (o *BulkRequestAllOf) SetRequests(v []BulkSubRequest) {
	o.Requests = v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkRequestAllOf) GetResults() []BulkApiResult {
	if o == nil {
		var ret []BulkApiResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkRequestAllOf) GetResultsOk() (*[]BulkApiResult, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return &o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *BulkRequestAllOf) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []BulkApiResult and assigns it to the Results field.
func (o *BulkRequestAllOf) SetResults(v []BulkApiResult) {
	o.Results = v
}

// GetSkipDuplicates returns the SkipDuplicates field value if set, zero value otherwise.
func (o *BulkRequestAllOf) GetSkipDuplicates() bool {
	if o == nil || o.SkipDuplicates == nil {
		var ret bool
		return ret
	}
	return *o.SkipDuplicates
}

// GetSkipDuplicatesOk returns a tuple with the SkipDuplicates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequestAllOf) GetSkipDuplicatesOk() (*bool, bool) {
	if o == nil || o.SkipDuplicates == nil {
		return nil, false
	}
	return o.SkipDuplicates, true
}

// HasSkipDuplicates returns a boolean if a field has been set.
func (o *BulkRequestAllOf) HasSkipDuplicates() bool {
	if o != nil && o.SkipDuplicates != nil {
		return true
	}

	return false
}

// SetSkipDuplicates gets a reference to the given bool and assigns it to the SkipDuplicates field.
func (o *BulkRequestAllOf) SetSkipDuplicates(v bool) {
	o.SkipDuplicates = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BulkRequestAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequestAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BulkRequestAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BulkRequestAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *BulkRequestAllOf) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequestAllOf) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *BulkRequestAllOf) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *BulkRequestAllOf) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *BulkRequestAllOf) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequestAllOf) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *BulkRequestAllOf) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *BulkRequestAllOf) SetUri(v string) {
	o.Uri = &v
}

// GetVerb returns the Verb field value if set, zero value otherwise.
func (o *BulkRequestAllOf) GetVerb() string {
	if o == nil || o.Verb == nil {
		var ret string
		return ret
	}
	return *o.Verb
}

// GetVerbOk returns a tuple with the Verb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequestAllOf) GetVerbOk() (*string, bool) {
	if o == nil || o.Verb == nil {
		return nil, false
	}
	return o.Verb, true
}

// HasVerb returns a boolean if a field has been set.
func (o *BulkRequestAllOf) HasVerb() bool {
	if o != nil && o.Verb != nil {
		return true
	}

	return false
}

// SetVerb gets a reference to the given string and assigns it to the Verb field.
func (o *BulkRequestAllOf) SetVerb(v string) {
	o.Verb = &v
}

// GetAsyncResults returns the AsyncResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkRequestAllOf) GetAsyncResults() []BulkSubRequestObjRelationship {
	if o == nil {
		var ret []BulkSubRequestObjRelationship
		return ret
	}
	return o.AsyncResults
}

// GetAsyncResultsOk returns a tuple with the AsyncResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkRequestAllOf) GetAsyncResultsOk() (*[]BulkSubRequestObjRelationship, bool) {
	if o == nil || o.AsyncResults == nil {
		return nil, false
	}
	return &o.AsyncResults, true
}

// HasAsyncResults returns a boolean if a field has been set.
func (o *BulkRequestAllOf) HasAsyncResults() bool {
	if o != nil && o.AsyncResults != nil {
		return true
	}

	return false
}

// SetAsyncResults gets a reference to the given []BulkSubRequestObjRelationship and assigns it to the AsyncResults field.
func (o *BulkRequestAllOf) SetAsyncResults(v []BulkSubRequestObjRelationship) {
	o.AsyncResults = v
}

// GetAsyncResultsFailed returns the AsyncResultsFailed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkRequestAllOf) GetAsyncResultsFailed() []BulkSubRequestObjRelationship {
	if o == nil {
		var ret []BulkSubRequestObjRelationship
		return ret
	}
	return o.AsyncResultsFailed
}

// GetAsyncResultsFailedOk returns a tuple with the AsyncResultsFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkRequestAllOf) GetAsyncResultsFailedOk() (*[]BulkSubRequestObjRelationship, bool) {
	if o == nil || o.AsyncResultsFailed == nil {
		return nil, false
	}
	return &o.AsyncResultsFailed, true
}

// HasAsyncResultsFailed returns a boolean if a field has been set.
func (o *BulkRequestAllOf) HasAsyncResultsFailed() bool {
	if o != nil && o.AsyncResultsFailed != nil {
		return true
	}

	return false
}

// SetAsyncResultsFailed gets a reference to the given []BulkSubRequestObjRelationship and assigns it to the AsyncResultsFailed field.
func (o *BulkRequestAllOf) SetAsyncResultsFailed(v []BulkSubRequestObjRelationship) {
	o.AsyncResultsFailed = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *BulkRequestAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequestAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *BulkRequestAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *BulkRequestAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise.
func (o *BulkRequestAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.WorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRequestAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.WorkflowInfo == nil {
		return nil, false
	}
	return o.WorkflowInfo, true
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *BulkRequestAllOf) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo != nil {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *BulkRequestAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo = &v
}

func (o BulkRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ActionOnError != nil {
		toSerialize["ActionOnError"] = o.ActionOnError
	}
	if o.Actions != nil {
		toSerialize["Actions"] = o.Actions
	}
	if o.CompletionTime != nil {
		toSerialize["CompletionTime"] = o.CompletionTime
	}
	if o.Headers != nil {
		toSerialize["Headers"] = o.Headers
	}
	if o.NumSubRequests != nil {
		toSerialize["NumSubRequests"] = o.NumSubRequests
	}
	if o.OrgMoid != nil {
		toSerialize["OrgMoid"] = o.OrgMoid
	}
	if o.RequestReceivedTime != nil {
		toSerialize["RequestReceivedTime"] = o.RequestReceivedTime
	}
	if o.Requests != nil {
		toSerialize["Requests"] = o.Requests
	}
	if o.Results != nil {
		toSerialize["Results"] = o.Results
	}
	if o.SkipDuplicates != nil {
		toSerialize["SkipDuplicates"] = o.SkipDuplicates
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusMessage != nil {
		toSerialize["StatusMessage"] = o.StatusMessage
	}
	if o.Uri != nil {
		toSerialize["Uri"] = o.Uri
	}
	if o.Verb != nil {
		toSerialize["Verb"] = o.Verb
	}
	if o.AsyncResults != nil {
		toSerialize["AsyncResults"] = o.AsyncResults
	}
	if o.AsyncResultsFailed != nil {
		toSerialize["AsyncResultsFailed"] = o.AsyncResultsFailed
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.WorkflowInfo != nil {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BulkRequestAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBulkRequestAllOf := _BulkRequestAllOf{}

	if err = json.Unmarshal(bytes, &varBulkRequestAllOf); err == nil {
		*o = BulkRequestAllOf(varBulkRequestAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActionOnError")
		delete(additionalProperties, "Actions")
		delete(additionalProperties, "CompletionTime")
		delete(additionalProperties, "Headers")
		delete(additionalProperties, "NumSubRequests")
		delete(additionalProperties, "OrgMoid")
		delete(additionalProperties, "RequestReceivedTime")
		delete(additionalProperties, "Requests")
		delete(additionalProperties, "Results")
		delete(additionalProperties, "SkipDuplicates")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusMessage")
		delete(additionalProperties, "Uri")
		delete(additionalProperties, "Verb")
		delete(additionalProperties, "AsyncResults")
		delete(additionalProperties, "AsyncResultsFailed")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "WorkflowInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkRequestAllOf struct {
	value *BulkRequestAllOf
	isSet bool
}

func (v NullableBulkRequestAllOf) Get() *BulkRequestAllOf {
	return v.value
}

func (v *NullableBulkRequestAllOf) Set(val *BulkRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkRequestAllOf(val *BulkRequestAllOf) *NullableBulkRequestAllOf {
	return &NullableBulkRequestAllOf{value: val, isSet: true}
}

func (v NullableBulkRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
