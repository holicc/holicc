package main

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
)

const (
	GREEN = "#66cd00"
	RED   = "#ee3b3b"
)

var (
	FGreen = style(GREEN)
	FRed   = style(RED)
	Ok     = FGreen("✔️ ")
	Error  = FRed("✖️ ")
	Warn   = FRed("⚠️ ")
)

type config struct {
	Name       string
	Shell      string
	ConfigFile string
	Doc        string
}

type App struct {
	spinner  spinner.Model
	t        *Template
	checkers []Checker
}

type Checker interface {
	Check(chan<- string)
}

func main() {

	s := initSpinner()
	app := &App{
		spinner: s,
		t:       NewTemplate(&s),
	}
	app.AddChecker(
	//WithSystemChecker(),
	)
	p := tea.NewProgram(app)
	if err := p.Start(); err != nil {
		log.Fatalln(err.Error())
	}
}

func (a *App) AddChecker(ck Checker, cks ...Checker) {
	a.checkers = append(a.checkers, ck)
	if len(cks) > 0 {
		a.checkers = append(a.checkers, cks...)
	}
}

func (a *App) Init() tea.Cmd {
	tea.HideCursor()
	return spinner.Tick
}

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return a, tea.Quit
		}
	default:
		var cmd tea.Cmd
		a.spinner, cmd = a.spinner.Update(msg)
		return a, cmd
	}
	return a, nil
}

func (a *App) View() string {
	return a.t.String()
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

func CheckingTemplate(name string) {

}
