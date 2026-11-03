//ff:func feature=codecspec type=helper control=sequence
//ff:what Spec.Of — the declared fields of a record's codec; false for an undeclared or ISSUER codec
package codecspec

import "github.com/geul-org/gogeul/pkg/core"

// Of returns the fields declared for r's codec. ISSUER records are not looked
// up yet: an Issuer codec's document is `<Issuer>/codecs/<코덱 번호>/`
// (grammar publishing.md), which no resolver reads yet.
func (s Spec) Of(r core.Record) ([]Field, bool) {
	if len(r.Issuer) > 0 {
		return nil, false
	}
	fs, ok := s[keyOf(r)]
	return fs, ok
}
