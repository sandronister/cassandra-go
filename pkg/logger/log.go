package logger

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	fileName string
}

func NewLogger(folder string) (*Logger, error) {
	fileName := fmt.Sprintf("%s/%s.log", folder, time.Now().Format("2006-01-02"))

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}

	return &Logger{fileName: fileName}, nil
}

func (l *Logger) Info(msg string) {
	file, err := os.OpenFile(l.fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("[INFO] %s\n", msg))
}
