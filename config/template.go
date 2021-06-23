package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	"strings"
)

type template struct {
	e error
	d bool

	name   string
	msg    string
	status int
}

func NewTemplate(name string) *template {
	return &template{
		name: name,
	}
}

func (t *template) Doing() *template {
	t.status = 0
	return t
}

func (t *template) Ok(info string) *template {
	t.status = 1
	t.msg = info
	return t
}

func (t *template) Warn(msg string) *template {
	t.status = 2
	t.msg = msg
	return t
}

func (t *template) Error(err error) *template {
	t.status = -1
	t.msg = err.Error()
	t.e = err
	return t
}

func (t *template) String(s *spinner.Model) string {
	var sb strings.Builder
	switch t.status {
	case -1:
		sb.WriteString(fmt.Sprintf(" > %s Error %s", Error, FRed(t.msg)))
	case 0:
		sb.WriteString(fmt.Sprintf("%sChecking %s...", s.View(), t.name))
	case 1:
		sb.WriteString(fmt.Sprintf("%sChecking %s...\n", Ok, t.name))
		sb.WriteString(fmt.Sprintf(" > %s\n", FGreen(t.msg)))
	case 2:
		sb.WriteString(fmt.Sprintf("%sChecking %s...\n", Warn, t.name))
		sb.WriteString(fmt.Sprintf(" > %s \n", FYellow(t.msg)))
	}
	return sb.String()
}
