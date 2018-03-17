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
	Verbose(v ...interface{})
	SetVerbose()
	SetQuiet()
	IsQuiet() bool
	IsVerbose() bool
}

type logMode int

const (
	LOG_NORMAL  logMode = 0
	LOG_VERBOSE logMode = 1
	LOG_QUIET   logMode = 2
)

type Log struct {
	mode logMode
}

func NewLog() ILog {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
	return &Log{LOG_NORMAL}
}

func (self *Log) SetVerbose() {
	self.mode = LOG_VERBOSE
}
func (self *Log) SetQuiet() {
	self.mode = LOG_QUIET
}

func (self *Log) IsQuiet() bool {
	return self.mode == LOG_QUIET
}

func (self *Log) IsVerbose() bool {
	return self.mode == LOG_VERBOSE
}

func (self *Log) Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func (self *Log) Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

func (self *Log) Println(v ...interface{}) {
	if !self.IsQuiet() {
		log.Println(v...)
	}
}

func (self *Log) Verbose(v ...interface{}) {
	if self.IsVerbose() {
		log.Println(v...)
	}
}
