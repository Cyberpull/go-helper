package log

import (
	"fmt"
	"log"
	"runtime"

	"github.com/fatih/color"
)

func Print(v ...any) {
	logger.Print(v...)
}

func Println(v ...any) {
	logger.Println(v...)
}

func Printf(format string, v ...any) {
	logger.Printf(format, v...)
}

func Printfln(format string, v ...any) {
	format = fmt.Sprintf(format, v...)
	logger.Println(format)
}

// =====================================

var logger *log.Logger

func init() {
	def := log.Default()

	logger = log.New(def.Writer(), def.Prefix(), def.Flags())

	if runtime.GOOS == "windows" {
		logger.SetOutput(color.Output)
	}
}
