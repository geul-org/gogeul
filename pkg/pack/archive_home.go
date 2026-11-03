//ff:func feature=pack type=writer control=sequence
//ff:what Archive.home — where a new path outside every Issuer goes: the top set, else Issuer 0x0000's, else the first, else a new top set
package pack

// home picks the set for new top-level paths (the Issuer list `issuers/`).
func (a *Archive) home() *packSet {
	for _, name := range []string{"", "0x00000000"} {
		if s := a.find(name); s != nil {
			return s
		}
	}
	if len(a.sets) > 0 {
		return a.sets[0]
	}
	return a.set("")
}
