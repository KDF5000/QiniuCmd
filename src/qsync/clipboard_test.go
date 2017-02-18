package main

import (
	"github.com/atotto/clipboard"
	"testing"
)

func TestClipboard(t *testing.T) {
	if err := clipboard.WriteAll(string("hello")); err != nil {
		panic(err)
	}
}
