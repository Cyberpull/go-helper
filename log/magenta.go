package log

func Magenta(v ...any) {
	Color(FgMagenta, v...)
}

func Magentaln(v ...any) {
	Colorln(FgMagenta, v...)
}

func Magentaf(format string, v ...any) {
	Colorf(FgMagenta, format, v...)
}

func Magentafln(format string, v ...any) {
	Colorfln(FgMagenta, format, v...)
}
