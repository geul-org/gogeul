//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what listParent — the PARENT hash a version list HEAD names, if any
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// listParent finds the PARENT Meta in the HEAD of a version list and
// returns the previous list's content hash it holds (a SHA-256-DIGEST Value).
// It is false when there is none or it does not read as one.
func listParent(head []core.Record) ([32]byte, bool) {
	var h [32]byte
	for _, r := range head {
		if r.Kind != core.KindMeta || r.Codec != core.MetaParent {
			continue
		}
		v, n, err := core.DecodeRecord(r.Payload)
		if err != nil || n != len(r.Payload) || v.Kind != core.KindValue || v.Codec != core.CodecSHA256Digest || len(v.Payload) != 16 {
			return h, false
		}
		copy(h[:], core.PutWords(v.Payload))
		return h, true
	}
	return h, false
}
