// Code generated by nexus. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"golang-appnet.eng.vmware.com/nexus-sdk/api/build/common"
)

// +k8s:openapi-gen=true
type Child struct {
	Group string `json:"group" yaml:"group"`
	Kind  string `json:"kind" yaml:"kind"`
	Name  string `json:"name" yaml:"name"`
}

// +k8s:openapi-gen=true
type Link struct {
	Group string `json:"group" yaml:"group"`
	Kind  string `json:"kind" yaml:"kind"`
	Name  string `json:"name" yaml:"name"`
}

/* ------------------- CRDs definitions ------------------- */

// +genclient
// +genclient:noStatus
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type OIDC struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata" yaml:"metadata"`
	Spec              OIDCSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

func (c *OIDC) CRDName() string {
	return "oidcs.authentication.nexus.org"
}

func (c *OIDC) DisplayName() string {
	if c.GetLabels() != nil {
		return c.GetLabels()[common.DISPLAY_NAME_LABEL]
	}
	return ""
}

// +k8s:openapi-gen=true
type OIDCSpec struct {
	Config          IDPConfig            `json:"config" yaml:"config"`
	ValidationProps ValidationProperties `json:"validationProps" yaml:"validationProps"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type OIDCList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []OIDC `json:"items" yaml:"items"`
}

// +k8s:openapi-gen=true
type IDPConfig struct {
	ClientId         string   `json:"clientId"`
	ClientSecret     string   `json:"clientSecret"`
	OAuthIssuerUrl   string   `json:"oAuthIssuerUrl"`
	Scopes           []string `json:"scopes"`
	OAuthRedirectUrl string   `json:"oAuthRedirectUrl"`
}

// +k8s:openapi-gen=true
type ValidationProperties struct {
	InsecureIssuerURLContext bool `json:"insecureIssuerURLContext"`
	SkipIssuerValidation     bool `json:"skipIssuerValidation"`
	SkipClientIdValidation   bool `json:"skipClientIdValidation"`
	SkipClientAudValidation  bool `json:"skipClientAudValidation"`
}
