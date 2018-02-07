package main

import (
	"github.com/Sirupsen/logrus"
)

func main() {
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
