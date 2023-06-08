package log

import (
	"fmt"
	"log"
)

func Print(v ...any) {
	log.Print(v...)
}

func Printf(format string, v ...any) {
	log.Printf(format, v...)
}

func Printfln(format string, v ...any) {
	format = fmt.Sprintf(format, v...)
	log.Println(format)
}

func Println(v ...any) {
	log.Println(v...)
}
