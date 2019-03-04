package main

import(
	"fmt"
	"strings"
	"text/scanner"
	"unicode"
)

func Example(){
	const src = `
		// This is scanned code.
		if a > 10 {
			someParsable = text
		}`
	//const src = "if a > 10 {someParsable = text}"
	
	fmt.Println(src)
	
	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "example"
	for tok := s.Scan(); tok != scanner.EOF;tok = s.Scan(){
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
		fmt.Println(tok)
	}
}

func Example_isIdentRune(){
	const src = "%var1 var2%"
	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "default"
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan(){
	    fmt.Printf("%s:%s\n", s.Position, s.TokenText())	
	}
	
	fmt.Println()
	s.Init(strings.NewReader(src))
	s.Filename = "percent"
	s.IsIdentRune = func(ch rune, i int) bool {
		return ch == '%' && i == 0 || unicode.IsLetter(ch) || unicode.IsDigit(ch) && i > 0
	}
	
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan(){
		fmt.Printf("%s:%s\n", s.Position, s.TokenText())
	}
}

func main(){
	Example()
	Example_isIdentRune()
}
