//ff:func feature=inspect type=decoder control=selection
//ff:what genericParts — a payload without a typed layout: a leading embedded record when the spec fixes one, then embedded records or the rest
package inspect

import "github.com/geul-org/gogeul/pkg/core"

// genericParts splits a payload whose fields are not known. An Edge payload
// opens with its LOCAL Edge Id (GEUL t002 OVERVIEW Edge) and a Meta Value
// payload is one embedded Value, so a leading record that decodes as such is
// its own part. After that, an Edge payload that is wholly a run of embedded
// Reference/Value records is shown record by record, so every embedded TAG WORD
// reads as one; anything else is one "payload" part.
func genericParts(e Entry) []Part {
	w := &walker{ws: e.Record.Payload}
	r, _, err := core.DecodeRecord(w.ws)
	switch {
	case err == nil && e.Record.Kind == core.KindEdge && r.Kind == core.KindReference && r.Codec == core.CodecLocal:
		w.rec("ID")
	case err == nil && e.Record.Kind == core.KindMeta && r.Kind == core.KindValue:
		w.rec("value")
	}
	if e.Record.Kind == core.KindEdge && recordRun(w.ws[w.pos:]) {
		recordParts(w)
	}
	if w.pos < len(w.ws) {
		w.words("payload", len(w.ws)-w.pos)
	}
	return w.out
}
