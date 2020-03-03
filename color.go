package color

import "fmt"

//colors
const (
	black   = "\u001b[30m%s\u001b[0m%s"
	red     = "\u001b[31m%s\u001b[0m%s"
	green   = "\u001b[32m%s\u001b[0m%s"
	yellow  = "\u001b[33m%s\u001b[0m%s"
	blue    = "\u001b[34m%s\u001b[0m%s"
	magenta = "\u001b[35m%s\u001b[0m%s"
	cyan    = "\u001b[36m%s\u001b[0m%s"
	white   = "\u001b[37m%s\u001b[0m%s"
)

//new line character for Linux
const (
	newLine = "\n"
)

func Black(text string) {
	print(black, text)
}

func Red(text string) {
	print(red, text)
}

func Green(text string) {
	print(green, text)
}

func Yellow(text string) {
	print(yellow, text)
}

func Blue(text string) {
	print(blue, text)
}

func Magenta(text string) {
	print(magenta, text)
}

func Cyan(text string) {
	print(cyan, text)
}

func White(text string) {
	print(white, text)
}

func print(color string, text string) {
	fmt.Print(fmt.Sprintf(color, text, newLine))
}
