//ff:func feature=cli type=helper control=sequence
//ff:what harness.languageSpec — test fixture: publish the LANGUAGE-SHORT-CODE codec spec, whose `code` field's one ENUM names 0x0784 en
package cli

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/resolve"
)

// languageSpec publishes the GEUL Language Registry's codec, 0x000007a0.
func (h *harness) languageSpec() {
	target := core.TargetWord(core.KindValue, core.CodecLanguageShortCode)
	id, _ := core.EncodeDocIDCodec([]core.Word{0, 0}, uint32(target))
	field, _ := core.EncodeBind(core.Bind{EdgeID: 3, Subject: core.EncodeLocalRef(2), Name: []byte("code")})
	enum, _ := core.EncodeBind(core.Bind{EdgeID: 5, Subject: core.EncodeLocalRef(4), Name: []byte("en")})
	body := []core.Record{
		core.EncodeDeclare(core.Declare{EdgeID: 1, Target: target}),
		core.EncodeFieldDecl(core.FieldDecl{EdgeID: 2, Declare: 1, Width: 1}),
		field,
		core.EncodeEnumDecl(core.EnumDecl{EdgeID: 4, Field: 2, Code: []core.Word{0x0784}}),
		enum,
	}
	doc, _ := core.EncodeDocument([]core.Record{core.EncodeSpecVersion(core.SpecVersion), id}, body)
	if _, err := (resolve.Dir{Root: h.data}).PublishCodec(doc); err != nil {
		h.t.Fatal(err)
	}
}
