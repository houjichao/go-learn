package main

import (
	"encoding/json"
	"fmt"
)

type CompanyDimUpdateDataStruct struct {
	Id      int    `json:"id"`
	Manager string `json:"manager"`
}

func main() {

	updateData := []CompanyDimUpdateDataStruct{{Id: int(111), Manager: "houjichao;banch"}}

	str, _ := json.Marshal(updateData)
	fmt.Println(string(str))

	updateData1 := []map[string]interface{}{{"id": int(111)}, {"manager": "houjichao;banch"}}
	str1, _ := json.Marshal(updateData1)
	fmt.Println(string(str1))

	var channelMark int

	fmt.Println(channelMark)

	// 正确示例
	x := map[string]string{"one": "2", "two": "", "three": "3"}
	if _, ok := x["two"]; !ok {
		fmt.Println("key two is no entry")

	}

	var flag bool
	fmt.Println(flag)

	var any1 any
	any1 = "1111"
	str2 := fmt.Sprintf("%v", any1)
	fmt.Println(str2)
}
