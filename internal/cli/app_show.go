//ff:func feature=cli type=command control=iteration dimension=1
//ff:what app.show — print one document's HEAD lines, BODY, and BODY lines, with only the codec specs it uses
package cli

import "fmt"

// show finds the one document (findDoc) and prints its lines, laid out by the
// specs of the codecs it uses (docSpec), or them as JSON.
func (a *app) show(arg string) error {
	b, err := a.findDoc(arg)
	if err != nil {
		return err
	}
	ls, err := linesWith(b, a.docSpec(b))
	if err != nil {
		return err
	}
	if done, err := a.json(ls); done {
		return err
	}
	for _, l := range ls.Head {
		fmt.Fprintln(a.out, l)
	}
	fmt.Fprintln(a.out, "BODY")
	for _, l := range ls.Body {
		fmt.Fprintln(a.out, l)
	}
	return nil
}
