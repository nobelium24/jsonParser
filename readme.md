# JSON Parser Tool

## Overview
The JSON Parser Tool is a lightweight Go-based utility designed to tokenize and parse JSON files. It features a custom lexer and parser implementation to process JSON data into structured Go types. The tool also performs basic validation to ensure the JSON conforms to the expected syntax.

---

## Features

1. **Lexer**:
   - Scans JSON files and generates tokens for parsing.
   - Supports parsing of JSON objects, arrays, strings, numbers, booleans, and null values.

2. **Parser**:
   - Converts tokens into structured Go data types (maps, slices, primitives).
   - Validates JSON syntax (e.g., matching braces, proper key-value pairs, etc.).

3. **Error Handling**:
   - Provides meaningful error messages for invalid JSON syntax.
   - Detects unterminated strings, malformed numbers, and other issues.

---

## Installation

1. Clone the repository:
   ```bash
   git clone <repository_url>
   cd <repository_directory>
   ```

2. Build the project:
   ```bash
   go build -o json_parser main.go
   ```

3. Run the tool with a JSON file:
   ```bash
   ./json_parser <path_to_json_file>
   ```

---

## Usage

### Input Example
```json
{
  "name": "John Doe",
  "age": 30,
  "isEmployed": true,
  "skills": ["Go", "Python", "JavaScript"],
  "address": {
    "street": "123 Main St",
    "city": "Anytown"
  }
}
```

### CLI Example
```bash
./json_parser example.json
```

### Output Example
```
Tokens: [{ name } : "John Doe" , age : 30 , isEmployed : true , skills [ "Go" , "Python" , "JavaScript" ] , address { street : "123 Main St" , city : "Anytown" } ]
Parsed tokens: map[name:John Doe age:30 isEmployed:true skills:[Go Python JavaScript] address:map[street:123 Main St city:Anytown]]
```

---

## Code Structure

### `lexer` Package
Responsible for tokenizing the input JSON file.

#### Functions
- **`Lexer(filePath string) []interface{}`**: Tokenizes the input file and returns a slice of tokens.
- **`AccNumber`**: Processes numeric tokens.
- **`AccKeyword`**: Handles boolean and null tokens.
- **`AccWord`**: Manages string tokens and escape sequences.

### `parser` Package
Converts tokens into structured Go types and validates JSON.

#### Functions
- **`Parser(tokens []interface{}, index *int) (interface{}, error)`**: Entry point for parsing tokens.
- **`parseObject`**: Parses JSON objects into `map[string]interface{}`.
- **`parseArray`**: Parses JSON arrays into `[]interface{}`.

### `main` Package
Handles the CLI interface and orchestrates the lexer and parser.

#### Functions
- **`main()`**: Accepts a file path as input, tokenizes the file, parses it, and outputs the result.

---

<!-- ## Development

### Testing
Run tests for the lexer and parser:
```bash
go test ./...
```

### Debugging
Enable debug prints in the lexer and parser to trace issues in JSON parsing:
```go
fmt.Printf("Debug info: %v", <variable>)
``` -->

---

## Known Limitations
- Limited support for Unicode escape sequences in strings.
- Error messages could be more detailed in complex scenarios.
- Does not handle deeply nested JSON structures efficiently.

---

## Contributing
Contributions are welcome! Feel free to submit pull requests or open issues.

---

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.

---

## Author
Developed by nobelium24.

