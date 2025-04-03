package console

// MessagePrinter print a colored message (Title: [tags] message).
type MessagePrinter interface {
	// Indent add indent to message.
	Indent() MessagePrinter

	// Red make message title red.
	Red(title string) MessagePrinter

	// Green make message title green.
	Green(title string) MessagePrinter

	// Yellow make message title yellow.
	Yellow(title string) MessagePrinter

	// Blue make message title blue.
	Blue(title string) MessagePrinter

	// Purple make message title purple.
	Purple(title string) MessagePrinter

	// Cyan make message title cyan.
	Cyan(title string) MessagePrinter

	// Underline add underline style to message details.
	Underline() MessagePrinter

	// Strike add strike style to message details.
	Strike() MessagePrinter

	// Italic add italic style to message details.
	Italic() MessagePrinter

	// Tags add tags to message.
	Tags(tags ...string) MessagePrinter

	// Print print message
	Print(message string)

	// Printf print message with fmt style
	Printf(pattern string, args ...any)
}

func Message() MessagePrinter {
	return new(message)
}

// implementation
type message struct {
	title  string
	color  string
	style  string
	indent string
	tags   []string
}

func (msg *message) Indent() MessagePrinter {
	msg.indent = msg.indent + "    "
	return msg
}

func (msg *message) Red(title string) MessagePrinter {
	msg.color = "r"
	msg.title = title
	return msg
}

func (msg *message) Green(title string) MessagePrinter {
	msg.color = "g"
	msg.title = title
	return msg
}

func (msg *message) Yellow(title string) MessagePrinter {
	msg.color = "y"
	msg.title = title
	return msg
}

func (msg *message) Blue(title string) MessagePrinter {
	msg.color = "b"
	msg.title = title
	return msg
}

func (msg *message) Purple(title string) MessagePrinter {
	msg.color = "p"
	msg.title = title
	return msg
}

func (msg *message) Cyan(title string) MessagePrinter {
	msg.color = "c"
	msg.title = title
	return msg
}

func (msg *message) Underline() MessagePrinter {
	msg.style = msg.style + "U"
	return msg
}

func (msg *message) Strike() MessagePrinter {
	msg.style = msg.style + "S"
	return msg
}

func (msg *message) Italic() MessagePrinter {
	msg.style = msg.style + "I"
	return msg
}

func (msg *message) Tags(tags ...string) MessagePrinter {
	msg.tags = append(msg.tags, tags...)
	return msg
}

func (msg *message) Print(message string) {
	title := ""
	tags := ""
	params := make([]any, 0)

	// Generate title style
	if msg.title != "" {
		title = "@B" + msg.color + "{%s:} "
		params = append(params, msg.title)
	}

	// Generate tags
	for _, tag := range msg.tags {
		if tag != "" {
			tags = tags + "[%s] "
			params = append(params, tag)
		}
	}

	// Generate message
	if msg.style != "" {
		message = "@" + msg.style + "{" + message + "}"
	}

	// Print message
	PrintF(msg.indent+title+tags+message+"\n", params...)
}

func (msg *message) Printf(pattern string, args ...any) {
	title := ""
	tags := ""
	params := make([]any, 0)

	// Generate title style
	if msg.title != "" {
		title = "@B" + msg.color + "{%s:} "
		params = append(params, msg.title)
	}

	// Generate tags
	for _, tag := range msg.tags {
		if tag != "" {
			tags = tags + "[%s] "
			params = append(params, tag)
		}
	}

	// Generate message
	if msg.style != "" {
		pattern = "@" + msg.style + "{" + pattern + "}"
	}

	// Print message
	params = append(params, args...)
	PrintF(msg.indent+title+tags+pattern+"\n", params...)
}
