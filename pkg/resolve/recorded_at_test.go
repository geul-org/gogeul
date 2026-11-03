//ff:func feature=resolve type=resolver control=sequence
//ff:what TestRecordedAt — the RECORDED-AT time; none or an unreadable one is zero
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRecordedAt(t *testing.T) {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	at := core.Record{Kind: core.KindMeta, Codec: core.MetaRecordedAt, Payload: enc(core.EncodePosixSeconds(1000))}
	if got := recordedAt([]core.Record{core.EncodeSpecVersion(2), at}); got.Unix() != 1000 {
		t.Fatalf("got %v", got)
	}
	if !recordedAt(nil).IsZero() || !recordedAt([]core.Record{{Kind: core.KindMeta, Codec: core.MetaRecordedAt, Payload: []core.Word{0, 2}}}).IsZero() {
		t.Fatal("want zero")
	}
}
