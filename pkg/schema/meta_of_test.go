//ff:func feature=schema type=helper control=sequence
//ff:what metaOf — test fixture: a Meta record embedding one Value
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
)

func metaOf(codec uint16, v core.Record) core.Record {
	return core.Record{Kind: core.KindMeta, Codec: codec, Payload: enc(v)}
}
