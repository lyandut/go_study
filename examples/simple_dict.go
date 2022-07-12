package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
}

type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func query(word string) {
	client := &http.Client{}
	//var data = strings.NewReader(`{"trans_type":"en2zh","source":"good"}`)
	request := DictRequest{TransType: "en2zh", Source: word}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)                                                           // json.Marshal返回一个byte数组
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data) // 创建http请求
	if err != nil {
		log.Fatal(err)
	}
	// 设置请求头
	//req.Header.Set("Connection", "keep-alive")
	//req.Header.Set("DNT", "1")
	//req.Header.Set("os-version", "")
	//req.Header.Set("sec-ch-ua-mobile", "?0")
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	//req.Header.Set("app-name", "xy")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	//req.Header.Set("Accept", "application/json, text/plain, */*")
	//req.Header.Set("device-id", "")
	//req.Header.Set("os-type", "web")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	//req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	//req.Header.Set("Sec-Fetch-Site", "cross-site")
	//req.Header.Set("Sec-Fetch-Mode", "cors")
	//req.Header.Set("Sec-Fetch-Dest", "empty")
	//req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	//req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	//req.Header.Set("Cookie", "_ym_uid=16456948721020430059; _ym_d=1645694872")

	resp, err := client.Do(req) // 发送请求
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body) // 读取响应，将body转换为byte数组
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	//fmt.Printf("%s\n", bodyText)
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%#v\n", dictResponse)
	fmt.Println(word, "UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)
	for _, item := range dictResponse.Dictionary.Explanations {
		fmt.Println(item)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: simpleDict WORD\nexample: simpleDict hello\n")
		os.Exit(1)
	}
	word := os.Args[1]
	query(word)
}
