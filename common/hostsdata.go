package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// ansibleとserverspec host_varsのinclude
type HostData struct {
	AnsibleData    `yaml:",inline"`
	ServerspecData `yaml:",inline"`
}

// ansibleとserverspecの全host_vars格納
type HostsData []HostData

// ansible host_varsより、必要項目のみ格納
type AnsibleData struct {
	Hostname string `yaml:"hostname"`
	I3env    string `yaml:"i3_env"`
	Ipaddr   string `yaml:"ip_addr"`
}

// serverspec host_varsより、必要項目のみ格納
type ServerspecData struct {
	CPU string `yaml:":cpu"`
	RAM string `yaml:":ram"`
	Hdd string `yaml:":hdd"`
	Os  string `yaml:":os"`
	If1 string `yaml:":if1"`
	If2 string `yaml:":if2"`
	If3 string `yaml:":if3"`
}

// ansible host_varsのディレクトリより読み込みHostsDataにセット
func (hostsdata *HostsData) AnsibleSetData(dir string) {
	err := filepath.Walk(dir, readAnsiblehost(hostsdata))
	if err != nil {
		fmt.Print("err")
	}
}

func readAnsiblehost(hostsdata *HostsData) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() {
			buf, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			hostdata := new(HostData)
			err = yaml.Unmarshal(buf, hostdata)
			*hostsdata = append(*hostsdata, *hostdata)
		}
		return nil
	}
}

// HostsDataに格納されているIPに対して
// serverspec host_varsのディレクトリより読み込みHostsDataにセット
func (hostsdata HostsData) ServerspecSetData(dir string) {
	var serverspec_buf []byte
	var err error
Pro_RowLoop:
	//for i, _ := range hostsdata {
	for i := range hostsdata {
		serverspec_buf, err = ioutil.ReadFile(dir + hostsdata[i].Ipaddr + ".yml")
		if err != nil {
			continue Pro_RowLoop
		}
		err = yaml.Unmarshal(serverspec_buf, &hostsdata[i])
	}
}
