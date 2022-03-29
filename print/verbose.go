package print

import "fmt"

var verbose bool = false

func SetVerbose(v bool) {
	verbose = v
}

func Verbose(a ...interface{}) (n int, err error) {
	if verbose {
		return fmt.Print(a...)
	}
	return 0, nil
}

func Verboseln(a ...interface{}) (n int, err error) {
	if verbose {
		return fmt.Println(a...)
	}
	return 0, nil
}

func Verbosef(format string, a ...interface{}) (n int, err error) {
	if verbose {
		return fmt.Printf(format, a...)
	}
	return 0, nil
}
