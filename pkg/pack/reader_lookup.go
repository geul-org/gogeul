//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.Lookup — the record of (Issuer, document, slot); slot 0 is versions.geul
package pack

import "context"

// Lookup reads a P record. issuer is canonical hex; an Issuer outside the
// manifest's table is NotFound.
func (r *Reader) Lookup(ctx context.Context, issuer string, doc uint64, slot uint32) ([]byte, error) {
	path := pPath(issuer, doc, slot)
	num, ok := r.inum[issuer]
	if !ok {
		return nil, notFound(path)
	}
	l, err := r.find(ctx, r.m.pLayers, pKey(num, doc, slot), path)
	if err != nil {
		return nil, err
	}
	return r.Record(ctx, l)
}
