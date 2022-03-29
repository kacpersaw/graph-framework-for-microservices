// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ServerType is the type of server.
type ServerType string

const (
	// ManagementClusterServerType is a management cluster server.
	ManagementClusterServerType ServerType = "managementcluster"

	// GlobalServerType is a global control plane server.
	GlobalServerType ServerType = "global"
)

// Server connection.
type Server struct {
	// Name of the server.
	Name string `json:"name,omitempty" yaml:"name"`

	// Type of the endpoint.
	Type ServerType `json:"type,omitempty" yaml:"type"`

	// GlobalOpts if the server is global.
	GlobalOpts *GlobalServer `json:"globalOpts,omitempty" yaml:"globalOpts"`

	// ManagementClusterOpts if the server is a management cluster.
	ManagementClusterOpts *ManagementClusterServer `json:"managementClusterOpts,omitempty" yaml:"managementClusterOpts"`
}

// ManagementClusterServer is the configuration for a management cluster kubeconfig.
type ManagementClusterServer struct {
	// Endpoint for the login.
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint"`

	// Path to the kubeconfig.
	Path string `json:"path,omitempty" yaml:"path"`

	// The context to use (if required), defaults to current.
	Context string `json:"context,omitempty" yaml:"context"`
}

// GlobalServer is the configuration for a global server.
type GlobalServer struct {
	// Endpoint for the server.
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint"`

	// Auth for the global server.
	Auth GlobalServerAuth `json:"auth,omitempty" yaml:"auth"`
}

// GlobalServerAuth is authentication for a global server.
type GlobalServerAuth struct {
	// Issuer url for IDP, compliant with OIDC Metadata Discovery.
	Issuer string `json:"issuer" yaml:"issuer"`

	// UserName is the authorized user the token is assigned to.
	UserName string `json:"userName" yaml:"userName"`

	// Permissions are roles assigned to the user.
	Permissions []string `json:"permissions" yaml:"permissions"`

	// AccessToken is the current access token based on the context.
	AccessToken string `json:"accessToken" yaml:"accessToken"`

	// IDToken is the current id token based on the context scoped to the CLI.
	IDToken string `json:"IDToken" yaml:"IDToken"`

	// RefreshToken will be stored only in case of api-token login flow.
	RefreshToken string `json:"refresh_token" yaml:"refresh_token"`

	// Expiration times of the token.
	Expiration metav1.Time `json:"expiration" yaml:"expiration"`

	// Type of the token (user or client).
	Type string `json:"type" yaml:"type"`
}

// ClientOptions are the client specific options.
type ClientOptions struct {
	// CLI options specific to the CLI.
	CLI *CLIOptions `json:"cli,omitempty" yaml:"cli"`
}

// CLIOptions are options for the CLI.
type CLIOptions struct {
	// Repositories are the plugin repositories.
	Repositories []PluginRepository `json:"repositories,omitempty" yaml:"repositories"`
}

// PluginRepository is a CLI plugin repository
type PluginRepository struct {
	// GCPPluginRepository is a plugin repository that utilizes GCP cloud storage.
	GCPPluginRepository *GCPPluginRepository `json:"gcpPluginRepository,omitempty" yaml:"gcpPluginRepository"`
}

// GCPPluginRepository is a plugin repository that utilizes GCP cloud storage.
type GCPPluginRepository struct {
	// Name of the repository.
	Name string `json:"name,omitempty" yaml:"name"`

	// BucketName is the name of the bucket.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName"`

	// RootPath within the bucket.
	RootPath string `json:"rootPath,omitempty" yaml:"rootPath"`
}

// +kubebuilder:object:root=true

// Config is the Schema for the configs API
type Config struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// KnownServers available.
	KnownServers []*Server `json:"servers,omitempty" yaml:"servers"`

	// CurrentServer in use.
	CurrentServer string `json:"current,omitempty" yaml:"current"`

	// ClientOptions are client specific options.
	ClientOptions *ClientOptions `json:"clientOptions,omitempty" yaml:"clientOptions"`
}

// +kubebuilder:object:root=true

// ConfigList contains a list of Config
type ConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Config `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Config{}, &ConfigList{})
}
