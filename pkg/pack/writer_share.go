//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.share — let a P or named record also serve its h/ copy when the bytes match; other h/ puts must be named by their content hash
package pack

import "fmt"

// share links each `<Issuer>/hashes/` put to a P or named put with the same bytes whose
// content hash names it under the same Issuer (flags bit0). An unlinked hashes/
// put becomes its own
// record and its name must be its content hash.
func (w *writer) share() error {
	hs := map[string]*put{}
	for _, p := range w.puts {
		if p.ref.space == spaceH {
			hs[p.path] = p
		}
	}
	for _, p := range w.puts {
		h, ok := hs[hPath(p.ref.issuer, string(p.ch[:]))]
		if p.ref.space != spaceH && p.chOK && ok && h.via == nil && h.sum == p.sum {
			p.share, h.via = h.path, p
			w.res.Shared++
		}
	}
	for _, p := range w.puts {
		if p.ref.space == spaceH && p.via == nil && (!p.chOK || string(p.ch[:]) != p.ref.hash) {
			return fmt.Errorf("%w: %s is not the content hash of its bytes", ErrCorrupt, p.path)
		}
	}
	return nil
}
