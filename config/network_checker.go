package main

import (
	"io/ioutil"
	"net/http"
)

type NetworkChecker struct {
	connected bool
	err       error
	t         *template
	done      bool
}

func WithNetworkChecker() *NetworkChecker {
	return &NetworkChecker{
		t: NewTemplate("Network"),
	}
}

func (n *NetworkChecker) Check() {
	resp, err := http.Get("https://baidu.com")
	if err != nil {
		n.err = err
		n.t.Error(err)
	} else {
		if resp.StatusCode == http.StatusOK {
			n.connected = true
			n.t.Ok("network is ok")
		} else {
			defer resp.Body.Close()
			d, _ := ioutil.ReadAll(resp.Body)
			n.t.Warn(string(d))
		}
	}
	n.done = true
}

func (n *NetworkChecker) IsDone() bool {
	return n.done
}

func (n *NetworkChecker) Template() *template {
	return n.t
}
