package meta_test

import (
	"github.com/stretchr/testify/mock"
	"strings"
)

//go:generate mockery -name=ITemplate
//go:generate mockery -name=ILanguage
//go:generate mockery -name=IRunner
//go:generate mockery -name=IUtil
//go:generate mockery -name=ILog

func contains(find string) interface{} {
	return mock.MatchedBy(func(cmd []string) bool { return strings.Contains(strings.Join(cmd, " "), find) })
}
