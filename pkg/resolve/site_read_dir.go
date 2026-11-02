//ff:func feature=resolve type=resolver control=sequence
//ff:what site.ReadDir — fs.ReadDirFS: a directory of a local site; an error on a remote one
package resolve

import "io/fs"

// ReadDir lists name.
func (s *site) ReadDir(name string) ([]fs.DirEntry, error) {
	l, err := s.listing()
	if err != nil {
		return nil, err
	}
	return l.ReadDir(name)
}
