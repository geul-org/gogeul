//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Reader.Documents — the document IDs (16-digit hex) with a live versions or version record under an Issuer
package pack

import (
	"context"
	"fmt"
)

// Documents lists issuer's documents in ID order; an unknown Issuer has none.
func (r *Reader) Documents(ctx context.Context, issuer string) ([]string, error) {
	num, ok := r.inum[issuer]
	if !ok {
		return nil, nil
	}
	s, err := r.live(ctx, r.m.pLayers, pKey(num, 0, 0), pKey(num+1, 0, 0))
	if err != nil {
		return nil, err
	}
	var out []string
	for _, k := range sortedKeys(s) {
		_, d, _ := splitPKey(k)
		id := fmt.Sprintf("0x%016X", d)
		if len(out) == 0 || out[len(out)-1] != id {
			out = append(out, id)
		}
	}
	return out, nil
}
