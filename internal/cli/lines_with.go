//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what linesWith — a document's records as notation lines, laid out by the given codec specs
package cli

import (
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/inspect"
	"github.com/geul-org/gogeul/pkg/notation"
)

// linesWith decodes b (inspect.DecodeWith) and writes each HEAD and BODY
// record as one line (notation.FormatEntry).
func linesWith(b []byte, s codecspec.Spec) (docLines, error) {
	doc, err := inspect.DecodeWith(b, s)
	if err != nil {
		return docLines{}, err
	}
	out := docLines{Head: []string{}, Body: []string{}}
	for _, e := range doc.Head {
		out.Head = append(out.Head, notation.FormatEntry(e))
	}
	for _, e := range doc.Body {
		out.Body = append(out.Body, notation.FormatEntry(e))
	}
	return out, nil
}
