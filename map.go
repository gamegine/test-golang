package main

import "fmt"

type A struct {
	A string
	B string
}
type X map[string]A
type Y map[string]*A

func main() {
	{
		v := X{
			"a": {A: "a", B: "b"},
		}
		fmt.Println(v)
		// v["a"].A = "z" //cannot assign to struct field v["a"].A in map
		v["a"] = A{A: "z", B: "z"}
		fmt.Println(v)
	}
	//
	//
	//
	{
		v := Y{
			"a": {A: "a", B: "b"},
		}
		fmt.Println(v["a"].A)
		v["a"].A = "z"
		fmt.Println(v["a"].A)
	}
}
