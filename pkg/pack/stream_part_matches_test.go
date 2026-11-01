//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestStreamPartMatches — a part matches by saved hash, a hash-less index by its P leaves, a hash-less sealed data pack by its inventory, a hash-less data part as it is; a missing or short file does not; an unreadable one is an error
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStreamPartMatches(t *testing.T) {
	s := streamTestNew(t, nil)
	streamTestPut(t, s, 1, 6)
	s.Checkpoint(nil)
	data, index := s.data.Name(), s.index.Name()
	withSum := func(o streamOpen, h []byte) streamOpen { o.Sum = h; return o }
	cases := map[string]struct {
		path   string
		kind   byte
		open   streamOpen
		sealed bool
		want   bool
	}{
		"hash":          {data, kindData, s.st.Data, false, true},
		"other hash":    {index, kindIndex, withSum(s.st.Index, s.st.Data.Sum), false, false},
		"legacy index":  {index, kindIndex, withSum(s.st.Index, nil), false, true},
		"legacy part":   {data, kindData, withSum(s.st.Data, nil), false, true},
		"legacy sealed": {data, kindData, withSum(s.st.Data, nil), true, false},
		"missing":       {filepath.Join(s.out, "none"), kindData, s.st.Data, false, false},
		"short":         {index, kindIndex, streamOpen{Num: 1, Off: 1 << 40}, false, false},
	}
	for name, c := range cases {
		if ok, err := s.partMatches(c.path, c.kind, c.open, c.sealed); ok != c.want || err != nil {
			t.Errorf("%s: %v %v", name, ok, err)
		}
	}
	locked := filepath.Join(s.out, "locked")
	os.WriteFile(locked, nil, 0)
	if _, err := s.partMatches(locked, kindData, s.st.Data, false); err == nil {
		t.Fatal("unreadable: want error")
	}
	if _, err := s.partMatches(data, kindData, withSum(s.st.Data, []byte("bad")), false); err == nil {
		t.Fatal("bad hash state: want error")
	}
}
