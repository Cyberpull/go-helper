package log

import (
	"log"
)

func Fatal(v ...any) {
	log.Fatal(ColorString(FgRed, v...))
}

func Fatalln(v ...any) {
	log.Fatalln(ColorString(FgRed, v...))
}

func Fatalf(format string, v ...any) {
	log.Fatalf(ColorStringF(FgRed, format, v...))
}

func Fatalfln(format string, v ...any) {
	log.Fatalln(ColorStringF(FgRed, format, v...))
}
