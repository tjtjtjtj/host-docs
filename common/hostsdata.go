package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// HostData ansibleとserverspec host_varsのinclude
type HostData struct {
	AnsibleData    `yaml:",inline"`
	ServerspecData `yaml:",inline"`
}

// HostsData ansibleとserverspecの全host_vars格納
type HostsData []HostData

// AnsibleData ansible host_varsより、必要項目のみ格納
type AnsibleData struct {
	Hostname string `yaml:"hostname"`
	I3env    string `yaml:"i3_env"`
	Ipaddr   string `yaml:"ip_addr"`
}

// ServerspecData serverspec host_varsより、必要項目のみ格納
type ServerspecData struct {
	CPU string `yaml:":cpu"`
	RAM string `yaml:":ram"`
	Hdd string `yaml:":hdd"`
	Os  string `yaml:":os"`
	If1 string `yaml:":if1"`
	If2 string `yaml:":if2"`
	If3 string `yaml:":if3"`
}

// AnsibleSetData ansible host_varsのディレクトリより読み込みHostsDataにセット
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

// ServerspecSetData HostsDataに格納されているIPに対して
// serverspec host_varsのディレクトリより読み込みHostsDataにセット
func (hostsdata HostsData) ServerspecSetData(dir string) {
	var serverspecbuf []byte
	var err error
Pro_RowLoop:
	//for i, _ := range hostsdata {
	for i := range hostsdata {
		serverspecbuf, err = ioutil.ReadFile(dir + hostsdata[i].Ipaddr + ".yml")
		if err != nil {
			continue Pro_RowLoop
		}
		err = yaml.Unmarshal(serverspecbuf, &hostsdata[i])
	}
}
