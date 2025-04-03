package console

import (
	"fmt"
	"strings"
)

const (
	resetStyle     = "\033[0m"
	boldStyle      = "\033[1m"
	underlineStyle = "\033[4m"
	strikeStyle    = "\033[9m"
	italicStyle    = "\033[3m"
)

const (
	redBack    = "\033[41m"
	greenBack  = "\033[42m"
	yellowBack = "\033[43m"
	blueBack   = "\033[44m"
	purpleBack = "\033[45m"
	cyanBack   = "\033[46m"
	whiteBack  = "\033[47m"
)

const (
	redColor    = "\033[31m"
	greenColor  = "\033[32m"
	yellowColor = "\033[33m"
	blueColor   = "\033[34m"
	purpleColor = "\033[35m"
	cyanColor   = "\033[36m"
	whiteColor  = "\033[37m"
)

// Printf applies ANSI escape codes to a given pattern string and prints the formatted string with the provided arguments.
// The pattern string can contain tags followed by content token for various styles and colors, which will be replaced by the corresponding ANSI codes.
// You can scape token with \@.
//
// Supported styles:
// B: BOLD.
// U: UNDERLINE.
// S: STRIKE.
// I: ITALIC.
//
// Supported backgrounds:
// rb: RED.
// gb: GREEN.
// yb: YELLOW.
// bb: BLUE.
// pb: PURPLE.
// cb: CYAN.
// wb: WHITE.
//
// Supported colors:
// r: RED.
// g: GREEN.
// y: YELLOW.
// b: BLUE.
// p: PURPLE.
// c: CYAN.
// w: WHITE.
//
// Arguments:
// - format: The string containing the standard go fmt format with styled tokens.
// - args: The arguments to be passed into the format string.
//
// code block:
//
// PrintF("@Bg{Bold Green Text} and @rb{Red %s}\n", "message")
func PrintF(format string, args ...any) {
	var started bool
	var inside bool
	var result strings.Builder
	var token strings.Builder
	var content strings.Builder
	replacer := strings.NewReplacer(
		// styles
		"B", boldStyle,
		"U", underlineStyle,
		"S", strikeStyle,
		"I", italicStyle,
		// background
		"rb", redBack,
		"gb", greenBack,
		"yb", yellowBack,
		"bb", blueBack,
		"pb", purpleBack,
		"cb", cyanBack,
		"wb", whiteBack,
		// colors
		"r", redColor,
		"g", greenColor,
		"y", yellowColor,
		"b", blueColor,
		"p", purpleColor,
		"c", cyanColor,
		"w", whiteColor,
	)

	for i := 0; i < len(format); i++ {
		if format[i] == '@' && (i == 0 || format[i-1] != '\\') {
			started = true
			token.Reset()
			content.Reset()
		} else if started && !inside {
			if format[i] == '{' && (i == 0 || format[i-1] != '\\') {
				inside = true
				content.Reset()
			} else {
				token.WriteByte(format[i])
			}
		} else if started && inside {
			if format[i] == '}' && (i == 0 || format[i-1] != '\\') {
				started = false
				inside = false
				result.WriteString(replacer.Replace(token.String()))
				result.WriteString(content.String())
				result.WriteString(resetStyle)
			} else {
				content.WriteByte(format[i])
			}
		} else {
			result.WriteByte(format[i])
		}
	}
	fmt.Printf(strings.NewReplacer("\\@", "@").Replace(result.String()), args...)
}
