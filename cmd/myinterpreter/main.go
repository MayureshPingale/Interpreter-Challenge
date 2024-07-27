package main

import (
	"fmt"
	"os"
)

const (
	LEFT_PAREN    = "LEFT_PAREN"
	RIGHT_PAREN   = "RIGHT_PAREN"
	LEFT_BRACE    = "LEFT_BRACE"
	RIGHT_BRACE   = "RIGHT_BRACE"
	COMMA         = "COMMA"
	DOT           = "DOT"
	MINUS         = "MINUS"
	PLUS          = "PLUS"
	SEMICOLON     = "SEMICOLON"
	STAR          = "STAR"
	BANG          = "BANG"
	BANG_EQUAL    = "BANG_EQUAL"
	EQUAL_EQUAL   = "EQUAL_EQUAL"
	LESS_EQUAL    = "LESS_EQUAL"
	GREATER_EQUAL = "GREATER_EQUAL"
	EQUAL         = "EQUAL"
	LESS          = "LESS"
	GREATER       = "GREATER"
)

var fileContent string = ""
var itr int = 0

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

	fileContent = string(fileContents)
	var hasError bool = scanToken()
	fmt.Println("EOF  null")

	if hasError {
		os.Exit(65)
	}
}

func scanToken() bool {
	var error bool = false
	var lineNo int = 1
	for ; itr < len(fileContent); itr++ {
		var c = rune(fileContent[itr])
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
		case '\n':
			lineNo++
		case '!':
			if match('=') {
				fmt.Println(BANG_EQUAL, "!=", "null")
			} else {
				fmt.Println(BANG, "=", "null")
			}
		case '=':
			if match('=') {
				fmt.Println(EQUAL_EQUAL, "==", "null")
			} else {
				fmt.Println(EQUAL, "=", "null")
			}
		case '<':
			if match('=') {
				fmt.Println(LESS_EQUAL, "<=", "null")
			} else {
				fmt.Println(LESS, ">", "null")
			}
		case '>':
			if match('=') {
				fmt.Println(GREATER_EQUAL, ">=", "null")
			} else {
				fmt.Println(GREATER, ">", "null")
			}
		default:
			fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n", lineNo, c)
			error = true
		}
	}

	return error
}

func match(exp rune) bool {
	if itr+1 >= len(fileContent) {
		return false
	}

	if rune(fileContent[itr+1]) != exp {
		return false
	}

	itr++
	return true
}
