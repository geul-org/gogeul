//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestRecordRun — a run of embedded records passes; empty, undecodable, implausible and trailing words fail
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRecordRun(t *testing.T) {
	if !recordRun([]core.Word{0x0051, 7, 0x0722, 0x6d61, 0x696e}) {
		t.Fatal("LOCAL + UTF8-TEXT: want true")
	}
	for i, ws := range [][]core.Word{nil, {0xE000}, {0x0060}, {0x0051, 7, 0x0060}, {0x0602}} {
		if recordRun(ws) {
			t.Fatalf("case %d: want false", i)
		}
	}
}
