//ff:func feature=pack type=writer control=sequence
//ff:what Stream.closedPut — a put after Close began: a key up to the last one was stored before and is acknowledged, a later one is an error
package pack

import "errors"

// closedPut answers Put on a stream whose Close has begun (resumed after
// Close stopped): the caller, going on from its last checkpoint, puts again
// records the stream already holds, so keys up to Last are accepted without
// storing anything; a key after Last cannot be added any more.
func (s *Stream) closedPut(key string) error {
	if s.st.Count > 0 && key <= s.st.Last {
		return nil
	}
	return errors.New("geul/pack: the stream is closing; no record can follow its last one")
}
