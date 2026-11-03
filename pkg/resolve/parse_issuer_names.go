//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what ParseIssuerNames — the name → Issuer map of the Issuer list (issuers/, BIND records); one name bound to two Issuers is an error
package resolve

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// ParseIssuerNames reads the Issuer list (`issuers/`, publishing.md Issuer 목록): every BIND whose
// Subject is ISSUER-POINTER(Issuer) maps its name to that Issuer in
// canonical hex. Other records are ignored. A name is bound to one Issuer
// only (publishing.md "이름 하나는 Issuer 하나에만 묶인다", 0x0A3L BIND), so the
// same name bound to another Issuer is an error rather than silently
// rebinding the name; binding it to the same Issuer again is harmless.
func ParseIssuerNames(b []byte) (map[string]string, error) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return nil, err
	}
	doc, err := core.SplitDocument(ws)
	if err != nil {
		return nil, err
	}
	names := map[string]string{}
	for _, r := range doc.Body {
		bind, err := core.DecodeBind(r)
		issuer, ok := issuerOf(bind.Subject)
		if err != nil || !ok {
			continue
		}
		if prev, seen := names[string(bind.Name)]; seen && prev != issuer {
			return nil, fmt.Errorf("geul/resolve: Issuer list binds %q to %s and %s", bind.Name, prev, issuer)
		}
		names[string(bind.Name)] = issuer
	}
	return names, nil
}
