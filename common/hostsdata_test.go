package common

import (
	"reflect"
	"testing"
)

func TestSetData(t *testing.T) {

	expected := HostsData{
		{AnsibleData: AnsibleData{Hostname: "vm001.example.co.jp", I3env: "production", Ip_addr: "127.0.0.1"},
			ServerspecData: ServerspecData{Cpu: "8", Ram: "4", Hdd: "45", Os: "6.5", If1: "eth0", If2: "eth1", If3: "eth2"}},
		{AnsibleData: AnsibleData{Hostname: "vm002.example.co.jp", I3env: "production", Ip_addr: "127.0.0.2"},
			ServerspecData: ServerspecData{Cpu: "4", Ram: "8", Hdd: "160", Os: "7.2", If1: "eth0", If2: "eth1", If3: "na"}},
		{AnsibleData: AnsibleData{Hostname: "vm003.example.co.jp", I3env: "staging", Ip_addr: "127.0.0.3"},
			ServerspecData: ServerspecData{Cpu: "2", Ram: "4", Hdd: "45", Os: "6.5", If1: "eth0", If2: "na", If3: "na"}},
	}

	hostsdata := new(HostsData)
	hostsdata.AnsibleSetData("ansible/host_vars/")
	hostsdata.ServerspecSetData("serverspec/host_vars/")

	if !reflect.DeepEqual(*hostsdata, expected) {
		t.Fatalf("It is different from the expected result \n%v != \n%v", *hostsdata, expected)
	}
}
