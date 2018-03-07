package meta_test

import (
	"bytes"
	"github.com/stretchr/testify/suite"
	"log"
	"meta/meta"
	"testing"
)

type LogTest struct {
	suite.Suite
	log meta.ILog
}

func (suite *LogTest) SetupTest() {
	suite.log = meta.NewLog()
}

func (suite *LogTest) TestPrintln() {
	buffer := new(bytes.Buffer)
	log.SetOutput(buffer)
	suite.log.Println("Hello")
	suite.Equal("Hello\n", buffer.String())
}

func TestLogSuite(t *testing.T) {
	suite.Run(t, new(LogTest))
}
