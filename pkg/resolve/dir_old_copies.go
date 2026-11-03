//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Dir.oldCopies — the h/ hashes a version list reaches: its versions, itself and every earlier list by PARENT
package resolve

import "context"

// oldCopies collects the content hashes of the `<Issuer>/hashes/` copies the
// version list lg of document k reaches: every version it names, the list itself and,
// following PARENT, every earlier list and the versions those name. An
// earlier list that is missing or does not decode ends the chain.
func (d Dir) oldCopies(lg versionLog, k DocKey) [][32]byte {
	var out [][32]byte
	seen := map[[32]byte]bool{}
	for lg.Present && !seen[lg.Hash] {
		seen[lg.Hash] = true
		out = append(out, lg.Hash)
		for _, e := range lg.Entries {
			out = append(out, e.Hash)
		}
		if !lg.HasParent {
			break
		}
		b, err := d.Fetch(context.Background(), hashPath(k.Issuer, lg.Parent))
		if err != nil {
			break
		}
		if lg, err = decodeVersionList(b, k.Issuer, k.Doc); err != nil {
			break
		}
	}
	return out
}
