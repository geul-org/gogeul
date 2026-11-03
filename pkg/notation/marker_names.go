//ff:type feature=notation type=model
//ff:what markerNames — the token of each slot marker WORD (the inverse of markerTokens)
package notation

import "github.com/geul-org/gogeul/pkg/core"

// markerNames maps each named slot marker WORD (reference/slot-marker) to its
// token; markerTokens is the other direction.
var markerNames = map[core.Word]string{core.Absent: "ABSENT", core.UnknownValue: "UNKNOWN", core.NoValue: "NONE"}
