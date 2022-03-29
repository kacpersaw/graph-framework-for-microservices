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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "gitlab.eng.vmware.com/nexus/compiler/_generated/apis/service_group.tsm.tanzu.vmware.com/v1"
	"gitlab.eng.vmware.com/nexus/compiler/_generated/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type Service_groupTsmV1Interface interface {
	RESTClient() rest.Interface
	SvcGroupsGetter
}

// Service_groupTsmV1Client is used to interact with features provided by the service_group.tsm.tanzu.vmware.com group.
type Service_groupTsmV1Client struct {
	restClient rest.Interface
}

func (c *Service_groupTsmV1Client) SvcGroups(namespace string) SvcGroupInterface {
	return newSvcGroups(c, namespace)
}

// NewForConfig creates a new Service_groupTsmV1Client for the given config.
func NewForConfig(c *rest.Config) (*Service_groupTsmV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &Service_groupTsmV1Client{client}, nil
}

// NewForConfigOrDie creates a new Service_groupTsmV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Service_groupTsmV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new Service_groupTsmV1Client for the given RESTClient.
func New(c rest.Interface) *Service_groupTsmV1Client {
	return &Service_groupTsmV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *Service_groupTsmV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
