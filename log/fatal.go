package log

func Fatal(v ...any) {
	logger.Fatal(ColorString(FgRed, v...))
}

func Fatalln(v ...any) {
	logger.Fatalln(ColorString(FgRed, v...))
}

func Fatalf(format string, v ...any) {
	logger.Fatalf(ColorStringF(FgRed, format, v...))
}

func Fatalfln(format string, v ...any) {
	logger.Fatalln(ColorStringF(FgRed, format, v...))
}
