package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	birthDay1, _ := time.ParseInLocation("2006-01-02 15:04:05", "1993-01-02 00:00:00", loc)
	birthDay2, err := time.ParseInLocation("2006-01-02 15:04:05", "1993-01-02 00:00:00", time.Local)
	fmt.Println(birthDay1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(birthDay2)
	//timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(timeObj)
}
