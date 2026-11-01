//ff:func feature=pack type=writer control=sequence
//ff:what TestCloseCodec — Zstd becomes a one-encoder copy at the same level with the same frames, its cached encoders dropped; other codecs stay
package pack

import (
	"bytes"
	"testing"
)

func TestCloseCodec(t *testing.T) {
	if c := closeCodec(Stored{}, nil); c != (Stored{}) {
		t.Fatal("Stored: want the same")
	}
	z := &Zstd{Level: 3, Workers: 4}
	d, err := z.Train(streamTestOptions().Samples, FirstDictID)
	if err != nil {
		t.Fatal(err)
	}
	src := bytes.Repeat(streamTestDoc(1, 1), 3)
	before, _ := z.Compress(src, nil)
	z.Compress(src, d)
	c := closeCodec(z, d)
	one, ok := c.(*Zstd)
	if !ok || one.Workers != 1 || one.Level != 3 || z.plain != nil || d.enc != nil {
		t.Fatalf("%+v", c)
	}
	if after, err := c.Compress(src, nil); err != nil || !bytes.Equal(after, before) {
		t.Fatal("frames differ")
	}
	if c := closeCodec(&Zstd{}, nil); c.(*Zstd).Workers != 1 {
		t.Fatal("no dictionary")
	}
}
