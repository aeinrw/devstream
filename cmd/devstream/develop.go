package main

import (
	"fmt"
	"github.com/devstream-io/devstream/cmd/devstream/options"
	"github.com/spf13/cobra"

	"github.com/devstream-io/devstream/internal/pkg/develop"
	"github.com/devstream-io/devstream/pkg/util/log"
)

var (
	name string
	all  bool
)

var developCMD = &cobra.Command{
	Use:   "develop",
	Short: "Develop is used for develop a new plugin",
	Long: `Develop is used for develop a new plugin.
Examples:
  dtm develop create-plugin --name=YOUR-PLUGIN-NAME,
  dtm develop validate-plugin --name=YOUR-PLUGIN-NAME`,
	Run: options.WithValidators(developCMDFunc, options.ArgsCountEqual(1), validateDevelopArgs),
}

func developCMDFunc(cmd *cobra.Command, args []string) {
	developAction := develop.Action(args[0])
	log.Debugf("The develop action is: %s.", developAction)
	if err := develop.ExecuteAction(developAction); err != nil {
		log.Fatal(err)
	}
}

func validateDevelopArgs(args []string) error {
	// "create-plugin" or "validate-plugin". Maybe it will be "delete-plugin"/"rename-plugin" in future.
	developAction := develop.Action(args[0])
	if !develop.IsValideAction(developAction) {
		return fmt.Errorf("invalide Develop Action")
	}
	return nil
}

func init() {
	developCMD.PersistentFlags().StringVarP(&name, "name", "n", "", "specify name with the new plugin")
	developCMD.PersistentFlags().BoolVarP(&all, "all", "a", false, "validate all plugins")
}
