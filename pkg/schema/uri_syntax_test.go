//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestURISyntax — RFC 3986 URIs (authority, IPv6 and IPvFuture literals, userinfo, port, query, fragment, rootless and empty paths, percent-encoding) match; no scheme, bad scheme, bad percent-encoding, spaces, brackets outside the host, a second "#" and empty do not
package schema

import "testing"

func TestURISyntax(t *testing.T) {
	good := []string{
		"https://geul.org/.well-known/geul/0x00000000/codecs/0x00000AA0/latest.geul",
		"http://user:pw@example.com:8080/a/b?c=d&e=f#g",
		"http://[2001:db8::7]/c", "http://[::1]:80", "ldap://[v7.fe80::a+en1]/x", "http://192.0.2.16/",
		"urn:oasis:names:specification:docbook:dtd:xml:4.1.2", "mailto:John.Doe@example.com", "news:comp.infosystems.www.servers.unix",
		"tel:+1-816-555-1212", "file:///etc/hosts", "a:", "x:/", "https://ko.wikipedia.org/wiki/%ED%95%9C", "foo://h/?/?#/?",
	}
	for _, s := range good {
		if !uriSyntax([]byte(s)) {
			t.Fatalf("good %q", s)
		}
	}
	bad := []string{"", "ab", "/relative/path", "1http://x", "http://x/%G1", "http://x/%2", "http://x/a b", "http://x/{y}", "http://x/a#b#c", "http://[::g]/", "http://x/\"q\"", "http://x/a|b", ":x"}
	for _, s := range bad {
		if uriSyntax([]byte(s)) {
			t.Fatalf("bad %q", s)
		}
	}
}
