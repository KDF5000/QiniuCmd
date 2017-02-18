package main

import (
	"fmt"
	"github.com/atotto/clipboard"
)

func TestClipboard() {
	if err := clipboard.WriteAll(string("hello")); err != nil {
		panic(err)
	}
	fmt.Println("su")
}
