//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestBCP47WellFormed — RFC 5646 appendix A examples and grandfathered tags are well-formed in any case; empty, underscores, overlong subtags, a lone singleton and misplaced subtags are not
package schema

import "testing"

func TestBCP47WellFormed(t *testing.T) {
	good := []string{
		"de", "fr", "ja", "i-enochian", "zh-Hant", "zh-Hans", "sr-Cyrl", "sr-Latn", "zh-cmn-Hans-CN", "cmn-Hans-CN", "zh-yue-HK", "yue-HK",
		"sr-Latn-RS", "sl-rozaj", "sl-rozaj-biske", "sl-nedis", "de-CH-1901", "sl-IT-nedis", "hy-Latn-IT-arevela", "de-DE", "en-US", "es-419",
		"de-CH-x-phonebk", "az-Arab-x-AZE-derbend", "x-whatever", "qaa-Qaaa-QM-x-southern", "de-Qaaa", "sr-Latn-QM", "en-US-u-islamcal",
		"zh-CN-a-myext-x-private", "en-a-myext-b-another", "EN-gb-OED", "zh-min-nan", "be-tarask", "be-x-old", "simple", "und", "zxx",
	}
	for _, s := range good {
		if !bcp47WellFormed([]byte(s)) {
			t.Fatalf("good %q", s)
		}
	}
	bad := []string{"", "en_US", "de-419-DE", "a-DE", "ar-a-aaa-b-bbb-a", "zh-classical", "en-", "-en", "en--US", "x", "en-x", "en US", "abcdefghi"}
	for _, s := range bad {
		if bcp47WellFormed([]byte(s)) {
			t.Fatalf("bad %q", s)
		}
	}
}
