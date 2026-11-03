//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what bindNames — E12: within a document one BIND name is bound to one Subject only
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// bindNames reports each BIND that binds a name (compared byte for byte) to
// another Subject than an earlier BIND of the same name in the document
// (grammar 0x0A3L "한 문서에서 같은 이름을 서로 다른 Subject에 묶지 않는다",
// edge.md). Binding the same Subject again, or several names to one Subject,
// is allowed. BINDs that do not decode are left to the layer that reads them.
func bindNames(all []rec) []Fact {
	subjects := map[string]string{}
	var fs []Fact
	for _, r := range all {
		b, err := core.DecodeBind(r.R)
		s := fmt.Sprintf("%d %d %s %s", b.Subject.Kind, b.Subject.Codec, core.FormatID(b.Subject.Issuer), core.FormatID(b.Subject.Payload))
		prev, seen := subjects[string(b.Name)]
		if err == nil && seen && prev != s {
			fs = append(fs, factAt("E12", Error, r, "이름", "이름 하나에 Subject 하나", fmt.Sprintf("%q 다른 Subject에 다시 묶음", b.Name)))
		}
		if err == nil && !seen {
			subjects[string(b.Name)] = s
		}
	}
	return fs
}
