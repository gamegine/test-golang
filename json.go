package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// json struct
// exported val start by uppercase
type Obj struct {
	A int64
}
type JSON struct {
	Int   int64 `json:"int"` // explicit JSON key
	Int2  int64 `json:"Int"`
	Str   string
	Bool  bool
	Array []int64
	Obj   Obj
}

func main() {
	// read jsonFile
	content, err := ioutil.ReadFile("json.json")
	// if read file returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
	// unmarshal
	var jsondata JSON
	err = json.Unmarshal(content, &jsondata)
	if err != nil {
		fmt.Println("Error during Unmarshal(): ", err)
	}
	fmt.Println(jsondata)
	// update val
	jsondata.Str = "update"
	fmt.Println(jsondata)

	// jsonByte, err := json.Marshal(jsondata)
	jsonByte, err := json.MarshalIndent(jsondata, "", "\t")
	if err != nil {
		fmt.Println("Error during Marshal(): ", err)
	}
	fmt.Println(string(jsonByte))
	err = ioutil.WriteFile("save.json", jsonByte, 0644) // 0644 Unix permission bits
	// if read file returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
}
