//ff:func feature=inspect type=decoder control=sequence
//ff:what TestPayloadParts — a codec spec lays out first; the fitting SPAN width is found; content without a fitting typed layout falls back to generic parts
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestPayloadParts(t *testing.T) {
	node, _ := toulmin.EncodeNode(toulmin.NodePacket{TID: 1, Type: toulmin.NodeCase, Span: toulmin.Span{Off: 70000}}, toulmin.Wire{Span: 2})
	r, _, _ := core.DecodeRecord(node)
	p := toulmin.NodePacket{TID: 1, Type: toulmin.NodeCase, Span: toulmin.Span{Off: 70000}}
	parts := payloadParts(Entry{Record: r, Words: node, Data: p})
	if len(parts) != 4 || parts[2].Name != "span.off" || len(parts[2].Words) != 2 {
		t.Fatalf("2-word span: %+v", parts)
	}
	r.Payload = r.Payload[:len(r.Payload)-1]
	if parts := payloadParts(Entry{Record: r, Data: p}); len(parts) != 2 || parts[1].Name != "payload" {
		t.Fatalf("no fitting width: %+v", parts)
	}
	tomb := core.EncodePublication(core.Publication{Codec: core.EdgeTombstone, EdgeID: 3, Issuer: []core.Word{0, 0}, Doc: []core.Word{0, 0, 0, 1}, Version: 1, Reason: 1})
	if parts := payloadParts(Entry{Record: tomb, Spec: tombstoneSpec()[0x09A0]}); len(parts) != 3 || parts[2].Enum != "license" {
		t.Fatalf("spec: %+v", parts)
	}
}
