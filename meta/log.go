package meta

import (
	"fmt"
	"log"
)

type logWriter struct{}

func (w logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print("Meta: ", string(bytes))
}

type ILog interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Println(v ...interface{})
}

type Log struct{}

func NewLog() ILog {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	return &Log{}
}

func (l *Log) Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func (l *Log) Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

func (l *Log) Println(v ...interface{}) {
	log.Println(v...)
}
