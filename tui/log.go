package tui

import "os"

func log(msg string) {

	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer func() { _ = file.Close() }()
	// Write the string to the file
	_, _ = file.WriteString(msg + "\n")
}
