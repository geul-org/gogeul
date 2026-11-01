//ff:func feature=pack type=writer control=sequence
//ff:what hPass.close — close the H pass's runs and head file
package pack

// close releases the pass's files (its writer is flushed by whoever saves).
func (p *hPass) close() {
	closeRuns(p.rs)
	p.heads.Close()
}
