package main

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
	"strings"
)

const (
	GREEN  = "#66cd00"
	RED    = "#ee3b3b"
	YELLOW = "#ffd700"
)

var (
	sb strings.Builder

	FGreen  = style(GREEN)
	FRed    = style(RED)
	FYellow = style(YELLOW)
	Ok      = FGreen("✔️ ")
	Error   = FRed("✖️ ")
	Warn    = FRed("⚠️ ")
)

type config struct {
	Name       string
	Shell      string
	ConfigFile string
	Doc        string
}

type App struct {
	spinner  spinner.Model
	checkers []Checker
	doneNum  int
}

type Checker interface {
	Check()
	IsDone() bool
	Template() *template
}

func main() {
	app := NewApp(initSpinner(),
		WithGitChecker(NewTemplate("git")),
	)
	p := tea.NewProgram(app)
	if err := p.Start(); err != nil {
		log.Fatalln(err.Error())
	}
}

func NewApp(s spinner.Model, cks ...Checker) App {
	return App{
		spinner:  initSpinner(),
		checkers: cks,
	}
}

func (a *App) AddChecker(ck Checker, cks ...Checker) {
	a.checkers = append(a.checkers, ck)
	if len(cks) > 0 {
		a.checkers = append(a.checkers, cks...)
	}
}

func (a App) Init() tea.Cmd {
	tea.HideCursor()
	return spinner.Tick
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return a, tea.Quit
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		a.spinner, cmd = a.spinner.Update(msg)
		return a, cmd

	}
	a.spinner.Finish()
	if a.doneNum == len(a.checkers) {
		return a, nil
	}
	for _, ck := range a.checkers {
		if !ck.IsDone() {
			go func(checker Checker) {
				defer func() {
					err := recover()
					if err != nil {
						log.Fatalln(err)
					}
					a.doneNum++
				}()
				checker.Check()
			}(ck)
		}
	}
	return a, nil
}

func (a App) View() string {
	sb.Reset()
	for _, ck := range a.checkers {
		sb.WriteString(ck.Template().String(&a.spinner))
	}
	return sb.String()
}

func style(color string) func(string) string {
	return func(s string) string {
		return lipgloss.NewStyle().
			Foreground(lipgloss.Color(color)).
			SetString(s).
			String()
	}
}

func initSpinner() spinner.Model {
	s := spinner.NewModel()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return s
}
