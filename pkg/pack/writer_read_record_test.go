//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterReadRecord — the encoded record; changed bytes, Get and codec errors fail
package pack

import (
	"errors"
	"testing"
)

func TestWriterReadRecord(t *testing.T) {
	data := map[string][]byte{"issuers/0x0000000a.geul": []byte("abc")}
	w := testWriter(t, t.TempDir(), Batch{Put: []string{"issuers/0x0000000a.geul"}, Get: memGet(data)})
	w.plan()
	p := w.puts[0]
	if e, err := w.readRecord(p); err != nil || e.enc != encRaw || string(e.b) != "abc" {
		t.Fatal(e, err)
	}
	data["issuers/0x0000000a.geul"] = []byte("abd")
	if _, err := w.readRecord(p); err == nil {
		t.Fatal("changed")
	}
	delete(data, "issuers/0x0000000a.geul")
	if _, err := w.readRecord(p); err == nil {
		t.Fatal("get")
	}
	data["issuers/0x0000000a.geul"] = []byte("abc")
	w.o.Codec = fakeCodec{err: errors.New("x")}
	if _, err := w.readRecord(p); err == nil {
		t.Fatal("codec")
	}
}
