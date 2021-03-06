package subdomain

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func methodHackerTarget(domain string) ([]string, error) {
	// Compose the URL
	url := fmt.Sprintf("https://api.hackertarget.com/hostsearch/?q=%s", domain)
	// Request the data
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New("HackerTarget is not available")
	}
	// Grab the content
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("HackerTarget does not respond correctly")
	}
	// Parse the Response
	subdomains := make([]string, 0)
	sc := bufio.NewScanner(bytes.NewReader(body))
	for sc.Scan() {
		splitter := strings.SplitN(sc.Text(), ",", 2)
		subdomains = append(subdomains, splitter[0])
	}
	return subdomains, nil
}
