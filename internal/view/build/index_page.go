//ff:type feature=build type=model
//ff:what indexPage — a directory listing: its URL path, parent link and entries
package build

// indexPage is the data of a directory listing. Up is the parent URL, empty at
// the top of the tree.
type indexPage struct {
	Path    string
	Up      string
	Entries []indexEntry
}
