//ff:func feature=cli type=helper control=sequence
//ff:what app.docSpec — the codec specs for the codecs one document uses, found under 0x00000000/codecs through the data folder and the source list; empty when they do not read
package cli

import (
	"context"
	"slices"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
)

// docSpec gives the specs b needs (codecspec.Targets) by their addresses
// (resolve.Resolver.CodecSpecs), so showing one document never lists an Issuer
// and a data folder without the GEUL Issuer still reads. Like spec, it falls
// back to none: a record of an undeclared codec is shown generically, and
// bytes that do not split are left for the decoder to report.
func (a *app) docSpec(b []byte) codecspec.Spec {
	ws, err := core.ReadWords(b)
	if err != nil {
		return codecspec.Spec{}
	}
	doc, err := core.SplitDocument(ws)
	if err != nil {
		return codecspec.Spec{}
	}
	r, err := a.resolver()
	if err != nil {
		return codecspec.Spec{}
	}
	ctx, cancel := context.WithTimeout(context.Background(), specTimeout)
	defer cancel()
	s, err := r.CodecSpecs(ctx, codecspec.Targets(slices.Concat(doc.Head, doc.Body)))
	if err != nil {
		return codecspec.Spec{}
	}
	return s
}
