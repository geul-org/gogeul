//ff:func feature=cli type=helper control=selection
//ff:what showArgs — geul show's arguments: one document or file, or the word versions and a document
package cli

import (
	"errors"

	"github.com/spf13/cobra"
)

// showArgs accepts a 공표 주소, a file, or `versions <문서>`. `versions` alone is
// an ordinary first argument, so a file of that name is still `./versions`.
func showArgs(_ *cobra.Command, args []string) error {
	switch {
	case len(args) == 1:
		return nil
	case len(args) == 2 && args[0] == "versions":
		return nil
	default:
		return errors.New("geul show takes one ADDRESS or FILE, or versions DOCUMENT")
	}
}
