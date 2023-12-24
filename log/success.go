package log

func Success(v ...any) {
	Color(FgGreen, v...)
}

func Successln(v ...any) {
	Colorln(FgGreen, v...)
}

func Successf(format string, v ...any) {
	Colorf(FgGreen, format, v...)
}

func Successfln(format string, v ...any) {
	Colorfln(FgGreen, format, v...)
}
