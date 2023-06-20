package errors

import (
	"regexp"
	"runtime"
)

type CustomError struct {
	s        string // error message
	Code     string // error code
	File     string
	Line     int
	FuncName string // caller function name
}

func New(s string, code string) error {
	// Get Stack info
	pc, file, line, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()

	cerr := &CustomError{s, code, formatFilename(file), line, formatFilename(funcName)}

	// Print Stack info
	Log(*cerr)

	return cerr
}

func (cerr *CustomError) Error() string {
	return cerr.s
}

func formatFilename(filename string) string {
	regex := regexp.MustCompile(`example-runtime-caller.+`)
	fname := regex.FindString(filename)
	if fname == "" {
		return filename
	}

	return fname
}
