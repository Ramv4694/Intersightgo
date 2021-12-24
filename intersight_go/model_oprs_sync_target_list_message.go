/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-12-24T09:42:08Z.
 *
 * API version: 0.0.1-37430
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

// OprsSyncTargetListMessage The targets sync messages are sent to assist and back to euclid for reconciliation.
type OprsSyncTargetListMessage struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The unique id of the assist which is associated with this message.
	AssistId *string `json:"AssistId,omitempty"`
	// The target type for which this sync message is meant for. * `` - The device reported an empty or unrecognized platform type. * `APIC` - An Application Policy Infrastructure Controller cluster. * `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * `IMC` - A standalone UCS Server Integrated Management Controller. * `IMCM4` - A standalone UCS M4 Server. * `IMCM5` - A standalone UCS M5 server. * `IMCRack` - A standalone UCS M6 and above server. * `UCSIOM` - An UCS Chassis IO module. * `HX` - A HyperFlex storage controller. * `HyperFlexAP` - A HyperFlex Application Platform. * `IWE` - An Intersight Workload Engine. * `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance. * `IntersightAssist` - A Cisco Intersight Assist. * `PureStorageFlashArray` - A Pure Storage FlashArray device. * `NexusDevice` - A generic platform type to support Nexus Network Device. This can also be extended to support all network devices later on. * `MDSDevice` - A platform type to support Nexus MDS devices. * `UCSC890` - A standalone Cisco UCSC890 server. * `NetAppOntap` - A NetApp ONTAP storage system. * `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager. * `EmcScaleIo` - An EMC ScaleIO storage system. * `EmcVmax` - An EMC VMAX storage system. * `EmcVplex` - An EMC VPLEX storage system. * `EmcXtremIo` - An EMC XtremIO storage system. * `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines. * `MicrosoftHyperV` - A Microsoft Hyper-V system that manages Virtual Machines. * `AppDynamics` - An AppDynamics controller that monitors applications. * `Dynatrace` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `NewRelic` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `ServiceNow` - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications. * `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `MicrosoftSqlServer` - A Microsoft SQL database server. * `Kubernetes` - A Kubernetes cluster that runs containerized applications. * `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost. * `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket. * `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions. * `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs. * `DellCompellent` - A Dell Compellent storage system. * `HPE3Par` - A HPE 3PAR storage system. * `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform. * `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers. * `IMCBlade` - An Intersight managed UCS Blade Server. * `TerraformCloud` - A Terraform Cloud account. * `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent. * `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * `AnsibleEndpoint` - An external endpoint added as Target that can be accessed through Ansible in Intersight Cloud Orchestrator automation workflow. * `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic, Bearer Token. * `SSHEndpoint` - An external endpoint added as Target that can be accessed through SSH in Intersight Cloud Orchestrator automation workflow. * `CiscoCatalyst` - A Cisco Catalyst networking switch device. * `PowerShellEndpoint` - A Windows machine on which PowerShell scripts can be executed remotely.
	TargetType *string  `json:"TargetType,omitempty"`
	Targets    []string `json:"Targets,omitempty"`
	// The time at which the event was generated. Date is accurate to Intersights clock. This time will be used to identify order of events.
	TimeStamp            *time.Time                           `json:"TimeStamp,omitempty"`
	Assist               *AssetDeviceRegistrationRelationship `json:"Assist,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OprsSyncTargetListMessage OprsSyncTargetListMessage

// NewOprsSyncTargetListMessage instantiates a new OprsSyncTargetListMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOprsSyncTargetListMessage(classId string, objectType string) *OprsSyncTargetListMessage {
	this := OprsSyncTargetListMessage{}
	this.ClassId = classId
	this.ObjectType = objectType
	var targetType string = ""
	this.TargetType = &targetType
	return &this
}

// NewOprsSyncTargetListMessageWithDefaults instantiates a new OprsSyncTargetListMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOprsSyncTargetListMessageWithDefaults() *OprsSyncTargetListMessage {
	this := OprsSyncTargetListMessage{}
	var classId string = "oprs.SyncTargetListMessage"
	this.ClassId = classId
	var objectType string = "oprs.SyncTargetListMessage"
	this.ObjectType = objectType
	var targetType string = ""
	this.TargetType = &targetType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OprsSyncTargetListMessage) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OprsSyncTargetListMessage) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OprsSyncTargetListMessage) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OprsSyncTargetListMessage) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OprsSyncTargetListMessage) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OprsSyncTargetListMessage) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAssistId returns the AssistId field value if set, zero value otherwise.
func (o *OprsSyncTargetListMessage) GetAssistId() string {
	if o == nil || o.AssistId == nil {
		var ret string
		return ret
	}
	return *o.AssistId
}

// GetAssistIdOk returns a tuple with the AssistId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsSyncTargetListMessage) GetAssistIdOk() (*string, bool) {
	if o == nil || o.AssistId == nil {
		return nil, false
	}
	return o.AssistId, true
}

// HasAssistId returns a boolean if a field has been set.
func (o *OprsSyncTargetListMessage) HasAssistId() bool {
	if o != nil && o.AssistId != nil {
		return true
	}

	return false
}

// SetAssistId gets a reference to the given string and assigns it to the AssistId field.
func (o *OprsSyncTargetListMessage) SetAssistId(v string) {
	o.AssistId = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *OprsSyncTargetListMessage) GetTargetType() string {
	if o == nil || o.TargetType == nil {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsSyncTargetListMessage) GetTargetTypeOk() (*string, bool) {
	if o == nil || o.TargetType == nil {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *OprsSyncTargetListMessage) HasTargetType() bool {
	if o != nil && o.TargetType != nil {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *OprsSyncTargetListMessage) SetTargetType(v string) {
	o.TargetType = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OprsSyncTargetListMessage) GetTargets() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OprsSyncTargetListMessage) GetTargetsOk() (*[]string, bool) {
	if o == nil || o.Targets == nil {
		return nil, false
	}
	return &o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *OprsSyncTargetListMessage) HasTargets() bool {
	if o != nil && o.Targets != nil {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []string and assigns it to the Targets field.
func (o *OprsSyncTargetListMessage) SetTargets(v []string) {
	o.Targets = v
}

// GetTimeStamp returns the TimeStamp field value if set, zero value otherwise.
func (o *OprsSyncTargetListMessage) GetTimeStamp() time.Time {
	if o == nil || o.TimeStamp == nil {
		var ret time.Time
		return ret
	}
	return *o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsSyncTargetListMessage) GetTimeStampOk() (*time.Time, bool) {
	if o == nil || o.TimeStamp == nil {
		return nil, false
	}
	return o.TimeStamp, true
}

// HasTimeStamp returns a boolean if a field has been set.
func (o *OprsSyncTargetListMessage) HasTimeStamp() bool {
	if o != nil && o.TimeStamp != nil {
		return true
	}

	return false
}

// SetTimeStamp gets a reference to the given time.Time and assigns it to the TimeStamp field.
func (o *OprsSyncTargetListMessage) SetTimeStamp(v time.Time) {
	o.TimeStamp = &v
}

// GetAssist returns the Assist field value if set, zero value otherwise.
func (o *OprsSyncTargetListMessage) GetAssist() AssetDeviceRegistrationRelationship {
	if o == nil || o.Assist == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Assist
}

// GetAssistOk returns a tuple with the Assist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsSyncTargetListMessage) GetAssistOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Assist == nil {
		return nil, false
	}
	return o.Assist, true
}

// HasAssist returns a boolean if a field has been set.
func (o *OprsSyncTargetListMessage) HasAssist() bool {
	if o != nil && o.Assist != nil {
		return true
	}

	return false
}

// SetAssist gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Assist field.
func (o *OprsSyncTargetListMessage) SetAssist(v AssetDeviceRegistrationRelationship) {
	o.Assist = &v
}

func (o OprsSyncTargetListMessage) MarshalJSON() ([]byte, error) {
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
	if o.AssistId != nil {
		toSerialize["AssistId"] = o.AssistId
	}
	if o.TargetType != nil {
		toSerialize["TargetType"] = o.TargetType
	}
	if o.Targets != nil {
		toSerialize["Targets"] = o.Targets
	}
	if o.TimeStamp != nil {
		toSerialize["TimeStamp"] = o.TimeStamp
	}
	if o.Assist != nil {
		toSerialize["Assist"] = o.Assist
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OprsSyncTargetListMessage) UnmarshalJSON(bytes []byte) (err error) {
	type OprsSyncTargetListMessageWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The unique id of the assist which is associated with this message.
		AssistId *string `json:"AssistId,omitempty"`
		// The target type for which this sync message is meant for. * `` - The device reported an empty or unrecognized platform type. * `APIC` - An Application Policy Infrastructure Controller cluster. * `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * `IMC` - A standalone UCS Server Integrated Management Controller. * `IMCM4` - A standalone UCS M4 Server. * `IMCM5` - A standalone UCS M5 server. * `IMCRack` - A standalone UCS M6 and above server. * `UCSIOM` - An UCS Chassis IO module. * `HX` - A HyperFlex storage controller. * `HyperFlexAP` - A HyperFlex Application Platform. * `IWE` - An Intersight Workload Engine. * `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance. * `IntersightAssist` - A Cisco Intersight Assist. * `PureStorageFlashArray` - A Pure Storage FlashArray device. * `NexusDevice` - A generic platform type to support Nexus Network Device. This can also be extended to support all network devices later on. * `MDSDevice` - A platform type to support Nexus MDS devices. * `UCSC890` - A standalone Cisco UCSC890 server. * `NetAppOntap` - A NetApp ONTAP storage system. * `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager. * `EmcScaleIo` - An EMC ScaleIO storage system. * `EmcVmax` - An EMC VMAX storage system. * `EmcVplex` - An EMC VPLEX storage system. * `EmcXtremIo` - An EMC XtremIO storage system. * `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines. * `MicrosoftHyperV` - A Microsoft Hyper-V system that manages Virtual Machines. * `AppDynamics` - An AppDynamics controller that monitors applications. * `Dynatrace` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `NewRelic` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `ServiceNow` - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications. * `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `MicrosoftSqlServer` - A Microsoft SQL database server. * `Kubernetes` - A Kubernetes cluster that runs containerized applications. * `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost. * `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket. * `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions. * `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs. * `DellCompellent` - A Dell Compellent storage system. * `HPE3Par` - A HPE 3PAR storage system. * `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform. * `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers. * `IMCBlade` - An Intersight managed UCS Blade Server. * `TerraformCloud` - A Terraform Cloud account. * `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent. * `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * `AnsibleEndpoint` - An external endpoint added as Target that can be accessed through Ansible in Intersight Cloud Orchestrator automation workflow. * `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic, Bearer Token. * `SSHEndpoint` - An external endpoint added as Target that can be accessed through SSH in Intersight Cloud Orchestrator automation workflow. * `CiscoCatalyst` - A Cisco Catalyst networking switch device. * `PowerShellEndpoint` - A Windows machine on which PowerShell scripts can be executed remotely.
		TargetType *string  `json:"TargetType,omitempty"`
		Targets    []string `json:"Targets,omitempty"`
		// The time at which the event was generated. Date is accurate to Intersights clock. This time will be used to identify order of events.
		TimeStamp *time.Time                           `json:"TimeStamp,omitempty"`
		Assist    *AssetDeviceRegistrationRelationship `json:"Assist,omitempty"`
	}

	varOprsSyncTargetListMessageWithoutEmbeddedStruct := OprsSyncTargetListMessageWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varOprsSyncTargetListMessageWithoutEmbeddedStruct)
	if err == nil {
		varOprsSyncTargetListMessage := _OprsSyncTargetListMessage{}
		varOprsSyncTargetListMessage.ClassId = varOprsSyncTargetListMessageWithoutEmbeddedStruct.ClassId
		varOprsSyncTargetListMessage.ObjectType = varOprsSyncTargetListMessageWithoutEmbeddedStruct.ObjectType
		varOprsSyncTargetListMessage.AssistId = varOprsSyncTargetListMessageWithoutEmbeddedStruct.AssistId
		varOprsSyncTargetListMessage.TargetType = varOprsSyncTargetListMessageWithoutEmbeddedStruct.TargetType
		varOprsSyncTargetListMessage.Targets = varOprsSyncTargetListMessageWithoutEmbeddedStruct.Targets
		varOprsSyncTargetListMessage.TimeStamp = varOprsSyncTargetListMessageWithoutEmbeddedStruct.TimeStamp
		varOprsSyncTargetListMessage.Assist = varOprsSyncTargetListMessageWithoutEmbeddedStruct.Assist
		*o = OprsSyncTargetListMessage(varOprsSyncTargetListMessage)
	} else {
		return err
	}

	varOprsSyncTargetListMessage := _OprsSyncTargetListMessage{}

	err = json.Unmarshal(bytes, &varOprsSyncTargetListMessage)
	if err == nil {
		o.MoBaseMo = varOprsSyncTargetListMessage.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AssistId")
		delete(additionalProperties, "TargetType")
		delete(additionalProperties, "Targets")
		delete(additionalProperties, "TimeStamp")
		delete(additionalProperties, "Assist")

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

type NullableOprsSyncTargetListMessage struct {
	value *OprsSyncTargetListMessage
	isSet bool
}

func (v NullableOprsSyncTargetListMessage) Get() *OprsSyncTargetListMessage {
	return v.value
}

func (v *NullableOprsSyncTargetListMessage) Set(val *OprsSyncTargetListMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableOprsSyncTargetListMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableOprsSyncTargetListMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOprsSyncTargetListMessage(val *OprsSyncTargetListMessage) *NullableOprsSyncTargetListMessage {
	return &NullableOprsSyncTargetListMessage{value: val, isSet: true}
}

func (v NullableOprsSyncTargetListMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOprsSyncTargetListMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
