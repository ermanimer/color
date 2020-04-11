# color
Go color

![Go](https://github.com/ermanimer/color/workflows/Go/badge.svg?branch=master)

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
	fmt.Println(color.Black("This is a black text."))

	//print red text
	fmt.Println(color.Red("This is a red text."))

	//print green text
	fmt.Println(color.Green("This is a green text."))

	//print yellow text
	fmt.Println(color.Yellow("This is a yellow text."))

	//print blue text
	fmt.Println(color.Blue("This is a blue text."))

	//print magenta text
	fmt.Println(color.Magenta("This is a magenta text."))

	//print cyan text
	fmt.Println(color.Cyan("This is a cyan text."))

	//print white text
	fmt.Println(color.White("This is a white text."))
}
```

## Terminal Output
![Terminal Output](/images/terminal_output.png)