package common

import (
	"os"
	"path/filepath"
)

type HostData struct {
	*AnsibleData
	*ServerspecData
}

type HostsData []HostData

type AnsibleData struct {
	Hostname string `yaml:"hostname"`
	I3env    string `yaml:"i3_env"`
	Ip_addr  string `yaml:"ip_addr"`
}

type ServerspecData struct {
	Cores string `yaml:":cpu"`
	Ram   string `yaml:":ram"`
	Hdd   string `yaml:":hdd"`
	Os    string `yaml:":os"`
	If1   string `yaml:":if1"`
	If2   string `yaml:":if2"`
	If3   string `yaml:":if3"`
}

func (HostsData) AnsibleSetData(dir string) {

	err := filepath.Walk(dir, readAnsiblehost)
	if err != nil {
		fmt.print("err")
	}

}

func readAnsiblehost(hostsdata HostsData) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		//hostdataに対するクロージャにする
	}
}
