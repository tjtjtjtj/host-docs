package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
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

func (h *HostsData) AnsibleSetData(dir string) {

	err := filepath.Walk(dir, readAnsiblehost(h))
	if err != nil {
		fmt.Print("err")
	}
}

func readAnsiblehost(hostsdata *HostsData) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() {
			fmt.Printf("path:%s\n", path)
			fmt.Printf("file:%s\n", info.Name())

			buf, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}
			hostdata := new(HostData)
			fmt.Println(hostdata)
			err = yaml.Unmarshal(buf, hostdata)
			fmt.Println(hostdata)
			*hostsdata = append(*hostsdata, *hostdata)
		} else {
			fmt.Printf("dir:%s\n", info.Name())
		}
		return nil
	}
}
