//ff:func feature=cli type=command control=sequence
//ff:what newNameCmd — geul name: bind a name to a published document
package cli

import (
	"errors"

	"github.com/spf13/cobra"
)

// newNameCmd is `geul name <문서> <이름>`. The document is given without a
// version: a name binds every version (edge/PUBLICATION 문서 이름).
func newNameCmd(a *app) *cobra.Command {
	return &cobra.Command{
		Use:   "name DOCUMENT NAME",
		Short: "Bind a name (a-z, 0-9, -) to a published document in its Issuer's document list (<Issuer>/docs/); old names keep working",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			d, err := a.dir()
			if err != nil {
				return err
			}
			k, err := a.key(args[0])
			if err != nil {
				return err
			}
			if k.Version != 0 {
				return errors.New("name: a name binds a document, not a version; drop @VERSION")
			}
			return d.Name(k.Issuer, k.Doc, args[1])
		},
	}
}
