//ff:func feature=pack type=writer control=sequence
//ff:what Stream.hPace — a safe point of the H pass: when Pace stops Close the pass position is saved first
package pack

// hPace asks Pace (Stream.pace) whether to go on. A stop saves the position
// (hSave), so it loses nothing of the pass.
func (s *Stream) hPace(p *hPass) error {
	err := s.pace("close:h")
	if err == nil {
		return nil
	}
	if serr := s.hSave(p); serr != nil {
		return serr
	}
	return err
}
