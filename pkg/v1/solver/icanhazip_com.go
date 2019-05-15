package solver

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// use curl icanhazip.com

//IcanhazipCom is a public IP solver with icanhazip.com
type IcanhazipCom struct {
}

// Run to get public ip
func (s *IcanhazipCom) Run() (string, error) {
	resp, err := http.Get("http://icanhazip.com")
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return strings.Trim(string(data), "\n"), nil

}
