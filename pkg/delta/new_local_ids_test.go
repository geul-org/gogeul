//ff:func feature=delta type=decoder control=sequence
//ff:what TestNewLocalIDs — no previous version starts at 1; a previous one starts above its largest LOCAL or the given next and carries its HEAD Edges; bad bytes and documents fail
package delta

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestNewLocalIDs(t *testing.T) {
	if l, err := NewLocalIDs(nil, 0); err != nil || l.Next() != 1 {
		t.Fatalf("first: %v", err)
	}
	prev := encodeDoc(nil, []core.Record{bindOf(1, 9, "a"), refNodeOf(2, 3, "http://x")})
	if l, err := NewLocalIDs(prev, 0); err != nil || l.Next() != 10 {
		t.Fatalf("prev: %d %v", l.Next(), err)
	}
	if l, _ := NewLocalIDs(prev, 50); l.Next() != 50 {
		t.Fatalf("kept next: %d", l.Next())
	}
	src := core.Record{Kind: core.KindEdge, Codec: core.EdgeSource, Payload: []core.Word{0x0051, 12, 0x0051, 0, 0x0141, 42}}
	withHead := encodeDoc([]core.Record{src}, []core.Record{bindOf(1, 9, "a")})
	l, err := NewLocalIDs(withHead, 0)
	if err != nil || l.Next() != 13 || l.Edge(core.EdgeSource, []core.Word{0x0051, 0, 0x0141, 42}) != 12 {
		t.Fatalf("HEAD SOURCE carried over: %v", err)
	}
	if _, err := NewLocalIDs([]byte{1}, 0); err == nil {
		t.Fatal("odd bytes: want error")
	}
	if _, err := NewLocalIDs([]byte{0x00, 0x05}, 0); err == nil {
		t.Fatal("not a document: want error")
	}
}
