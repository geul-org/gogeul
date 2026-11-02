//ff:type feature=resolve type=resolver
//ff:what overlay — a Store over the loose files beside it: the stored paths win, a path it lacks is read as a loose file
package resolve

// overlay makes a folder that keeps some Issuers in pack sets and others as
// files read as one tree. Writes go to stored; a path stored does not have is
// looked for, and removed, in loose.
type overlay struct {
	stored Store
	loose  dirStore
}
