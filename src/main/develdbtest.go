package main

import (
	dbutil "github.com/tendermint/tmlibs/db"
	"fmt"
)
func main()  {

	//windows的话在c盘根目录
	db := dbutil.NewDB("test","leveldb","/tmp/data")
	defer db.Close()


	db.Set([]byte("testkey"),[]byte("This is a test !"))

	value := db.Get([]byte("testkey"))
	if value == nil{
		return
	}
	fmt.Println(string(value))
	fmt.Println("11111111")
	db.Delete([]byte("testkey"))
	fmt.Println("2222222")
	value1 := db.Get([]byte("testkey"))
	if value1 == nil{

		fmt.Println("33333333")
	}
	fmt.Println(string(value1))

}
