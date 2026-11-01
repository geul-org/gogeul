//ff:func feature=cli type=helper control=selection
//ff:what app.findDoc — a document's bytes from a file path, else from the data folder by its 공표 주소
package cli

import "os"

// findDoc reads arg as a file when one exists there, else finds it as a 공표
// 주소 (resolve.Dir.Find).
func (a *app) findDoc(arg string) ([]byte, error) {
	if fi, err := os.Stat(arg); err == nil && !fi.IsDir() {
		return os.ReadFile(arg)
	}
	s, err := a.localIssuer(arg)
	if err != nil {
		return nil, err
	}
	d, err := a.dir()
	if err != nil {
		return nil, err
	}
	return d.Find(s)
}
