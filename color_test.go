package color

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

const (
	testSentence = "this is sentence has random colors and random text decorations for visual test"
)

func TestSprintFunction(t *testing.T) {
	foreground := getRandomColor()
	hasBackGround := getRandomBool()
	background := getRandomColor()
	isBold := getRandomBool()
	isUnderlined := getRandomBool()
	isReversed := getRandomBool()
	color := New(foreground, hasBackGround, background, isBold, isUnderlined, isReversed)
	function := color.SprintFunction()
	fmt.Fprintf(os.Stdout, "foreground: %s, hasBackground: %t, background: %s, isBold: %t, isUnderlined: %t, isReversed: %t\n", getColorName(foreground), hasBackGround, getColorName(background), isBold, isUnderlined, isReversed)
	fmt.Fprintln(os.Stdout, function(testSentence))
}

func TestSprintfFunction(t *testing.T) {
	foreground := getRandomColor()
	hasBackGround := getRandomBool()
	background := getRandomColor()
	isBold := getRandomBool()
	isUnderlined := getRandomBool()
	isReversed := getRandomBool()
	color := New(foreground, hasBackGround, background, isBold, isUnderlined, isReversed)
	function := color.SprintfFunction()
	fmt.Fprintf(os.Stdout, "foreground: %s, hasBackground: %t, background: %s, isBold: %t, isUnderlined: %t, isReversed: %t\n", getColorName(foreground), hasBackGround, getColorName(background), isBold, isUnderlined, isReversed)
	fmt.Fprintln(os.Stdout, function("%s", testSentence))
}

func TestPrintFunction(t *testing.T) {
	foreground := getRandomColor()
	hasBackGround := getRandomBool()
	background := getRandomColor()
	isBold := getRandomBool()
	isUnderlined := getRandomBool()
	isReversed := getRandomBool()
	color := New(foreground, hasBackGround, background, isBold, isUnderlined, isReversed)
	function := color.PrintFunction()
	fmt.Fprintf(os.Stdout, "foreground: %s, hasBackground: %t, background: %s, isBold: %t, isUnderlined: %t, isReversed: %t\n", getColorName(foreground), hasBackGround, getColorName(background), isBold, isUnderlined, isReversed)
	function(testSentence)
	fmt.Fprint(os.Stdout, "\n")
}

func TestPrintfFunction(t *testing.T) {
	foreground := getRandomColor()
	hasBackGround := getRandomBool()
	background := getRandomColor()
	isBold := getRandomBool()
	isUnderlined := getRandomBool()
	isReversed := getRandomBool()
	color := New(foreground, hasBackGround, background, isBold, isUnderlined, isReversed)
	function := color.PrintfFunction()
	fmt.Fprintf(os.Stdout, "foreground: %s, hasBackground: %t, background: %s, isBold: %t, isUnderlined: %t, isReversed: %t\n", getColorName(foreground), hasBackGround, getColorName(background), isBold, isUnderlined, isReversed)
	function("%s", testSentence)
	fmt.Fprint(os.Stdout, "\n")
}

func TestPrintlnFunction(t *testing.T) {
	foreground := getRandomColor()
	hasBackGround := getRandomBool()
	background := getRandomColor()
	isBold := getRandomBool()
	isUnderlined := getRandomBool()
	isReversed := getRandomBool()
	color := New(foreground, hasBackGround, background, isBold, isUnderlined, isReversed)
	function := color.PrintlnFunction()
	fmt.Fprintf(os.Stdout, "foreground: %s, hasBackground: %t, background: %s, isBold: %t, isUnderlined: %t, isReversed: %t\n", getColorName(foreground), hasBackGround, getColorName(background), isBold, isUnderlined, isReversed)
	function(testSentence)
}

func TestPrintflnFunction(t *testing.T) {
	foreground := getRandomColor()
	hasBackGround := getRandomBool()
	background := getRandomColor()
	isBold := getRandomBool()
	isUnderlined := getRandomBool()
	isReversed := getRandomBool()
	color := New(foreground, hasBackGround, background, isBold, isUnderlined, isReversed)
	function := color.PrintflnFunction()
	fmt.Fprintf(os.Stdout, "foreground: %s, hasBackground: %t, background: %s, isBold: %t, isUnderlined: %t, isReversed: %t\n", getColorName(foreground), hasBackGround, getColorName(background), isBold, isUnderlined, isReversed)
	function("%s", testSentence)
	fmt.Fprint(os.Stdout, "\n")
}

func TestFprintFunction(t *testing.T) {
	foreground := White
	hasBackGround := getRandomBool()
	background := Yellow
	isBold := getRandomBool()
	isUnderlined := getRandomBool()
	isReversed := getRandomBool()
	color := New(foreground, hasBackGround, background, isBold, isUnderlined, isReversed)
	function := color.FprintFunction()
	fmt.Fprintf(os.Stdout, "foreground: %s, hasBackground: %t, background: %s, isBold: %t, isUnderlined: %t, isReversed: %t\n", getColorName(foreground), hasBackGround, getColorName(background), isBold, isUnderlined, isReversed)
	function(os.Stdout, testSentence)
	fmt.Fprint(os.Stdout, "\n")
}

func TestFprintfFunction(t *testing.T) {
	foreground := getRandomColor()
	hasBackGround := getRandomBool()
	background := getRandomColor()
	isBold := getRandomBool()
	isUnderlined := getRandomBool()
	isReversed := getRandomBool()
	color := New(foreground, hasBackGround, background, isBold, isUnderlined, isReversed)
	function := color.FprintfFunction()
	fmt.Fprintf(os.Stdout, "foreground: %s, hasBackground: %t, background: %s, isBold: %t, isUnderlined: %t, isReversed: %t\n", getColorName(foreground), hasBackGround, getColorName(background), isBold, isUnderlined, isReversed)
	function(os.Stdout, "%s", testSentence)
	fmt.Fprint(os.Stdout, "\n")
}

func TestFprintlnFunction(t *testing.T) {
	foreground := getRandomColor()
	hasBackGround := getRandomBool()
	background := getRandomColor()
	isBold := getRandomBool()
	isUnderlined := getRandomBool()
	isReversed := getRandomBool()
	color := New(foreground, hasBackGround, background, isBold, isUnderlined, isReversed)
	function := color.FprintlnFunction()
	fmt.Fprintf(os.Stdout, "foreground: %s, hasBackground: %t, background: %s, isBold: %t, isUnderlined: %t, isReversed: %t\n", getColorName(foreground), hasBackGround, getColorName(background), isBold, isUnderlined, isReversed)
	function(os.Stdout, testSentence)
}

func TestFprintflnFunction(t *testing.T) {
	foreground := getRandomColor()
	hasBackGround := getRandomBool()
	background := getRandomColor()
	isBold := getRandomBool()
	isUnderlined := getRandomBool()
	isReversed := getRandomBool()
	color := New(foreground, hasBackGround, background, isBold, isUnderlined, isReversed)
	function := color.FprintflnFunction()
	fmt.Fprintf(os.Stdout, "foreground: %s, hasBackground: %t, background: %s, isBold: %t, isUnderlined: %t, isReversed: %t\n", getColorName(foreground), hasBackGround, getColorName(background), isBold, isUnderlined, isReversed)
	function(os.Stdout, "%s", testSentence)
	fmt.Fprint(os.Stdout, "\n")
}

func getRandomColor() byte {
	rand.Seed(time.Now().UnixNano())
	return byte(rand.Intn(16))
}

func getRandomBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func getColorName(code byte) string {
	switch code {
	case 0:
		return "Black"
	case 1:
		return "Red"
	case 2:
		return "Green"
	case 3:
		return "Yelow"
	case 4:
		return "Blue"
	case 5:
		return "Magenta"
	case 6:
		return "Cyan"
	case 7:
		return "White"
	case 8:
		return "BrightBlack"
	case 9:
		return "BrightRed"
	case 10:
		return "BrightGreen"
	case 11:
		return "BrightYellow"
	case 12:
		return "BrightBlue"
	case 13:
		return "BrightMagenta"
	case 14:
		return "BrightCyan"
	case 15:
		return "BrightWhite"
	}
	return "unknown color"
}
