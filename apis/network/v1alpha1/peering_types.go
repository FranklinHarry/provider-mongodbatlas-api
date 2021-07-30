/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Peering struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PeeringSpec   `json:"spec,omitempty"`
	Status            PeeringStatus `json:"status,omitempty"`
}

type PeeringSpec struct {
	State *PeeringSpecResource `json:"state,omitempty" tf:"-"`

	Resource PeeringSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type PeeringSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccepterRegionName *string `json:"accepterRegionName,omitempty" tf:"accepter_region_name"`
	// +optional
	AtlasCIDRBlock *string `json:"atlasCIDRBlock,omitempty" tf:"atlas_cidr_block"`
	// +optional
	AtlasGcpProjectID *string `json:"atlasGcpProjectID,omitempty" tf:"atlas_gcp_project_id"`
	// +optional
	AtlasID *string `json:"atlasID,omitempty" tf:"atlas_id"`
	// +optional
	AtlasVpcName *string `json:"atlasVpcName,omitempty" tf:"atlas_vpc_name"`
	// +optional
	AwsAccountID *string `json:"awsAccountID,omitempty" tf:"aws_account_id"`
	// +optional
	AzureDirectoryID *string `json:"azureDirectoryID,omitempty" tf:"azure_directory_id"`
	// +optional
	AzureSubscriptionID *string `json:"azureSubscriptionID,omitempty" tf:"azure_subscription_id"`
	// +optional
	ConnectionID *string `json:"connectionID,omitempty" tf:"connection_id"`
	ContainerID  *string `json:"containerID" tf:"container_id"`
	// +optional
	ErrorMessage *string `json:"errorMessage,omitempty" tf:"error_message"`
	// +optional
	ErrorState *string `json:"errorState,omitempty" tf:"error_state"`
	// +optional
	ErrorStateName *string `json:"errorStateName,omitempty" tf:"error_state_name"`
	// +optional
	GcpProjectID *string `json:"gcpProjectID,omitempty" tf:"gcp_project_id"`
	// +optional
	NetworkName *string `json:"networkName,omitempty" tf:"network_name"`
	// +optional
	PeerID       *string `json:"peerID,omitempty" tf:"peer_id"`
	ProjectID    *string `json:"projectID" tf:"project_id"`
	ProviderName *string `json:"providerName" tf:"provider_name"`
	// +optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name"`
	// +optional
	RouteTableCIDRBlock *string `json:"routeTableCIDRBlock,omitempty" tf:"route_table_cidr_block"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	StatusName *string `json:"statusName,omitempty" tf:"status_name"`
	// +optional
	VnetName *string `json:"vnetName,omitempty" tf:"vnet_name"`
	// +optional
	VpcID *string `json:"vpcID,omitempty" tf:"vpc_id"`
}

type PeeringStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PeeringList is a list of Peerings
type PeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Peering CRD objects
	Items []Peering `json:"items,omitempty"`
}