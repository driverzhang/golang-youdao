package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/koron/go-dproxy"
)

const (
	KEY     = "1945325576"
	KEYFROM = "Youdao-dict-v21"
	URL     = "http://fanyi.youdao.com/openapi.do?keyfrom=%s&key=%s&type=data&doctype=json&version=1.1&q=%s"
	VERSION = "1.0.0"
)

// 打印版本号
func ShowVersion() {
	fmt.Printf("youdao-go %s\n", VERSION)
}

// 请求连接并返回响应
func Request(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
}

// 处理返回的 json 数据
func Parser(words string) {
	// 翻译的内容需要进行 urlencode 编码
	escapeeWords := url.QueryEscape(words) // 对 words 进行转码使之可以安全的用在URL查询里
	queryUrl := fmt.Sprintf(URL, KEYFROM, KEY, escapeeWords)

	resp, err := Request(queryUrl)

	var v interface{}
	err = json.Unmarshal([]byte(resp), &v)
	if err != nil {
		return
	}

	// 使用 go-dproxy 库解析 json 数据
	p := dproxy.New(v)
	errorCode, err := p.M("errorCode").Int64()
	if errorCode != 0 || err != nil {
		fmt.Println(errorCode)
		fmt.Println("=>>>  Exception: The words can't be found, please check your spelling")
		return
	}
	// 词意翻译
	translation, _ := p.M("translation").A(0).String()
	fmt.Printf("=>>>   %s: %s\n\n", words, translation)

	// 发音
	usPhonetic, err := p.M("basic").M("us-phonetic").String()
	if err == nil {
		ukPhonetic, _ := p.M("basic").M("uk-phonetic").String()
		fmt.Printf("    美:[%s]  英:[%s]\n\n", usPhonetic, ukPhonetic)
	}

	// 基本释义
	explains, err := p.M("basic").M("explains").Array()
	if err == nil {
		for _, value := range explains {
			fmt.Printf("    %s\n", value)
		}
		fmt.Println()
	}

	// 网络释义
	web, err := p.M("web").Array()
	if err == nil {
		for _, value := range web {
			item := value.(map[string]interface{})
			fmt.Printf("    %s%s\n", item["key"], item["value"])
		}
	}
}
