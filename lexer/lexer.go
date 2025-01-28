package lexer

import (
	"bufio"
	"fmt"
	"os"

	// "regexp"
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
		fmt.Printf("Scanned item: %s\n", item)
		if strings.TrimSpace(item) == "" {
			continue
		}
		if item == "{" || item == "}" || item == "[" || item == "]" || item == ":" || item == "," {
			tokens = append(tokens, item)
			continue
		} else if item == "\"" {
			word, err := AccWord(data)
			if err != nil {

				return nil
			}
			tokens = append(tokens, word)
		} else if item >= "0" && item <= "9" || item == "-" {
			word, err := AccNumber(data, item)
			if err != nil {
				return nil
			}
			tokens = append(tokens, word)
		} else if item >= "a" && item <= "z" || item >= "A" && item <= "Z" {
			word, err := AccKeyword(data, item)
			if err != nil {
				return nil
			}
			tokens = append(tokens, word)
		}
	}
	return tokens
}

func AccNumber(scanner *bufio.Scanner, firstChar string) (interface{}, error) {
	var buffer strings.Builder
	buffer.WriteString(firstChar)
	for scanner.Scan() {
		char := scanner.Text()
		if char < "0" || char > "9" && char != "." && char != "e" && char != "E" && char != "+" && char != "-" {
			break
		}
		buffer.WriteString(char)
	}
	return ParseNumber(buffer.String())
}

// func AccKeyword(scanner *bufio.Scanner, firstChar string) (interface{}, error) {
// 	var buffer strings.Builder
// 	buffer.WriteString(firstChar)
// 	for scanner.Scan() {
// 		char := scanner.Text()
// 		if char < "a" || char > "z" && char < "A" || char > "Z" {
// 			break
// 		}
// 		buffer.WriteString(char)
// 	}
// 	word := buffer.String()
// 	switch word {
// 	case "true":
// 		return true, nil
// 	case "false":
// 		return false, nil
// 	case "null":
// 		return nil, nil
// 	default:
// 		return nil, fmt.Errorf("invalid token: %q", word)
// 	}
// }

func AccKeyword(scanner *bufio.Scanner, firstChar string) (interface{}, error) {
	var buffer strings.Builder
	buffer.WriteString(firstChar)
	for scanner.Scan() {
		char := scanner.Text()
		if char < "a" || char > "z" && char < "A" || char > "Z" {
			break
		}
		buffer.WriteString(char)
	}
	word := buffer.String()
	switch word {
	case "true":
		return true, nil
	case "false":
		return false, nil
	case "null":
		return nil, nil
	default:
		return nil, fmt.Errorf("invalid token: %q", word)
	}
}

func AccWord(scanner *bufio.Scanner) (interface{}, error) {
	var buffer strings.Builder
	for scanner.Scan() {
		char := scanner.Text()
		if char == `"` {
			return buffer.String(), nil
		}

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
			} else if escapeChar == "u" {
				if !scanner.Scan() {
					return nil, fmt.Errorf("Unterminated Unicode escape sequence in JSON")
				}
				unicodeSeq := scanner.Text()
				if len(unicodeSeq) != 4 {
					return nil, fmt.Errorf("Invalid Unicode escape sequence: %q", unicodeSeq)
				}
				// Convert the Unicode sequence to a rune and then to a string
				unicodeValue, err := strconv.ParseInt(unicodeSeq, 16, 32)
				if err != nil {
					return nil, fmt.Errorf("Invalid Unicode escape sequence: %q", unicodeSeq)
				}
				buffer.WriteRune(rune(unicodeValue))
				continue
			} else {
				return nil, fmt.Errorf("Invalid escape character %q", escapeChar)
			}
			continue
		}
		buffer.WriteString(char)
	}
	return nil, fmt.Errorf("unterminated string")
}

func ParseNumber(s string) (interface{}, error) {
	if strings.Contains(s, ".") || strings.ContainsAny(s, "eE") {
		return strconv.ParseFloat(s, 64)
	}
	return strconv.Atoi(s)
}

// func AccWord(scanner *bufio.Scanner) (interface{}, error) {
// 	var buffer strings.Builder
// 	isString := false

// 	if scanner.Text() == `"` {
// 		isString = true
// 	}
// 	numberRe := regexp.MustCompile(`^-?(0|[1-9][0-9]*)(\.[0-9]+)?([eE][+-]?[0-9]+)?$`)
// 	keywordRe := regexp.MustCompile(`^(true|false|null)$`)
// 	for scanner.Scan() {
// 		char := scanner.Text()
// 		if isString {
// 			if char == "\\" {
// 				if !scanner.Scan() {
// 					return nil, fmt.Errorf("Unterminated escape sequence in JSON")
// 				}
// 				escapeChar := scanner.Text()
// 				if escapeChar == `"` || escapeChar == `\` || escapeChar == "/" {
// 					buffer.WriteString(escapeChar)
// 				} else if escapeChar == "b" {
// 					buffer.WriteByte('\b')
// 				} else if escapeChar == "f" {
// 					buffer.WriteByte('\f')
// 				} else if escapeChar == "n" {
// 					buffer.WriteByte('\n')
// 				} else if escapeChar == "r" {
// 					buffer.WriteByte('\r')
// 				} else if escapeChar == "t" {
// 					buffer.WriteByte('\t')
// 				} else {
// 					return nil, fmt.Errorf("Invalid escape character %q", escapeChar)
// 				}
// 				continue
// 			}
// 			if char == `"` {
// 				return buffer.String(), nil
// 			}
// 		} else {
// 			if char < "a" || char > "z" {
// 				break
// 			}
// 			buffer.WriteString(char)
// 		}
// 	}
// 	word := buffer.String()
// 	if isString {
// 		return nil, fmt.Errorf("unterminated string")
// 	}
// 	if keywordRe.MatchString(word) {
// 		switch word {
// 		case "true":
// 			return true, nil
// 		case "false":
// 			return false, nil
// 		case "null":
// 			return nil, nil
// 		default:
// 			return word, nil // Unrecognized word
// 		}
// 	} else if numberRe.MatchString(word) {
// 		return ParseNumber(word)
// 	}
// 	return nil, fmt.Errorf("invalid token: %q", word)
// }
