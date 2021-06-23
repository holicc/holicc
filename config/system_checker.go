package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	"runtime"
)

type SystemType struct {
	osName string
}

func WithSystemChecker() *SystemType {
	return &SystemType{}
}

func (s *SystemType) View(spinner spinner.Model) string {
	str := " %sChecking system os\n"
	if s.osName == "" {
		s.checkSystemOs()
		return fmt.Sprintf(str, spinner.View())
	} else {
		return fmt.Sprintf("%s  > %s\n", fmt.Sprintf(str, Ok), FGreen(s.osName))
	}
}

func (s *SystemType) checkSystemOs() {
	s.osName = runtime.GOOS
}
