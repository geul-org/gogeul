//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestStreamMergeWidth — MergeRuns with no MergeWidth; what MergeWidth says kept between 2 and MergeRuns
package pack

import "testing"

func TestStreamMergeWidth(t *testing.T) {
	s := &Stream{o: StreamOptions{MergeRuns: 10}}
	if got := s.mergeWidth(); got != 10 {
		t.Fatalf("no hook: %d", got)
	}
	want := map[int]int{1: 2, 5: 5, 99: 10}
	for in, out := range want {
		s.o.MergeWidth = func() int { return in }
		if got := s.mergeWidth(); got != out {
			t.Fatalf("%d: %d", in, got)
		}
	}
}
