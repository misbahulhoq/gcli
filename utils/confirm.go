package utils

import (
	"bufio"
	"os"
	"strings"

	"github.com/fatih/color"
)

func Confirm(message string) bool {
	// Ask for confirmation
	color.New(color.FgMagenta).Print(message)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(strings.ToLower(input))
	// We check for "y", "yes", or empty string (Enter to mean Yes)
	if input == "y" || input == "yes" || input == "" {
		return true
	}

	return false
}
