package parser

import (
	"fmt"
	"os"
)

func Parser(args []interface{}) {
	if len(args) > 2 && args[0] != "{" && args[len(args)-1] != "}" {
		fmt.Print("Invalid JSON")
		os.Exit(1)
	} else {
		fmt.Print("Valid JSON")
		os.Exit(0)
	}

}
