package color

import (
	"fmt"
	"strings"
)

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

func Black(values ...interface{}) string {
	return paint(black, values...)
}

func Blackf(messageFormat string, values ...interface{}) string {
	return paintf(black, messageFormat, values...)
}

func Red(values ...interface{}) string {
	return paint(red, values...)
}

func Redf(messageFormat string, values ...interface{}) string {
	return paintf(red, messageFormat, values...)
}

func Green(values ...interface{}) string {
	return paint(green, values...)
}

func Greenf(messageFormat string, values ...interface{}) string {
	return paintf(green, messageFormat, values...)
}

func Yellow(values ...interface{}) string {
	return paint(yellow, values...)
}

func Yellowf(messageFormat string, values ...interface{}) string {
	return paintf(yellow, messageFormat, values...)
}

func Blue(values ...interface{}) string {
	return paint(blue, values...)
}

func Bluef(messageFormat string, values ...interface{}) string {
	return paintf(blue, messageFormat, values...)
}

func Magenta(values ...interface{}) string {
	return paint(magenta, values...)
}

func Magentaf(messageFormat string, values ...interface{}) string {
	return paintf(magenta, messageFormat, values...)
}

func Cyan(values ...interface{}) string {
	return paint(cyan, values...)
}

func Cyanf(messageFormat string, values ...interface{}) string {
	return paintf(cyan, messageFormat, values...)
}

func White(values ...interface{}) string {
	return paint(white, values...)
}

func Whitef(messageFormat string, values ...interface{}) string {
	return paintf(white, messageFormat, values...)
}

func paint(color string, values ...interface{}) string {
	messageFormat := createMessageFormat(values...)
	message := fmt.Sprintf(messageFormat, values...)
	return fmt.Sprintf(color, message)
}

func paintf(color string, messageFormat string, values ...interface{}) string {
	message := fmt.Sprintf(messageFormat, values...)
	return fmt.Sprintf(color, message)
}

func createMessageFormat(values ...interface{}) string {
	messageFormat := strings.Repeat("%v, ", len(values))
	messageFormat = strings.Trim(messageFormat, ", ")
	return messageFormat
}
