//ff:func feature=cli type=helper control=sequence
//ff:what app.docSpec — the data folder's codec specs for the codecs one document uses, read by address; empty without a data folder or when they do not read
package cli

import (
	"slices"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
)

// docSpec gives the specs b needs (codecspec.Targets) from the data folder
// (resolve.Dir.CodecSpecs), so showing one document never lists an Issuer.
// Like spec, it falls back to none: a record of an undeclared codec is shown
// generically, and bytes that do not split are left for the decoder to report.
func (a *app) docSpec(b []byte) codecspec.Spec {
	d, err := a.dir()
	if err != nil {
		return codecspec.Spec{}
	}
	ws, err := core.ReadWords(b)
	if err != nil {
		return codecspec.Spec{}
	}
	doc, err := core.SplitDocument(ws)
	if err != nil {
		return codecspec.Spec{}
	}
	s, err := d.CodecSpecs(codecspec.Targets(slices.Concat(doc.Head, doc.Body)))
	if err != nil {
		return codecspec.Spec{}
	}
	return s
}
