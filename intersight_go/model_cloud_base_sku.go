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

// CloudBaseSku Contains the details of an instance type.
type CloudBaseSku struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType       string                  `json:"ObjectType"`
	CustomAttributes []CloudCustomAttributes `json:"CustomAttributes,omitempty"`
	// Any additional description for the instance type.
	Description *string `json:"Description,omitempty"`
	// Flag to indicate if this SKU is active or not.
	IsActive *bool `json:"IsActive,omitempty"`
	// Flag to indicate if SKU is discovered during inventory collection.
	IsAutoDiscovered *bool `json:"IsAutoDiscovered,omitempty"`
	// The display name for instance type.
	Name *string `json:"Name,omitempty"`
	// The platformType for this SKU. * `` - The device reported an empty or unrecognized platform type. * `APIC` - An Application Policy Infrastructure Controller cluster. * `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * `IMC` - A standalone UCS Server Integrated Management Controller. * `IMCM4` - A standalone UCS M4 Server. * `IMCM5` - A standalone UCS M5 server. * `IMCRack` - A standalone UCS M6 and above server. * `UCSIOM` - An UCS Chassis IO module. * `HX` - A HyperFlex storage controller. * `HyperFlexAP` - A HyperFlex Application Platform. * `IWE` - An Intersight Workload Engine. * `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance. * `IntersightAssist` - A Cisco Intersight Assist. * `PureStorageFlashArray` - A Pure Storage FlashArray device. * `UCSC890` - A standalone Cisco UCSC890 server. * `NetAppOntap` - A NetApp ONTAP storage system. * `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager. * `EmcScaleIo` - An EMC ScaleIO storage system. * `EmcVmax` - An EMC VMAX storage system. * `EmcVplex` - An EMC VPLEX storage system. * `EmcXtremIo` - An EMC XtremIO storage system. * `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines. * `MicrosoftHyperV` - A Microsoft Hyper-V system that manages Virtual Machines. * `AppDynamics` - An AppDynamics controller that monitors applications. * `Dynatrace` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `NewRelic` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications. * `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `MicrosoftSqlServer` - A Microsoft SQL database server. * `Kubernetes` - A Kubernetes cluster that runs containerized applications. * `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost. * `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket. * `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions. * `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs. * `DellCompellent` - A Dell Compellent storage system. * `HPE3Par` - A HPE 3PAR storage system. * `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform. * `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers. * `IMCBlade` - An Intersight managed UCS Blade Server. * `TerraformCloud` - A Terraform Cloud account. * `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent. * `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * `AnsibleEndpoint` - An external endpoint added as Target that can be accessed through Ansible in Intersight Cloud Orchestrator automation workflow. * `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic, Bearer Token. * `SSHEndpoint` - An external endpoint added as Target that can be accessed through SSH in Intersight Cloud Orchestrator automation workflow. * `CiscoCatalyst` - A Cisco Catalyst networking switch device.
	PlatformType *string `json:"PlatformType,omitempty"`
	// Indicates if this sku belongs to Compute, Storage, Database or Network category. * `Compute` - Compute service offered by cloud provider. * `Storage` - Storage service offered by cloud provider. * `Database` - Database service offered by cloud provider. * `Network` - Network service offered by cloud provider.
	ServiceCategory *string `json:"ServiceCategory,omitempty"`
	// Property to identify the family of service that the sku belongs to.
	ServiceFamily *string `json:"ServiceFamily,omitempty"`
	// Any display name for the ServiceCategory if available.
	ServiceName          *string                  `json:"ServiceName,omitempty"`
	Target               *AssetTargetRelationship `json:"Target,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudBaseSku CloudBaseSku

// NewCloudBaseSku instantiates a new CloudBaseSku object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBaseSku(classId string, objectType string) *CloudBaseSku {
	this := CloudBaseSku{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isActive bool = true
	this.IsActive = &isActive
	var platformType string = ""
	this.PlatformType = &platformType
	var serviceCategory string = "Compute"
	this.ServiceCategory = &serviceCategory
	return &this
}

// NewCloudBaseSkuWithDefaults instantiates a new CloudBaseSku object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBaseSkuWithDefaults() *CloudBaseSku {
	this := CloudBaseSku{}
	var isActive bool = true
	this.IsActive = &isActive
	var platformType string = ""
	this.PlatformType = &platformType
	var serviceCategory string = "Compute"
	this.ServiceCategory = &serviceCategory
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudBaseSku) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudBaseSku) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudBaseSku) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudBaseSku) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudBaseSku) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudBaseSku) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCustomAttributes returns the CustomAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseSku) GetCustomAttributes() []CloudCustomAttributes {
	if o == nil {
		var ret []CloudCustomAttributes
		return ret
	}
	return o.CustomAttributes
}

// GetCustomAttributesOk returns a tuple with the CustomAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseSku) GetCustomAttributesOk() (*[]CloudCustomAttributes, bool) {
	if o == nil || o.CustomAttributes == nil {
		return nil, false
	}
	return &o.CustomAttributes, true
}

// HasCustomAttributes returns a boolean if a field has been set.
func (o *CloudBaseSku) HasCustomAttributes() bool {
	if o != nil && o.CustomAttributes != nil {
		return true
	}

	return false
}

// SetCustomAttributes gets a reference to the given []CloudCustomAttributes and assigns it to the CustomAttributes field.
func (o *CloudBaseSku) SetCustomAttributes(v []CloudCustomAttributes) {
	o.CustomAttributes = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudBaseSku) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseSku) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudBaseSku) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudBaseSku) SetDescription(v string) {
	o.Description = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *CloudBaseSku) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseSku) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *CloudBaseSku) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *CloudBaseSku) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsAutoDiscovered returns the IsAutoDiscovered field value if set, zero value otherwise.
func (o *CloudBaseSku) GetIsAutoDiscovered() bool {
	if o == nil || o.IsAutoDiscovered == nil {
		var ret bool
		return ret
	}
	return *o.IsAutoDiscovered
}

// GetIsAutoDiscoveredOk returns a tuple with the IsAutoDiscovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseSku) GetIsAutoDiscoveredOk() (*bool, bool) {
	if o == nil || o.IsAutoDiscovered == nil {
		return nil, false
	}
	return o.IsAutoDiscovered, true
}

// HasIsAutoDiscovered returns a boolean if a field has been set.
func (o *CloudBaseSku) HasIsAutoDiscovered() bool {
	if o != nil && o.IsAutoDiscovered != nil {
		return true
	}

	return false
}

// SetIsAutoDiscovered gets a reference to the given bool and assigns it to the IsAutoDiscovered field.
func (o *CloudBaseSku) SetIsAutoDiscovered(v bool) {
	o.IsAutoDiscovered = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudBaseSku) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseSku) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudBaseSku) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudBaseSku) SetName(v string) {
	o.Name = &v
}

// GetPlatformType returns the PlatformType field value if set, zero value otherwise.
func (o *CloudBaseSku) GetPlatformType() string {
	if o == nil || o.PlatformType == nil {
		var ret string
		return ret
	}
	return *o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseSku) GetPlatformTypeOk() (*string, bool) {
	if o == nil || o.PlatformType == nil {
		return nil, false
	}
	return o.PlatformType, true
}

// HasPlatformType returns a boolean if a field has been set.
func (o *CloudBaseSku) HasPlatformType() bool {
	if o != nil && o.PlatformType != nil {
		return true
	}

	return false
}

// SetPlatformType gets a reference to the given string and assigns it to the PlatformType field.
func (o *CloudBaseSku) SetPlatformType(v string) {
	o.PlatformType = &v
}

// GetServiceCategory returns the ServiceCategory field value if set, zero value otherwise.
func (o *CloudBaseSku) GetServiceCategory() string {
	if o == nil || o.ServiceCategory == nil {
		var ret string
		return ret
	}
	return *o.ServiceCategory
}

// GetServiceCategoryOk returns a tuple with the ServiceCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseSku) GetServiceCategoryOk() (*string, bool) {
	if o == nil || o.ServiceCategory == nil {
		return nil, false
	}
	return o.ServiceCategory, true
}

// HasServiceCategory returns a boolean if a field has been set.
func (o *CloudBaseSku) HasServiceCategory() bool {
	if o != nil && o.ServiceCategory != nil {
		return true
	}

	return false
}

// SetServiceCategory gets a reference to the given string and assigns it to the ServiceCategory field.
func (o *CloudBaseSku) SetServiceCategory(v string) {
	o.ServiceCategory = &v
}

// GetServiceFamily returns the ServiceFamily field value if set, zero value otherwise.
func (o *CloudBaseSku) GetServiceFamily() string {
	if o == nil || o.ServiceFamily == nil {
		var ret string
		return ret
	}
	return *o.ServiceFamily
}

// GetServiceFamilyOk returns a tuple with the ServiceFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseSku) GetServiceFamilyOk() (*string, bool) {
	if o == nil || o.ServiceFamily == nil {
		return nil, false
	}
	return o.ServiceFamily, true
}

// HasServiceFamily returns a boolean if a field has been set.
func (o *CloudBaseSku) HasServiceFamily() bool {
	if o != nil && o.ServiceFamily != nil {
		return true
	}

	return false
}

// SetServiceFamily gets a reference to the given string and assigns it to the ServiceFamily field.
func (o *CloudBaseSku) SetServiceFamily(v string) {
	o.ServiceFamily = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *CloudBaseSku) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseSku) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *CloudBaseSku) HasServiceName() bool {
	if o != nil && o.ServiceName != nil {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *CloudBaseSku) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *CloudBaseSku) GetTarget() AssetTargetRelationship {
	if o == nil || o.Target == nil {
		var ret AssetTargetRelationship
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseSku) GetTargetOk() (*AssetTargetRelationship, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *CloudBaseSku) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given AssetTargetRelationship and assigns it to the Target field.
func (o *CloudBaseSku) SetTarget(v AssetTargetRelationship) {
	o.Target = &v
}

func (o CloudBaseSku) MarshalJSON() ([]byte, error) {
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
	if o.CustomAttributes != nil {
		toSerialize["CustomAttributes"] = o.CustomAttributes
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.IsActive != nil {
		toSerialize["IsActive"] = o.IsActive
	}
	if o.IsAutoDiscovered != nil {
		toSerialize["IsAutoDiscovered"] = o.IsAutoDiscovered
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PlatformType != nil {
		toSerialize["PlatformType"] = o.PlatformType
	}
	if o.ServiceCategory != nil {
		toSerialize["ServiceCategory"] = o.ServiceCategory
	}
	if o.ServiceFamily != nil {
		toSerialize["ServiceFamily"] = o.ServiceFamily
	}
	if o.ServiceName != nil {
		toSerialize["ServiceName"] = o.ServiceName
	}
	if o.Target != nil {
		toSerialize["Target"] = o.Target
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudBaseSku) UnmarshalJSON(bytes []byte) (err error) {
	type CloudBaseSkuWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType       string                  `json:"ObjectType"`
		CustomAttributes []CloudCustomAttributes `json:"CustomAttributes,omitempty"`
		// Any additional description for the instance type.
		Description *string `json:"Description,omitempty"`
		// Flag to indicate if this SKU is active or not.
		IsActive *bool `json:"IsActive,omitempty"`
		// Flag to indicate if SKU is discovered during inventory collection.
		IsAutoDiscovered *bool `json:"IsAutoDiscovered,omitempty"`
		// The display name for instance type.
		Name *string `json:"Name,omitempty"`
		// The platformType for this SKU. * `` - The device reported an empty or unrecognized platform type. * `APIC` - An Application Policy Infrastructure Controller cluster. * `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * `IMC` - A standalone UCS Server Integrated Management Controller. * `IMCM4` - A standalone UCS M4 Server. * `IMCM5` - A standalone UCS M5 server. * `IMCRack` - A standalone UCS M6 and above server. * `UCSIOM` - An UCS Chassis IO module. * `HX` - A HyperFlex storage controller. * `HyperFlexAP` - A HyperFlex Application Platform. * `IWE` - An Intersight Workload Engine. * `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance. * `IntersightAssist` - A Cisco Intersight Assist. * `PureStorageFlashArray` - A Pure Storage FlashArray device. * `UCSC890` - A standalone Cisco UCSC890 server. * `NetAppOntap` - A NetApp ONTAP storage system. * `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager. * `EmcScaleIo` - An EMC ScaleIO storage system. * `EmcVmax` - An EMC VMAX storage system. * `EmcVplex` - An EMC VPLEX storage system. * `EmcXtremIo` - An EMC XtremIO storage system. * `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines. * `MicrosoftHyperV` - A Microsoft Hyper-V system that manages Virtual Machines. * `AppDynamics` - An AppDynamics controller that monitors applications. * `Dynatrace` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `NewRelic` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications. * `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `MicrosoftSqlServer` - A Microsoft SQL database server. * `Kubernetes` - A Kubernetes cluster that runs containerized applications. * `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost. * `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket. * `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions. * `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs. * `DellCompellent` - A Dell Compellent storage system. * `HPE3Par` - A HPE 3PAR storage system. * `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform. * `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers. * `IMCBlade` - An Intersight managed UCS Blade Server. * `TerraformCloud` - A Terraform Cloud account. * `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent. * `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * `AnsibleEndpoint` - An external endpoint added as Target that can be accessed through Ansible in Intersight Cloud Orchestrator automation workflow. * `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic, Bearer Token. * `SSHEndpoint` - An external endpoint added as Target that can be accessed through SSH in Intersight Cloud Orchestrator automation workflow. * `CiscoCatalyst` - A Cisco Catalyst networking switch device.
		PlatformType *string `json:"PlatformType,omitempty"`
		// Indicates if this sku belongs to Compute, Storage, Database or Network category. * `Compute` - Compute service offered by cloud provider. * `Storage` - Storage service offered by cloud provider. * `Database` - Database service offered by cloud provider. * `Network` - Network service offered by cloud provider.
		ServiceCategory *string `json:"ServiceCategory,omitempty"`
		// Property to identify the family of service that the sku belongs to.
		ServiceFamily *string `json:"ServiceFamily,omitempty"`
		// Any display name for the ServiceCategory if available.
		ServiceName *string                  `json:"ServiceName,omitempty"`
		Target      *AssetTargetRelationship `json:"Target,omitempty"`
	}

	varCloudBaseSkuWithoutEmbeddedStruct := CloudBaseSkuWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudBaseSkuWithoutEmbeddedStruct)
	if err == nil {
		varCloudBaseSku := _CloudBaseSku{}
		varCloudBaseSku.ClassId = varCloudBaseSkuWithoutEmbeddedStruct.ClassId
		varCloudBaseSku.ObjectType = varCloudBaseSkuWithoutEmbeddedStruct.ObjectType
		varCloudBaseSku.CustomAttributes = varCloudBaseSkuWithoutEmbeddedStruct.CustomAttributes
		varCloudBaseSku.Description = varCloudBaseSkuWithoutEmbeddedStruct.Description
		varCloudBaseSku.IsActive = varCloudBaseSkuWithoutEmbeddedStruct.IsActive
		varCloudBaseSku.IsAutoDiscovered = varCloudBaseSkuWithoutEmbeddedStruct.IsAutoDiscovered
		varCloudBaseSku.Name = varCloudBaseSkuWithoutEmbeddedStruct.Name
		varCloudBaseSku.PlatformType = varCloudBaseSkuWithoutEmbeddedStruct.PlatformType
		varCloudBaseSku.ServiceCategory = varCloudBaseSkuWithoutEmbeddedStruct.ServiceCategory
		varCloudBaseSku.ServiceFamily = varCloudBaseSkuWithoutEmbeddedStruct.ServiceFamily
		varCloudBaseSku.ServiceName = varCloudBaseSkuWithoutEmbeddedStruct.ServiceName
		varCloudBaseSku.Target = varCloudBaseSkuWithoutEmbeddedStruct.Target
		*o = CloudBaseSku(varCloudBaseSku)
	} else {
		return err
	}

	varCloudBaseSku := _CloudBaseSku{}

	err = json.Unmarshal(bytes, &varCloudBaseSku)
	if err == nil {
		o.MoBaseMo = varCloudBaseSku.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CustomAttributes")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "IsActive")
		delete(additionalProperties, "IsAutoDiscovered")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PlatformType")
		delete(additionalProperties, "ServiceCategory")
		delete(additionalProperties, "ServiceFamily")
		delete(additionalProperties, "ServiceName")
		delete(additionalProperties, "Target")

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

type NullableCloudBaseSku struct {
	value *CloudBaseSku
	isSet bool
}

func (v NullableCloudBaseSku) Get() *CloudBaseSku {
	return v.value
}

func (v *NullableCloudBaseSku) Set(val *CloudBaseSku) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudBaseSku) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudBaseSku) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudBaseSku(val *CloudBaseSku) *NullableCloudBaseSku {
	return &NullableCloudBaseSku{value: val, isSet: true}
}

func (v NullableCloudBaseSku) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudBaseSku) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
