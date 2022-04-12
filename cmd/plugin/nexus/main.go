package main

import (
	"os"

	"github.com/spf13/cobra"
	"gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/cli.git/pkg/servicemesh"
)

var enableDebug = false

var rootCmd = &cobra.Command{
	Use:   "nexus",
	Short: "nexus cli",
	Long:  "nexus cli to execute tsm operations",
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(
		servicemesh.ApplyCmd,
		servicemesh.DeleteCmd,
		servicemesh.LoginCmd,
		servicemesh.RuntimeCmd,
		servicemesh.DataModelCmd,
		servicemesh.AppCmd,
		servicemesh.OperatorCmd,
		servicemesh.VersionCmd,
	)
	rootCmd.PersistentFlags().BoolVarP(&enableDebug, "debug", "", false, "Enables extra logging")
}
