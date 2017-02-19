### QiniuCmd
`QinuCmd`的目的是在七牛SDK的基础上封装一个跨平台的命令行工具，用于向七牛云上传下载文件。之所以做这个工具，源自一个很简单的需求，平常写博客或者记笔记都是使用`Markdown`，博客的所有图片都在七牛云存储，每次写的时候遇到插入图片都是先登录网页版的七牛，然后上传图片，复制图片的外链粘贴到`Markdown`文件，写这个工具就是为了简化这个流程，在本地终端直接上传图片，然后生成的外链自动添加到剪切板，这样就可以直接`Ctrl+v或者Cmd+v`粘贴到文章里了。

### RoadMap
- [x] 上传单个文件
- [ ] 批量文件上传
- [x] 生成外链，自动添加到系统的剪切板
- [ ] 监控某个目录，增量上传

### Usage
#### 使用`go get`工具获取
```
#切换到你想要放置代码的目录
$export GOPATH=`pwd`
$ go get github.com/KDF5000/QiniuCmd/src/qupload
```
`get`成功后，`GOPATH`目录下会有自动编译好的可执行程序，然后复制`GOPATH/src/github.com/KDF5000/QiniuCmd/src/qupload `下的`conf.json`文件复制到当前二进制文件目录或者在运新程序时候通过`-c`指定也可。
#### 运行
运行程序的方法如下:
```
./qupload [-r] [-c configure_file_path] file_name
```
* `-r`添加这个参数则文件上传到七牛云的文件名是随机的
* `-c`用于指定配置文件，也就是上面提到的conf.json的路径，也可以配置文件放到二进制文件的相同目录下，这样就可以省略这个参数

下面是一个样例：
```
#进入GOPATH/bin,复制conf.json到当期目录，上传qupload文件本身
$./qupload qupload
http://7sbpmg.com1.z0.glb.clouddn.com/blog/images/qupload
$
```
成功后，`http://7sbpmg.com1.z0.glb.clouddn.com/blog/images/qupload`已经在系统的剪切板了，可以`Cmd+v`到任何地方了。
