package api

import (
	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
	"golang-appnet.eng.vmware.com/nexus-sdk/api/config"
	"golang-appnet.eng.vmware.com/nexus-sdk/api/runtime"
)

// Nexus is the root node for Nexus infra/runtime datamodel.
//
// This hosts the graph that will consist of user configuration,
// runtime state, inventory and other state essential to the
// functioning of Nexus SDK and runtime.
type Nexus struct {
	nexus.Node

	// Configuration.
	Config  config.Config   `nexus:"child"`
	Runtime runtime.Runtime `nexus:"child" json:"runtime,omitempty"`
}
