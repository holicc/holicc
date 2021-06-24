package main

import (
	"runtime"
)

type SystemType struct {
	osName string
	t      *template
}

func WithSystemChecker() *SystemType {
	return &SystemType{
		t: NewTemplate("System-Os"),
	}
}

func (s *SystemType) Check() {
	s.osName = runtime.GOOS
	s.t.Ok(s.osName)
}

func (s *SystemType) IsDone() bool {
	return s.osName != ""
}

func (s *SystemType) Template() *template {
	return s.t
}
