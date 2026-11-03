//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestParseRecord — every token form reads back to what formatRecord writes; empty and unreadable tokens (the old _ marker among them) fail
package notation

import "testing"

func TestParseRecord(t *testing.T) {
	ctx := Context{Lookup: fakeLookup}
	for _, s := range []string{"ABSENT", "UNKNOWN", "NONE", "#3", "Q42", "P31", `"a b"`, `"x"@ko`, "<https://geul.org/>", "1.20", "-3", "0x00000000/0x0000000000000806@1", "0x00000000/0x0000000000000806#2", "GREGORIAN-DATE(1952-03-11)", "QUANTITY(652, Q828224, ABSENT, ABSENT)", "0x00A10005"} {
		r, err := parseRecord(s, ctx)
		if err != nil || formatRecord(r) != s {
			t.Errorf("%s: %s %v", s, formatRecord(r), err)
		}
	}
	for _, s := range []string{"", "_", "word", "Q", "1.2.3"} {
		if _, err := parseRecord(s, ctx); err == nil {
			t.Errorf("%q: want error", s)
		}
	}
}
