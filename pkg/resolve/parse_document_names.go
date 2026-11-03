//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what ParseDocumentNames — the name → document map of an Issuer's document list (BIND records)
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// ParseDocumentNames reads an Issuer's document list (`<Issuer>/docs/`, the
// old documents.geul; GEUL-ORG 문서 이름): every BIND whose Subject is
// DOC-POINTER(Issuer, document, 0) or CODEC-POINTER(Issuer, codec, 0) maps
// its name to that document, and a Subject that is a codec TAG WORD maps it to
// that codec's codec document (codecKey). Other records are ignored. A
// later BIND of the same name wins.
func ParseDocumentNames(b []byte) (map[string]DocKey, error) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return nil, err
	}
	doc, err := core.SplitDocument(ws)
	if err != nil {
		return nil, err
	}
	names := map[string]DocKey{}
	for _, r := range doc.Body {
		bind, err := core.DecodeBind(r)
		if err != nil {
			continue
		}
		if d, ok := docOf(bind.Subject); ok {
			names[string(bind.Name)] = d
		}
		if t, ok := core.CodecSubject(bind.Subject); ok {
			names[string(bind.Name)] = codecKey(t)
		}
	}
	return names, nil
}
