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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/api.git/build/apis/gateway.nexus.org/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GatewayLister helps list Gateways.
// All objects returned here must be treated as read-only.
type GatewayLister interface {
	// List lists all Gateways in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Gateway, err error)
	// Get retrieves the Gateway from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Gateway, error)
	GatewayListerExpansion
}

// gatewayLister implements the GatewayLister interface.
type gatewayLister struct {
	indexer cache.Indexer
}

// NewGatewayLister returns a new GatewayLister.
func NewGatewayLister(indexer cache.Indexer) GatewayLister {
	return &gatewayLister{indexer: indexer}
}

// List lists all Gateways in the indexer.
func (s *gatewayLister) List(selector labels.Selector) (ret []*v1.Gateway, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Gateway))
	})
	return ret, err
}

// Get retrieves the Gateway from the index for a given name.
func (s *gatewayLister) Get(name string) (*v1.Gateway, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("gateway"), name)
	}
	return obj.(*v1.Gateway), nil
}
