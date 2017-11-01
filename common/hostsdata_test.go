package common

import (
	"reflect"
	"testing"
)

func TestSetData(t *testing.T) {

	expected := HostsData{
		{AnsibleData: AnsibleData{Hostname: "vm001.example.co.jp", I3env: "production", Ipaddr: "127.0.0.1"},
			ServerspecData: ServerspecData{CPU: "8", RAM: "4", Hdd: "45", Os: "6.5", If1: "eth0", If2: "eth1", If3: "eth2"}},
		{AnsibleData: AnsibleData{Hostname: "vm002.example.co.jp", I3env: "production", Ipaddr: "127.0.0.2"},
			ServerspecData: ServerspecData{CPU: "4", RAM: "8", Hdd: "160", Os: "7.2", If1: "eth0", If2: "eth1", If3: "na"}},
		{AnsibleData: AnsibleData{Hostname: "vm003.example.co.jp", I3env: "staging", Ipaddr: "127.0.0.3"},
			ServerspecData: ServerspecData{CPU: "2", RAM: "4", Hdd: "45", Os: "6.5", If1: "eth0", If2: "na", If3: "na"}},
	}

	hostsdata := new(HostsData)
	hostsdata.AnsibleSetData("ansible/host_vars/")
	hostsdata.ServerspecSetData("serverspec/host_vars/")

	if !reflect.DeepEqual(*hostsdata, expected) {
		t.Fatalf("It is different from the expected result \n%v != \n%v", *hostsdata, expected)
	}
}
