package color

import (
	"bytes"
	"fmt"
	"io"
)

//Standard Colors
const (
	Black         byte = 0
	Red           byte = 1
	Green         byte = 2
	Yellow        byte = 3
	Blue          byte = 4
	Magenta       byte = 5
	Cyan          byte = 6
	White         byte = 7
	BrightBlack   byte = 8
	BrightRed     byte = 9
	BrightGreen   byte = 10
	BrightYellow  byte = 11
	BrightBlue    byte = 12
	BrightMagenta byte = 13
	BrightCyan    byte = 14
	BrightWhite   byte = 15
)

//ANSI Escape Codes
const (
	ecForegroundColor = "\u001b[38;5;%dm"
	ecBackgroundColor = "\u001b[48;5;%dm"
	ecBold            = "\u001b[1m"
	ecUnderline       = "\u001b[4m"
	ecReversed        = "\u001b[7m"
	ecReset           = "\u001b[0m"
)

type Color struct {
	Foreground    byte
	HasBackground bool
	Background    byte
	IsBold        bool
	IsUnderlined  bool
	IsReversed    bool
}

func New(foreground byte, hasBackground bool, background byte, isBold bool, isUnderlined bool, isReversed bool) *Color {
	return &Color{
		Foreground:    foreground,
		HasBackground: hasBackground,
		Background:    background,
		IsBold:        isBold,
		IsUnderlined:  isUnderlined,
		IsReversed:    isReversed,
	}
}

func (c *Color) Sprint(values ...interface{}) string {
	text := fmt.Sprint(values...)
	return c.paint(text)
}

func (c *Color) Sprintf(format string, values ...interface{}) string {
	text := fmt.Sprintf(format, values...)
	return c.paint(text)
}

func (c *Color) Print(values ...interface{}) (n int, err error) {
	text := fmt.Sprint(values...)
	return fmt.Print(c.paint(text))
}

func (c *Color) Printf(format string, values ...interface{}) (n int, err error) {
	text := fmt.Sprintf(format, values...)
	return fmt.Print(c.paint(text))
}

func (c *Color) Println(values ...interface{}) (n int, err error) {
	text := fmt.Sprint(values...)
	return fmt.Println(c.paint(text))
}

func (c *Color) Printfln(format string, values ...interface{}) (n int, err error) {
	text := fmt.Sprintf(format, values...)
	return fmt.Println(c.paint(text))
}

func (c *Color) Fprint(writer io.Writer, values ...interface{}) (n int, err error) {
	text := fmt.Sprint(values...)
	return fmt.Fprint(writer, c.paint(text))
}

func (c *Color) Fprintf(writer io.Writer, format string, values ...interface{}) (n int, err error) {
	text := fmt.Sprintf(format, values...)
	return fmt.Fprint(writer, c.paint(text))
}

func (c *Color) Fprintln(writer io.Writer, values ...interface{}) (n int, err error) {
	text := fmt.Sprint(values...)
	return fmt.Fprintln(writer, c.paint(text))
}

func (c *Color) Fprintfln(writer io.Writer, format string, values ...interface{}) (n int, err error) {
	text := fmt.Sprintf(format, values...)
	return fmt.Fprintln(writer, c.paint(text))
}

func (c *Color) SprintFunction() func(values ...interface{}) {
	return func(values ...interface{}) {
		c.Sprint(values...)
	}
}

func (c *Color) SprintfFunction() func(format string, values ...interface{}) {
	return func(format string, values ...interface{}) {
		c.Sprintf(format, values...)
	}
}

func (c *Color) PrintFunction() func(values ...interface{}) {
	return func(values ...interface{}) {
		c.Print(values...)
	}
}

func (c *Color) PrintfFunction() func(format string, values ...interface{}) {
	return func(format string, values ...interface{}) {
		c.Printf(format, values...)
	}
}

func (c *Color) PrintlnFunction() func(values ...interface{}) {
	return func(values ...interface{}) {
		c.Println(values...)
	}
}

func (c *Color) PrintflnFunction() func(format string, values ...interface{}) {
	return func(format string, values ...interface{}) {
		c.Printfln(format, values...)
	}
}

func (c *Color) FprintFunction() func(writer io.Writer, values ...interface{}) {
	return func(writer io.Writer, values ...interface{}) {
		c.Fprint(writer, values...)
	}
}

func (c *Color) FprintfFunction() func(writer io.Writer, format string, values ...interface{}) {
	return func(writer io.Writer, format string, values ...interface{}) {
		c.Fprintf(writer, format, values...)
	}
}

func (c *Color) FprintlnFunction() func(writer io.Writer, values ...interface{}) {
	return func(writer io.Writer, values ...interface{}) {
		c.Fprintln(writer, values...)
	}
}

func (c *Color) FprintflnFunction() func(writer io.Writer, format string, values ...interface{}) {
	return func(writer io.Writer, format string, values ...interface{}) {
		c.Fprintfln(writer, format, values...)
	}
}

func (c *Color) paint(text string) string {
	//create buffer
	b := new(bytes.Buffer)
	//add foreground color
	foreGround := fmt.Sprintf(ecForegroundColor, c.Foreground)
	b.WriteString(foreGround)
	//add background color
	if c.HasBackground {
		backGround := fmt.Sprintf(ecBackgroundColor, c.Background)
		b.WriteString(backGround)
	}
	//add bold decoration
	if c.IsBold {
		b.WriteString(ecBold)
	}
	//add underline decoration
	if c.IsUnderlined {
		b.WriteString(ecUnderline)
	}
	//add reversed decoration
	if c.IsReversed {
		b.WriteString(ecReversed)
	}
	//add text
	b.WriteString(text)
	//add reset escape code
	b.WriteString(ecReset)
	return b.String()
}
