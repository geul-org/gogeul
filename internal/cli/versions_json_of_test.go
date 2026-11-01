//ff:func feature=cli type=helper control=sequence
//ff:what TestVersionsJSONOf — hex hash, RFC 3339 time, state; an unknown time is empty; no versions is an empty list
package cli

import (
	"testing"
	"time"

	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestVersionsJSONOf(t *testing.T) {
	h := resolve.History{Next: 5, Entries: []resolve.LogEntry{
		{VersionEntry: resolve.VersionEntry{Version: 1, Hash: [32]byte{0xab}, Deprecated: true}, At: time.Unix(0, 0)},
		{VersionEntry: resolve.VersionEntry{Version: 2}},
	}}
	got := versionsJSONOf(h)
	v := got.Versions
	if got.NextLocalID != 5 || len(v) != 2 || v[0].Hash[:6] != "0xab00" || len(v[0].Hash) != 66 || v[0].At != "1970-01-01T00:00:00Z" || v[0].State != "deprecated" {
		t.Fatalf("%+v", got)
	}
	if v[1].At != "" || v[1].State != "" {
		t.Fatalf("%+v", v[1])
	}
	if got := versionsJSONOf(resolve.History{}); got.Versions == nil || len(got.Versions) != 0 {
		t.Fatalf("%+v", got)
	}
}
