package log

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var (
	// 1P
	Prefix1P            = "[" + WhiteBgBlackText.Render("1P") + "]"
	WhiteBgBlackText = lipgloss.NewStyle().
				Background(lipgloss.Color("15")).
				Foreground(lipgloss.Color("0"))
	
	// BV
	PrefixBV  = "[" + BlueB.Render("B") + GreenV.Render("V") + "]"
	BlueB  = lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Bold(true)
	GreenV = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)

	// Function
	FunctionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("202"))
)

func ApplyStyles(l *log.Logger) *log.Logger {
	styles := log.DefaultStyles()

	styles.Prefix = lipgloss.NewStyle()

	styles.Keys["element"] = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575"))

	l.SetStyles(styles)
	return l
}
