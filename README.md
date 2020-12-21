# color
Go color

![Go](https://github.com/ermanimer/color/workflows/Go/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/ermanimer/color)](https://goreportcard.com/report/github.com/ermanimer/color)

## Features
color prints 8-bit colored text for Linux terminal.

## Methods
- Sprint
- Sprintf
- Print
- Printf
- Println
- Printfln
- Fprint
- Fprintf
- Fprintln
- Fprintfln

## Functions
- SprintFunction
- SprintfFunction
- PrintFunction
- PrintfFunction
- PrintlnFunction
- PrintflnFunction
- FprintFunction
- FprintfFunction
- FprintlnFunction
- FprintflnFunction

## Installation
```bash
go get -u github.com/ermanimer/color
```

## Usage
```go
package main

import (
	"os"

	"github.com/ermanimer/color/v2"
)

func main() {
	//create color with struct
	//use standard color names or 8-bit color codes for foreground and background color codes
	color1 := &color.Color{
		Foreground:    color.Blue,
		HasBackground: true,
		Background:    240,
		IsBold:        true,
		IsUnderlined:  false,
		IsReversed:    false,
	}

	//create color with New function
	//func New(foreground byte, hasBackground bool, background byte, isBold bool, isUnderlined bool, isReversed bool) *color.Color
	color2 := color.New(color.Red, true, color.White, false, false, false)

	//use methods
	color1.Print("color 1")

        //use functions
	printColor2 := color2.PrintFunction()
	printColor2("color 2")
}

```

## Standard Colors
|Name|Value|Color|Name|Value|Color|
|:--|:-:|:-:|:-:|:--|:-:|
|Black|0|![#000000](https://singlecolorimage.com/get/000000/24x24)|BrightBlack|8|![#808080](https://singlecolorimage.com/get/808080/24x24)|
|Red|1|![#800000](https://singlecolorimage.com/get/800000/24x24)|BrightRed|9|![#ff0000](https://singlecolorimage.com/get/ff0000/24x24)|
|Green|2|![#008000](https://singlecolorimage.com/get/008000/24x24)|BrightGreen|10|![#00ff00](https://singlecolorimage.com/get/00ff00/24x24)|
|Yelow|3|![#808000](https://singlecolorimage.com/get/808000/24x24)|BrightYellow|11|![#ffff00](https://singlecolorimage.com/get/ffff00/24x24)|
|Blue|4|![#000080](https://singlecolorimage.com/get/000080/24x24)|BrightBlue|12|![#0000ff](https://singlecolorimage.com/get/0000ff/24x24)|
|Magenta|5|![#800080](https://singlecolorimage.com/get/800080/24x24)|BrightMagenta|13|![#ff00ff](https://singlecolorimage.com/get/ff00ff/24x24)|
|Cyan|6|![#008080](https://singlecolorimage.com/get/008080/24x24)|BrightCyan|14|![#00ffff](https://singlecolorimage.com/get/00ffff/24x24)|
|White|7|![#c0c0c0](https://singlecolorimage.com/get/c0c0c0/24x24)|BrightWhite|15|![#ffffff](https://singlecolorimage.com/get/ffffff/24x24)|

## 8-Bit Colors
|Value|Color|Value|Color|Value|Color|Value|Color|
|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
|0|![#000000](https://singlecolorimage.com/get/000000/24x24)|64|![#5f8700](https://singlecolorimage.com/get/5f8700/24x24)|128|![#af00d7](https://singlecolorimage.com/get/af00d7/24x24)|192|![#d7ff87](https://singlecolorimage.com/get/d7ff87/24x24)|
|1|![#800000](https://singlecolorimage.com/get/800000/24x24)|65|![#5f875f](https://singlecolorimage.com/get/5f875f/24x24)|129|![#af00ff](https://singlecolorimage.com/get/af00ff/24x24)|193|![#d7ffaf](https://singlecolorimage.com/get/d7ffaf/24x24)|
|2|![#008000](https://singlecolorimage.com/get/008000/24x24)|66|![#5f8787](https://singlecolorimage.com/get/5f8787/24x24)|130|![#af5f00](https://singlecolorimage.com/get/af5f00/24x24)|194|![#d7ffd7](https://singlecolorimage.com/get/d7ffd7/24x24)|
|3|![#808000](https://singlecolorimage.com/get/808000/24x24)|67|![#5f87af](https://singlecolorimage.com/get/5f87af/24x24)|131|![#af5f5f](https://singlecolorimage.com/get/af5f5f/24x24)|195|![#d7ffff](https://singlecolorimage.com/get/d7ffff/24x24)|
|4|![#000080](https://singlecolorimage.com/get/000080/24x24)|68|![#5f87d7](https://singlecolorimage.com/get/5f87d7/24x24)|132|![#af5f87](https://singlecolorimage.com/get/af5f87/24x24)|196|![#ff0000](https://singlecolorimage.com/get/ff0000/24x24)|
|5|![#800080](https://singlecolorimage.com/get/800080/24x24)|69|![#5f87ff](https://singlecolorimage.com/get/5f87ff/24x24)|133|![#af5faf](https://singlecolorimage.com/get/af5faf/24x24)|197|![#ff005f](https://singlecolorimage.com/get/ff005f/24x24)|
|6|![#008080](https://singlecolorimage.com/get/008080/24x24)|70|![#5faf00](https://singlecolorimage.com/get/5faf00/24x24)|134|![#af5fd7](https://singlecolorimage.com/get/af5fd7/24x24)|198|![#ff0087](https://singlecolorimage.com/get/ff0087/24x24)|
|7|![#c0c0c0](https://singlecolorimage.com/get/c0c0c0/24x24)|71|![#5faf5f](https://singlecolorimage.com/get/5faf5f/24x24)|135|![#af5fff](https://singlecolorimage.com/get/af5fff/24x24)|199|![#ff00af](https://singlecolorimage.com/get/ff00af/24x24)|
|8|![#808080](https://singlecolorimage.com/get/808080/24x24)|72|![#5faf87](https://singlecolorimage.com/get/5faf87/24x24)|136|![#af8700](https://singlecolorimage.com/get/af8700/24x24)|200|![#ff00d7](https://singlecolorimage.com/get/ff00d7/24x24)|
|9|![#ff0000](https://singlecolorimage.com/get/ff0000/24x24)|73|![#5fafaf](https://singlecolorimage.com/get/5fafaf/24x24)|137|![#af875f](https://singlecolorimage.com/get/af875f/24x24)|201|![#ff00ff](https://singlecolorimage.com/get/ff00ff/24x24)|
|10|![#00ff00](https://singlecolorimage.com/get/00ff00/24x24)|74|![#5fafd7](https://singlecolorimage.com/get/5fafd7/24x24)|138|![#af8787](https://singlecolorimage.com/get/af8787/24x24)|202|![#ff5f00](https://singlecolorimage.com/get/ff5f00/24x24)|
|11|![#ffff00](https://singlecolorimage.com/get/ffff00/24x24)|75|![#5fafff](https://singlecolorimage.com/get/5fafff/24x24)|139|![#af87af](https://singlecolorimage.com/get/af87af/24x24)|203|![#ff5f5f](https://singlecolorimage.com/get/ff5f5f/24x24)|
|12|![#0000ff](https://singlecolorimage.com/get/0000ff/24x24)|76|![#5fd700](https://singlecolorimage.com/get/5fd700/24x24)|140|![#af87d7](https://singlecolorimage.com/get/af87d7/24x24)|204|![#ff5f87](https://singlecolorimage.com/get/ff5f87/24x24)|
|13|![#ff00ff](https://singlecolorimage.com/get/ff00ff/24x24)|77|![#5fd75f](https://singlecolorimage.com/get/5fd75f/24x24)|141|![#af87ff](https://singlecolorimage.com/get/af87ff/24x24)|205|![#ff5faf](https://singlecolorimage.com/get/ff5faf/24x24)|
|14|![#00ffff](https://singlecolorimage.com/get/00ffff/24x24)|78|![#5fd787](https://singlecolorimage.com/get/5fd787/24x24)|142|![#afaf00](https://singlecolorimage.com/get/afaf00/24x24)|206|![#ff5fd7](https://singlecolorimage.com/get/ff5fd7/24x24)|
|15|![#ffffff](https://singlecolorimage.com/get/ffffff/24x24)|79|![#5fd7af](https://singlecolorimage.com/get/5fd7af/24x24)|143|![#afaf5f](https://singlecolorimage.com/get/afaf5f/24x24)|207|![#ff5fff](https://singlecolorimage.com/get/ff5fff/24x24)|
|16|![#000000](https://singlecolorimage.com/get/000000/24x24)|80|![#5fd7d7](https://singlecolorimage.com/get/5fd7d7/24x24)|144|![#afaf87](https://singlecolorimage.com/get/afaf87/24x24)|208|![#ff8700](https://singlecolorimage.com/get/ff8700/24x24)|
|17|![#00005f](https://singlecolorimage.com/get/00005f/24x24)|81|![#5fd7ff](https://singlecolorimage.com/get/5fd7ff/24x24)|145|![#afafaf](https://singlecolorimage.com/get/afafaf/24x24)|209|![#ff875f](https://singlecolorimage.com/get/ff875f/24x24)|
|18|![#000087](https://singlecolorimage.com/get/000087/24x24)|82|![#5fff00](https://singlecolorimage.com/get/5fff00/24x24)|146|![#afafd7](https://singlecolorimage.com/get/afafd7/24x24)|210|![#ff8787](https://singlecolorimage.com/get/ff8787/24x24)|
|19|![#0000af](https://singlecolorimage.com/get/0000af/24x24)|83|![#5fff5f](https://singlecolorimage.com/get/5fff5f/24x24)|147|![#afafff](https://singlecolorimage.com/get/afafff/24x24)|211|![#ff87af](https://singlecolorimage.com/get/ff87af/24x24)|
|20|![#0000d7](https://singlecolorimage.com/get/0000d7/24x24)|84|![#5fff87](https://singlecolorimage.com/get/5fff87/24x24)|148|![#afd700](https://singlecolorimage.com/get/afd700/24x24)|212|![#ff87d7](https://singlecolorimage.com/get/ff87d7/24x24)|
|21|![#0000ff](https://singlecolorimage.com/get/0000ff/24x24)|85|![#5fffaf](https://singlecolorimage.com/get/5fffaf/24x24)|149|![#afd75f](https://singlecolorimage.com/get/afd75f/24x24)|213|![#ff87ff](https://singlecolorimage.com/get/ff87ff/24x24)|
|22|![#005f00](https://singlecolorimage.com/get/005f00/24x24)|86|![#5fffd7](https://singlecolorimage.com/get/5fffd7/24x24)|150|![#afd787](https://singlecolorimage.com/get/afd787/24x24)|214|![#ffaf00](https://singlecolorimage.com/get/ffaf00/24x24)|
|23|![#005f5f](https://singlecolorimage.com/get/005f5f/24x24)|87|![#5fffff](https://singlecolorimage.com/get/5fffff/24x24)|151|![#afd7af](https://singlecolorimage.com/get/afd7af/24x24)|215|![#ffaf5f](https://singlecolorimage.com/get/ffaf5f/24x24)|
|24|![#005f87](https://singlecolorimage.com/get/005f87/24x24)|88|![#870000](https://singlecolorimage.com/get/870000/24x24)|152|![#afd7d7](https://singlecolorimage.com/get/afd7d7/24x24)|216|![#ffaf87](https://singlecolorimage.com/get/ffaf87/24x24)|
|25|![#005faf](https://singlecolorimage.com/get/005faf/24x24)|89|![#87005f](https://singlecolorimage.com/get/87005f/24x24)|153|![#afd7ff](https://singlecolorimage.com/get/afd7ff/24x24)|217|![#ffafaf](https://singlecolorimage.com/get/ffafaf/24x24)|
|26|![#005fd7](https://singlecolorimage.com/get/005fd7/24x24)|90|![#870087](https://singlecolorimage.com/get/870087/24x24)|154|![#afff00](https://singlecolorimage.com/get/afff00/24x24)|218|![#ffafd7](https://singlecolorimage.com/get/ffafd7/24x24)|
|27|![#005fff](https://singlecolorimage.com/get/005fff/24x24)|91|![#8700af](https://singlecolorimage.com/get/8700af/24x24)|155|![#afff5f](https://singlecolorimage.com/get/afff5f/24x24)|219|![#ffafff](https://singlecolorimage.com/get/ffafff/24x24)|
|28|![#008700](https://singlecolorimage.com/get/008700/24x24)|92|![#8700d7](https://singlecolorimage.com/get/8700d7/24x24)|156|![#afff87](https://singlecolorimage.com/get/afff87/24x24)|220|![#ffd700](https://singlecolorimage.com/get/ffd700/24x24)|
|29|![#00875f](https://singlecolorimage.com/get/00875f/24x24)|93|![#8700ff](https://singlecolorimage.com/get/8700ff/24x24)|157|![#afffaf](https://singlecolorimage.com/get/afffaf/24x24)|221|![#ffd75f](https://singlecolorimage.com/get/ffd75f/24x24)|
|30|![#008787](https://singlecolorimage.com/get/008787/24x24)|94|![#875f00](https://singlecolorimage.com/get/875f00/24x24)|158|![#afffd7](https://singlecolorimage.com/get/afffd7/24x24)|222|![#ffd787](https://singlecolorimage.com/get/ffd787/24x24)|
|31|![#0087af](https://singlecolorimage.com/get/0087af/24x24)|95|![#875f5f](https://singlecolorimage.com/get/875f5f/24x24)|159|![#afffff](https://singlecolorimage.com/get/afffff/24x24)|223|![#ffd7af](https://singlecolorimage.com/get/ffd7af/24x24)|
|32|![#0087d7](https://singlecolorimage.com/get/0087d7/24x24)|96|![#875f87](https://singlecolorimage.com/get/875f87/24x24)|160|![#d70000](https://singlecolorimage.com/get/d70000/24x24)|224|![#ffd7d7](https://singlecolorimage.com/get/ffd7d7/24x24)|
|33|![#0087ff](https://singlecolorimage.com/get/0087ff/24x24)|97|![#875faf](https://singlecolorimage.com/get/875faf/24x24)|161|![#d7005f](https://singlecolorimage.com/get/d7005f/24x24)|225|![#ffd7ff](https://singlecolorimage.com/get/ffd7ff/24x24)|
|34|![#00af00](https://singlecolorimage.com/get/00af00/24x24)|98|![#875fd7](https://singlecolorimage.com/get/875fd7/24x24)|162|![#d70087](https://singlecolorimage.com/get/d70087/24x24)|226|![#ffff00](https://singlecolorimage.com/get/ffff00/24x24)|
|35|![#00af5f](https://singlecolorimage.com/get/00af5f/24x24)|99|![#875fff](https://singlecolorimage.com/get/875fff/24x24)|163|![#d700af](https://singlecolorimage.com/get/d700af/24x24)|227|![#ffff5f](https://singlecolorimage.com/get/ffff5f/24x24)|
|36|![#00af87](https://singlecolorimage.com/get/00af87/24x24)|100|![#878700](https://singlecolorimage.com/get/878700/24x24)|164|![#d700d7](https://singlecolorimage.com/get/d700d7/24x24)|228|![#ffff87](https://singlecolorimage.com/get/ffff87/24x24)|
|37|![#00afaf](https://singlecolorimage.com/get/00afaf/24x24)|101|![#87875f](https://singlecolorimage.com/get/87875f/24x24)|165|![#d700ff](https://singlecolorimage.com/get/d700ff/24x24)|229|![#ffffaf](https://singlecolorimage.com/get/ffffaf/24x24)|
|38|![#00afd7](https://singlecolorimage.com/get/00afd7/24x24)|102|![#878787](https://singlecolorimage.com/get/878787/24x24)|166|![#d75f00](https://singlecolorimage.com/get/d75f00/24x24)|230|![#ffffd7](https://singlecolorimage.com/get/ffffd7/24x24)|
|39|![#00afff](https://singlecolorimage.com/get/00afff/24x24)|103|![#8787af](https://singlecolorimage.com/get/8787af/24x24)|167|![#d75f5f](https://singlecolorimage.com/get/d75f5f/24x24)|231|![#ffffff](https://singlecolorimage.com/get/ffffff/24x24)|
|40|![#00d700](https://singlecolorimage.com/get/00d700/24x24)|104|![#8787d7](https://singlecolorimage.com/get/8787d7/24x24)|168|![#d75f87](https://singlecolorimage.com/get/d75f87/24x24)|232|![#080808](https://singlecolorimage.com/get/080808/24x24)|
|41|![#00d75f](https://singlecolorimage.com/get/00d75f/24x24)|105|![#8787ff](https://singlecolorimage.com/get/8787ff/24x24)|169|![#d75faf](https://singlecolorimage.com/get/d75faf/24x24)|233|![#121212](https://singlecolorimage.com/get/121212/24x24)|
|42|![#00d787](https://singlecolorimage.com/get/00d787/24x24)|106|![#87af00](https://singlecolorimage.com/get/87af00/24x24)|170|![#d75fd7](https://singlecolorimage.com/get/d75fd7/24x24)|234|![#1c1c1c](https://singlecolorimage.com/get/1c1c1c/24x24)|
|43|![#00d7af](https://singlecolorimage.com/get/00d7af/24x24)|107|![#87af5f](https://singlecolorimage.com/get/87af5f/24x24)|171|![#d75fff](https://singlecolorimage.com/get/d75fff/24x24)|235|![#262626](https://singlecolorimage.com/get/262626/24x24)|
|44|![#00d7d7](https://singlecolorimage.com/get/00d7d7/24x24)|108|![#87af87](https://singlecolorimage.com/get/87af87/24x24)|172|![#d78700](https://singlecolorimage.com/get/d78700/24x24)|236|![#303030](https://singlecolorimage.com/get/303030/24x24)|
|45|![#00d7ff](https://singlecolorimage.com/get/00d7ff/24x24)|109|![#87afaf](https://singlecolorimage.com/get/87afaf/24x24)|173|![#d7875f](https://singlecolorimage.com/get/d7875f/24x24)|237|![#3a3a3a](https://singlecolorimage.com/get/3a3a3a/24x24)|
|46|![#00ff00](https://singlecolorimage.com/get/00ff00/24x24)|110|![#87afd7](https://singlecolorimage.com/get/87afd7/24x24)|174|![#d78787](https://singlecolorimage.com/get/d78787/24x24)|238|![#444444](https://singlecolorimage.com/get/444444/24x24)|
|47|![#00ff5f](https://singlecolorimage.com/get/00ff5f/24x24)|111|![#87afff](https://singlecolorimage.com/get/87afff/24x24)|175|![#d787af](https://singlecolorimage.com/get/d787af/24x24)|239|![#4e4e4e](https://singlecolorimage.com/get/4e4e4e/24x24)|
|48|![#00ff87](https://singlecolorimage.com/get/00ff87/24x24)|112|![#87d700](https://singlecolorimage.com/get/87d700/24x24)|176|![#d787d7](https://singlecolorimage.com/get/d787d7/24x24)|240|![#585858](https://singlecolorimage.com/get/585858/24x24)|
|49|![#00ffaf](https://singlecolorimage.com/get/00ffaf/24x24)|113|![#87d75f](https://singlecolorimage.com/get/87d75f/24x24)|177|![#d787ff](https://singlecolorimage.com/get/d787ff/24x24)|241|![#626262](https://singlecolorimage.com/get/626262/24x24)|
|50|![#00ffd7](https://singlecolorimage.com/get/00ffd7/24x24)|114|![#87d787](https://singlecolorimage.com/get/87d787/24x24)|178|![#d7af00](https://singlecolorimage.com/get/d7af00/24x24)|242|![#6c6c6c](https://singlecolorimage.com/get/6c6c6c/24x24)|
|51|![#00ffff](https://singlecolorimage.com/get/00ffff/24x24)|115|![#87d7af](https://singlecolorimage.com/get/87d7af/24x24)|179|![#d7af5f](https://singlecolorimage.com/get/d7af5f/24x24)|243|![#767676](https://singlecolorimage.com/get/767676/24x24)|
|52|![#5f0000](https://singlecolorimage.com/get/5f0000/24x24)|116|![#87d7d7](https://singlecolorimage.com/get/87d7d7/24x24)|180|![#d7af87](https://singlecolorimage.com/get/d7af87/24x24)|244|![#808080](https://singlecolorimage.com/get/808080/24x24)|
|53|![#5f005f](https://singlecolorimage.com/get/5f005f/24x24)|117|![#87d7ff](https://singlecolorimage.com/get/87d7ff/24x24)|181|![#d7afaf](https://singlecolorimage.com/get/d7afaf/24x24)|245|![#8a8a8a](https://singlecolorimage.com/get/8a8a8a/24x24)|
|54|![#5f0087](https://singlecolorimage.com/get/5f0087/24x24)|118|![#87ff00](https://singlecolorimage.com/get/87ff00/24x24)|182|![#d7afd7](https://singlecolorimage.com/get/d7afd7/24x24)|246|![#949494](https://singlecolorimage.com/get/949494/24x24)|
|55|![#5f00af](https://singlecolorimage.com/get/5f00af/24x24)|119|![#87ff5f](https://singlecolorimage.com/get/87ff5f/24x24)|183|![#d7afff](https://singlecolorimage.com/get/d7afff/24x24)|247|![#9e9e9e](https://singlecolorimage.com/get/9e9e9e/24x24)|
|56|![#5f00d7](https://singlecolorimage.com/get/5f00d7/24x24)|120|![#87ff87](https://singlecolorimage.com/get/87ff87/24x24)|184|![#d7d700](https://singlecolorimage.com/get/d7d700/24x24)|248|![#a8a8a8](https://singlecolorimage.com/get/a8a8a8/24x24)|
|57|![#5f00ff](https://singlecolorimage.com/get/5f00ff/24x24)|121|![#87ffaf](https://singlecolorimage.com/get/87ffaf/24x24)|185|![#d7d75f](https://singlecolorimage.com/get/d7d75f/24x24)|249|![#b2b2b2](https://singlecolorimage.com/get/b2b2b2/24x24)|
|58|![#5f5f00](https://singlecolorimage.com/get/5f5f00/24x24)|122|![#87ffd7](https://singlecolorimage.com/get/87ffd7/24x24)|186|![#d7d787](https://singlecolorimage.com/get/d7d787/24x24)|250|![#bcbcbc](https://singlecolorimage.com/get/bcbcbc/24x24)|
|59|![#5f5f5f](https://singlecolorimage.com/get/5f5f5f/24x24)|123|![#87ffff](https://singlecolorimage.com/get/87ffff/24x24)|187|![#d7d7af](https://singlecolorimage.com/get/d7d7af/24x24)|251|![#c6c6c6](https://singlecolorimage.com/get/c6c6c6/24x24)|
|60|![#5f5f87](https://singlecolorimage.com/get/5f5f87/24x24)|124|![#af0000](https://singlecolorimage.com/get/af0000/24x24)|188|![#d7d7d7](https://singlecolorimage.com/get/d7d7d7/24x24)|252|![#d0d0d0](https://singlecolorimage.com/get/d0d0d0/24x24)|
|61|![#5f5faf](https://singlecolorimage.com/get/5f5faf/24x24)|125|![#af005f](https://singlecolorimage.com/get/af005f/24x24)|189|![#d7d7ff](https://singlecolorimage.com/get/d7d7ff/24x24)|253|![#dadada](https://singlecolorimage.com/get/dadada/24x24)|
|62|![#5f5fd7](https://singlecolorimage.com/get/5f5fd7/24x24)|126|![#af0087](https://singlecolorimage.com/get/af0087/24x24)|190|![#d7ff00](https://singlecolorimage.com/get/d7ff00/24x24)|254|![#e4e4e4](https://singlecolorimage.com/get/e4e4e4/24x24)|
|63|![#5f5fff](https://singlecolorimage.com/get/5f5fff/24x24)|127|![#af00af](https://singlecolorimage.com/get/af00af/24x24)|191|![#d7ff5f](https://singlecolorimage.com/get/d7ff5f/24x24)|255|![#eeeeee](https://singlecolorimage.com/get/eeeeee/24x24)|

## References
- [ANSI Escape Codes, Wikipedia](https://en.wikipedia.org/wiki/ANSI_escape_code)
- [Build Your Own Command Line With ANSI Escape Codes, Haoyi's Programming Blog](https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html)
