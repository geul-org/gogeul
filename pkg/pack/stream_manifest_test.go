//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamManifest — generation, objects, Issuers, NAMED, layers and the dictionary entry (none without one)
package pack

import "testing"

func TestStreamManifest(t *testing.T) {
	s := streamTestNew(t, nil)
	s.st.Objects = []streamObject{{Kind: kindData, Size: 9}}
	m := s.manifest([]namedEntry{{path: "a"}}, []layer{{count: 1}}, nil)
	if m.gen != 1 || m.created != 1_800_000_000 || len(m.objects) != 1 || m.objects[0].size != 9 || m.issuers[0] != "0x00000001" || len(m.named) != 1 || len(m.pLayers) != 1 || m.dicts != nil {
		t.Fatalf("%+v", m)
	}
	s.st.Dict = streamDict{ID: FirstDictID}
	if m := s.manifest(nil, nil, nil); len(m.dicts) != 1 || m.dicts[0].id != FirstDictID {
		t.Fatal("dictionary")
	}
}
