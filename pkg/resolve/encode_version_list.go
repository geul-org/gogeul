//ff:func feature=resolve type=encoder control=iteration dimension=1
//ff:what encodeVersionList — a version list document: SPEC-VERSION, PARENT, RECORDED-AT, NEXT-LOCAL-ID, then the records
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// encodeVersionList writes a version list document (edge/PUBLICATION): HEAD
// SPEC-VERSION, PARENT (the previous list's content hash, when there is one)
// RECORDED-AT (seconds) and NEXT-LOCAL-ID (when next > 0,
// meta/lineage/NEXT-LOCAL-ID), then the BODY records in order.
func encodeVersionList(records []core.Record, parent *[32]byte, at int64, next uint64) []byte {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	in := enc(core.EncodeSpecVersion(core.SpecVersion))
	if parent != nil {
		h, _ := core.ReadWords(parent[:])
		in = append(in, enc(core.Record{Kind: core.KindMeta, Codec: core.MetaParent, Payload: enc(core.Record{Kind: core.KindValue, Codec: core.CodecSHA256Digest, Payload: h})})...)
	}
	in = append(in, enc(core.Record{Kind: core.KindMeta, Codec: core.MetaRecordedAt, Payload: enc(core.EncodePosixSeconds(at))})...)
	if next > 0 {
		in = append(in, enc(core.EncodeNextLocalID(next))...)
	}
	in = append(in, core.TargetWord(core.KindMeta, core.MetaBody))
	for _, r := range records {
		in = append(in, enc(r)...)
	}
	return core.PutWords(core.WrapDocument(append(in, core.TargetWord(core.KindMeta, core.MetaEnd))))
}
