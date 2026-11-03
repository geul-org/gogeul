//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what TestExampleHash — the grammar example.md document in new WORDs (gen12.json ex_new) splits and hashes to the generated new content hash
package core

import (
	"encoding/hex"
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestExampleHash(t *testing.T) {
	b, err := os.ReadFile("testdata/migration1/gen12.json")
	if err != nil {
		t.Fatal(err)
	}
	var g struct {
		ExNew   []string `json:"ex_new"`
		NewHash string   `json:"new_hash"`
	}
	if err := json.Unmarshal(b, &g); err != nil {
		t.Fatal(err)
	}
	var ws []Word
	for _, h := range strings.Fields(strings.Join(g.ExNew, " ")) {
		v, _ := strconv.ParseUint(h, 16, 16)
		ws = append(ws, Word(v))
	}
	if len(ws) != 63 {
		t.Fatalf("%d words", len(ws))
	}
	h, err := ContentHash(ws)
	if err != nil || hex.EncodeToString(h[:]) != g.NewHash {
		t.Fatalf("%x %v", h, err)
	}
	d, err := SplitDocument(ws)
	if err != nil || len(d.Head) != 2 || len(d.Body) != 4 {
		t.Fatalf("%+v %v", d, err)
	}
	if iss, doc, err := DecodeDocID(d.Head[1]); err != nil || iss[1] != 1 || doc[0] != 0x5100 {
		t.Fatalf("DOC-ID %04X %04X %v", iss, doc, err)
	}
}
