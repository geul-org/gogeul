//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestStreamEncode — a closing stream leaves records raw; a document compresses and carries its content hash; Stored keeps bytes raw; incompressible bytes stay raw; non-documents have no hash; a failing codec fails
package pack

import (
	"crypto/sha256"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestStreamEncode(t *testing.T) {
	doc := streamTestDoc(1, 1)
	ws, _ := core.ReadWords(doc)
	h, _ := core.ContentHash(ws)
	r, err := streamTestNew(t, nil).Encode(doc)
	if err != nil || r.Enc != encZstd || len(r.Bytes) >= len(doc) || r.Raw != len(doc) || !r.HashOK || r.Hash != h {
		t.Fatalf("%+v %v", r.Enc, err)
	}
	if r, _ := streamTestNew(t, Stored{}).Encode(doc); r.Enc != encRaw || string(r.Bytes) != string(doc) {
		t.Fatal("Stored: want raw")
	}
	var noise []byte
	for i := 0; i < 20; i++ {
		s := sha256.Sum256([]byte{byte(i)})
		noise = append(noise, s[:]...)
	}
	if r, err := streamTestNew(t, nil).Encode(noise); err != nil || r.Enc != encRaw || r.HashOK {
		t.Fatalf("noise: %+v %v", r, err)
	}
	cl := streamTestNew(t, failCodec{})
	cl.st.Close.Phase = closeStarted
	if r, err := cl.Encode(doc); err != nil || r.Enc != encRaw || !r.HashOK {
		t.Fatalf("closing: %+v %v", r.Enc, err)
	}
	if _, err := streamTestNew(t, failCodec{}).Encode(doc); err == nil {
		t.Fatal("failing codec: want error")
	}
}
