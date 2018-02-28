package main_test

import (
	"github.com/stretchr/testify/mock"
	"strings"
)

func contains(find string) interface{} {
	return mock.MatchedBy(func(cmd []string) bool { return strings.Contains(strings.Join(cmd, " "), find) })
}
