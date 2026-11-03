//ff:func feature=codecspec type=decoder control=sequence
//ff:what TestLanguages — the code field's ENUMs of the 0848 spec; nil without that spec, or when its code field is not 1-WORD raw
package codecspec

import "testing"

func TestLanguages(t *testing.T) {
	enums := map[string]string{"0x0784": "en", "0x4130": "pt-BR"}
	s := Spec{0x07A0: {{Name: "code", Width: 1, Enums: enums}}}
	if got := s.Languages(); got["0x4130"] != "pt-BR" || len(got) != 2 {
		t.Fatalf("got %v", got)
	}
	if got := (Spec{}).Languages(); got != nil {
		t.Fatalf("no spec: %v", got)
	}
	if got := (Spec{0x07A0: {{Name: "code", Width: 2, Enums: enums}}}).Languages(); got != nil {
		t.Fatalf("2-WORD field: %v", got)
	}
}
