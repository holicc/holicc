package main

import (
	"os/exec"
	"strings"
)

type Git struct {
	path    string
	version string
	err     error
}

func WithGitChecker() *Git {
	return &Git{}
}

func (g *Git) View(t chan string) {
	if g.err != nil {
		return g.t.Error(g.err)
	}
	if g.path == "" {
		go g.checkGit()
		return g.t
	} else {
		return g.t.Ok(g.version)
	}
}

func (g *Git) checkGit() {
	path, err := exec.LookPath("git")
	if err != nil {
		g.err = err
	} else {
		g.path = path
		cmd := exec.Command("git", "version")
		version, err := cmd.CombinedOutput()
		if err != nil {
			g.err = err
		} else {
			g.version = strings.ReplaceAll(string(version), "\n", "")
		}
	}
}
