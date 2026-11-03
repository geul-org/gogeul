//ff:func feature=resolve type=helper control=sequence
//ff:what TestAddrCodec — an 8-digit codec ID with a stored tail is AddrCodec; a 16-digit ID or a bad tail is bad
package resolve

import (
	"errors"
	"testing"
)

func TestAddrCodec(t *testing.T) {
	bad := errors.New("bad")
	a, err := addrCodec(Addr{}, "0x00000980", "0x00000002.geul", bad)
	if err != nil || a.Kind != AddrCodec || a.ID[1] != 0x0980 || a.Version != 2 {
		t.Fatal(a, err)
	}
	if _, err := addrCodec(Addr{}, "0x0000000000000980", "latest.geul", bad); err != bad {
		t.Fatal(err)
	}
	if _, err := addrCodec(Addr{}, "0x00000980", "0x00000000.geul", bad); err != bad {
		t.Fatal(err)
	}
}
