//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckRecords — HEAD and BODY records both checked, with the HEAD's IMPORT IDs
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCheckRecords(t *testing.T) {
	head := []rec{recOf(core.Record{Kind: core.KindMeta, Codec: 0x2A})}
	body := []rec{recOf(core.Record{Kind: core.KindReference, Codec: core.CodecExternalPointer1Word, Payload: []core.Word{3, 1}})}
	if fs := checkRecords(docParts{Head: head, Body: body}); rulesOf(fs) != "M1,R3" {
		t.Fatalf("got %+v", fs)
	}
}
