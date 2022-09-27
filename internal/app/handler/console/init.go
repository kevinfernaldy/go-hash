package console

import "runtime"

type Console struct {
	environment string
	newline     string
}

func StartConsole() {
	environment := runtime.GOOS
	var newline string

	switch environment {
	case "linux":
		{
			newline = "\n"
		}
	case "windows":
		{
			newline = "\r\n"
		}
	default:
		{
			panic("sorry, your OS is not supported. Supported OS: Windows, Linux")
		}
	}

	console := Console{environment: environment, newline: newline}

	console.handleSimpleHash()
}
