//ff:func feature=cli type=command control=sequence
//ff:what newPublishCmd — geul publish: a draft becomes the next version in the data folder
package cli

import "github.com/spf13/cobra"

// newPublishCmd is `geul publish <파일> [--amend]` (git commit).
func newPublishCmd(a *app) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "publish FILE",
		Short: "Publish a document as its next version (versions.geul, NEXT-LOCAL-ID and <Issuer>/hashes/ are updated); --amend replaces the last",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			amend, _ := cmd.Flags().GetBool("amend")
			return a.publish(args[0], amend)
		},
	}
	cmd.Flags().Bool("amend", false, "replace the versions with this one as version 1 (early development only)")
	return cmd
}
