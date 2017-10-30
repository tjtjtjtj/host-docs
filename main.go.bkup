package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

type Data struct {
	Hostname string `yaml:"hostname"`
	I3env    string `yaml:"i3_env"`
	Ip_addr  string `yaml:"ip_addr"`
	Cores    string `yaml:":cpu"`
	Ram      string `yaml:":ram"`
	Hdd      string `yaml:":hdd"`
	Os       string `yaml:":os"`
	If1      string `yaml:":if1"`
	If2      string `yaml:":if2"`
	If3      string `yaml:":if3"`
	//ここ構造体のネスト
}

type Serverspec_Data struct {
	Cores string `yaml:":cpu"`
	Ram   string `yaml:":ram"`
	Hdd   string `yaml:":hdd"`
	Os    string `yaml:":os"`
	If1   string `yaml:":if1"`
	If2   string `yaml:":if2"`
	If3   string `yaml:":if3"`
}

func main() {
	fmt.Println("vim-go")
	dirwalk("/home/jenkins/ansible/host_vars")
}

func dirwalk(dir string) {
	staging_file, err := os.Create("./staging.md")
	production_file, err := os.Create("./production.md")
	if err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	production_data := make([]Data, 0, len(files))
	staging_data := make([]Data, 0, len(files))
	var d Data

	for _, file := range files {

		if file.Mode().IsRegular() {
			fmt.Printf("file:%s\n", file.Name())

			buf, err := ioutil.ReadFile(dir + "/" + file.Name())
			if err != nil {
				panic(err)
			}

			err = yaml.Unmarshal(buf, &d)
			fmt.Println(d)

			switch d.I3env {
			case "production":
				production_data = append(production_data, d)
			case "staging":
				staging_data = append(staging_data, d)
			}

		} else {
			fmt.Printf("dir:%s\n", file.Name())
		}
	}

	var serverspec_buf []byte
	var s_d Serverspec_Data
Pro_RowLoop:
	for k, host := range production_data {
		serverspec_buf, err = ioutil.ReadFile("/home/jenkins/serverspec/host_vars/" + host.Ip_addr + ".yml")
		if err != nil {
			fmt.Printf("not found %s \n", host.Hostname)
			production_data[k].Cores = "-"
			production_data[k].Ram = "-"
			production_data[k].Hdd = "-"
			production_data[k].Os = "-"
			production_data[k].If1 = "-"
			production_data[k].If2 = "-"
			production_data[k].If3 = "-"
			continue Pro_RowLoop
		}
		//fmt.Printf("serverspec: %v", string(serverspec_buf))
		err = yaml.Unmarshal(serverspec_buf, &s_d)
		//fmt.Printf("serverspec: %v\n", s_d)
		production_data[k].Cores = s_d.Cores
		production_data[k].Ram = s_d.Ram
		production_data[k].Hdd = s_d.Hdd
		production_data[k].Os = s_d.Os
		production_data[k].If1 = s_d.If1
		production_data[k].If2 = s_d.If2
		production_data[k].If3 = s_d.If3
		//fmt.Printf("hostcores: %s\n", host.Cores)
	}

Stg_RowLoop:
	for k, host := range staging_data {
		serverspec_buf, err = ioutil.ReadFile("/home/jenkins/serverspec/host_vars/" + host.Ip_addr + ".yml")
		if err != nil {
			fmt.Printf("not found %s \n", host.Hostname)
			staging_data[k].Cores = "-"
			staging_data[k].Ram = "-"
			staging_data[k].Hdd = "-"
			staging_data[k].Os = "-"
			staging_data[k].If1 = "-"
			staging_data[k].If2 = "-"
			staging_data[k].If3 = "-"
			continue Stg_RowLoop
		}
		//fmt.Printf("serverspec: %v", string(serverspec_buf))
		err = yaml.Unmarshal(serverspec_buf, &s_d)
		//fmt.Printf("serverspec: %v\n", s_d)
		staging_data[k].Cores = s_d.Cores
		staging_data[k].Cores = s_d.Cores
		staging_data[k].Ram = s_d.Ram
		staging_data[k].Hdd = s_d.Hdd
		staging_data[k].Os = s_d.Os
		staging_data[k].If1 = s_d.If1
		staging_data[k].If2 = s_d.If2
		staging_data[k].If3 = s_d.If3
		//fmt.Printf("hostcores: %s\n", host.Cores)
	}

	tmpl := template.Must(template.ParseFiles("serverlist.tmpl"))

	fmt.Printf("%v", production_data)

	err = tmpl.Execute(production_file, production_data)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(staging_file, staging_data)
	if err != nil {
		panic(err)
	}

	return
}

/*
func main() {
  buf, err := ioutil.ReadFile("test.yml")
  if err != nil {
    panic(err)
  }

  var d Data
  err = yaml.Unmarshal(buf, &d)
  fmt.Println(d)
}
*/
