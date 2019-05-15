package solver

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// Use curl ident.me

//IdentMe is a public IP solver with Ident.me
type IdentMe struct {
}

// Run to get public ip
func (s *IdentMe) Run() (string, error) {
	resp, err := http.Get("http://ident.me")
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(data), "\n"), nil

}
