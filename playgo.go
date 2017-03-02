package playgo

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var playgroundHost = "https://play.golang.org"

// ShareAndOpen func
func ShareAndOpen() (string, error) {
	flag.Parse()
	path := flag.Arg(0)

	url, shareErr := Share(path)
	if shareErr != nil {
		return "", shareErr
	}

	clipboard.WriteAll(url)
	openErr := open.Start(url)
	if openErr != nil {
		return "", openErr
	}

	return url, nil
}

// Share func
func Share(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	ext := filepath.Ext(path)
	if ext != ".go" {
		return "", fmt.Errorf("File %s is not a .go file", path)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}
	if fileInfo.Size() == 0 {
		return "", fmt.Errorf("File %s is empty", path)
	}

	req, err := http.NewRequest("POST", playgroundHost+"/share", file)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "playgo/1.0")

	c := new(http.Client)
	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}

	respBody, respErr := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if respErr != nil {
		return "", respErr
	}

	return fmt.Sprintf("%s/p/%s", playgroundHost, string(respBody)), nil
}
