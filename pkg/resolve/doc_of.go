//ff:func feature=resolve type=decoder control=sequence
//ff:what docOf — the document a Reference stands for: DOC-POINTER(Issuer, document, version 0)
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// docOf returns the document, as a DocKey in canonical hex, of a DOC-POINTER
// (or a codec's CODEC-POINTER) Reference with version ID 0: a name or title binds the document, not one of
// its versions.
func docOf(subj core.Record) (DocKey, bool) {
	issuer, doc, ver, err := core.DecodeTargetPointer(subj)
	if err != nil || ver[0]|ver[1] != 0 {
		return DocKey{}, false
	}
	return DocKey{Issuer: core.FormatID(issuer), Doc: core.FormatID(doc)}, true
}
