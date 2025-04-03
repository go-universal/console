package console_test

import (
	"testing"

	"github.com/go-universal/console"
)

func TestColors(t *testing.T) {
	tokens := []string{
		"@B{Bold} ",
		"@U{Underline} ",
		"@S{Strike} ",
		"@I{Italic} ",
		"\n",

		"@rb{ Red Background } ",
		"@gb{ Green Background } ",
		"@yb{ Yellow Background } ",
		"@bb{ Blue Background } ",
		"@pb{ Purple Background } ",
		"@cb{ Cyan Background } ",
		"@wb{ White Background } ",
		"\n",

		"@r{Red Foreground} ",
		"@g{Green Foreground} ",
		"@y{Yellow Foreground} ",
		"@b{Blue Foreground} ",
		"@p{Purple Foreground} ",
		"@c{Cyan Foreground} ",
		"@w{White Foreground} ",
		"\n",
	}

	for _, token := range tokens {
		console.PrintF(token)
	}

	console.Message().
		Green("Migrate contents").
		Italic().
		Strike().
		Print("")

	console.Message().
		Blue("Message").
		Italic().
		Underline().
		Indent().
		Printf("Welcome %s", "John")
	console.Message().
		Red("Error").
		Italic().
		Strike().
		Indent().
		Tags("One", "Two").
		Print("some data is invalid")
}
