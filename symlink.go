package main

import (
	"fmt"
	"os"
)

func main() {
	const someSymlink = "symlink"
	// const someSymlink = "websocket"

	//get file info
	fileInfo, err := os.Lstat(someSymlink)
	//handle error
	if err != nil {
		panic(err)
	}

	// print file info
	fmt.Print("simlink: ")
	fmt.Println(fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink)
	if fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink {
		resolvedLink, err := os.Readlink(someSymlink)
		if err != nil {
			fmt.Println("error Readlink: ", err)
		}
		fmt.Println("->", resolvedLink)
	}
}
