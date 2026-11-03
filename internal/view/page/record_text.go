//ff:func feature=page type=helper control=selection
//ff:what recordText — one record as — (Absent), a slot marker name, #n/Qn/Pn, its decoded Value, or its codec name and payload hex
package page

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// recordText shows r: the Absent record of a left-out optional field as —,
// the other named slot markers by name (UNKNOWN-VALUE, NO-VALUE),
// a LOCAL or Wikidata Reference as #n, Qn or Pn
// (refText), a Value the library decodes (inspect.ParseValue: text, language
// text, URI, language tag, EDTF, numbers, dates …) by valueText, else its
// record name (core.RecordName: a Value ISSUER format by its format name;
// "codec 0x42" for an unnamed or retired codec) and payload hex
// (a lone TAG WORD, such as a codec BIND Subject, as the codec name only).
func recordText(r core.Record) string {
	if core.IsAbsent(r) {
		return "—"
	}
	if core.IsSlotMarker(r) && core.SlotMarkerName(r.Codec) != "" {
		return core.SlotMarkerName(r.Codec)
	}
	if s, ok := refText(r); ok {
		return s
	}
	v, ok := inspect.ParseValue(r)
	name := core.RecordName(r)
	if name == "" {
		name = fmt.Sprintf("codec 0x%02X", r.Codec)
	}
	switch {
	case r.Kind == core.KindValue && ok:
		return valueText(v)
	case len(r.Payload) == 0:
		return name
	default:
		return name + " " + core.FormatID(r.Payload)
	}
}
