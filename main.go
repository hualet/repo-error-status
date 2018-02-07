package main

import "github.com/Sirupsen/logrus"

func main() {
	repo := &Repo{
		Name:         "Debian Stable",
		Host:         "http://ftp.cn.debian.org/debian",
		Distribution: "stable",
		Arch:         "amd64",
	}

	err := inspect(repo)
	if err != nil {
		logrus.Fatal(err)
	}
}
