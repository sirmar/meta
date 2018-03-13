package meta

import (
	"fmt"
	"github.com/blang/semver"
	"strings"
)

//go:generate mockery -name=IRelease
type IRelease interface {
	Create(newVersion, message string)
	List()
	Unreleased()
}

type Release struct {
	util   IUtil
	runner IRunner
	log    ILog
}

func NewRelease(util IUtil, runner IRunner, log ILog) IRelease {
	return &Release{util, runner, log}
}

func (self *Release) Create(level, message string) {
	nextVersion := self.nextVersion(self.currentVersion(), level)
	if self.util.YesNo(fmt.Sprintf("Create release %s? ", nextVersion)) {
		self.runner.Run("git", []string{"tag", nextVersion, "-a", "-m", message})
		self.runner.Run("git", []string{"push", "origin", nextVersion})
	}
}

func (self *Release) List() {
	self.runner.Run("git", []string{"tag", "-n", "--sort", "version:refname", "-l", "[0-9]*"})
}

func (self *Release) Unreleased() {
	if self.currentVersion() != "0.0.0" {
		self.runner.Run("git", []string{"log", self.currentVersion(), "..HEAD", "--oneline"})
	}
}

func (self *Release) nextVersion(currentVersion, level string) string {
	cur := self.version(currentVersion)
	switch level {
	case "major":
		return fmt.Sprintf("%d.0.0", cur.Major+1)
	case "minor":
		return fmt.Sprintf("%d.%d.0", cur.Major, cur.Minor+1)
	default:
		return fmt.Sprintf("%d.%d.%d", cur.Major, cur.Minor, cur.Patch+1)
	}
}

func (self *Release) currentVersion() string {
	output := self.runner.Output("git", []string{"tag", "--sort", "version:refname", "-l", "[0-9]*"})
	if len(output) > 0 {
		tags := strings.Split(output, "\n")
		return tags[len(tags)-1]
	}
	return "0.0.0"
}

func (self *Release) version(versionString string) semver.Version {
	version, err := semver.Make(versionString)
	if err != nil {
		self.log.Fatal(err)
	}
	return version
}
