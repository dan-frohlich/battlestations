package tui

import (
	"fmt"
	"os"
	"time"
)

func logWithTS(ts time.Time, msg string) {
	log(fmt.Sprintf("[%s] - %s)", time.Now().Format(time.RFC3339Nano), msg))
}

func log(msg string) {

	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer func() { _ = file.Close() }()
	// Write the string to the file
	_, _ = file.WriteString(msg + "\n")
}
