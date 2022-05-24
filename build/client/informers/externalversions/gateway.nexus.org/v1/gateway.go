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

	gatewaynexusorgv1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/api.git/build/apis/gateway.nexus.org/v1"
	versioned "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/api.git/build/client/clientset/versioned"
	internalinterfaces "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/api.git/build/client/informers/externalversions/internalinterfaces"
	v1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/api.git/build/client/listers/gateway.nexus.org/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GatewayInformer provides access to a shared informer and lister for
// Gateways.
type GatewayInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.GatewayLister
}

type gatewayInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewGatewayInformer constructs a new informer for Gateway type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGatewayInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGatewayInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredGatewayInformer constructs a new informer for Gateway type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGatewayInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GatewayNexusV1().Gateways().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GatewayNexusV1().Gateways().Watch(context.TODO(), options)
			},
		},
		&gatewaynexusorgv1.Gateway{},
		resyncPeriod,
		indexers,
	)
}

func (f *gatewayInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGatewayInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *gatewayInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&gatewaynexusorgv1.Gateway{}, f.defaultInformer)
}

func (f *gatewayInformer) Lister() v1.GatewayLister {
	return v1.NewGatewayLister(f.Informer().GetIndexer())
}
