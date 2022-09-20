package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	ERR_OS_STDIN_STAT = "Failed to read STDIN info"
	ERR_PIPED_INPUT   = "Failed to read piped input"
)

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(ERR_OS_STDIN_STAT)
	}

	var parsedStrings []string
	if info.Mode()&os.ModeNamedPipe != 0 {
		parsedStrings, err = handlePiped(os.Stdin)
		if err != nil {
			log.Fatal(ERR_PIPED_INPUT)
		}
	} else {
		// handle arguments passed in via the command line
		parsedStrings = os.Args[1:]
	}

	for _, str := range parsedStrings {
		logResult(str, isPalindrome(str))
	}
}

func handlePiped(stdIn io.Reader) ([]string, error) {
	reader := bufio.NewReader(stdIn)
	output := []string{""}
	for {
		input, _, err := reader.ReadRune()

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				return output, err
			}
		}

		if input == rune('\n') {
			output = append(output, "")
		} else {
			output[len(output)-1] += string(input)
		}
	}
	return output, nil
}

func isPalindrome(str string) bool {
	if len(str) <= 1 {
		return false
	}

	chars := strings.Split(str, "")

	for i := 0; i <= len(chars)/2; i++ {
		if chars[i] != chars[len(chars)-(i+1)] {
			return false
		}
	}
	return true
}

func logResult(str string, result bool) {
	s := "is not"
	if result {
		s = "is"
	}

	fmt.Printf("\"%s\" %s a palindrome\n", strings.Trim(str, ""), s)
}
