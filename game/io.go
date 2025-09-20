package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var arr = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.ANSIColor(2))

var warn = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.ANSIColor(1))

func message(s string) {
	fmt.Printf("%s %s\n", arr.Render("->"), s)
}

func messageItemPick(min, max int) {
	fmt.Printf("%s Take your reward %s.\n",
		arr.Render("->"),
		warn.Render(fmt.Sprintf("(%d/%d)", min, max)))
}

func listCommands(s ...string) {
	fmt.Println("Available commands:")
	for _, v := range s {
		fmt.Printf("-%s\n", v)
	}
}

func warning(s string) {
	fmt.Printf("%s\n", warn.Render(s))
}

func blank() {
	fmt.Println()
}

func dial() string {
	reader := bufio.NewReader(os.Stdin)

	warning("Command?")
	line, _, _ := reader.ReadLine()
	blank()

	formattedLine := strings.TrimSpace(
		strings.ToLower(string(line)))

	return formattedLine
}

const (
	PICKUPGOLD   = "pickupgold"
	PICKUPSHOVEL = "pickupshovel"
	PICKUPROPE   = "pickuprope"

	USEGOLD   = "pickupgold"
	USESHOVEL = "pickupshovel"
	USEROPE   = "pickuprope"

	GOSOUTH = "gosouth"
	GONORTH = "gonorth"
	GOEAST  = "goeast"
	GOWEST  = "gowest"

	DIGHOLE  = "dighole"
	DROPGOLD = "dropgold"
	FILLHOLE = "fillhole"

	HELP  = "help"
	CHEAT = "cheat"
)
