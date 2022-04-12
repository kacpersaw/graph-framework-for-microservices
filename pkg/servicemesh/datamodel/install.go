package datamodel

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/cli.git/pkg/utils"
)

var Namespace string

func Install(cmd *cobra.Command, args []string) error {
	envList := []string{}
	fmt.Print("Checking if the tenant-apiserver is reachable for installing datamodel crds\n")

	if err := utils.GoToNexusDirectory(); err != nil {
		return err
	}

	if DatamodelName != "" {
		envList = append(envList, fmt.Sprintf("DATAMODEL=%s", DatamodelName))
		if err := utils.CheckDatamodelDirExists(DatamodelName); err != nil {
			return err
		}
	} else {
		DatamodelName, err := utils.GetCurrentDatamodel()
		if err != nil {
			return err
		}
		fmt.Printf("Installing datamodel %s\n", DatamodelName)
		envList = append(envList, fmt.Sprintf("DATAMODEL=%s", DatamodelName))
	}

	if Namespace != "" {
		envList = append(envList, fmt.Sprintf("NAMESPACE=%s", Namespace))
	}

	err := utils.SystemCommand(envList, !utils.IsDebug(cmd), "make", "datamodel_install")
	if err != nil {
		return err
	}

	fmt.Printf("\u2713 Datamodel %s install successful\n", DatamodelName)
	return nil
}

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install specified datamodel's generated CRDs to the specified namespace",
	//Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		return nil
	},
	RunE: Install,
}

func init() {
	InstallCmd.Flags().StringVarP(&Namespace, "namespace",
		"r", "", "name of the namespace to install to")
	InstallCmd.Flags().StringVarP(&DatamodelName, "name",
		"n", "", "name of the datamodel to install")

}
