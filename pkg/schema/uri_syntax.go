//ff:func feature=schema type=rule control=sequence
//ff:what uriSyntax — whether bytes match RFC 3986 `URI = scheme ":" hier-part [ "?" query ] [ "#" fragment ]` (scheme and ":" required, percent-encoding as pct-encoded)
package schema

import "regexp"

// The RFC 3986 ABNF (section 3 and appendix A) as a regular expression. An
// IPv4 address is also a reg-name, so host is IP-literal or reg-name.
const (
	uriH16    = `[0-9A-Fa-f]{1,4}`
	uriDec    = `(?:25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])`
	uriIPv4   = uriDec + `\.` + uriDec + `\.` + uriDec + `\.` + uriDec
	uriLs32   = `(?:` + uriH16 + `:` + uriH16 + `|` + uriIPv4 + `)`
	uriIPv6   = `(?:(?:` + uriH16 + `:){6}` + uriLs32 + `|::(?:` + uriH16 + `:){5}` + uriLs32 + `|(?:` + uriH16 + `)?::(?:` + uriH16 + `:){4}` + uriLs32 + `|(?:(?:` + uriH16 + `:){0,1}` + uriH16 + `)?::(?:` + uriH16 + `:){3}` + uriLs32 + `|(?:(?:` + uriH16 + `:){0,2}` + uriH16 + `)?::(?:` + uriH16 + `:){2}` + uriLs32 + `|(?:(?:` + uriH16 + `:){0,3}` + uriH16 + `)?::` + uriH16 + `:` + uriLs32 + `|(?:(?:` + uriH16 + `:){0,4}` + uriH16 + `)?::` + uriLs32 + `|(?:(?:` + uriH16 + `:){0,5}` + uriH16 + `)?::` + uriH16 + `|(?:(?:` + uriH16 + `:){0,6}` + uriH16 + `)?::)`
	uriPct    = `%[0-9A-Fa-f]{2}`
	uriUnSub  = `A-Za-z0-9\-._~!$&'()*+,;=`
	uriPchar  = `(?:[` + uriUnSub + `:@]|` + uriPct + `)`
	uriSeg    = uriPchar + `*`
	uriSegNz  = uriPchar + `+`
	uriHost   = `(?:\[(?:` + uriIPv6 + `|v[0-9A-Fa-f]+\.[` + uriUnSub + `:]+)\]|(?:[` + uriUnSub + `]|` + uriPct + `)*)`
	uriAuth   = `(?:(?:[` + uriUnSub + `:]|` + uriPct + `)*@)?` + uriHost + `(?::[0-9]*)?`
	uriHier   = `(?://` + uriAuth + `(?:/` + uriSeg + `)*|/(?:` + uriSegNz + `(?:/` + uriSeg + `)*)?|` + uriSegNz + `(?:/` + uriSeg + `)*|)`
	uriQuery  = `(?:` + uriPchar + `|[/?])*`
	uriSchema = `[A-Za-z][A-Za-z0-9+\-.]*`
)

var uriRE = regexp.MustCompile(`^` + uriSchema + `:` + uriHier + `(?:\?` + uriQuery + `)?(?:#` + uriQuery + `)?$`)

// uriSyntax reports whether b is a URI by the RFC 3986 grammar the URI Value
// requires (grammar 0x074L "해석과 유효성"): a scheme and ":", then hier-part,
// query and fragment of allowed ASCII characters and pct-encoded octets.
// A relative reference is no URI (URI-RELATIVE-REF is another codec).
func uriSyntax(b []byte) bool {
	return uriRE.Match(b)
}
