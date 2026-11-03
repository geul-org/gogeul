//ff:func feature=resolve type=resolver control=sequence
//ff:what TestResolverFetchWith — order, not-found fall-through, `|` fall-through, `,` stop, private skip, off, verify stop
package resolve

import (
	"context"
	"errors"
	"testing"
)

func TestResolverFetchWith(t *testing.T) {
	ctx := context.Background()
	tree := Entry{Name: "tree", Source: sampleTree(t)}
	broken := Entry{Name: "broken", Source: HTTP{Base: "http://\x7f"}}
	empty := Entry{Name: "empty", Source: Memory{}}
	k := sampleKey(1, 0)
	fetch := func(es ...Entry) (*Node, error) {
		return (&Resolver{Config: Config{Proxy: es}}).fetchWith(ctx, k, fetchFrom)
	}
	if n, err := fetch(empty, tree); err != nil || n.Source != "tree" {
		t.Fatalf("not found then tree: %+v %v", n, err)
	}
	pipe := broken
	pipe.Fallthrough = true
	if n, err := fetch(pipe, tree); err != nil || n.Source != "tree" {
		t.Fatalf("| fall-through: %+v %v", n, err)
	}
	if _, err := fetch(broken, tree); err == nil || errors.Is(err, ErrNotFound) {
		t.Fatalf(", stop: %v", err)
	}
	if _, err := fetch(empty, Entry{Off: true}, tree); !errors.Is(err, ErrNotFound) {
		t.Fatalf("off: %v", err)
	}
	pub := tree
	pub.Public = true
	r := &Resolver{Config: Config{Proxy: []Entry{pub}, Private: []Pattern{{Issuer: "0x00000000"}}}}
	if _, err := r.fetchWith(ctx, k, fetchFrom); !errors.Is(err, ErrNotFound) {
		t.Fatalf("private: %v", err)
	}
	bad := Entry{Name: "bad", Source: Memory{listPath(k.DocKey()): []byte("junk")}, Fallthrough: true}
	if _, err := fetch(bad, tree); !errors.Is(err, ErrVerify) {
		t.Fatalf("verify stop: %v", err)
	}
}
