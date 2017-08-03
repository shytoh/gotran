// test project main.go
package main

import (
	"encoding/json"
	"fmt"
)

type MyData struct {
	ID   int         //id
	Over int         //阈值(多少个链接发出)
	Next map[int]int //上个连接 id-id
	Pre  map[int]int //下个连接 id-id
	//Input [100]chan int //输入
	Depth int //深度
}

func init() {
	var detail MyData
	var js MyData
	detail.Next = make(map[int]int)
	detail.Pre = make(map[int]int)
	detail.ID = 1
	detail.Over = 2
	detail.Next[1] = 2
	detail.Pre[2] = 2
	fmt.Println(detail)
	body, err := json.Marshal(detail)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(body))
	err = json.Unmarshal(body, &js)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(js)
}
