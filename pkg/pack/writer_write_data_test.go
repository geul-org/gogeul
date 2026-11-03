//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterWriteData — records go into packs split at PackSize (an oversized record alone); a failing record or dictionary sample stops it
package pack

import (
	"testing"
)

func TestWriterWriteData(t *testing.T) {
	out := t.TempDir()
	get := memGet(map[string][]byte{"issuers/0x0000000a.geul": make([]byte, 100), "issuers/0x0000000b.geul": make([]byte, 10), "issuers/0x0000000c.geul": make([]byte, 10)})
	w := testWriter(t, out, Batch{Put: []string{"issuers/0x0000000a.geul", "issuers/0x0000000b.geul", "issuers/0x0000000c.geul"}, Get: get})
	w.o.PackSize = 64 + 30
	if err := w.plan(); err != nil {
		t.Fatal(err)
	}
	never().MkdirAll(out + "/o")
	if err := w.writeData(); err != nil || len(w.objs) != 2 || w.res.Records != 3 {
		t.Fatal(err, len(w.objs))
	}
	w = testWriter(t, out, Batch{Put: []string{"issuers/0x0000000a.geul"}, Get: get})
	w.plan()
	w.b.Get = memGet(nil)
	if err := w.writeData(); err == nil {
		t.Fatal("get")
	}
	paths, sget := synth(40)
	w = testWriter(t, out, Batch{Put: paths, Get: sget})
	w.o.Codec = &Zstd{DictSize: 4096}
	w.plan()
	w.b.Get = memGet(nil)
	if err := w.writeData(); err == nil {
		t.Fatal("dictionary samples")
	}
}
