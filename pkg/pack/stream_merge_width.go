//ff:func feature=pack type=writer control=sequence
//ff:what Stream.mergeWidth — how many runs the next merge reads at once: MergeRuns, or what MergeWidth says, between 2 and MergeRuns
package pack

// mergeWidth asks StreamOptions.MergeWidth, when set, and keeps the answer
// between 2 and MergeRuns.
func (s *Stream) mergeWidth() int {
	if s.o.MergeWidth == nil {
		return s.o.MergeRuns
	}
	return min(s.o.MergeRuns, max(2, s.o.MergeWidth()))
}
