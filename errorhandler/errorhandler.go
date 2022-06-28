package errorhandler

import (
	"fmt"
	"log"
	"os"
)

var exitHandler func()

func SetExitHandler(f func()) {
	exitHandler = f
}

func errorHandler(logger *log.Logger) func(interface{}, ...int) {
	return func(err interface{}, exitCode ...int) {
		if err != nil {
			logger.Println(err)
			if exitHandler != nil {
				exitHandler()
			}
			if len(exitCode) == 1 {
				os.Exit(exitCode[0])
				return
			}
			os.Exit(1)
		}
	}
}

func CreateExitIfError(msg string) func(interface{}, ...int) {
	info := fmt.Sprintf("[%s]: ", msg)
	logger := log.New(os.Stderr, info, 0)
	exitIfError := errorHandler(logger)
	return exitIfError
}
