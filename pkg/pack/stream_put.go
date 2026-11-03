//ff:func feature=pack type=writer control=sequence
//ff:what Stream.Put — store the record of path (issuer, document ID, slot) after the previous key: P leaf entry, inventory entry and, for a document, its h/ key
package pack

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// Put stores r as the P record (issuer, doc, slot) — slot 0 is versions.geul,
// v ≥ 1 version v (§3.4) — and, when r is a document, under
// `<Issuer>/hashes/<r.Hash>.geul` too (flags bit0: the bytes are shared, never copied).
// Keys must come in strictly ascending order (Issuer number, document ID,
// slot) and the Issuer must be one of StreamOptions.Issuers. Once Close has
// begun (a resumed stream whose Close stopped) closedPut answers instead. The
// reserved Issuer list ID (core.DocIDIssuers) is refused: the list is named.
func (s *Stream) Put(issuer []core.Word, doc uint64, slot uint32, r StreamRecord) error {
	if n, ok := core.IssuerNormal(issuer); ok {
		issuer = n
	}
	if doc == core.DocIDIssuers {
		return fmt.Errorf("geul/pack: 0x%016X is the reserved Issuer list ID (PutNamed `issuers/…`)", doc)
	}
	num, ok := s.inum[core.FormatID(issuer)]
	if !ok {
		return fmt.Errorf("geul/pack: Issuer %s is not one of the stream's", core.FormatID(issuer))
	}
	key := pKey(num, doc, slot)
	if s.st.Close.Phase > closeNone {
		return s.closedPut(key)
	}
	if s.st.Count > 0 && key <= s.st.Last {
		return fmt.Errorf("geul/pack: %s/0x%016X slot %d is not after the previous record", core.FormatID(issuer), doc, slot)
	}
	if len(s.st.PLeaf) == s.o.Build.PLeaf {
		if err := s.flushLeaf(); err != nil {
			return err
		}
	}
	loc, err := s.place(r.Bytes, r.Enc)
	if err != nil {
		return err
	}
	e := streamInv{Kind: invP, Enc: r.Enc, Ref: num, Doc: doc, Slot: slot, N: uint32(len(r.Bytes)), Off: loc.Off}
	s.st.PLeaf = append(s.st.PLeaf, streamKV{Key: key, Loc: loc})
	s.st.Last, s.st.Count = key, s.st.Count+1
	s.count(r)
	if r.HashOK {
		e.Flags = flagHash
		s.hbuf = append(s.hbuf, hEntry{key: r.Hash, loc: loc})
	}
	s.st.Inv = append(s.st.Inv, e)
	if len(s.hbuf) >= s.o.RunEntries {
		return s.spill()
	}
	return nil
}
