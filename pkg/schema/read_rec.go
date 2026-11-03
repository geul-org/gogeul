//ff:func feature=schema type=decoder control=sequence
//ff:what readRec — frame one record (F2~F5); false when its length cannot be trusted
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// readRec frames the record at ws[at]: prefix (F4), Length (F2, F3) and the
// ISSUER extension (Issuer and 2-word codec number) and the canonical width
// of its Issuer (F5). The bool is false when the
// Length cannot be trusted; framing stops there.
func readRec(ws []core.Word, at, index int) (rec, []Fact, bool) {
	tag := ws[at]
	r := rec{Index: index, Word: at}
	var fs []Fact
	if p := tag >> 13; p == 1 {
		fs = append(fs, factAt("F4", Warn, r, "TAG WORD", "접두사 000 (Proposal)", "001 (Standard, 미배정)"))
	} else if p != 0 {
		fs = append(fs, factAt("F4", Error, r, "TAG WORD", "접두사 000", fmt.Sprintf("%03b", p)))
	}
	total, head, canonical, ok := readLength(ws, at)
	if !ok || at+total > len(ws) {
		return r, append(fs, factAt("F2", Error, r, "Length", "입력 안에 들어가는 길이", fmt.Sprintf("%d WORD (남은 입력 %d WORD)", total, len(ws)-at))), false
	}
	if !canonical {
		fs = append(fs, factAt("F3", Error, r, "Length", "가장 짧은 Length 표현", fmt.Sprintf("%d WORD를 %d WORD TAG로 표현", total, head)))
	}
	r.Total = total
	k, codec := core.SplitTarget(tag)
	r.R = core.Record{Kind: k, Codec: codec}
	body := ws[at+head : at+total]
	w := core.IssuerWidth(r.R.Kind, r.R.Codec)
	if w > 0 && len(body) < w+core.FormatWords {
		r.R.Payload = body
		return r, append(fs, factAt("F5", Error, r, "ISSUER 확장", fmt.Sprintf("Issuer %d WORD + 코덱 번호 %d WORD", w, core.FormatWords), fmt.Sprintf("%d WORD", len(body)))), true
	}
	if w > 0 {
		r.R.Issuer, r.R.Format, body = body[:w], uint32(body[w])<<16|uint32(body[w+1]), body[w+core.FormatWords:]
		if !core.IssuerCanonical(r.R.Issuer) {
			fs = append(fs, factAt("F5", Error, r, "Issuer", "값이 들어가는 가장 짧은 폭", fmt.Sprintf("ISSUER-%dWORD", w)))
		}
	}
	r.R.Payload = body
	return r, fs, true
}
