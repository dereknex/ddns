package solver

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// use https://ifconfig.co/

//IfconfigCo is a public IP solver with ifconfig.co
type IfconfigCo struct {
}

// Run to get public ip
func (s *IfconfigCo) Run() (string, error) {
	resp, err := http.Get("http://ifconfig.co")
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(data), "\n"), nil

}
