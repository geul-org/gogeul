//ff:func feature=pack type=writer control=sequence
//ff:what Stream.closeSave — sync the open data pack and index object, then save the state (a point Close goes on from)
package pack

// closeSave makes the open objects' bytes durable and saves the state with
// their hashes, so a resume finds what the state counts.
func (s *Stream) closeSave() error {
	if err := syncFile(s.data); err != nil {
		return err
	}
	if err := syncFile(s.index); err != nil {
		return err
	}
	return s.saveState()
}
