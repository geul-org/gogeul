//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamLayersOf — count, encoding and stored root are kept; no layers stay none
package pack

import "testing"

func TestStreamLayersOf(t *testing.T) {
	got := streamLayersOf([]layer{{count: 3, enc: encZstd, root: []byte("r"), blk: block{kind: blkPLeaf}}})
	if len(got) != 1 || got[0].Count != 3 || got[0].Enc != encZstd || string(got[0].Root) != "r" {
		t.Fatalf("%+v", got)
	}
	if streamLayersOf(nil) != nil {
		t.Fatal("none")
	}
}
