//ff:func feature=pack type=writer control=sequence
//ff:what streamPacedClose — the crash test stream closed with a Pace that fails with stop at safe point n, noting the phases in seen
package pack

import "testing"

// streamPacedClose writes the crash test stream to out and closes it; Pace
// returns stop at its call number n (from 0) and records each phase.
func streamPacedClose(t *testing.T, out string, n int, stop error, seen map[string]bool) error {
	t.Helper()
	s, _ := streamCrashStart(t, out, false)
	streamCrashRest(t, s)
	calls := 0
	s.o.Pace = func(phase string) error {
		seen[phase] = true
		calls++
		if calls == n+1 {
			return stop
		}
		return nil
	}
	_, err := s.Close()
	return err
}
