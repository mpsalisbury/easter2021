package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	AnsiPrintf("Welcome to {red}E{yellow}a{orange}s{blue}t{green}e{red}r{reset} Puzzles 2021!\n")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		AnsiPrintf("\nEnter answer code: {bold}")
		if !scanner.Scan() {
			break
		}
		AnsiPrintf("{reset}")
		answer := scanner.Text()
		clue, ok := checkAnswer(answer)
		if ok {
			AnsiPrintf("{red}Correct!{reset} ")
			AnsiPrintf("Your next puzzle location clue: ")
			AnsiPrintf("{bold}%s{reset}\n", clue)
		} else {
			AnsiPrintf("Sorry, answer %q is not right\n", answer)
		}
	}
}

func AnsiPrintf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(Ansify(format), a...)
}

// Replaces ansi commands (e.g. "{bold}") with their ansi code equivalents.
func Ansify(s string) string {
	codes := map[string]string{
		"reset":  "0m",
		"bold":   "1m",
		"black":  "30m",
		"red":    "31m",
		"green":  "32m",
		"yellow": "33m",
		"blue":   "94m",
		"orange": "48;2;255;165;0m",
	}
	for cmd, code := range codes {
		s = strings.ReplaceAll(s,
			fmt.Sprintf("{%s}", cmd), fmt.Sprintf("\x1b[%s", code))
	}
	return s
}

// Returns the clue corresponding to the given answer, or return false.
func checkAnswer(answer string) (string, bool) {
	var clues = map[string]string{
		"bunny":       "Under spoons",
		"omega":       "Inside Apples",
		"gohan":       "Pull ups",
		"sandy":       "Under rocker stool",
		"subjective":  "With Piggy",
		"mclaren":     "Under filter",
		"spike":       "Ticklish ivories",
		"cocoa":       "With Calvin",
		"beethoven":   "Underwater",
		"pitbull":     "Martian",
		"chamberlain": "Under rodent",
		"6593":        "Underweight",
		"century":     "In a vacuum",
		"shuri":       "Under table",
		"happyland":   "Around a dull blade",
		"superman":    "Where Life Begins & Love Never Ends",
	}
	key := strings.ToLower(strings.Replace(answer, " ", "", -1))
	clue, ok := clues[key]
	return clue, ok
}
