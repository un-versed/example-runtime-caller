package errors

import (
	"fmt"
)

func Log(err CustomError) {
	// Print Stack info to stdout
	// TODO: Replace with your own logger
	fmt.Printf("File: %s, Line: %d, Function: %s, Code %s \r\n", err.File, err.Line, err.FuncName, err.Code)
}
