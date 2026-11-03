//ff:func feature=toulmin type=helper control=iteration dimension=1
//ff:what TestRecordNameVector — the TOULMIN direct codes of the migration1 reference vector have the same names here (grammar = code)
package toulmin

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRecordNameVector(t *testing.T) {
	f, err := os.Open("../core/testdata/migration1/vector.tsv")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	missing := map[string]bool{"SUPERSEDE-CHAIN": true, "VERDICT": true, "EVIDENCE": true}
	seen := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		a := strings.Split(sc.Text(), "\t")
		w, err := strconv.ParseUint(a[1], 16, 16)
		if err != nil || !strings.HasPrefix(a[2], "E:") || w&0xF != 0 {
			continue
		}
		k, c := core.SplitTarget(core.Word(w))
		name := strings.TrimSuffix(strings.TrimPrefix(a[2], "E:"), "/L0")
		if k != core.KindEdge || c < 0x60 || c > 0x7B || missing[name] {
			continue
		}
		seen++
		if got := RecordName(CodecRef{Codec: c}); got != name {
			t.Fatalf("Edge 0x%02X: %q want %q", c, got, name)
		}
	}
	if seen != 15 {
		t.Fatalf("TOULMIN direct codes %d, want 15", seen)
	}
}
