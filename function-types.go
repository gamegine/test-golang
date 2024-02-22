package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

type ExecFunc func(x int, y int) int

func exec(fn ExecFunc, a, b int) int {
	return fn(a, b)
}

var fn map[string]ExecFunc

func execMap(f string, a, b int) int {
	return fn[f](a, b)
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(sub(42, 13))

	fmt.Println(exec(add, 42, 13))
	fmt.Println(exec(sub, 42, 13))
	fmt.Println(exec(func(x, y int) int { return x }, 42, 13))

	// init map
	fn = map[string]ExecFunc{}
	// register func
	fn["add"] = add
	fn["sub"] = sub
	fn["x"] = func(x, y int) int { return x }

	fmt.Println(execMap("add", 42, 13))
	fmt.Println(execMap("sub", 42, 13))
	fmt.Println(execMap("x", 42, 13))
	// fmt.Println(execMap("z", 42, 13)) // make crash
}
