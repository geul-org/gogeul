//ff:func feature=cli type=command control=sequence
//ff:what TestAppShowLanguage — a language code is named by the registry, whose codec only appears embedded in a text
package cli

import (
	"strings"
	"testing"
)

func TestAppShowLanguage(t *testing.T) {
	h := newHarness(t)
	h.languageSpec()
	addr, draft := h.doc()
	h.must("add", draft, "ALIAS", "Q42", `"Doug"@@en`)
	h.must("publish", draft)
	if out := h.must("show", addr); !strings.Contains(out, `"Doug"@@en`) || strings.Contains(out, "0x0784") {
		t.Fatal(out)
	}
}
