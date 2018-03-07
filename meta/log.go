package meta

import (
	"fmt"
	"log"
)

type logWriter struct{}

func (w logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print("Meta: ", string(bytes))
}

//go:generate mockery -name=ILog
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

func (self *Log) Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func (self *Log) Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

func (self *Log) Println(v ...interface{}) {
	log.Println(v...)
}
