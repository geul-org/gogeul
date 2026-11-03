//ff:func feature=page type=builder control=sequence
//ff:what TestBindNames — BIND Subjects by their words, the last BIND winning; non-BIND records skipped
package page

import (
	"fmt"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestBindNames(t *testing.T) {
	body := []inspect.Entry{bindEntry(1, 7, "a"), bindEntry(2, 7, "b"), bindEntry(3, 1, "c"), inspectEntry(core.EncodeUintValue(5), nil)}
	ns := bindNames(body)
	seven := fmt.Sprint([]core.Word{0x0051, 7})
	if len(ns) != 2 || ns[seven] != "b" || ns[fmt.Sprint([]core.Word{0x0051, 1})] != "c" {
		t.Fatalf("got %v", ns)
	}
}
