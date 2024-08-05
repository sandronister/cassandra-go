package logger

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	fileName string
}

func createFolder(folder string) error {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err := os.Mkdir(folder, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func createFile(fileName string) error {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

func NewLogger(folder string) (*Logger, error) {
	folder = fmt.Sprintf("../%s", folder)
	err := createFolder(folder)

	if err != nil {
		return nil, err
	}

	fileName := fmt.Sprintf("%s/%s.log", folder, time.Now().Format("2006-01-02"))

	err = createFile(fileName)

	if err != nil {
		return nil, err
	}

	return &Logger{fileName: fileName}, nil
}

func (l *Logger) Info(infoErr error) {
	if infoErr == nil {
		return
	}

	formatErr := fmt.Sprintf("[INFO] %s", infoErr)
	file, err := os.OpenFile(l.fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(formatErr)
}
