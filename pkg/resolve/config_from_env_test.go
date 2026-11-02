//ff:func feature=resolve type=decoder control=sequence
//ff:what TestConfigFromEnv — default geul.org,registry; custom file:// list with private patterns; bad values fail
package resolve

import (
	"testing"
)

func TestConfigFromEnv(t *testing.T) {
	env := func(m map[string]string) func(string) string { return func(k string) string { return m[k] } }
	c, err := ConfigFromEnv(env(nil))
	if err != nil || len(c.Proxy) != 2 || c.Proxy[0].Name != "https://geul.org" || !c.Proxy[0].Public || c.Proxy[1].Name != "registry" || c.Private != nil {
		t.Fatalf("default: %+v %v", c, err)
	}
	c, err = ConfigFromEnv(env(map[string]string{"GEULPROXY": "file:///srv,off", "GEULPRIVATE": "0x0001"}))
	if err != nil || c.Proxy[0].Source.(Dir).Root != "/srv" || !c.Proxy[1].Off || len(c.Private) != 1 {
		t.Fatalf("custom: %+v %v", c, err)
	}
	if _, err := ConfigFromEnv(env(map[string]string{"GEULPROXY": "http://x"})); err == nil {
		t.Fatal("bad proxy: want error")
	}
	if _, err := ConfigFromEnv(env(map[string]string{"GEULPRIVATE": "bad"})); err == nil {
		t.Fatal("bad private: want error")
	}
}
