package ip

import (
	"io/ioutil"
	"net/http"
)

/*
External returns the public IP address for this machine found by hitting the myexternalip service
*/
func External() (string, error) {
	resp, err := http.Get("http://ipv4.myexternalip.com/raw")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}
