//ff:func feature=cli type=command control=sequence
//ff:what newRootCmd — the geul command with its persistent flags and every subcommand
package cli

import "github.com/spf13/cobra"

// newRootCmd builds the command tree for a. The persistent flags --data,
// --format and --config feed GEULDATA, GEULFORMAT and GEULENV (load). After a
// command succeeds, pack writes are committed (commit).
func newRootCmd(a *app) *cobra.Command {
	root := &cobra.Command{
		Use:                "geul",
		Short:              "Read, write and publish GEUL documents (.geul)",
		SilenceUsage:       true,
		SilenceErrors:      true,
		PersistentPreRunE:  func(cmd *cobra.Command, _ []string) error { return a.load(cmd) },
		PersistentPostRunE: func(*cobra.Command, []string) error { return a.commit() },
	}
	root.PersistentFlags().String("data", "", "data folder holding .well-known/geul/ (GEULDATA)")
	root.PersistentFlags().String("format", "", "output format: text or json (GEULFORMAT)")
	root.PersistentFlags().String("config", "", "settings file (GEULENV)")
	root.SetOut(a.out)
	root.SetIn(a.in)
	root.AddCommand(newEnvCmd(a), newLsCmd(a), newShowCmd(a), newGetCmd(a), newCatCmd(a), newHashCmd(a), newDiffCmd(a), newCheckCmd(a))
	root.AddCommand(newBlameCmd(a), newGrepCmd(a), newResolveCmd(a), newServeCmd(a), newBuildCmd(a))
	root.AddCommand(newPublishCmd(a), newDeltaCmd(a), newDeprecateCmd(a), newTombstoneCmd(a), newNameCmd(a))
	root.AddCommand(newNewCmd(a), newOpenCmd(a), newAddCmd(a), newDefCmd(a), newSetCmd(a), newRmCmd(a), newMetaCmd(a), newIssuerCmd(a))
	root.AddCommand(newPackCmd(a))
	return root
}
