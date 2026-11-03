//ff:func feature=toulmin type=helper control=iteration dimension=1
//ff:what TestS314Names — every Edge row of the §3.14 table (pkg/core testdata) is a direct TOULMIN codec with that name, encoded as TAG then Payload
package toulmin

import (
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestS314Names(t *testing.T) {
	b, err := os.ReadFile("../core/testdata/migration1/s314_p2.tsv")
	if err != nil {
		t.Fatal(err)
	}
	edges := 0
	for _, line := range strings.Split(strings.TrimSpace(string(b)), "\n")[1:] {
		a := strings.Split(line, "\t")
		tag, err := strconv.ParseUint(strings.TrimSuffix(a[3], "L")+"0", 0, 16)
		if a[0] != "Edge" || err != nil {
			continue
		}
		edges++
		k, c := core.SplitTarget(core.Word(tag))
		ws, err := EdgeRecord(CodecRef{Codec: c}, []core.Word{9})
		if k != core.KindEdge || RecordName(CodecRef{Codec: c}) != a[2] || err != nil || !slices.Equal(ws, []core.Word{core.Word(tag) | 1, 9}) {
			t.Fatalf("%s: %v %X %q %04X %v", a[2], k, c, RecordName(CodecRef{Codec: c}), ws, err)
		}
	}
	if edges != 21 {
		t.Fatalf("%d Edge rows, want 21", edges)
	}
}
