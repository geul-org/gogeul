//ff:func feature=resolve type=resolver control=sequence
//ff:what site.Open — fs.FS: a file or directory of a local site; an error on a remote one
package resolve

import "io/fs"

// Open opens name.
func (s *site) Open(name string) (fs.File, error) {
	l, err := s.listing()
	if err != nil {
		return nil, err
	}
	return l.Open(name)
}
