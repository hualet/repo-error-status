package main

type config struct {
	InfluxDBHost string `json:"influxdb_host"`
	InfluxDBPort uint16 `json:"influxdb_port"`

	Repos []Repo `json:"repos"`
}

var defaultConfig = &config{
	InfluxDBHost: "10.0.10.65",
	InfluxDBPort: 8806,

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
