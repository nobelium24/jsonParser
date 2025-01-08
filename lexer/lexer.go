package lexer

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Lexer(filePath string) (tokens []interface{}) {
	file, err := os.Open(filePath)
	if err != nil {
		return tokens
	}

	defer file.Close()
	data := bufio.NewScanner(file)

	data.Split(bufio.ScanRunes)
	for data.Scan() {
		item := data.Text()
		if item == "{" || item == "}" || item == "[" || item == "]" || item == ":" || item == "," {
			tokens = append(tokens, item)
			continue
		} else if item == "\"" {
			word, err := AccWord(data)
			if err != nil {

				return nil
			}
			tokens = append(tokens, word)
		}
	}
	return tokens
}

func AccWord(scanner *bufio.Scanner) (interface{}, error) {
	var buffer strings.Builder
	isString := false

	if scanner.Text() == `"` {
		isString = true
	}
	numberRe := regexp.MustCompile(`^-?(0|[1-9][0-9]*)(\.[0-9]+)?([eE][+-]?[0-9]+)?$`)
	keywordRe := regexp.MustCompile(`^(true|false|null)$`)
	for scanner.Scan() {
		char := scanner.Text()
		if isString {
			if char == "\\" {
				if !scanner.Scan() {
					return nil, fmt.Errorf("Unterminated escape sequence in JSON")
				}
				escapeChar := scanner.Text()
				if escapeChar == `"` || escapeChar == `\` || escapeChar == "/" {
					buffer.WriteString(escapeChar)
				} else if escapeChar == "b" {
					buffer.WriteByte('\b')
				} else if escapeChar == "f" {
					buffer.WriteByte('\f')
				} else if escapeChar == "n" {
					buffer.WriteByte('\n')
				} else if escapeChar == "r" {
					buffer.WriteByte('\r')
				} else if escapeChar == "t" {
					buffer.WriteByte('\t')
				} else {
					return nil, fmt.Errorf("Invalid escape character %q", escapeChar)
				}
				continue
			}
			if char == `"` {
				return buffer.String(), nil
			}
		} else {
			if char < "a" || char > "z" {
				break
			}
			buffer.WriteString(char)
		}
	}
	word := buffer.String()
	if isString {
		return nil, fmt.Errorf("unterminated string")
	}
	if keywordRe.MatchString(word) {
		switch word {
		case "true":
			return true, nil
		case "false":
			return false, nil
		case "null":
			return nil, nil
		default:
			return word, nil // Unrecognized word
		}
	} else if numberRe.MatchString(word) {
		return ParseNumber(word)
	}
	return nil, fmt.Errorf("invalid token: %q", word)
}

func ParseNumber(s string) (interface{}, error) {
	if strings.Contains(s, ".") || strings.ContainsAny(s, "eE") {
		return strconv.ParseFloat(s, 64)
	}
	return strconv.Atoi(s)
}
