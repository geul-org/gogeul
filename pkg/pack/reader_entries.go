//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Reader.entries — every live path of the generation with its record location (layers and NAMED)
package pack

import "context"

// entries maps each live path to its Loc, as the index answers.
func (r *Reader) entries(ctx context.Context) (map[string]Loc, error) {
	ps, err := r.live(ctx, r.m.pLayers, "", "")
	if err != nil {
		return nil, err
	}
	hs, err := r.live(ctx, r.m.hLayers, "", "")
	if err != nil {
		return nil, err
	}
	out := make(map[string]Loc, len(ps)+len(hs)+len(r.m.named))
	for k, l := range ps {
		i, d, s := splitPKey(k)
		if int(i) >= len(r.m.issuers) {
			return nil, corrupt("P key Issuer number out of range")
		}
		out[pPath(r.m.issuers[i], d, s)] = l
	}
	for k, l := range hs {
		p, err := r.hashEntryPath(ctx, k, l)
		if err != nil {
			return nil, err
		}
		out[p] = l
	}
	for _, n := range r.m.named {
		out[n.path] = n.loc
	}
	return out, nil
}
