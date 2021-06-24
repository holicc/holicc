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

type checkDoneMsg int

type App struct {
	spinner   spinner.Model
	checkers  []Checker
	choices   []string
	selected  []map[int]struct{}
	doneNum   int
	checkDone chan checkDoneMsg
}

type Checker interface {
	Check()
	IsDone() bool
	Template() *template
}

func main() {
	app := NewApp(initSpinner(),
		WithSystemChecker(),
		WithNetworkChecker(),
	)
	p := tea.NewProgram(&app)
	if err := p.Start(); err != nil {
		log.Fatalln(err.Error())
	}
}

func NewApp(s spinner.Model, cks ...Checker) App {
	return App{
		spinner:   s,
		checkers:  cks,
		checkDone: make(chan checkDoneMsg, 1),
		choices: []string{
			"Git",
		},
	}
}

func (a *App) Init() tea.Cmd {
	tea.HideCursor()
	a.check()
	return tea.Batch(
		spinner.Tick,
		a.waitCheck,
	)
}

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return a.handleKeyboard(msg)
	case checkDoneMsg:
		a.doneNum++
		if a.doneNum == len(a.checkers) {
			return a, nil
		}
		return a, a.waitCheck
	case spinner.TickMsg:
		var cmd tea.Cmd
		a.spinner, cmd = a.spinner.Update(msg)
		return a, cmd
	}
	return a, nil
}

func (a *App) View() string {
	sb.Reset()
	for _, ck := range a.checkers {
		sb.WriteString(ck.Template().String(&a.spinner))
	}
	return sb.String()
}

/// private method
func (a *App) handleKeyboard(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return a, tea.Quit
	default:
		return a, nil
	}
}

func (a *App) check() {
	for _, ck := range a.checkers {
		if !ck.IsDone() {
			go func(checker Checker) {
				defer func() {
					if err := recover(); err != nil {
						log.Println("err", err)
						return
					}
					a.checkDone <- checkDoneMsg(1)
				}()
				checker.Check()
			}(ck)
		}
	}
}

func (a *App) waitCheck() tea.Msg {
	return <-a.checkDone
}

/// private function
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
