//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Dir.namedIn — the folders of dir named as canonical IDs of the given width
package resolve

// namedIn lists the folders of dir (`<Issuer>/docs` or `<Issuer>/codecs`) whose
// names are canonical hex IDs of words words, in name order. A missing folder
// lists none.
func (d Dir) namedIn(dir string, words int) []Named {
	ents, _ := d.store().ReadDir(dir)
	var out []Named
	for _, e := range ents {
		if _, ok := storedHexID(e.Name(), words); ok && e.IsDir() {
			out = append(out, Named{ID: e.Name()})
		}
	}
	return out
}
