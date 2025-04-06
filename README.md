# Colorize Console

![GitHub Tag](https://img.shields.io/github/v/tag/go-universal/console?sort=semver&label=version)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-universal/console.svg)](https://pkg.go.dev/github.com/go-universal/console)
[![License](https://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/go-universal/console/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-universal/console)](https://goreportcard.com/report/github.com/go-universal/console)
![Contributors](https://img.shields.io/github/contributors/go-universal/console)
![Issues](https://img.shields.io/github/issues/go-universal/console)

A token based (`@Style{msg}`) printer library for styling and formatting console output with ANSI escape codes. It supports rich formatting options such as text styles, colors, and backgrounds to make your terminal output more expressive and readable.

## Features

- **Text Styling:** Bold, Underline, Italic, and Strike-through.
- **Text Colors:** Red, Green, Yellow, Blue, Purple, Cyan, White.
- **Background Colors:** Red, Green, Yellow, Blue, Purple, Cyan, White.
- **Chainable Message** API for complex message composition.

## Installation

Add the library to your Go project using:

```sh
go get github.com/go-universal/console
```

## Usage

### PrintF

The PrintF function allows you to apply inline styles and colors directly in your formatted strings. Use `\@` to escape tokens in your patterns.

```go
PrintF(format string, args ...any)

// Usage
PrintF("@Br{Bold Red Text} and @gb{Green Background %s}\n", "message")
```

#### Parameters

- `format`: A string containing golang standart print pattern with text and styled tokens.
- `args`: Additional arguments to replace format specifiers in the string.

#### Supported Tokens

- **Styles (Uppercase):**
  - `B`: Bold
  - `U`: Underline
  - `I`: Italic
  - `S`: Strike-through
- **Text Colors (Lowercase):**
  - `r`: Red
  - `g`: Green
  - `y`: Yellow
  - `b`: Blue
  - `p`: Purple
  - `c`: Cyan
  - `w`: White
- **Background Colors (Lowercase):**
  - `rb`: Red
  - `gb`: Green
  - `yb`: Yellow
  - `bb`: Blue
  - `pb`: Purple
  - `cb`: Cyan
  - `wb`: White

### Message

The `Message` simplifies creating and styling complex messages with a title, tags, and formatted content.

#### Methods

- `Indent()`: Indents the message. For make long indent you can call this method multiple times.
- `Red(title string)`: Makes the message title red.
- `Green(title string)`: Makes the message title green.
- `Yellow(title string)`: Makes the message title yellow.
- `Blue(title string)`: Makes the message title blue.
- `Purple(title string)`: Makes the message title purple.
- `Cyan(title string)`: Makes the message title cyan.
- `Underline()`: Adds underline style to the message body.
- `Strike()`: Adds a strike-through style to the message body.
- `Italic()`: Adds italic style to the message body.
- `Tags(tags ...string)`: Adds tags to the message.
- `Print(pattern string, args ...any)`: Prints the formatted message.

```go
Message().
    Red("Error").
    Italic().
    Tags("CRITICAL", "SYSTEM").
    Print("System encountered an unexpected error: %s", "Disk full")
```
