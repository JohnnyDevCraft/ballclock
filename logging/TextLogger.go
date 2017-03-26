package logging

import (
	"fmt"
	"os"
	"time"
)

type TextLogger struct {
	Path    string
	logFile *os.File
	err     error
}

func (tl *TextLogger) Initialize(path string) {
	tl.Path = path
	tl.logFile, tl.err = os.Create(path)

	if tl.err != nil {
		panic(tl.err)
	}
}

func (tl TextLogger) Info(data string) {
	tl.logFile.WriteString(fmt.Sprintf("%s:: INFO - %s", tl.GetTime(), data))
}

func (tl TextLogger) GetTime() string {
	return time.Now().Format(time.Stamp)
}

func (tl TextLogger) Warn(data string) {
	tl.logFile.WriteString(fmt.Sprintf("%s:: WARN - %s", tl.GetTime(), data))
}

func (tl TextLogger) Err(data string) {
	tl.logFile.WriteString(fmt.Sprintf("%s:: ERROR - %s", tl.GetTime(), data))
}

func (tl *TextLogger) Dispose() {
	tl.logFile.Close()
}
