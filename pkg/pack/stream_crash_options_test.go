//ff:func feature=pack type=writer control=sequence
//ff:what streamCrashOptions — the stream test options with merges two runs at a time, a save every two H leaves and one level-3 encoder, so every Close step happens several times
package pack

// streamCrashOptions are streamTestOptions with MergeRuns 2 and saveLeaves 2
// and one Zstd encoder at level 3 (fast; the steps are the same).
func streamCrashOptions() StreamOptions {
	o := streamTestOptions()
	o.MergeRuns, o.saveLeaves, o.Build.Codec = 2, 2, &Zstd{Level: 3, Workers: 1}
	return o
}
