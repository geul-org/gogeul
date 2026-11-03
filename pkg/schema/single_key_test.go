//ff:func feature=schema type=rule control=sequence
//ff:what TestSingleKey — Edges never count; MESSAGE keys by language or not at all; single Meta by code; repeatable Meta not single
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestSingleKey(t *testing.T) {
	if _, _, ok := singleKey(recOf(edgeOf(core.EdgeSource, 1))); ok {
		t.Fatal("edge")
	}
	if rule, key, ok := singleKey(recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaVerdict})); !ok || rule != "M8" || key != "0x29" {
		t.Fatal(rule, key, ok)
	}
	if _, _, ok := singleKey(recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaCoAuthor})); ok {
		t.Fatal("CO-AUTHOR may repeat")
	}
	if _, _, ok := singleKey(recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaMessage})); ok {
		t.Fatal("unreadable MESSAGE")
	}
}
