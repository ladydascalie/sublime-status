package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTime(t *testing.T) {
	time, err := getTime()
	if err != nil {
		t.Error("Error getting time")
	}
	if time != true {
		t.Error("Did not print")
	}
}
func testGetDir(t *testing.T) {
	if getDir() == "" {
		t.Error("Did not print")
	}
}

func testCountDirtyFiles(t *testing.T) {
	if countDirtyFiles() == "" {
		t.Error("did not print")
	}
}

func testGetBranch(t *testing.T) {
	branch, err := getBranch()
	if err != nil {
		assert.Equal(t, "Not a Git repo ◀", err)
		t.Error("Not a git repo")
	}

	if branch == "* master" {
		assert.Equal(t, "* master", branch)
	}
}
