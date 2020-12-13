# color
Go color

![Go](https://github.com/ermanimer/color/workflows/Go/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/ermanimer/color)](https://goreportcard.com/report/github.com/ermanimer/color)

## Features
color prints 8-bit colored text for Linux terminal.

## Installation
```bash
go get -u github.com/ermanimer/color/v2
```

## Usage
```go
package main

import (
	"fmt"

	"github.com/ermanimer/color/v2"
)

func main() {
	//...
}
```

## Standard Colors
|Name|Value|Color|Hex Code|Name|Value|Color|Hex Code|
|:--|:-:|:-:|:-:|:--|:-:|:-:|:-:|
|Black|0|![#000000](https://singlecolorimage.com/get/000000/24x24)|#000000|BrightBlack|8|![#808080](https://singlecolorimage.com/get/808080/24x24)|#808080|
|Red|1|![#800000](https://singlecolorimage.com/get/800000/24x24)|#800000|BrightRed|9|![#ff0000](https://singlecolorimage.com/get/ff0000/24x24)|#ff0000|
|Green|2|![#008000](https://singlecolorimage.com/get/008000/24x24)|#008000|BrightGreen|10|![#00ff00](https://singlecolorimage.com/get/00ff00/24x24)|#00ff00|
|Yelow|3|![#808000](https://singlecolorimage.com/get/808000/24x24)|#808000|BrightYellow|11|![#ffff00](https://singlecolorimage.com/get/ffff00/24x24)|#ffff00|
|Blue|4|![#000080](https://singlecolorimage.com/get/000080/24x24)|#000080|BrightBlue|12|![#0000ff](https://singlecolorimage.com/get/0000ff/24x24)|#0000ff|
|Magenta|5|![#800080](https://singlecolorimage.com/get/800080/24x24)|#800080|BrightMagenta|13|![#ff00ff](https://singlecolorimage.com/get/ff00ff/24x24)|#ff00ff|
|Cyan|6|![#008080](https://singlecolorimage.com/get/008080/24x24)|#008080|BrightCyan|14|![#00ffff](https://singlecolorimage.com/get/00ffff/24x24)|#00ffff|
|White|7|![#c0c0c0](https://singlecolorimage.com/get/c0c0c0/24x24)|#c0c0c0|BrightWhite|15|![#ffffff](https://singlecolorimage.com/get/ffffff/24x24)|#ffffff|

## All 8-Bit Colors
|Value|Color|Hex Code|Value|Color|Hex Code|Value|Color|Hex Code|Value|Color|Hex Code|
|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
|0|![#000000](https://singlecolorimage.com/get/000000/24x24)|#000000|64|![#5f8700](https://singlecolorimage.com/get/5f8700/24x24)|#5f8700|128|![#af00d7](https://singlecolorimage.com/get/af00d7/24x24)|#af00d7|192|![#d7ff87](https://singlecolorimage.com/get/d7ff87/24x24)|#d7ff87|
|1|![#800000](https://singlecolorimage.com/get/800000/24x24)|#800000|65|![#5f875f](https://singlecolorimage.com/get/5f875f/24x24)|#5f875f|129|![#af00ff](https://singlecolorimage.com/get/af00ff/24x24)|#af00ff|193|![#d7ffaf](https://singlecolorimage.com/get/d7ffaf/24x24)|#d7ffaf|
|2|![#008000](https://singlecolorimage.com/get/008000/24x24)|#008000|66|![#5f8787](https://singlecolorimage.com/get/5f8787/24x24)|#5f8787|130|![#af5f00](https://singlecolorimage.com/get/af5f00/24x24)|#af5f00|194|![#d7ffd7](https://singlecolorimage.com/get/d7ffd7/24x24)|#d7ffd7|
|3|![#808000](https://singlecolorimage.com/get/808000/24x24)|#808000|67|![#5f87af](https://singlecolorimage.com/get/5f87af/24x24)|#5f87af|131|![#af5f5f](https://singlecolorimage.com/get/af5f5f/24x24)|#af5f5f|195|![#d7ffff](https://singlecolorimage.com/get/d7ffff/24x24)|#d7ffff|
|4|![#000080](https://singlecolorimage.com/get/000080/24x24)|#000080|68|![#5f87d7](https://singlecolorimage.com/get/5f87d7/24x24)|#5f87d7|132|![#af5f87](https://singlecolorimage.com/get/af5f87/24x24)|#af5f87|196|![#ff0000](https://singlecolorimage.com/get/ff0000/24x24)|#ff0000|
|5|![#800080](https://singlecolorimage.com/get/800080/24x24)|#800080|69|![#5f87ff](https://singlecolorimage.com/get/5f87ff/24x24)|#5f87ff|133|![#af5faf](https://singlecolorimage.com/get/af5faf/24x24)|#af5faf|197|![#ff005f](https://singlecolorimage.com/get/ff005f/24x24)|#ff005f|
|6|![#008080](https://singlecolorimage.com/get/008080/24x24)|#008080|70|![#5faf00](https://singlecolorimage.com/get/5faf00/24x24)|#5faf00|134|![#af5fd7](https://singlecolorimage.com/get/af5fd7/24x24)|#af5fd7|198|![#ff0087](https://singlecolorimage.com/get/ff0087/24x24)|#ff0087|
|7|![#c0c0c0](https://singlecolorimage.com/get/c0c0c0/24x24)|#c0c0c0|71|![#5faf5f](https://singlecolorimage.com/get/5faf5f/24x24)|#5faf5f|135|![#af5fff](https://singlecolorimage.com/get/af5fff/24x24)|#af5fff|199|![#ff00af](https://singlecolorimage.com/get/ff00af/24x24)|#ff00af|
|8|![#808080](https://singlecolorimage.com/get/808080/24x24)|#808080|72|![#5faf87](https://singlecolorimage.com/get/5faf87/24x24)|#5faf87|136|![#af8700](https://singlecolorimage.com/get/af8700/24x24)|#af8700|200|![#ff00d7](https://singlecolorimage.com/get/ff00d7/24x24)|#ff00d7|
|9|![#ff0000](https://singlecolorimage.com/get/ff0000/24x24)|#ff0000|73|![#5fafaf](https://singlecolorimage.com/get/5fafaf/24x24)|#5fafaf|137|![#af875f](https://singlecolorimage.com/get/af875f/24x24)|#af875f|201|![#ff00ff](https://singlecolorimage.com/get/ff00ff/24x24)|#ff00ff|
|10|![#00ff00](https://singlecolorimage.com/get/00ff00/24x24)|#00ff00|74|![#5fafd7](https://singlecolorimage.com/get/5fafd7/24x24)|#5fafd7|138|![#af8787](https://singlecolorimage.com/get/af8787/24x24)|#af8787|202|![#ff5f00](https://singlecolorimage.com/get/ff5f00/24x24)|#ff5f00|
|11|![#ffff00](https://singlecolorimage.com/get/ffff00/24x24)|#ffff00|75|![#5fafff](https://singlecolorimage.com/get/5fafff/24x24)|#5fafff|139|![#af87af](https://singlecolorimage.com/get/af87af/24x24)|#af87af|203|![#ff5f5f](https://singlecolorimage.com/get/ff5f5f/24x24)|#ff5f5f|
|12|![#0000ff](https://singlecolorimage.com/get/0000ff/24x24)|#0000ff|76|![#5fd700](https://singlecolorimage.com/get/5fd700/24x24)|#5fd700|140|![#af87d7](https://singlecolorimage.com/get/af87d7/24x24)|#af87d7|204|![#ff5f87](https://singlecolorimage.com/get/ff5f87/24x24)|#ff5f87|
|13|![#ff00ff](https://singlecolorimage.com/get/ff00ff/24x24)|#ff00ff|77|![#5fd75f](https://singlecolorimage.com/get/5fd75f/24x24)|#5fd75f|141|![#af87ff](https://singlecolorimage.com/get/af87ff/24x24)|#af87ff|205|![#ff5faf](https://singlecolorimage.com/get/ff5faf/24x24)|#ff5faf|
|14|![#00ffff](https://singlecolorimage.com/get/00ffff/24x24)|#00ffff|78|![#5fd787](https://singlecolorimage.com/get/5fd787/24x24)|#5fd787|142|![#afaf00](https://singlecolorimage.com/get/afaf00/24x24)|#afaf00|206|![#ff5fd7](https://singlecolorimage.com/get/ff5fd7/24x24)|#ff5fd7|
|15|![#ffffff](https://singlecolorimage.com/get/ffffff/24x24)|#ffffff|79|![#5fd7af](https://singlecolorimage.com/get/5fd7af/24x24)|#5fd7af|143|![#afaf5f](https://singlecolorimage.com/get/afaf5f/24x24)|#afaf5f|207|![#ff5fff](https://singlecolorimage.com/get/ff5fff/24x24)|#ff5fff|
|16|![#000000](https://singlecolorimage.com/get/000000/24x24)|#000000|80|![#5fd7d7](https://singlecolorimage.com/get/5fd7d7/24x24)|#5fd7d7|144|![#afaf87](https://singlecolorimage.com/get/afaf87/24x24)|#afaf87|208|![#ff8700](https://singlecolorimage.com/get/ff8700/24x24)|#ff8700|
|17|![#00005f](https://singlecolorimage.com/get/00005f/24x24)|#00005f|81|![#5fd7ff](https://singlecolorimage.com/get/5fd7ff/24x24)|#5fd7ff|145|![#afafaf](https://singlecolorimage.com/get/afafaf/24x24)|#afafaf|209|![#ff875f](https://singlecolorimage.com/get/ff875f/24x24)|#ff875f|
|18|![#000087](https://singlecolorimage.com/get/000087/24x24)|#000087|82|![#5fff00](https://singlecolorimage.com/get/5fff00/24x24)|#5fff00|146|![#afafd7](https://singlecolorimage.com/get/afafd7/24x24)|#afafd7|210|![#ff8787](https://singlecolorimage.com/get/ff8787/24x24)|#ff8787|
|19|![#0000af](https://singlecolorimage.com/get/0000af/24x24)|#0000af|83|![#5fff5f](https://singlecolorimage.com/get/5fff5f/24x24)|#5fff5f|147|![#afafff](https://singlecolorimage.com/get/afafff/24x24)|#afafff|211|![#ff87af](https://singlecolorimage.com/get/ff87af/24x24)|#ff87af|
|20|![#0000d7](https://singlecolorimage.com/get/0000d7/24x24)|#0000d7|84|![#5fff87](https://singlecolorimage.com/get/5fff87/24x24)|#5fff87|148|![#afd700](https://singlecolorimage.com/get/afd700/24x24)|#afd700|212|![#ff87d7](https://singlecolorimage.com/get/ff87d7/24x24)|#ff87d7|
|21|![#0000ff](https://singlecolorimage.com/get/0000ff/24x24)|#0000ff|85|![#5fffaf](https://singlecolorimage.com/get/5fffaf/24x24)|#5fffaf|149|![#afd75f](https://singlecolorimage.com/get/afd75f/24x24)|#afd75f|213|![#ff87ff](https://singlecolorimage.com/get/ff87ff/24x24)|#ff87ff|
|22|![#005f00](https://singlecolorimage.com/get/005f00/24x24)|#005f00|86|![#5fffd7](https://singlecolorimage.com/get/5fffd7/24x24)|#5fffd7|150|![#afd787](https://singlecolorimage.com/get/afd787/24x24)|#afd787|214|![#ffaf00](https://singlecolorimage.com/get/ffaf00/24x24)|#ffaf00|
|23|![#005f5f](https://singlecolorimage.com/get/005f5f/24x24)|#005f5f|87|![#5fffff](https://singlecolorimage.com/get/5fffff/24x24)|#5fffff|151|![#afd7af](https://singlecolorimage.com/get/afd7af/24x24)|#afd7af|215|![#ffaf5f](https://singlecolorimage.com/get/ffaf5f/24x24)|#ffaf5f|
|24|![#005f87](https://singlecolorimage.com/get/005f87/24x24)|#005f87|88|![#870000](https://singlecolorimage.com/get/870000/24x24)|#870000|152|![#afd7d7](https://singlecolorimage.com/get/afd7d7/24x24)|#afd7d7|216|![#ffaf87](https://singlecolorimage.com/get/ffaf87/24x24)|#ffaf87|
|25|![#005faf](https://singlecolorimage.com/get/005faf/24x24)|#005faf|89|![#87005f](https://singlecolorimage.com/get/87005f/24x24)|#87005f|153|![#afd7ff](https://singlecolorimage.com/get/afd7ff/24x24)|#afd7ff|217|![#ffafaf](https://singlecolorimage.com/get/ffafaf/24x24)|#ffafaf|
|26|![#005fd7](https://singlecolorimage.com/get/005fd7/24x24)|#005fd7|90|![#870087](https://singlecolorimage.com/get/870087/24x24)|#870087|154|![#afff00](https://singlecolorimage.com/get/afff00/24x24)|#afff00|218|![#ffafd7](https://singlecolorimage.com/get/ffafd7/24x24)|#ffafd7|
|27|![#005fff](https://singlecolorimage.com/get/005fff/24x24)|#005fff|91|![#8700af](https://singlecolorimage.com/get/8700af/24x24)|#8700af|155|![#afff5f](https://singlecolorimage.com/get/afff5f/24x24)|#afff5f|219|![#ffafff](https://singlecolorimage.com/get/ffafff/24x24)|#ffafff|
|28|![#008700](https://singlecolorimage.com/get/008700/24x24)|#008700|92|![#8700d7](https://singlecolorimage.com/get/8700d7/24x24)|#8700d7|156|![#afff87](https://singlecolorimage.com/get/afff87/24x24)|#afff87|220|![#ffd700](https://singlecolorimage.com/get/ffd700/24x24)|#ffd700|
|29|![#00875f](https://singlecolorimage.com/get/00875f/24x24)|#00875f|93|![#8700ff](https://singlecolorimage.com/get/8700ff/24x24)|#8700ff|157|![#afffaf](https://singlecolorimage.com/get/afffaf/24x24)|#afffaf|221|![#ffd75f](https://singlecolorimage.com/get/ffd75f/24x24)|#ffd75f|
|30|![#008787](https://singlecolorimage.com/get/008787/24x24)|#008787|94|![#875f00](https://singlecolorimage.com/get/875f00/24x24)|#875f00|158|![#afffd7](https://singlecolorimage.com/get/afffd7/24x24)|#afffd7|222|![#ffd787](https://singlecolorimage.com/get/ffd787/24x24)|#ffd787|
|31|![#0087af](https://singlecolorimage.com/get/0087af/24x24)|#0087af|95|![#875f5f](https://singlecolorimage.com/get/875f5f/24x24)|#875f5f|159|![#afffff](https://singlecolorimage.com/get/afffff/24x24)|#afffff|223|![#ffd7af](https://singlecolorimage.com/get/ffd7af/24x24)|#ffd7af|
|32|![#0087d7](https://singlecolorimage.com/get/0087d7/24x24)|#0087d7|96|![#875f87](https://singlecolorimage.com/get/875f87/24x24)|#875f87|160|![#d70000](https://singlecolorimage.com/get/d70000/24x24)|#d70000|224|![#ffd7d7](https://singlecolorimage.com/get/ffd7d7/24x24)|#ffd7d7|
|33|![#0087ff](https://singlecolorimage.com/get/0087ff/24x24)|#0087ff|97|![#875faf](https://singlecolorimage.com/get/875faf/24x24)|#875faf|161|![#d7005f](https://singlecolorimage.com/get/d7005f/24x24)|#d7005f|225|![#ffd7ff](https://singlecolorimage.com/get/ffd7ff/24x24)|#ffd7ff|
|34|![#00af00](https://singlecolorimage.com/get/00af00/24x24)|#00af00|98|![#875fd7](https://singlecolorimage.com/get/875fd7/24x24)|#875fd7|162|![#d70087](https://singlecolorimage.com/get/d70087/24x24)|#d70087|226|![#ffff00](https://singlecolorimage.com/get/ffff00/24x24)|#ffff00|
|35|![#00af5f](https://singlecolorimage.com/get/00af5f/24x24)|#00af5f|99|![#875fff](https://singlecolorimage.com/get/875fff/24x24)|#875fff|163|![#d700af](https://singlecolorimage.com/get/d700af/24x24)|#d700af|227|![#ffff5f](https://singlecolorimage.com/get/ffff5f/24x24)|#ffff5f|
|36|![#00af87](https://singlecolorimage.com/get/00af87/24x24)|#00af87|100|![#878700](https://singlecolorimage.com/get/878700/24x24)|#878700|164|![#d700d7](https://singlecolorimage.com/get/d700d7/24x24)|#d700d7|228|![#ffff87](https://singlecolorimage.com/get/ffff87/24x24)|#ffff87|
|37|![#00afaf](https://singlecolorimage.com/get/00afaf/24x24)|#00afaf|101|![#87875f](https://singlecolorimage.com/get/87875f/24x24)|#87875f|165|![#d700ff](https://singlecolorimage.com/get/d700ff/24x24)|#d700ff|229|![#ffffaf](https://singlecolorimage.com/get/ffffaf/24x24)|#ffffaf|
|38|![#00afd7](https://singlecolorimage.com/get/00afd7/24x24)|#00afd7|102|![#878787](https://singlecolorimage.com/get/878787/24x24)|#878787|166|![#d75f00](https://singlecolorimage.com/get/d75f00/24x24)|#d75f00|230|![#ffffd7](https://singlecolorimage.com/get/ffffd7/24x24)|#ffffd7|
|39|![#00afff](https://singlecolorimage.com/get/00afff/24x24)|#00afff|103|![#8787af](https://singlecolorimage.com/get/8787af/24x24)|#8787af|167|![#d75f5f](https://singlecolorimage.com/get/d75f5f/24x24)|#d75f5f|231|![#ffffff](https://singlecolorimage.com/get/ffffff/24x24)|#ffffff|
|40|![#00d700](https://singlecolorimage.com/get/00d700/24x24)|#00d700|104|![#8787d7](https://singlecolorimage.com/get/8787d7/24x24)|#8787d7|168|![#d75f87](https://singlecolorimage.com/get/d75f87/24x24)|#d75f87|232|![#080808](https://singlecolorimage.com/get/080808/24x24)|#080808|
|41|![#00d75f](https://singlecolorimage.com/get/00d75f/24x24)|#00d75f|105|![#8787ff](https://singlecolorimage.com/get/8787ff/24x24)|#8787ff|169|![#d75faf](https://singlecolorimage.com/get/d75faf/24x24)|#d75faf|233|![#121212](https://singlecolorimage.com/get/121212/24x24)|#121212|
|42|![#00d787](https://singlecolorimage.com/get/00d787/24x24)|#00d787|106|![#87af00](https://singlecolorimage.com/get/87af00/24x24)|#87af00|170|![#d75fd7](https://singlecolorimage.com/get/d75fd7/24x24)|#d75fd7|234|![#1c1c1c](https://singlecolorimage.com/get/1c1c1c/24x24)|#1c1c1c|
|43|![#00d7af](https://singlecolorimage.com/get/00d7af/24x24)|#00d7af|107|![#87af5f](https://singlecolorimage.com/get/87af5f/24x24)|#87af5f|171|![#d75fff](https://singlecolorimage.com/get/d75fff/24x24)|#d75fff|235|![#262626](https://singlecolorimage.com/get/262626/24x24)|#262626|
|44|![#00d7d7](https://singlecolorimage.com/get/00d7d7/24x24)|#00d7d7|108|![#87af87](https://singlecolorimage.com/get/87af87/24x24)|#87af87|172|![#d78700](https://singlecolorimage.com/get/d78700/24x24)|#d78700|236|![#303030](https://singlecolorimage.com/get/303030/24x24)|#303030|
|45|![#00d7ff](https://singlecolorimage.com/get/00d7ff/24x24)|#00d7ff|109|![#87afaf](https://singlecolorimage.com/get/87afaf/24x24)|#87afaf|173|![#d7875f](https://singlecolorimage.com/get/d7875f/24x24)|#d7875f|237|![#3a3a3a](https://singlecolorimage.com/get/3a3a3a/24x24)|#3a3a3a|
|46|![#00ff00](https://singlecolorimage.com/get/00ff00/24x24)|#00ff00|110|![#87afd7](https://singlecolorimage.com/get/87afd7/24x24)|#87afd7|174|![#d78787](https://singlecolorimage.com/get/d78787/24x24)|#d78787|238|![#444444](https://singlecolorimage.com/get/444444/24x24)|#444444|
|47|![#00ff5f](https://singlecolorimage.com/get/00ff5f/24x24)|#00ff5f|111|![#87afff](https://singlecolorimage.com/get/87afff/24x24)|#87afff|175|![#d787af](https://singlecolorimage.com/get/d787af/24x24)|#d787af|239|![#4e4e4e](https://singlecolorimage.com/get/4e4e4e/24x24)|#4e4e4e|
|48|![#00ff87](https://singlecolorimage.com/get/00ff87/24x24)|#00ff87|112|![#87d700](https://singlecolorimage.com/get/87d700/24x24)|#87d700|176|![#d787d7](https://singlecolorimage.com/get/d787d7/24x24)|#d787d7|240|![#585858](https://singlecolorimage.com/get/585858/24x24)|#585858|
|49|![#00ffaf](https://singlecolorimage.com/get/00ffaf/24x24)|#00ffaf|113|![#87d75f](https://singlecolorimage.com/get/87d75f/24x24)|#87d75f|177|![#d787ff](https://singlecolorimage.com/get/d787ff/24x24)|#d787ff|241|![#626262](https://singlecolorimage.com/get/626262/24x24)|#626262|
|50|![#00ffd7](https://singlecolorimage.com/get/00ffd7/24x24)|#00ffd7|114|![#87d787](https://singlecolorimage.com/get/87d787/24x24)|#87d787|178|![#d7af00](https://singlecolorimage.com/get/d7af00/24x24)|#d7af00|242|![#6c6c6c](https://singlecolorimage.com/get/6c6c6c/24x24)|#6c6c6c|
|51|![#00ffff](https://singlecolorimage.com/get/00ffff/24x24)|#00ffff|115|![#87d7af](https://singlecolorimage.com/get/87d7af/24x24)|#87d7af|179|![#d7af5f](https://singlecolorimage.com/get/d7af5f/24x24)|#d7af5f|243|![#767676](https://singlecolorimage.com/get/767676/24x24)|#767676|
|52|![#5f0000](https://singlecolorimage.com/get/5f0000/24x24)|#5f0000|116|![#87d7d7](https://singlecolorimage.com/get/87d7d7/24x24)|#87d7d7|180|![#d7af87](https://singlecolorimage.com/get/d7af87/24x24)|#d7af87|244|![#808080](https://singlecolorimage.com/get/808080/24x24)|#808080|
|53|![#5f005f](https://singlecolorimage.com/get/5f005f/24x24)|#5f005f|117|![#87d7ff](https://singlecolorimage.com/get/87d7ff/24x24)|#87d7ff|181|![#d7afaf](https://singlecolorimage.com/get/d7afaf/24x24)|#d7afaf|245|![#8a8a8a](https://singlecolorimage.com/get/8a8a8a/24x24)|#8a8a8a|
|54|![#5f0087](https://singlecolorimage.com/get/5f0087/24x24)|#5f0087|118|![#87ff00](https://singlecolorimage.com/get/87ff00/24x24)|#87ff00|182|![#d7afd7](https://singlecolorimage.com/get/d7afd7/24x24)|#d7afd7|246|![#949494](https://singlecolorimage.com/get/949494/24x24)|#949494|
|55|![#5f00af](https://singlecolorimage.com/get/5f00af/24x24)|#5f00af|119|![#87ff5f](https://singlecolorimage.com/get/87ff5f/24x24)|#87ff5f|183|![#d7afff](https://singlecolorimage.com/get/d7afff/24x24)|#d7afff|247|![#9e9e9e](https://singlecolorimage.com/get/9e9e9e/24x24)|#9e9e9e|
|56|![#5f00d7](https://singlecolorimage.com/get/5f00d7/24x24)|#5f00d7|120|![#87ff87](https://singlecolorimage.com/get/87ff87/24x24)|#87ff87|184|![#d7d700](https://singlecolorimage.com/get/d7d700/24x24)|#d7d700|248|![#a8a8a8](https://singlecolorimage.com/get/a8a8a8/24x24)|#a8a8a8|
|57|![#5f00ff](https://singlecolorimage.com/get/5f00ff/24x24)|#5f00ff|121|![#87ffaf](https://singlecolorimage.com/get/87ffaf/24x24)|#87ffaf|185|![#d7d75f](https://singlecolorimage.com/get/d7d75f/24x24)|#d7d75f|249|![#b2b2b2](https://singlecolorimage.com/get/b2b2b2/24x24)|#b2b2b2|
|58|![#5f5f00](https://singlecolorimage.com/get/5f5f00/24x24)|#5f5f00|122|![#87ffd7](https://singlecolorimage.com/get/87ffd7/24x24)|#87ffd7|186|![#d7d787](https://singlecolorimage.com/get/d7d787/24x24)|#d7d787|250|![#bcbcbc](https://singlecolorimage.com/get/bcbcbc/24x24)|#bcbcbc|
|59|![#5f5f5f](https://singlecolorimage.com/get/5f5f5f/24x24)|#5f5f5f|123|![#87ffff](https://singlecolorimage.com/get/87ffff/24x24)|#87ffff|187|![#d7d7af](https://singlecolorimage.com/get/d7d7af/24x24)|#d7d7af|251|![#c6c6c6](https://singlecolorimage.com/get/c6c6c6/24x24)|#c6c6c6|
|60|![#5f5f87](https://singlecolorimage.com/get/5f5f87/24x24)|#5f5f87|124|![#af0000](https://singlecolorimage.com/get/af0000/24x24)|#af0000|188|![#d7d7d7](https://singlecolorimage.com/get/d7d7d7/24x24)|#d7d7d7|252|![#d0d0d0](https://singlecolorimage.com/get/d0d0d0/24x24)|#d0d0d0|
|61|![#5f5faf](https://singlecolorimage.com/get/5f5faf/24x24)|#5f5faf|125|![#af005f](https://singlecolorimage.com/get/af005f/24x24)|#af005f|189|![#d7d7ff](https://singlecolorimage.com/get/d7d7ff/24x24)|#d7d7ff|253|![#dadada](https://singlecolorimage.com/get/dadada/24x24)|#dadada|
|62|![#5f5fd7](https://singlecolorimage.com/get/5f5fd7/24x24)|#5f5fd7|126|![#af0087](https://singlecolorimage.com/get/af0087/24x24)|#af0087|190|![#d7ff00](https://singlecolorimage.com/get/d7ff00/24x24)|#d7ff00|254|![#e4e4e4](https://singlecolorimage.com/get/e4e4e4/24x24)|#e4e4e4|
|63|![#5f5fff](https://singlecolorimage.com/get/5f5fff/24x24)|#5f5fff|127|![#af00af](https://singlecolorimage.com/get/af00af/24x24)|#af00af|191|![#d7ff5f](https://singlecolorimage.com/get/d7ff5f/24x24)|#d7ff5f|255|![#eeeeee](https://singlecolorimage.com/get/eeeeee/24x24)|#eeeeee|

## References
- [ANSI Escape Codes, Wikipedia](https://en.wikipedia.org/wiki/ANSI_escape_code)
- [Build Your Own Command Line With ANSI Escape Codes, Haoyi's Programming Blog](https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html)
