//ff:func feature=pack type=writer control=sequence
//ff:what TestLayersOf — saved layers come back with count, encoding and root, and encode as the originals did
package pack

import (
	"bytes"
	"testing"
)

func TestLayersOf(t *testing.T) {
	orig := []layer{{count: 3, enc: encZstd, root: []byte("r")}, {count: 1, enc: encRaw, root: []byte("s")}}
	got := layersOf(streamLayersOf(orig))
	if len(got) != 2 || got[1].count != 1 || !bytes.Equal(encodeLayers(got), encodeLayers(orig)) {
		t.Fatalf("%+v", got)
	}
	if layersOf(nil) != nil {
		t.Fatal("none")
	}
}
