//ff:func feature=pack type=writer control=sequence
//ff:what StreamOptions.withDefaults — block and pack sizes as for Write, Zstd when no codec, the default run size and merge width
package pack

// withDefaults fills o.Build as Write does, makes a nil codec Zstd and a run
// size below 1 DefaultRunEntries, a merge width below 2 DefaultMergeRuns and
// a save interval below 1 hSaveLeaves.
func (o StreamOptions) withDefaults() StreamOptions {
	o.Build = o.Build.withDefaults()
	if o.Build.Codec == nil {
		o.Build.Codec = &Zstd{}
	}
	o.RunEntries = defaultInt(o.RunEntries, 1, DefaultRunEntries)
	o.MergeRuns = defaultInt(o.MergeRuns, 2, DefaultMergeRuns)
	o.saveLeaves = defaultInt(o.saveLeaves, 1, hSaveLeaves)
	return o
}
