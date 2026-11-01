//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what Stream.closedPutNamed — a named put after Close began: a path the stream holds is acknowledged, a new one is an error
package pack

import "fmt"

// closedPutNamed answers PutNamed on a stream whose Close has begun: a path
// put before is accepted without storing (the caller is putting it again
// after a resume); a new path is an error.
func (s *Stream) closedPutNamed(path string) error {
	for _, n := range s.st.Named {
		if n.Key == path {
			return nil
		}
	}
	return fmt.Errorf("geul/pack: the stream is closing; %q cannot be added", path)
}
