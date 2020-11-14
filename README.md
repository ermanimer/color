# color
Go color

![Go](https://github.com/ermanimer/color/workflows/Go/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/ermanimer/color)](https://goreportcard.com/report/github.com/ermanimer/color)

## Features
color creates colored text for Linux terminal.

## Installation
```bash
go get -u github.com/ermanimer/color
```

## Colors
| Constant | Value                   | Description         |
| :------- | :-------------------- | :------------------ |
| Black    | \u001b[30m%s\u001b[0m | Returns black text   |
| Red      | \u001b[31m%s\u001b[0m | Returns red text     |
| Green    | \u001b[32m%s\u001b[0m | Returns green text   |
| Yellow   | \u001b[33m%s\u001b[0m | Returns yellow text  |
| Blue     | \u001b[34m%s\u001b[0m | Returns blue text    |
| Magenta  | \u001b[35m%s\u001b[0m | Returns magenta text |
| Cyan     | \u001b[36m%s\u001b[0m | Returns cyan text    |
| White    | \u001b[37m%s\u001b[0m | Returns white text   |

## Usage
```go
package main

import (
	"fmt"

	"github.com/ermanimer/color"
)

func main() {
	//print black text
	fmt.Println(color.Black("this is a black text"))

	//print formatted black text
	fmt.Println(color.Blackf("this is a %v black text", "formatted"))

	//print red text
	fmt.Println(color.Red("this is a red text"))

	//print formatted red text
	fmt.Println(color.Redf("this is a %v red text", "formatted"))

	//print green text
	fmt.Println(color.Green("this is a green text"))

	//print formatted green text
	fmt.Println(color.Greenf("this is a %v green text", "formatted"))

	//print yellow text
	fmt.Println(color.Yellow("this is a yellow text"))

	//print formatted yellow text
	fmt.Println(color.Yellowf("this is a %v yellow text", "formatted"))

	//print blue text
	fmt.Println(color.Blue("this is a blue text"))

	//print formatted blue text
	fmt.Println(color.Bluef("this is a %v blue text", "formatted"))

	//print magenta text
	fmt.Println(color.Magenta("this is a magenta text"))

	//print formatted magenta text
	fmt.Println(color.Magentaf("this is a %v magenta text", "formatted"))

	//print cyan text
	fmt.Println(color.Cyan("this is a cyan text"))

	//print formatted cyan text
	fmt.Println(color.Cyanf("this is a %v cyan text", "formatted"))

	//print white text
	fmt.Println(color.White("this is a white text"))

	//print formatted white text
	fmt.Println(color.Whitef("this is a %v white text", "formatted"))
}
```

## Terminal Output
![Terminal Output](/images/terminal_output.png)
