//ff:func feature=core type=decoder control=sequence
//ff:what CodecSubject — the codec a BIND Subject stands for when it is a lone TAG WORD (LLLL = 0, no payload)
package core

// CodecSubject reports whether a BIND Subject is a lone TAG WORD — a direct
// codec's TAG with LLLL = 0 and no payload — and returns that codec's target
// word (TargetWord). Such a Subject stands for the codec itself, that is its
// codec document `0x00000000/codecs/<codec ID>/` whose codec ID is that target
// word as 32 bits (GEUL edge/BIND).
func CodecSubject(r Record) (Word, bool) {
	if len(r.Payload) > 0 || len(r.Issuer) > 0 {
		return 0, false
	}
	return TargetWord(r.Kind, r.Codec), true
}
