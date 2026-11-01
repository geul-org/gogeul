//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamOptionsWithDefaults — block sizes as Write's, Zstd for no codec, the default run size, merge width and save interval; given values stay
package pack

import "testing"

func TestStreamOptionsWithDefaults(t *testing.T) {
	o := StreamOptions{}.withDefaults()
	if _, ok := o.Build.Codec.(*Zstd); !ok || o.RunEntries != DefaultRunEntries || o.MergeRuns != DefaultMergeRuns || o.saveLeaves != hSaveLeaves || o.Build.PLeaf != DefaultPLeaf || o.Build.PackSize != DefaultPackSize {
		t.Fatalf("%+v", o)
	}
	o = StreamOptions{Build: BuildOptions{Codec: Stored{}}, RunEntries: 5, MergeRuns: 3}.withDefaults()
	if _, ok := o.Build.Codec.(Stored); !ok || o.RunEntries != 5 || o.MergeRuns != 3 {
		t.Fatalf("%+v", o)
	}
}
