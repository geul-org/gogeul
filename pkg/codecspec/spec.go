//ff:type feature=codecspec type=model
//ff:what Spec — field lists keyed by the declared codec's target word
package codecspec

import "github.com/geul-org/gogeul/pkg/core"

// Spec maps a direct codec's target word (core.TargetWord, each code its own:
// a 5Q Reference keeps its Quantifier) to its declared fields in payload order. An Edge's leading
// Edge Id is implicit and not listed.
type Spec map[core.Word][]Field
