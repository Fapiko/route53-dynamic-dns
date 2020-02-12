package ip

import (
	"io/ioutil"
	"net/http"
	"strings"
)

/*
External returns the public IP address for this machine found by hitting the myexternalip service
*/
func External() (string, error) {
	resp, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(ip)), nil
}
