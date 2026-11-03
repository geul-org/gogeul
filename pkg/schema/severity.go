//ff:type feature=schema type=model
//ff:what Severity — how serious a Fact is: Error breaks the spec, Warn may be a newer spec
package schema

// Severity grades a Fact. Error is a spec violation. Warn marks what may be
// valid under a newer spec;
// Options.Strict turns every Warn into Error.
type Severity int
