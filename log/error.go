package log

func Error(v ...any) {
	Color(FgRed, v...)
}

func Errorf(format string, v ...any) {
	Colorf(FgRed, format, v...)
}

func Errorfln(format string, v ...any) {
	Colorfln(FgRed, format, v...)
}

func Errorln(v ...any) {
	Colorln(FgRed, v...)
}
