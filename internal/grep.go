package internal

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

func Grep(pattern, source []byte, re bool) {
	if re {
		grepE(pattern, source)

	} else {
		grep(pattern, source)
	}
}

func grep(pattern, source []byte) {
	for _, line := range bytes.Split(source, []byte("\n")) {
		if bytes.Contains(line, pattern) {
			fmt.Printf("%s\n", strings.ReplaceAll(string(line), string(pattern), color.RedString("%s", pattern)))
		}
	}
}

// with regexp support
func grepE(pattern, source []byte) {
	pat := regexp.MustCompile(string(pattern))
	for _, line := range bytes.Split(source, []byte("\n")) {
		if pat.Match(line) {
			fmt.Printf("%s\n", pat.ReplaceAllString(string(line), color.RedString("%s", pattern)))
		}
	}
}
