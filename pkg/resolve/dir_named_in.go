//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Dir.namedIn — the folders of dir named as canonical IDs of the given width, each with the names byDoc gives it
package resolve

// namedIn lists the folders of dir (`<Issuer>/docs` or `<Issuer>/codecs`) whose
// names are canonical hex IDs of words words, in name order, each with its
// names from byDoc. A missing folder lists none.
func (d Dir) namedIn(dir string, words int, byDoc map[string][]string) []Named {
	ents, _ := d.store().ReadDir(dir)
	var out []Named
	for _, e := range ents {
		if _, ok := parseHexID(e.Name(), words); ok && e.IsDir() {
			out = append(out, Named{ID: e.Name(), Names: byDoc[e.Name()]})
		}
	}
	return out
}
