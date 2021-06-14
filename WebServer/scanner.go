package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

var s scanner.Scanner

func main() {
	const src = `
// This is scanned code.
if a > 10 {
	someParsable = text
}`
	s.Init(strings.NewReader(src))

	s.Filename = "example"
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}
}
