//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterReadRecord — the encoded record; changed bytes, Get and codec errors fail
package pack

import (
	"errors"
	"testing"
)

func TestWriterReadRecord(t *testing.T) {
	data := map[string][]byte{"issuers/0x0000000A.geul": []byte("abc"), "issuers/0x0000000b.geul": []byte("lower")}
	w := testWriter(t, t.TempDir(), Batch{Put: []string{"issuers/0x0000000A.geul", "issuers/0x0000000b.geul"}, Get: memGet(data)})
	w.plan()
	p := w.puts[0]
	if e, err := w.readRecord(w.puts[1]); err != nil || string(e.b) != "lower" || w.puts[1].path != "issuers/0x0000000B.geul" {
		t.Fatal("source spelling:", e, err)
	}
	if e, err := w.readRecord(p); err != nil || e.enc != encRaw || string(e.b) != "abc" {
		t.Fatal(e, err)
	}
	data["issuers/0x0000000A.geul"] = []byte("abd")
	if _, err := w.readRecord(p); err == nil {
		t.Fatal("changed")
	}
	delete(data, "issuers/0x0000000A.geul")
	if _, err := w.readRecord(p); err == nil {
		t.Fatal("get")
	}
	data["issuers/0x0000000A.geul"] = []byte("abc")
	w.o.Codec = fakeCodec{err: errors.New("x")}
	if _, err := w.readRecord(p); err == nil {
		t.Fatal("codec")
	}
}
