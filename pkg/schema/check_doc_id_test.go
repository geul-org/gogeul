//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckDocID — a DOC-POINTER and a CODEC-POINTER DOC-ID pass; a raw payload and the reserved list ID F…F fail M2
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCheckDocID(t *testing.T) {
	ok, _ := core.EncodeDocID([]core.Word{1, 1}, []core.Word{0, 0, 0, 5})
	raw := core.Record{Kind: core.KindMeta, Codec: core.MetaDocID, Payload: make([]core.Word, 5)}
	if checkDocID(recOf(ok)) != nil || rulesOf(checkDocID(recOf(raw))) != "M2" {
		t.Fatal("wrong")
	}
	list, _ := core.EncodeDocID([]core.Word{0, 0}, []core.Word{0xffff, 0xffff, 0xffff, 0xffff})
	if rulesOf(checkDocID(recOf(list))) != "M2" {
		t.Fatal("reserved list ID as DOC-ID")
	}
	codec, _ := core.EncodeDocIDCodec([]core.Word{0, 0}, 0x0AA0)
	if checkDocID(recOf(codec)) != nil {
		t.Fatal("codec document")
	}
}
