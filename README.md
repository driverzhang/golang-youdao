# Golang 命令行翻译器（有道词典）

> Golang 编写的有道词典命令行版本，支持单词或句子中英互译。

利用 强大的 go build 命令参数

修改并添加因为golang 版本更新强制使用 go mod 作为包管理
 

**获取源码**

```bash
$ git clone https://github.com/driverzhang/golang-youdao.git
```

**编译**

* Windows
```golang
$ go build -ldflags "-s -w" -o yd.exe cli.go core.go
```

* Linux/MacOS
```golang
$ go build -ldflags "-s -w" -o yd cli.go core.go // 压缩二进制文件 并 按照名为 yd 进行打包输出包名
```

然后将你编译好的二进制可执行程序 拷贝的 你的工作目录 go/bin 中即可


### 使用
```
$ yd --help
yd <words> [options]

Query words meanings via the command line.

Example:
  words could be word or sentence.

  yd hello
  yd php is the best language in the world
  yd 你好

Usage:
  yd <words>...
  yd -h | --help
  yd -v | --version

Options:
  -h --help         show this help message and exit.
  -v  --version     displays the current version of youdao-go.
```

**查询单词**
```bash
$ yd coder
>>  coder: 编码器

    美:['kəudə]  英:['kəʊdə]

    n. 编码器；编码员
    n. (Coder)人名；(法)科代；(英、葡)科德尔

    coder[编码员 编码器 程序员]
    speech coder[语言编码器 音频编码器 语音编码器]
    incremental coder[增量编码器]
```

**查询句子**
```bash
$ yd talk is cheap, show me the code
>>  talk is cheap, show me the code: 说说很简单,给我代码
```

**中文翻译**
```
$ yd php是世界上最好的语言
>>  php是世界上最好的语言: PHP is one of the best language in the world
```