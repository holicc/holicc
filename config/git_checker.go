package main

import (
	"os/exec"
	"strings"
)

type Git struct {
	t       *template
	path    string
	version string
	err     error
	done    bool
}

func WithGitChecker() *Git {
	return &Git{
		t: NewTemplate("Git"),
	}
}

func (g *Git) Check() {
	defer func() {
		g.done = true
	}()
	path, err := exec.LookPath("git")
	if err != nil {
		g.err = err
	} else {
		g.path = path
		cmd := exec.Command("git", "version")
		version, err := cmd.CombinedOutput()
		if err != nil {
			g.err = err
			g.t.Error(g.err)
		} else {
			g.version = strings.ReplaceAll(string(version), "\n", "")
			g.t.Ok(g.version)
		}
	}
}

func (g *Git) IsDone() bool {
	return g.done
}

func (g *Git) Template() *template { return g.t }
