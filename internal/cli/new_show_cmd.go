//ff:func feature=cli type=command control=sequence
//ff:what newShowCmd — geul show: one document as text, one record per line, or that document's versions
package cli

import "github.com/spf13/cobra"

// newShowCmd is `geul show <공표 주소> | <파일>` (git show) and
// `geul show versions <문서>`.
func newShowCmd(a *app) *cobra.Command {
	return &cobra.Command{
		Use:   "show ADDRESS | FILE | versions DOCUMENT",
		Short: "Print one document as text (HEAD records, BODY, then one Edge per line), or list its versions",
		Long: `Print one document as text: HEAD records, BODY, then one Edge per line (#id NAME fields).

  geul show <Issuer>/<document>.geul              the latest version
  geul show <Issuer>/<document>/<version>.geul    a given version
  geul show <Issuer>/<document>/versions.geul     the version list document
  geul show <Issuer>/documents.geul               the Issuer's document names
  geul show issuers.geul                          the Issuer names
  geul show h/<content hash>.geul                 a document by its hash
  geul show ./draft.geul                          a local file
  geul show versions <Issuer>/<document>          number, time, content hash, state, NEXT-LOCAL-ID

Issuer may be a domain and document a name. .geul may be left off.

A file named versions is ./versions. Records are laid out by the spec documents
of the codecs the document uses, found under 0x0000 by TAG in the data folder,
then through GEULPROXY; a codec with none is shown generically. Nothing is listed: show reads this document and those only.`,
		Args: showArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 2 {
				return a.versions(args[1])
			}
			return a.show(args[0])
		},
	}
}
