/*
Copyright The Kubernetes Authors.

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

package v1

import (
	"context"
	time "time"

	connectnexusvmwarecomv1 "golang-appnet.eng.vmware.com/nexus-sdk/api/build/apis/connect.nexus.vmware.com/v1"
	versioned "golang-appnet.eng.vmware.com/nexus-sdk/api/build/client/clientset/versioned"
	internalinterfaces "golang-appnet.eng.vmware.com/nexus-sdk/api/build/client/informers/externalversions/internalinterfaces"
	v1 "golang-appnet.eng.vmware.com/nexus-sdk/api/build/client/listers/connect.nexus.vmware.com/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ReplicationConfigInformer provides access to a shared informer and lister for
// ReplicationConfigs.
type ReplicationConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ReplicationConfigLister
}

type replicationConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewReplicationConfigInformer constructs a new informer for ReplicationConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewReplicationConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredReplicationConfigInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredReplicationConfigInformer constructs a new informer for ReplicationConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredReplicationConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConnectNexusV1().ReplicationConfigs().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConnectNexusV1().ReplicationConfigs().Watch(context.TODO(), options)
			},
		},
		&connectnexusvmwarecomv1.ReplicationConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *replicationConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredReplicationConfigInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *replicationConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&connectnexusvmwarecomv1.ReplicationConfig{}, f.defaultInformer)
}

func (f *replicationConfigInformer) Lister() v1.ReplicationConfigLister {
	return v1.NewReplicationConfigLister(f.Informer().GetIndexer())
}
