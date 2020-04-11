package color

import "fmt"

//colors
const (
	black   = "\u001b[30m%s\u001b[0m"
	red     = "\u001b[31m%s\u001b[0m"
	green   = "\u001b[32m%s\u001b[0m"
	yellow  = "\u001b[33m%s\u001b[0m"
	blue    = "\u001b[34m%s\u001b[0m"
	magenta = "\u001b[35m%s\u001b[0m"
	cyan    = "\u001b[36m%s\u001b[0m"
	white   = "\u001b[37m%s\u001b[0m"
)

func Black(text string) string {
	return fmt.Sprintf(black, text)
}

func Red(text string) string {
	return fmt.Sprintf(red, text)
}

func Green(text string) string {
	return fmt.Sprintf(green, text)
}

func Yellow(text string) string {
	return fmt.Sprintf(yellow, text)
}

func Blue(text string) string {
	return fmt.Sprintf(blue, text)
}

func Magenta(text string) string {
	return fmt.Sprintf(magenta, text)
}

func Cyan(text string) string {
	return fmt.Sprintf(cyan, text)
}

func White(text string) string {
	return fmt.Sprintf(white, text)
}
