//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestStreamUnknownObjects — objects of the kind the state has not closed are listed in name order, closed ones and other kinds not; no o/ is an error
package pack

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestStreamUnknownObjects(t *testing.T) {
	s := streamTestNew(t, nil)
	streamTestPut(t, s, 1, 6)
	known := objectName(object{kind: s.st.Objects[0].Kind, hash: s.st.Objects[0].Hash})
	for _, n := range []string{"b.gpk", "a.gpk", "c.gpx"} {
		os.WriteFile(filepath.Join(s.out, "o", n), nil, 0o644)
	}
	got, err := s.unknownObjects(kindData)
	if err != nil || !slices.Equal(got, []string{"a.gpk", "b.gpk"}) || filepath.Ext(known) != ".gpk" {
		t.Fatalf("%v %v", got, err)
	}
	if got, _ := s.unknownObjects(kindIndex); !slices.Equal(got, []string{"c.gpx"}) {
		t.Fatalf("index %v", got)
	}
	os.RemoveAll(filepath.Join(s.out, "o"))
	if _, err := s.unknownObjects(kindData); err == nil {
		t.Fatal("no o/: want error")
	}
}
