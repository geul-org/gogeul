//ff:func feature=core type=helper control=iteration dimension=1
//ff:what headEdgeFixture — an Edge record of a codec from its Edge Id and LOCAL/raw fields, for the HEAD placement tests
package core

// headEdgeFixture encodes codec(Edge Id id, fields…): each field is a
// record (a LOCAL Reference for an ID).
func headEdgeFixture(codec uint16, id uint64, fields ...Record) Record {
	ws, _ := EncodeRecord(EncodeLocalRef(id))
	for _, f := range fields {
		fw, _ := EncodeRecord(f)
		ws = append(ws, fw...)
	}
	return Record{Kind: KindEdge, Codec: codec, Payload: ws}
}
