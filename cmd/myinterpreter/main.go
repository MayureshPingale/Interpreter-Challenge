package main

import (
	"fmt"
	"os"
)

const (
	LEFT_PAREN  = "LEFT_PAREN"
	RIGHT_PAREN = "RIGHT_PAREN"
	LEFT_BRACE  = "LEFT_BRACE"
	RIGHT_BRACE = "RIGHT_BRACE"
	COMMA       = "COMMA"
	DOT         = "DOT"
	MINUS       = "MINUS"
	PLUS        = "PLUS"
	SEMICOLON   = "SEMICOLON"
	STAR        = "STAR"
)

var current int = 0
var fileContent string = ""

func check(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", e)
		os.Exit(1)
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	check(err)

	if len(fileContents) > 0 {
		fileContent = string(fileContents)
		scanToken()
		fmt.Println("EOF  null")
	} else {
		panic("Unable to read fileContents")
	}
}

func scanToken() {
	for _, c := range fileContent {
		switch c {
		case '(':
			fmt.Println(LEFT_PAREN, "(", "null")
		case ')':
			fmt.Println(RIGHT_PAREN, ")", "null")
		case '{':
			fmt.Println(LEFT_BRACE, "{", "null")
		case '}':
			fmt.Println(RIGHT_BRACE, "}", "null")
		case ',':
			fmt.Println(COMMA, ",", "null")
		case '.':
			fmt.Println(DOT, ".", "null")
		case '-':
			fmt.Println(MINUS, "-", "null")
		case '+':
			fmt.Println(PLUS, "+", "null")
		case ';':
			fmt.Println(SEMICOLON, ";", "null")
		case '*':
			fmt.Println(STAR, "*", "null")
		}
	}
}
