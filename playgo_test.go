package playgo

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	emptyFileName = ".empty.go"

	shares = []struct {
		path string
		ok   bool
	}{
		{"", false},
		{"iujo7y9n86gn6gjiwef", false},
		{"README.md", false},
		{emptyFileName, false},
		{"cmd/playgo/main.go", true},
	}
)

var mockServer *httptest.Server

func init() {
	// create test empty file
	os.OpenFile(emptyFileName, os.O_CREATE, 0755)

	mux := http.NewServeMux()
	mux.HandleFunc("/share", func(w http.ResponseWriter, req *http.Request) {})

	mockServer = httptest.NewServer(mux)
	playgroundHost = mockServer.URL
}

func TestShare(t *testing.T) {
	defer os.Remove(emptyFileName)

	for _, s := range shares {
		url, err := Share(s.path)
		if (err == nil) != s.ok {
			t.Errorf("Share(%s) expected ok=%t, got error=%v", s.path, s.ok, err)
		}

		if s.ok && url == "" {
			t.Errorf("Share(%s) expected non-empty url", s.path)
		}

		if s.ok && url[:4] != "http" {
			t.Errorf("Share(%s) expected valid url, got url %s", s.path, url)
		}
	}
}
