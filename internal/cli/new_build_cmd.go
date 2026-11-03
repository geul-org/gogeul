//ff:func feature=cli type=command control=sequence
//ff:what newBuildCmd — geul build: bake a .well-known/geul tree (the data folder by default) into static files under an output root
package cli

import "github.com/spf13/cobra"

// newBuildCmd is `geul build -o <out> [DIR]`.
func newBuildCmd(a *app) *cobra.Command {
	var dst string
	cmd := &cobra.Command{
		Use:   "build -o OUT [DIR]",
		Short: "Bake a .well-known/geul tree (the data folder by default) into static files for a web host",
		Long: `build writes a static copy of the tree DIR (the data folder when omitted)
holds, or is, to <OUT>/.well-known/geul/ at the public paths: each <path>.geul
is copied and gets its viewer page at <path>, each directory gets index.html
and index.json, and the graph scripts go to _static/.`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			target, err := a.target(args)
			if err != nil {
				return err
			}
			return buildCmd(target, dst, a.out)
		},
	}
	cmd.Flags().StringVarP(&dst, "out", "o", "", "output root (for example the site's public/)")
	cmd.MarkFlagRequired("out")
	return cmd
}
