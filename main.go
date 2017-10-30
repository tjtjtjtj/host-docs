package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"

	"github.com/tjtjtjtj/host-docs/common"
	"github.com/urfave/cli"
)

var (
	hash      string
	builddate string
	goversion string
	version   = "1.0.0"
)

var envlist = [...]string{"production", "staging", "stress"}

func main() {

	app := cli.NewApp()
	app.Name = "hostvars-batch"
	app.Usage = "make host-vars markdown list"
	app.Version = version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "ansibledir",
			Usage: "ansible host-vars dir",
			Value: "./common/ansible/host_vars/",
		},
		cli.StringFlag{
			Name:  "serverspecdir",
			Usage: "serverspec host_vars dir",
			Value: "./common/serverspec/host_vars/",
		},
		cli.StringFlag{
			Name:  "outputdir",
			Usage: "markdown list ouput dir",
			Value: "/tmp/",
		},
	}

	app.Action = func(c *cli.Context) error {
		hostsdata := new(common.HostsData)
		hostsdata.AnsibleSetData(c.String("ansibledir"))
		hostsdata.ServerspecSetData(c.String("serverspecdir"))
		var outputdate common.HostsData

		for _, env := range envlist {
			// todo:環境で出力を分ける(ここテンプレート上での分岐がいい気が）

			file, err := os.Create(c.String("outputdir") + env + ".md")
			tmpl := template.Must(template.ParseFiles("serverlist.tmpl"))

			err = tmpl.Execute(file, hostsdata)
			if err != nil {
				return err
			}
		}
		return nil
	}

	cli.VersionPrinter = func(c *cli.Context) {
		pv := struct {
			Version   string
			Hash      string
			BuildDate string
			GoVersion string
		}{
			Version:   app.Version,
			Hash:      hash,
			BuildDate: builddate,
			GoVersion: goversion,
		}
		pv_str, _ := json.Marshal(pv)
		fmt.Println(string(pv_str))
	}

	app.Run(os.Args)

}
