//ff:func feature=cli type=command control=sequence
//ff:what newGetCmd — geul get: the pure GEUL bytes of a 공표 주소 (file:// or https://) on standard output; other forms and missing addresses fail
package cli

import "github.com/spf13/cobra"

// newGetCmd is `geul get <주소>`: resolve.ParseAddr → resolve.Fetch → standard
// output. The address is one of the ten 공표 주소 forms under
// `file://<dir>/.well-known/geul/` or `https://<host>/.well-known/geul/`; files,
// pack sets and zstd pack sets answer alike. Any other form and a missing
// address end with the not-found error.
func newGetCmd(a *app) *cobra.Command {
	return &cobra.Command{
		Use:   "get ADDRESS",
		Short: "Write the pure GEUL bytes at a 공표 주소 (file:// or https://) to standard output",
		Long: `Write the bytes at one 공표 주소 to standard output, unchanged:

  issuers/latest.geul | issuers/<version>.geul | issuers/versions.geul
  <Issuer>/codecs/<codec ID>/latest.geul | <version>.geul | versions.geul
  <Issuer>/docs/<document ID>/latest.geul | <version>.geul | versions.geul
  <Issuer>/hashes/<content hash>.geul

each under file://<dir>/.well-known/geul/ or https://<host>/.well-known/geul/.
The Issuer may be written with 8, 16 or 64 hex digits (the same Issuer). Any
other form, or an address that is not there, is "not found". For a text view
use geul show.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			b, err := getAddr(cmd.Context(), args[0])
			if err != nil {
				return err
			}
			_, err = a.out.Write(b)
			return err
		},
	}
}
