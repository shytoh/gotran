// test project main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var cells map[int]cell

type cell struct {
	ID   int
	Over int
	Next []weight
	Pre  []weight
	Ch   [100]chan int
	Head map[int]weight
}

type weight struct {
	ID     int
	Weight int
}

func hello() {
	fmt.Println("Hello World!")
	return
}
func httpGet() {
	resp, err := http.Get("https://translate.google.cn/translate_a/single?client=t&sl=en&tl=zh-CN&hl=zh-CN&dt=at&dt=bd&dt=ex&dt=ld&dt=md&dt=qc&dt=rw&dt=rm&dt=ss&dt=t&ie=UTF-8&oe=UTF-8&source=sel&tk=97753.468883&q=when")
	if err != nil {
		fmt.Println("服务器无法访问！")
		return
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	//fmt.Println(string(body))
	ret, err := json2map(body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)

}
func json2map(str []byte) (s interface{}, err error) {
	var result interface{}
	if err := json.Unmarshal(str, &result); err != nil {
		return nil, err
	}
	return result, nil
}
