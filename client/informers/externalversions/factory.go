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
	reflect "reflect"
	sync "sync"
	time "time"

	versioned "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned"
	alert "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/alert"
	auditing "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/auditing"
	cloud "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/cloud"
	cluster "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/cluster"
	custom "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/custom"
	data "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/data"
	database "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/database"
	encryption "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/encryption"
	event "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/event"
	global "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/global"
	internalinterfaces "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/internalinterfaces"
	ldap "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/ldap"
	maintenance "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/maintenance"
	network "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/network"
	online "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/online"
	private "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/private"
	privatelink "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/privatelink"
	project "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/project"
	search "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/search"
	team "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/team"
	teams "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/teams"
	third "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/third"
	x509 "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/x509"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// SharedInformerOption defines the functional option type for SharedInformerFactory.
type SharedInformerOption func(*sharedInformerFactory) *sharedInformerFactory

type sharedInformerFactory struct {
	client           versioned.Interface
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	lock             sync.Mutex
	defaultResync    time.Duration
	customResync     map[reflect.Type]time.Duration

	informers map[reflect.Type]cache.SharedIndexInformer
	// startedInformers is used for tracking which informers have been started.
	// This allows Start() to be called multiple times safely.
	startedInformers map[reflect.Type]bool
}

// WithCustomResyncConfig sets a custom resync period for the specified informer types.
func WithCustomResyncConfig(resyncConfig map[v1.Object]time.Duration) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		for k, v := range resyncConfig {
			factory.customResync[reflect.TypeOf(k)] = v
		}
		return factory
	}
}

// WithTweakListOptions sets a custom filter on all listers of the configured SharedInformerFactory.
func WithTweakListOptions(tweakListOptions internalinterfaces.TweakListOptionsFunc) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		factory.tweakListOptions = tweakListOptions
		return factory
	}
}

// WithNamespace limits the SharedInformerFactory to the specified namespace.
func WithNamespace(namespace string) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		factory.namespace = namespace
		return factory
	}
}

// NewSharedInformerFactory constructs a new instance of sharedInformerFactory for all namespaces.
func NewSharedInformerFactory(client versioned.Interface, defaultResync time.Duration) SharedInformerFactory {
	return NewSharedInformerFactoryWithOptions(client, defaultResync)
}

// NewFilteredSharedInformerFactory constructs a new instance of sharedInformerFactory.
// Listers obtained via this SharedInformerFactory will be subject to the same filters
// as specified here.
// Deprecated: Please use NewSharedInformerFactoryWithOptions instead
func NewFilteredSharedInformerFactory(client versioned.Interface, defaultResync time.Duration, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) SharedInformerFactory {
	return NewSharedInformerFactoryWithOptions(client, defaultResync, WithNamespace(namespace), WithTweakListOptions(tweakListOptions))
}

// NewSharedInformerFactoryWithOptions constructs a new instance of a SharedInformerFactory with additional options.
func NewSharedInformerFactoryWithOptions(client versioned.Interface, defaultResync time.Duration, options ...SharedInformerOption) SharedInformerFactory {
	factory := &sharedInformerFactory{
		client:           client,
		namespace:        v1.NamespaceAll,
		defaultResync:    defaultResync,
		informers:        make(map[reflect.Type]cache.SharedIndexInformer),
		startedInformers: make(map[reflect.Type]bool),
		customResync:     make(map[reflect.Type]time.Duration),
	}

	// Apply all options
	for _, opt := range options {
		factory = opt(factory)
	}

	return factory
}

// Start initializes all requested informers.
func (f *sharedInformerFactory) Start(stopCh <-chan struct{}) {
	f.lock.Lock()
	defer f.lock.Unlock()

	for informerType, informer := range f.informers {
		if !f.startedInformers[informerType] {
			go informer.Run(stopCh)
			f.startedInformers[informerType] = true
		}
	}
}

// WaitForCacheSync waits for all started informers' cache were synced.
func (f *sharedInformerFactory) WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool {
	informers := func() map[reflect.Type]cache.SharedIndexInformer {
		f.lock.Lock()
		defer f.lock.Unlock()

		informers := map[reflect.Type]cache.SharedIndexInformer{}
		for informerType, informer := range f.informers {
			if f.startedInformers[informerType] {
				informers[informerType] = informer
			}
		}
		return informers
	}()

	res := map[reflect.Type]bool{}
	for informType, informer := range informers {
		res[informType] = cache.WaitForCacheSync(stopCh, informer.HasSynced)
	}
	return res
}

// InternalInformerFor returns the SharedIndexInformer for obj using an internal
// client.
func (f *sharedInformerFactory) InformerFor(obj runtime.Object, newFunc internalinterfaces.NewInformerFunc) cache.SharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerType := reflect.TypeOf(obj)
	informer, exists := f.informers[informerType]
	if exists {
		return informer
	}

	resyncPeriod, exists := f.customResync[informerType]
	if !exists {
		resyncPeriod = f.defaultResync
	}

	informer = newFunc(f.client, resyncPeriod)
	f.informers[informerType] = informer

	return informer
}

// SharedInformerFactory provides shared informers for resources in all known
// API group versions.
type SharedInformerFactory interface {
	internalinterfaces.SharedInformerFactory
	ForResource(resource schema.GroupVersionResource) (GenericInformer, error)
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool

	Alert() alert.Interface
	Auditing() auditing.Interface
	Cloud() cloud.Interface
	Cluster() cluster.Interface
	Custom() custom.Interface
	Data() data.Interface
	Database() database.Interface
	Encryption() encryption.Interface
	Event() event.Interface
	Global() global.Interface
	Ldap() ldap.Interface
	Maintenance() maintenance.Interface
	Network() network.Interface
	Online() online.Interface
	Private() private.Interface
	Privatelink() privatelink.Interface
	Project() project.Interface
	Search() search.Interface
	Team() team.Interface
	Teams() teams.Interface
	Third() third.Interface
	X509() x509.Interface
}

func (f *sharedInformerFactory) Alert() alert.Interface {
	return alert.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Auditing() auditing.Interface {
	return auditing.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cloud() cloud.Interface {
	return cloud.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cluster() cluster.Interface {
	return cluster.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Custom() custom.Interface {
	return custom.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Data() data.Interface {
	return data.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Database() database.Interface {
	return database.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Encryption() encryption.Interface {
	return encryption.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Event() event.Interface {
	return event.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Global() global.Interface {
	return global.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ldap() ldap.Interface {
	return ldap.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Maintenance() maintenance.Interface {
	return maintenance.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Network() network.Interface {
	return network.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Online() online.Interface {
	return online.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Private() private.Interface {
	return private.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Privatelink() privatelink.Interface {
	return privatelink.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Project() project.Interface {
	return project.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Search() search.Interface {
	return search.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Team() team.Interface {
	return team.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Teams() teams.Interface {
	return teams.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Third() third.Interface {
	return third.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) X509() x509.Interface {
	return x509.New(f, f.namespace, f.tweakListOptions)
}
