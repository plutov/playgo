package playgo

import (
	"testing"
)

var shares = []struct {
	path string
	ok   bool
}{
	{"", false},
	{"iujo7y9n86gn6gjiwef", false},
	{"README.md", false},
	{".empty", false},
	{"cmd/playgo/main.go", true},
}

func TestShare(t *testing.T) {
	for _, s := range shares {
		url, err := Share(s.path)
		if (err == nil) != s.ok {
			t.Errorf("Share(%s) expected ok=%t, got error=%v", s.path, s.ok, err)
		}

		if s.ok && url == "" {
			t.Errorf("Share(%s) expected non-empty url", s.path)
		}

		if s.ok && url[:5] != "https" {
			t.Errorf("Share(%s) expected valid url, got url %s", s.path, url)
		}
	}
}
