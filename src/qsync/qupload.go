package qsync

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
)

func LoadConf(file string) QiniuConf {
	if _, err := FileExist(file); err == false {
		fmt.Println("Error: Please confirm the default configure file(conf.josn in current directory) or specify a configure file with option -c ")
		os.Exit(1)
	}

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "qsync:%v\n", err)
		os.Exit(1)
	}

	qc := QiniuConf{}
	if err := json.Unmarshal(bytes, &qc); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		os.Exit(1)
	}
	// fmt.Println(qc)
	return qc
}

// 构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

type QiniuConf struct {
	AccessKey  string
	SecretKey  string
	Bucket     string
	Domain     string
	PathPrefix string
}

func PutFile(filepath string, qc QiniuConf, randKey bool) string {
	// 初始化AK，SK
	conf.ACCESS_KEY = qc.AccessKey
	conf.SECRET_KEY = qc.SecretKey
	bucket := qc.Bucket
	domain := qc.Domain
	pathPrefix := qc.PathPrefix

	// 创建一个Client
	c := kodo.New(0, nil)
	// 设置上传的策略
	scope := bucket
	key := pathPrefix + path.Base(filepath)
	if !randKey {
		//生成路径前缀+文件名作为key
		scope += ":" + key
	}
	policy := &kodo.PutPolicy{
		Scope: scope,
		// 设置Token过期时间
		Expires: 3600,
	}
	// 生成一个上传token
	token := c.MakeUptoken(policy)
	//构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)
	var ret PutRet
	// 设置上传文件的路径
	// 调用PutFile方式上传，这里的key需要和上传指定的key一致
	var res interface{}
	if randKey {
		res = uploader.PutFileWithoutKey(nil, &ret, token, filepath, nil)
	} else {
		res = uploader.PutFile(nil, &ret, token, key, filepath, nil)
	}
	// 打印出错信息
	if res != nil {
		fmt.Fprintf(os.Stderr, "qsync:%v\n", res)
		return ""
	}
	fmt.Println(domain + ret.Key)
	return domain + ret.Key
}
