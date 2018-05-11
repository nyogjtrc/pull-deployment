package main

import (
	"os"
	"testing"
)

func TestExecPrinting(t *testing.T) {
	err := execPrinting("ls", "-a")
	if err != nil {
		t.Error(err)
	}
}

func TestFindRepoDir(t *testing.T) {
	err := findAndCreateDir(".")
	if err != nil {
		t.Error(err)
	}

	path := "abc"
	err = findAndCreateDir(path)
	if err != nil {
		t.Error(err)
	}

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		t.Error(err)
	}

	err = os.Remove(path)
	if err != nil {
		t.Error(err)
	}
}
