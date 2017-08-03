// test project main.go
package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	//fmt.Println(tk())
	db, err := leveldb.OpenFile("./db", nil)
	fmt.Println(err)
	data, err := db.Get([]byte("key"), nil)
	fmt.Println(string(data))
	err = db.Put([]byte("key"), []byte("value"), nil)
	err = db.Delete([]byte("key"), nil)
	defer db.Close()
}
