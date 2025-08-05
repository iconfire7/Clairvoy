package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

// PromptLine reads a line of text.
func PromptLine(label string) string {
	fmt.Print(label)
	r := bufio.NewReader(os.Stdin)
	text, _ := r.ReadString('\n')
	return strings.TrimSpace(text)
}

// PromptSecret reads a password without echo.
func PromptSecret(label string) (string, error) {
	fmt.Print(label)
	pass, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	return string(pass), err
}

// PromptMultiline reads until EOF.
func PromptMultiline(desc string) string {
	fmt.Printf("%s (Ctrl+D to finish):\n", desc)
	var b strings.Builder
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		b.WriteString(scanner.Text() + "\n")
	}
	return b.String()
}
