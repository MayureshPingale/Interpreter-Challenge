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
	SLASH         = "SLASH"
	STRING        = "STRING"
)

var fileContent string = ""
var itr int = 0
var lineNo int = 1
var scanningError bool = false

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

	scanToken()
	fmt.Println("EOF  null")

	if scanningError {
		os.Exit(65)
	}
}

func scanToken() {
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
			checkDualOperator('!', '=', BANG_EQUAL, BANG)
		case '=':
			checkDualOperator('=', '=', EQUAL_EQUAL, EQUAL)
		case '<':
			checkDualOperator('<', '=', LESS_EQUAL, LESS)
		case '>':
			checkDualOperator('>', '=', GREATER_EQUAL, GREATER)
		case '/':
			if match('/') {
				for itr < len(fileContent) && rune(fileContent[itr]) != '\n' {
					itr++
				}
				lineNo++
			} else {
				fmt.Println(SLASH, "/", "null")
			}
		case ' ':
		case '\r':
		case '\t':
		case '"':
			scanString()
		default:
			fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n", lineNo, c)
			scanningError = true
		}
	}
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

func checkDualOperator(firstChar rune, secondChar rune, dualOperatorToken string, singleCharToken string) {
	if match(secondChar) {
		fmt.Println(dualOperatorToken, string(firstChar)+string(secondChar), "null")
	} else {
		fmt.Println(singleCharToken, string(firstChar), "null")
	}
}

func scanString() {
	var start int = itr
	itr++

	for itr < len(fileContent) && fileContent[itr] != '"' {
		if fileContent[itr] == '\n' {
			lineNo++
		}

		itr++
	}

	if itr == len(fileContent) {
		fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.\n", lineNo)
		scanningError = true
	} else {
		var str = fileContent[start : itr+1]
		fmt.Println(STRING, str, fileContent[start+1:itr])
	}
}
