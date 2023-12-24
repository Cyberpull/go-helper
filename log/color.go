package log

import (
	"github.com/fatih/color"
)

type ColorAttribute color.Attribute

const (
	FgGreen   ColorAttribute = ColorAttribute(color.FgGreen)
	FgRed     ColorAttribute = ColorAttribute(color.FgRed)
	FgCyan    ColorAttribute = ColorAttribute(color.FgCyan)
	FgMagenta ColorAttribute = ColorAttribute(color.FgMagenta)
	FgBlack   ColorAttribute = ColorAttribute(color.FgBlack)
	FgBlue    ColorAttribute = ColorAttribute(color.FgBlue)
)

func ColorString(a ColorAttribute, v ...any) string {
	c := color.New(color.Attribute(a))
	return c.SprintFunc()(v...)
}

func ColorStringF(a ColorAttribute, format string, v ...any) string {
	c := color.New(color.Attribute(a))
	return c.SprintfFunc()(format, v...)
}

// ======================================

func Color(a ColorAttribute, v ...any) {
	Print(ColorString(a, v...))
}

func Colorln(a ColorAttribute, v ...any) {
	Println(ColorString(a, v...))
}

func Colorf(a ColorAttribute, format string, v ...any) {
	Printf(ColorStringF(a, format, v...))
}

func Colorfln(a ColorAttribute, format string, v ...any) {
	Printfln(ColorStringF(a, format, v...))
}
