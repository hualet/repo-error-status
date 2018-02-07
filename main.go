package main

import (
	"flag"
	"os"

	"github.com/Sirupsen/logrus"
)

var (
	configFile string
)

func main() {
	flag.StringVar(&configFile, "config", "",
		"configuration file path, default values will be used if not provided.")
	flag.Parse()

	if flag.Arg(0) == "dump" {
		err := dumpDefaultConfig()
		if err != nil {
			logrus.Fatal(err)
		}
		os.Exit(0)
	}

	repos := defaultConfig.Repos

	for _, repo := range repos {
		// not using goroutine here because r-inspector consumes too much system
		// resource, no need to parallize two instances.
		result, err := inspect(&repo)

		if err != nil {
			logrus.Errorf("failed to inspect repo: %s, %s", repo.Name, err)
			continue
		}

		err = report(result)
		if err != nil {
			logrus.Errorf("failed to report inspect result of %s: %s", repo.Name, err)
			continue
		}
	}
}
