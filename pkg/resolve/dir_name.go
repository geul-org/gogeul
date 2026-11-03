//ff:func feature=resolve type=resolver control=selection
//ff:what Dir.Name — bind a name to a published document in its Issuer's document list (docs/); a name already bound elsewhere is refused
package resolve

import (
	"errors"
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// Name binds name to document doc of issuer in the Issuer's document list
// (`<Issuer>/docs/`, namesKey; GEUL-ORG 문서 이름): a BIND whose Subject is
// DOC-POINTER(issuer, doc, 0) (CODEC-POINTER for a codec) is appended. Names are only bound to published
// documents, one name to one document; an old name keeps pointing at its
// document and the last BIND is the main name. Binding a name the document
// already has writes nothing.
func (d Dir) Name(issuer, doc []core.Word, name string) error {
	if !validName(name) {
		return fmt.Errorf("geul/resolve: %q is not a document name (a-z, 0-9, '-')", name)
	}
	k := Key{Issuer: issuer, Doc: doc}.DocKey()
	lg, err := d.readList(k, k.Doc)
	switch {
	case err != nil:
		return err
	case len(lg.Entries) == 0:
		return errors.New("geul/resolve: a name is only bound to a published document")
	}
	head, body, names, err := d.namesOf(namesKey(k.Issuer))
	if err != nil {
		return err
	}
	if old, ok := names[name]; ok {
		if old == k {
			return nil
		}
		return fmt.Errorf("geul/resolve: name %q is bound to %s/%s", name, old.Issuer, old.Doc)
	}
	subj, _ := core.EncodeTargetPointer(issuer, doc, []core.Word{0, 0})
	b, _ := namesDoc(head, body, core.Bind{EdgeID: 1, Subject: subj, Name: []byte(name)})
	_, err = d.PublishList(issuer, core.DocIDDocs, b)
	return err
}
