package main

import (
	"flag"
	"fmt"
	"github.com/KDF5000/QiniuCmd/src/qsync"
	"github.com/atotto/clipboard"
	"os"
	"os/exec"
	"path/filepath"
)

func usage() {
	fmt.Println("Usage: qsync [-r] [-c configure_file_path] file_name")
	os.Exit(1)
}

func main() {
	var randKey bool
	var configFile string
	// curPath, err := os.Getwd()
	exePath, err := exec.LookPath(os.Args[0])

	if err != nil {
		usage()
	}
	curPath := filepath.Dir(exePath)
	// var key string
	flag.BoolVar(&randKey, "r", false, "generate key randomly, if not using the filename as key")
	flag.StringVar(&configFile, "c", curPath+"/conf.json", "configure file")
	flag.Parse()

	if flag.NArg() == 0 {
		usage()
	}

	//load config, return qiniu config(QiniuConf)
	qc := qsync.LoadConf(configFile)
	file := flag.Arg(0) // flag之后的第一个为文件名
	absFile, err2 := qsync.FileExist(file)
	if err2 == false {
		fmt.Fprintf(os.Stderr, "qsync:%v\n", err2)
		os.Exit(1)
	}

	url := qsync.PutFile(absFile, qc, randKey)
	//写入剪贴版
	if err := clipboard.WriteAll(string(url)); err != nil {
		panic(err)
	}
}
