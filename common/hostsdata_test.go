package common

import (
	"fmt"
	"testing"
)

func TestSetData(t *testing.T) {

	hostsdata := new(HostsData)
	hostsdata.AnsibleSetData("./ansible/host_vars/")
	for _, v := range *hostsdata {
		fmt.Printf("\n\ntttttt:%v", v.Hostname)
		fmt.Printf("\n\ntttttt:%v", v.I3env)
		fmt.Printf("\n\ntttttt:%v", v.Ip_addr)
	}

	hostsdata.ServerspecSetData("./serverspec/host_vars/")
	for _, v := range *hostsdata {
		fmt.Printf("\n\ntttttt:%v", v.Cores)
		fmt.Printf("\n\ntttttt:%v", v.Ram)
		fmt.Printf("\n\ntttttt:%v", v.If1)
	}

	t.Fatal("errrrrrrrrrrrrrr")

}
