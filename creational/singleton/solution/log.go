package main

import "fmt"

type Logger struct {
	filename  string
	instances map[string]Logger
}

func (Logger) Write(msg string) {
	fmt.Println("writing message to the log")
}

func (l *Logger) GetInstance(fileName string) Logger {
	if _, ok := l.instances[fileName]; !ok {
		l.instances[fileName] = Logger{filename: fileName}
	}
	return l.instances[fileName]
}
