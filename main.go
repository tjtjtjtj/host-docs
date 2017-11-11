package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/tjtjtjtj/host-docs/assets"
	"github.com/tjtjtjtj/host-docs/common"
	"github.com/urfave/cli"
)

//go:generate go-bindata -o assets/assets.go -pkg assets assets/

var (
	hash      string
	builddate string
	goversion string
	version   = "1.0.0"
)

var envlist = [...]string{"production", "staging", "stress"}

type serverlistdata struct {
	HostsData *common.HostsData
	Env       string
}

func main() {

	app := cli.NewApp()
	app.Name = "hostvars-batch"
	app.Usage = "make host-vars markdown list"
	app.Version = version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "ansibledir",
			Usage: "ansible host-vars dir",
			Value: "common/ansible/host_vars",
		},
		cli.StringFlag{
			Name:  "serverspecdir",
			Usage: "serverspec host_vars dir",
			Value: "common/serverspec/host_vars",
		},
		cli.StringFlag{
			Name:  "outputdir",
			Usage: "markdown list ouput dir",
			Value: "/tmp",
		},
	}

	app.Action = func(c *cli.Context) error {
		s := new(serverlistdata)
		s.HostsData = new(common.HostsData)
		err := s.HostsData.AnsibleSetData(filepath.Clean(c.String("ansibledir")))
		if err != nil {
			return err
		}
		err := s.HostsData.ServerspecSetData(filepath.Clean(c.String("serverspecdir")))
		if err != nil {
			return err
		}

		for _, env := range envlist {
			s.Env = env
			file, err := os.Create(filepath.Join(filepath.Clean(c.String("outputdir")), env+".md"))
			if err != nil {
				return err
			}
			tmplData, err := assets.Asset("assets/serverlist.tmpl")
			if err != nil {
				return err
			}

			tmpl := template.Must(template.New("md").Parse(string(tmplData)))
			err = tmpl.Execute(file, s)
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
		pvStr, _ := json.Marshal(pv)
		fmt.Println(string(pvStr))
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}
}
