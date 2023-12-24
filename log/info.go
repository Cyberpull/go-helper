package log

func Info(v ...any) {
	Color(FgCyan, v...)
}

func Infoln(v ...any) {
	Colorln(FgCyan, v...)
}

func Infof(format string, v ...any) {
	Colorf(FgCyan, format, v...)
}

func Infofln(format string, v ...any) {
	Colorfln(FgCyan, format, v...)
}
