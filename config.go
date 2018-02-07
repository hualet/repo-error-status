package main

import (
	"encoding/json"
	"fmt"
)

type config struct {
	InfluxDBHost     string `json:"influxdb_host"`
	InfluxDBPort     uint16 `json:"influxdb_port"`
	InfluxDBUserName string `json:"influxdb_username"`
	InfluxDBPassword string `json:"influxdb_password"`
	InfluxDBName     string `json:"influxdb_name"`

	Repos []Repo `json:"repos"`
}

var defaultConfig = &config{
	InfluxDBHost:     "",
	InfluxDBPort:     8806,
	InfluxDBUserName: "",
	InfluxDBPassword: "",
	InfluxDBName:     "",

	Repos: []Repo{
		Repo{
			Name:         "Debian Stable",
			Host:         "http://ftp.cn.debian.org/debian",
			Distribution: "stable",
			Arch:         "amd64",
		},
		Repo{
			Name:         "Debian Unstable",
			Host:         "http://ftp.cn.debian.org/debian",
			Distribution: "unstable",
			Arch:         "amd64",
		},
		Repo{
			Name:         "Debian Testing",
			Host:         "http://ftp.cn.debian.org/debian",
			Distribution: "testing",
			Arch:         "amd64",
		},
	},
}

func dumpDefaultConfig() error {
	data, err := json.MarshalIndent(defaultConfig, "", "    ")
	if err != nil {
		return err
	}
	_, err = fmt.Print(string(data))
	return err
}
