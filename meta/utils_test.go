package meta_test

import (
	"github.com/stretchr/testify/mock"
	"meta/meta"
	"meta/meta/mocks"
	"strings"
)

//go:generate mockery -name=ITemplate
//go:generate mockery -name=ICommand
//go:generate mockery -name=IRunner
//go:generate mockery -name=IUtil
//go:generate mockery -name=ILog

func contains(find string) interface{} {
	return mock.MatchedBy(func(cmd []string) bool { return strings.Contains(strings.Join(cmd, " "), find) })
}

func mockDotMeta() *meta.DotMeta {
	return &meta.DotMeta{&meta.MetaYml{"name", "language"}, "/root", new(mocks.IUtil)}
}

func mockLanguageYml() *meta.LanguageYml {
	return &meta.LanguageYml{
		[]string{"build"},
		[]string{"test"},
		[]string{"lint"},
		[]string{"coverage"},
		[]string{"ci1", "ci2"},
	}
}
