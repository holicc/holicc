package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	"strings"
)

type Template struct {
	s *spinner.Model
	b strings.Builder
	d bool
}

func NewTemplate(s *spinner.Model) *Template {
	var b strings.Builder
	return &Template{
		b: b,
		s: s,
	}
}

func (t *Template) Checker(name string) *Template {
	t.b.WriteString(fmt.Sprintf(" %sChecking %s...\n", t.s.View(), name))
	return t
}

func (t *Template) Ok(info string) *Template {
	t.b.WriteString(fmt.Sprintf("%s > %s\n", Ok, info))
	return t
}

func (t *Template) Warn(msg string) *Template {
	t.b.WriteString(fmt.Sprintf("%s > %s\n", Warn, msg))
	return t
}

func (t *Template) Error(err error) *Template {
	t.b.WriteString(fmt.Sprintf(" %s Error %v", Error, err.Error()))
	return t
}

func (t *Template) String() string {
	return t.b.String()
}
