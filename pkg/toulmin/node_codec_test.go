//ff:func feature=toulmin type=encoder control=iteration dimension=1
//ff:what TestNodeCodec — every node type round-trips through NodeFromCodec; unknown types fail
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestNodeCodec(t *testing.T) {
	for _, nt := range nodeTypes {
		ref, err := NodeCodec(nt, RoleGeneral)
		if err != nil {
			t.Fatalf("%d: %v", nt, err)
		}
		got, _, _, ok := NodeFromCodec(ref)
		if !ok || got != nt {
			t.Fatalf("%d: round trip got %d ok=%v", nt, got, ok)
		}
	}
	if _, err := NodeCodec(0x00FF, RoleGeneral); err == nil {
		t.Fatal("unknown NodeType must fail")
	}
	if ref, _ := NodeCodec(NodeCase, RoleGeneral); ref.Codec != 0x60 || ref.Format != 0 {
		t.Fatalf("case codec: %+v", ref)
	}
	if ref, _ := NodeCodec(NodeAlias, RoleGeneral); ref.Codec != core.CodecEdgeIssuer2Word || ref.Format != FormatAlias {
		t.Fatalf("alias codec: %+v", ref)
	}
}
