package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type HostData struct {
	AnsibleData    `yaml:",inline"`
	ServerspecData `yaml:",inline"`
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

func (hostsdata *HostsData) AnsibleSetData(dir string) {

	err := filepath.Walk(dir, readAnsiblehost(hostsdata))
	if err != nil {
		fmt.Print("err")
	}
}

func readAnsiblehost(hostsdata *HostsData) filepath.WalkFunc {
	var i = 1
	return func(path string, info os.FileInfo, err error) error {
		fmt.Printf("\nno:%v\n", i)
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
			fmt.Printf("hostsdata:%v", hostsdata)
		} else {
			fmt.Printf("dir:%s\n", info.Name())
		}
		i++
		return nil
	}
}

func (hostsdata HostsData) ServerspecSetData(dir string) {
	var serverspec_buf []byte
	var err error
Pro_RowLoop:
	for i, _ := range hostsdata {
		serverspec_buf, err = ioutil.ReadFile(dir + hostsdata[i].Ip_addr + ".yml")
		if err != nil {
			//fmt.Printf("not found %s \n", h.Hostname)
			continue Pro_RowLoop
		}
		fmt.Printf("\nsssssssssssss:%v", hostsdata[i])
		err = yaml.Unmarshal(serverspec_buf, &hostsdata[i])
		fmt.Printf("\nsssssssssssss:%v", hostsdata[i])
	}

	fmt.Printf("\n\n\nss:%v", hostsdata)

}
