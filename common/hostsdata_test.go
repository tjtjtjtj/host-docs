package common

import (
	"testing"
)

func TestAnsibleSetData(t *testing.T) {

	h := make(HostsData, 0, 10)
	h.AnsibleSetData("./host_vars")
	if h[0].Hostname != "vm001.example.co.jp" {
		t.Fatal("ちゃんとセットされていない")
	}
}
