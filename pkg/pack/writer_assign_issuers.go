//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.assignIssuers — append the put Issuers the table lacks (numbers never change, §4.3 ISSUERS)
package pack

import "sort"

// assignIssuers extends the Issuer table with new Issuers in resolve.Dir order.
func (w *writer) assignIssuers() {
	var add []string
	for _, p := range w.puts {
		if _, ok := w.inum[p.ref.issuer]; p.ref.issuer != "" && !ok {
			w.inum[p.ref.issuer] = 0
			add = append(add, p.ref.issuer)
		}
	}
	sort.Slice(add, func(i, j int) bool { return issuerLess(add[i], add[j]) })
	for _, id := range add {
		w.inum[id] = uint32(len(w.issuers))
		w.issuers = append(w.issuers, id)
	}
}
