//ff:func feature=pack type=writer control=sequence
//ff:what Stream.pace — a safe point of Close: call StreamOptions.Pace with the phase; its error comes back as ErrStopped too
package pack

import "fmt"

// pace calls StreamOptions.Pace, if set. Its error stops Close, whose saved
// progress ResumeStream goes on from.
func (s *Stream) pace(phase string) error {
	if s.o.Pace == nil {
		return nil
	}
	if err := s.o.Pace(phase); err != nil {
		return fmt.Errorf("%w at %s: %w", ErrStopped, phase, err)
	}
	return nil
}
