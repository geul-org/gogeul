//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what specOfDoc — the codec specs a spec document's BODY declares
package resolve

import (
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// specOfDoc collects the BODY of a codec spec document.
func specOfDoc(doc inspect.Document) codecspec.Spec {
	body := make([]core.Record, len(doc.Body))
	for i, e := range doc.Body {
		body[i] = e.Record
	}
	return codecspec.Collect(body)
}
