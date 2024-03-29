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

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/alert/v1alpha1"
	auditingv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/auditing/v1alpha1"
	cloudv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/cloud/v1alpha1"
	clusterv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/cluster/v1alpha1"
	customv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/custom/v1alpha1"
	datav1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/data/v1alpha1"
	databasev1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/database/v1alpha1"
	encryptionv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/encryption/v1alpha1"
	eventv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/event/v1alpha1"
	globalv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/global/v1alpha1"
	ldapv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/ldap/v1alpha1"
	maintenancev1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/maintenance/v1alpha1"
	networkv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/network/v1alpha1"
	onlinev1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/online/v1alpha1"
	privatev1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/private/v1alpha1"
	privatelinkv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/privatelink/v1alpha1"
	projectv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/project/v1alpha1"
	searchv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/search/v1alpha1"
	teamv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/team/v1alpha1"
	teamsv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/teams/v1alpha1"
	thirdv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/third/v1alpha1"
	x509v1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/x509/v1alpha1"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=alert.mongodbatlas.kubeform.com, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("configurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Alert().V1alpha1().Configurations().Informer()}, nil

		// Group=auditing.mongodbatlas.kubeform.com, Version=v1alpha1
	case auditingv1alpha1.SchemeGroupVersion.WithResource("auditings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Auditing().V1alpha1().Auditings().Informer()}, nil

		// Group=cloud.mongodbatlas.kubeform.com, Version=v1alpha1
	case cloudv1alpha1.SchemeGroupVersion.WithResource("provideraccesses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cloud().V1alpha1().ProviderAccesses().Informer()}, nil
	case cloudv1alpha1.SchemeGroupVersion.WithResource("provideraccessauthorizations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cloud().V1alpha1().ProviderAccessAuthorizations().Informer()}, nil
	case cloudv1alpha1.SchemeGroupVersion.WithResource("provideraccesssetups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cloud().V1alpha1().ProviderAccessSetups().Informer()}, nil
	case cloudv1alpha1.SchemeGroupVersion.WithResource("providersnapshots"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cloud().V1alpha1().ProviderSnapshots().Informer()}, nil
	case cloudv1alpha1.SchemeGroupVersion.WithResource("providersnapshotbackuppolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cloud().V1alpha1().ProviderSnapshotBackupPolicies().Informer()}, nil
	case cloudv1alpha1.SchemeGroupVersion.WithResource("providersnapshotrestorejobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cloud().V1alpha1().ProviderSnapshotRestoreJobs().Informer()}, nil

		// Group=cluster.mongodbatlas.kubeform.com, Version=v1alpha1
	case clusterv1alpha1.SchemeGroupVersion.WithResource("clusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cluster().V1alpha1().Clusters().Informer()}, nil

		// Group=custom.mongodbatlas.kubeform.com, Version=v1alpha1
	case customv1alpha1.SchemeGroupVersion.WithResource("dbroles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Custom().V1alpha1().DbRoles().Informer()}, nil
	case customv1alpha1.SchemeGroupVersion.WithResource("dnsconfigurationclusterawses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Custom().V1alpha1().DnsConfigurationClusterAwses().Informer()}, nil

		// Group=data.mongodbatlas.kubeform.com, Version=v1alpha1
	case datav1alpha1.SchemeGroupVersion.WithResource("lakes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Data().V1alpha1().Lakes().Informer()}, nil

		// Group=database.mongodbatlas.kubeform.com, Version=v1alpha1
	case databasev1alpha1.SchemeGroupVersion.WithResource("users"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Database().V1alpha1().Users().Informer()}, nil

		// Group=encryption.mongodbatlas.kubeform.com, Version=v1alpha1
	case encryptionv1alpha1.SchemeGroupVersion.WithResource("atrests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Encryption().V1alpha1().AtRests().Informer()}, nil

		// Group=event.mongodbatlas.kubeform.com, Version=v1alpha1
	case eventv1alpha1.SchemeGroupVersion.WithResource("triggers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Event().V1alpha1().Triggers().Informer()}, nil

		// Group=global.mongodbatlas.kubeform.com, Version=v1alpha1
	case globalv1alpha1.SchemeGroupVersion.WithResource("clusterconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Global().V1alpha1().ClusterConfigs().Informer()}, nil

		// Group=ldap.mongodbatlas.kubeform.com, Version=v1alpha1
	case ldapv1alpha1.SchemeGroupVersion.WithResource("configurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ldap().V1alpha1().Configurations().Informer()}, nil
	case ldapv1alpha1.SchemeGroupVersion.WithResource("verifies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ldap().V1alpha1().Verifies().Informer()}, nil

		// Group=maintenance.mongodbatlas.kubeform.com, Version=v1alpha1
	case maintenancev1alpha1.SchemeGroupVersion.WithResource("windows"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Maintenance().V1alpha1().Windows().Informer()}, nil

		// Group=network.mongodbatlas.kubeform.com, Version=v1alpha1
	case networkv1alpha1.SchemeGroupVersion.WithResource("containers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Network().V1alpha1().Containers().Informer()}, nil
	case networkv1alpha1.SchemeGroupVersion.WithResource("peerings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Network().V1alpha1().Peerings().Informer()}, nil

		// Group=online.mongodbatlas.kubeform.com, Version=v1alpha1
	case onlinev1alpha1.SchemeGroupVersion.WithResource("archives"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Online().V1alpha1().Archives().Informer()}, nil

		// Group=private.mongodbatlas.kubeform.com, Version=v1alpha1
	case privatev1alpha1.SchemeGroupVersion.WithResource("ipmodes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Private().V1alpha1().IpModes().Informer()}, nil

		// Group=privatelink.mongodbatlas.kubeform.com, Version=v1alpha1
	case privatelinkv1alpha1.SchemeGroupVersion.WithResource("endpoints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Privatelink().V1alpha1().Endpoints().Informer()}, nil
	case privatelinkv1alpha1.SchemeGroupVersion.WithResource("endpointservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Privatelink().V1alpha1().EndpointServices().Informer()}, nil

		// Group=project.mongodbatlas.kubeform.com, Version=v1alpha1
	case projectv1alpha1.SchemeGroupVersion.WithResource("ipaccesslists"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Project().V1alpha1().IpAccessLists().Informer()}, nil
	case projectv1alpha1.SchemeGroupVersion.WithResource("projects"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Project().V1alpha1().Projects().Informer()}, nil

		// Group=search.mongodbatlas.kubeform.com, Version=v1alpha1
	case searchv1alpha1.SchemeGroupVersion.WithResource("indexes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Search().V1alpha1().Indexes().Informer()}, nil

		// Group=team.mongodbatlas.kubeform.com, Version=v1alpha1
	case teamv1alpha1.SchemeGroupVersion.WithResource("teams"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Team().V1alpha1().Teams().Informer()}, nil

		// Group=teams.mongodbatlas.kubeform.com, Version=v1alpha1
	case teamsv1alpha1.SchemeGroupVersion.WithResource("teamses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Teams().V1alpha1().Teamses().Informer()}, nil

		// Group=third.mongodbatlas.kubeform.com, Version=v1alpha1
	case thirdv1alpha1.SchemeGroupVersion.WithResource("partyintegrations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Third().V1alpha1().PartyIntegrations().Informer()}, nil

		// Group=x509.mongodbatlas.kubeform.com, Version=v1alpha1
	case x509v1alpha1.SchemeGroupVersion.WithResource("authenticationdatabaseusers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.X509().V1alpha1().AuthenticationDatabaseUsers().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
