//ff:func feature=inspect type=decoder control=sequence
//ff:what rawData — the Value a raw field's words are the Payload of, under its one CODEC
package inspect

import "github.com/geul-org/gogeul/pkg/core"

// rawData reads the words of a raw field as the Payload of its CODEC target
// (edge/CODEC: the TAG is left out, Kind·Codec from the CODEC, Length from
// the field width): the decoded Value, or nil without a single CODEC or when
// the target is a wildcard, not a Value, or does not decode.
func rawData(codecs []core.Word, ws []core.Word) any {
	if len(codecs) != 1 || codecs[0]&0xF != 0 || core.KindOf(codecs[0]) != core.KindValue {
		return nil
	}
	_, c := core.SplitTarget(codecs[0])
	r := core.Record{Kind: core.KindValue, Codec: c, Payload: ws}
	return asData(ParseValue(r))
}
