//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what streamTestPut — put the list and version of test documents from..to into a stream, in key order
package pack

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// streamTestPut encodes and puts streamTestDoc(d, 0) and (d, 1) for d in
// from..to under Issuer 0x00010000.
func streamTestPut(t *testing.T, s *Stream, from, to uint64) {
	t.Helper()
	for i := 2 * from; i < 2*to+2; i++ {
		r, err := s.Encode(streamTestDoc(i/2, uint32(i%2)))
		if err == nil {
			err = s.Put([]core.Word{0, 1}, 0x5100000000000000|i/2, uint32(i%2), r)
		}
		if err != nil {
			t.Fatal(err)
		}
	}
}
