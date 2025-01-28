package parser

import (
	"fmt"
)

// func Parser(args []interface{}) {
// if len(args) > 2 && args[0] != "{" && args[len(args)-1] != "}" {
// 	fmt.Print("Invalid JSON")
// 	os.Exit(1)
// } else {
// 	fmt.Print("Valid JSON")
// 	os.Exit(0)
// }
// 	for i, t := range args {

// 	}

// }

// func parseObject(tokens []interface{}, index *int) (map[string]interface{}, error) {
// 	json := make(map[string]interface{})
// 	if tokens[0] != "{" && tokens[len(tokens)-1] != "}" {
// 		return nil, fmt.Errorf("Invalid JSON object: missing opening or closing brace")
// 	}
// 	if len(tokens) == 2 {
// 		return make(map[string]interface{}), nil
// 	} else {
// 		tokens = tokens[1:]
// 		tokens = tokens[:len(tokens)-1]

// 		var index int
// 		for i := 0; i < len(tokens); i++ {
// 			key, ok := tokens[i].(string)
// 			if !ok {
// 				return nil, fmt.Errorf("Invalid JSON object")
// 			}
// 			index = i
// 			if tokens[index+1] != ":" {
// 				return nil, fmt.Errorf("Invalid JSON object")
// 			}
// 			json[key] = tokens[index+2]
// 			index++
// 		}
// 	}
// 	return json, nil

// }

// func Parser(tokens []interface{}, index *int) (map[string]interface{}, error) {
// 	json := make(map[string]interface{})
// 	if tokens[0] != "{" || tokens[len(tokens)-1] != "}" {
// 		return nil, fmt.Errorf("Invalid JSON object: missing opening or closing brace")
// 	}

// 	if len(tokens) == 2 {
// 		return json, nil
// 	}
// 	*index++
// 	for *index < len(tokens)-1 {
// 		key, ok := tokens[*index].(string)
// 		if !ok {
// 			return nil, fmt.Errorf("Invalid JSON object: expected a string key at position %d", *index)
// 		}
// 		*index++
// 		if tokens[*index] != ":" {
// 			return nil, fmt.Errorf("Invalid JSON object: expected ':' after key at position %d", *index)
// 		}
// 		*index++

// 		value := tokens[*index]
// 		json[key] = value

// 		*index++
// 		if tokens[*index] == "," {
// 			*index++
// 		} else if tokens[*index] == "}" {
// 			break
// 		} else {
// 			return nil, fmt.Errorf("Invalid JSON object: expected ',' or '}' at position %d", *index)
// 		}
// 	}
// 	return json, nil
// }

func Parser(tokens []interface{}, index *int) (interface{}, error) {
	token := tokens[*index]
	switch token {
	case "{":
		return parseObject(tokens, index)
	case "[":
		return parseArray(tokens, index)
	default:
		*index++
		return token, nil
	}
}

func parseArray(tokens []interface{}, index *int) ([]interface{}, error) {
	array := []interface{}{}
	*index++
	for *index < len(tokens) {
		if tokens[*index] == "]" {
			*index++
			return array, nil
		}
		value, err := Parser(tokens, index)
		if err != nil {
			return nil, err
		}
		array = append(array, value)
		if *index < len(tokens) && tokens[*index] == "," {
			*index++
		} else if *index < len(tokens) && tokens[*index] != "]" {
			return nil, fmt.Errorf("Expected ',' or ']' at position %d", *index)
		}
	}
	return nil, fmt.Errorf("Unterminated array")
}

func parseObject(tokens []interface{}, index *int) (map[string]interface{}, error) {
	json := make(map[string]interface{})
	if tokens[*index] != "{" {
		return nil, fmt.Errorf("Invalid JSON object: missing opening brace")
	}
	*index++
	for *index < len(tokens) {
		if tokens[*index] == "}" {
			*index++
			return json, nil
		}
		key, ok := tokens[*index].(string)
		if !ok {
			return nil, fmt.Errorf("Invalid JSON object: expected a string key at position %d", *index)
		}
		*index++
		if tokens[*index] != ":" {
			return nil, fmt.Errorf("Invalid JSON object: expected ':' after key at position %d", *index)
		}
		*index++
		value, err := Parser(tokens, index)
		if err != nil {
			return nil, err
		}
		json[key] = value
		if *index < len(tokens) && tokens[*index] == "," {
			*index++
		} else if *index < len(tokens) && tokens[*index] != "}" {
			return nil, fmt.Errorf("Invalid JSON object: expected ',' or '}' at position %d", *index)
		}
	}
	return nil, fmt.Errorf("Unterminated object")
}
