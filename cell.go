// test project main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

var cells map[int]*cell
var chs map[int]chan int

func init() {
	cells = make(map[int]*cell)
	chs = make(map[int]chan int)
}

type cell struct {
	ID       int         //id
	Over     int         //阈值(多少个链接发出)
	Nextwith map[int]int //上个连接 id-id
	Prewith  map[int]int //下个连接 id-id
	//需要一个集合
	//Nextwit	 map[]
	//Input [100]chan int //输入
	Depth int //深度
}

func upcell(id int) {
	cells[id] = new(cell)
	ch := make(chan int, 10)
	var sig int
	//for {
	sig = <-ch
	//}
	fmt.Println(sig)
	return
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
	for v := range ret.([]interface{}) {

		//fmt.Println(ret.([]interface{})[v])
		if ret.([]interface{})[v] == nil {
			continue
		}
		fmt.Print("\t")
		fmt.Println(reflect.TypeOf(ret.([]interface{})[v]).String())
	}

	//fmt.Println(ret)
	//fmt.Println(reflect.TypeOf(ret).String())
}
func json2map(str []byte) (s interface{}, err error) {
	var result interface{}
	if err := json.Unmarshal(str, &result); err != nil {
		return nil, err
	}
	return result, nil
}
