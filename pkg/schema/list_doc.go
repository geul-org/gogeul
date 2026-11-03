//ff:func feature=schema type=decoder control=sequence
//ff:what listDoc — the document a version list lists: the Issuer and document ID of its first PUBLICATION
package schema

import "github.com/geul-org/gogeul/pkg/core"

// listDoc returns, in canonical hex, the Issuer and the document ID of the
// version list's first record (a PUBLICATION, as isVersionList requires);
// "" and "" when it does not decode.
func listDoc(body []rec) (string, string) {
	p, err := core.DecodePublication(body[0].R)
	if err != nil {
		return "", ""
	}
	return core.FormatID(p.Issuer), core.FormatID(p.Doc)
}
