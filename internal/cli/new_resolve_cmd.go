//ff:func feature=cli type=command control=sequence
//ff:what newResolveCmd — geul resolve: follow a document's IMPORTs and print the versions they bind to
package cli

import "github.com/spf13/cobra"

// newResolveCmd is `geul resolve <문서>[@<버전>] | <파일>`.
func newResolveCmd(a *app) *cobra.Command {
	return &cobra.Command{
		Use:   "resolve DOCUMENT[@VERSION] | FILE",
		Short: "Resolve a document's IMPORT-DOC dependencies: the data folder first, then GEULPROXY",
		Long: `Resolve a document's IMPORT-DOC dependencies and print each bound version and its source.

Sources are tried in order: the data folder, GEULPACK, then GEULPROXY. A file://
or https:// source may hold .geul files, raw pack sets or zstd pack sets, in any
mix per Issuer; latest.txt tells them apart as they are read. A source that lacks
a document passes to the next. One that is broken stops the search, unless | follows it.
Pack sets over https:// are read by HTTP Range, a few requests per document.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.resolve(args[0])
		},
	}
}
