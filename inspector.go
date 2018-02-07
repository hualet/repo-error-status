package main

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/Sirupsen/logrus"
)

// Repo represents a debian repository to be inspected.
// Basically it's all information needed to feed r-inspector.
type Repo struct {
	Name         string `json:"name"`
	Host         string `json:"host"`
	Distribution string `json:"dist"`
	Arch         string `json:"arch"`
}

// Result is a inspect result, contains all repo error information.
type Result struct {
	Repo               *Repo
	ErrBreakDependInfo uint32
	ErrBreakDepends    uint32
	ErrInvalidBinary   uint32
	ErrMismatchBinary  uint32
	ErrNoSrc           uint32
	ErrNoSrc2          uint32
}

func inspect(repo *Repo) (*Result, error) {
	cmd := exec.Command("/usr/bin/r-inspector", "-host", repo.Host,
		"-suite", repo.Distribution, "-arch", repo.Arch)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to pipe stdout of r-inspector: %s", err)
	}
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to run r-inspector: %s", err)
	}

	logrus.Infof("inspecting repo %s...", repo.Name)

	var inspectResult struct {
		Summary struct {
			ProblemPackages Result `json:"problemPackages"`
		} `json:"summary"`
	}
	if err := json.NewDecoder(stdout).Decode(&inspectResult); err != nil {
		return nil, fmt.Errorf("parsing output of r-inspector error: %s", err)
	}
	if err := cmd.Wait(); err != nil {
		return nil, fmt.Errorf("wait for r-inspector error: %s", err)
	}

	inspectResult.Summary.ProblemPackages.Repo = repo

	return &inspectResult.Summary.ProblemPackages, nil
}
