package main

import (
	"github.com/charmbracelet/bubbles/spinner"
)

type NetworkChecker struct {
	connected bool
	err       error
}

func WithNetworkChecker() *NetworkChecker {
	return &NetworkChecker{}
}

func (n *NetworkChecker) View(model spinner.Model) string {
	return ""
}
