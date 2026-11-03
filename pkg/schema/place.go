//ff:func feature=schema type=rule control=selection
//ff:what place — D2~D4: route one record into HEAD or BODY by the placement rules
package schema

import "github.com/geul-org/gogeul/pkg/core"

// place applies one record after the DOC record to st and parts (D2~D4).
// Before BODY, Meta and the HEAD Edges (a SOURCE on the whole document and
// the QUALIFIERs on it, headEdgeRec) go to the HEAD. A BODY after an empty
// HEAD is not the canonical form: an empty HEAD omits BODY (document.md, 0x1C3L
// BODY). Misplaced records are reported and still kept so their own rules are
// checked.
func place(st *placeState, parts *docParts, r rec) []Fact {
	meta := r.R.Kind == core.KindMeta
	switch {
	case st.ended:
		return []Fact{factAt("D2", Error, r, "", "END가 마지막 레코드", "END 뒤 레코드")}
	case meta && r.R.Codec == core.MetaEnd:
		st.ended = true
		return checkOneWord(r, "END")
	case meta && r.R.Codec == core.MetaBody && st.inBody:
		return append(checkOneWord(r, "BODY"), factAt("D3", Error, r, "", "BODY 한 번", "BODY 중복 또는 BODY 레코드 뒤"))
	case meta && r.R.Codec == core.MetaBody && len(parts.Head) == 0:
		st.inBody = true
		return append(checkOneWord(r, "BODY"), factAt("D3", Error, r, "", "HEAD가 비면 BODY 생략(정규형)", "빈 HEAD 뒤 BODY"))
	case meta && r.R.Codec == core.MetaBody:
		st.inBody = true
		return checkOneWord(r, "BODY")
	case meta && r.R.Codec <= core.MetaDoc4Word:
		return []Fact{factAt("D4", Error, r, "", "중첩하지 않는 문서", "DOC-* 레코드")}
	case meta && st.inBody:
		parts.Body = append(parts.Body, r)
		return []Fact{factAt("D3", Error, r, "", "BODY에는 Meta가 아닌 레코드", "Meta")}
	case meta, !st.inBody && headEdgeRec(st, r):
		parts.Head = append(parts.Head, r)
		return nil
	case !st.inBody && len(parts.Head) > 0:
		parts.Body = append(parts.Body, r)
		return []Fact{factAt("D3", Error, r, "", "HEAD에는 Meta와 문서 전체 SOURCE·그 QUALIFIER만 (BODY 누락)", "그 밖의 레코드")}
	default:
		st.inBody = true
		parts.Body = append(parts.Body, r)
		return nil
	}
}
