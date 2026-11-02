//ff:func feature=resolve type=resolver control=selection
//ff:what Dir.store — the Dir's Store: the directory tree under Root, the given Store, or that Store over the loose files under Root
package resolve

// store returns the Dir's reading and writing side (see Dir).
func (d Dir) store() Store {
	switch {
	case d.Store == nil:
		return dirStore{root: d.file("")}
	case d.Root == "":
		return d.Store
	default:
		return overlay{stored: d.Store, loose: dirStore{root: d.file("")}}
	}
}
