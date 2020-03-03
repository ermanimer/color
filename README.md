# color
Go color

[![Build Status](https://travis-ci.org/ermanimer/color.svg?branch=master)](https://travis-ci.org/ermanimer/color)

## Features
color prints colored text on Linux terminal.

## Installation
```bash
go get -u github.com/ermanimer/color
```

## Colors
| Constant | Value                   | Description         |
| :------- | :---------------------- | :------------------ |
| Black    | \u001b[30m%s\u001b[0m%s | Prints black text   |
| Red      | \u001b[31m%s\u001b[0m%s | Prints red text     |
| Green    | \u001b[32m%s\u001b[0m%s | Prints green text   |
| Yellow   | \u001b[33m%s\u001b[0m%s | Prints yellow text  |
| Blue     | \u001b[34m%s\u001b[0m%s | Prints blue text    |
| Magenta  | \u001b[35m%s\u001b[0m%s | Prints magenta text |
| Cyan     | \u001b[36m%s\u001b[0m%s | Prints cyan text    |
| White    | \u001b[37m%s\u001b[0m%s | Prints white text   |

## New Line
| Constant | Value | Description                  |
| :------- | :-----| :--------------------------- |
| newLine  | \n    | New line character for Linux |


## Usage
```go
package main

import "github.com/ermanimer/color"

func main() {
	//print black text
	color.Black("This is a black text.")
	
	//print red text
	color.Red("This is a red text.")
	
	//print green text
	color.Green("This is a green text.")
	

	//print yellow text
	color.Yellow("This is a yellow text.")
	
	//print blue text
	color.Blue("This is a blue text.")
	
	//print magenta text
	color.Magenta("This is a magenta text.")
	
	//print cyan text
	color.Cyan("This is a cyan text.")
	
	//print white text
	color.White("This is a white text.")
}
```

## Terminal Output
![Terminal Output](/images/terminal_output.png)

