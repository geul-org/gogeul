//ff:func feature=pack type=encoder control=sequence
//ff:what TestZstdSetWorkers — the kept encoders are made again for the new count with the same frames, the dictionary's too; below 1 changes nothing
package pack

import (
	"bytes"
	"testing"
)

func TestZstdSetWorkers(t *testing.T) {
	z := &Zstd{Level: 3, Workers: 4}
	d, err := z.Train(streamTestOptions().Samples, FirstDictID)
	if err != nil {
		t.Fatal(err)
	}
	src := bytes.Repeat(streamTestDoc(1, 1), 3)
	plain, _ := z.Compress(src, nil)
	dict, _ := z.Compress(src, d)
	e, enc := z.plain, d.enc
	z.SetWorkers(0)
	if z.Workers != 4 || z.plain != e {
		t.Fatal("below 1: changed")
	}
	z.SetWorkers(2)
	if z.Workers != 2 || z.plain != nil {
		t.Fatal("encoder kept")
	}
	if b, err := z.Compress(src, nil); err != nil || !bytes.Equal(b, plain) {
		t.Fatal("plain frames differ")
	}
	if b, err := z.Compress(src, d); err != nil || !bytes.Equal(b, dict) || d.enc == enc || d.used != z.gen {
		t.Fatal("dictionary frames differ or encoder kept")
	}
	again := d.enc
	if _, err := z.Compress(src, d); err != nil || d.enc != again {
		t.Fatal("encoder not reused")
	}
}
