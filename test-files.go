package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		fmt.Println("	ModeDir", file.Mode()&fs.ModeDir == fs.ModeDir)                      // d: is a directory
		fmt.Println("	ModeAppend", file.Mode()&fs.ModeAppend == fs.ModeAppend)             // a: append-only
		fmt.Println("	ModeExclusive", file.Mode()&fs.ModeExclusive == fs.ModeExclusive)    // l: exclusive use
		fmt.Println("	ModeTemporary", file.Mode()&fs.ModeTemporary == fs.ModeTemporary)    // T: temporary file; Plan 9 only
		fmt.Println("	ModeSymlink", file.Mode()&fs.ModeSymlink == fs.ModeSymlink)          // L: symbolic link
		fmt.Println("	ModeDevice", file.Mode()&fs.ModeDevice == fs.ModeDevice)             // D: device file
		fmt.Println("	ModeNamedPipe", file.Mode()&fs.ModeNamedPipe == fs.ModeNamedPipe)    // p: named pipe (FIFO)
		fmt.Println("	ModeSocket", file.Mode()&fs.ModeSocket == fs.ModeSocket)             // S: Unix domain socket
		fmt.Println("	ModeSetuid", file.Mode()&fs.ModeSetuid == fs.ModeSetuid)             // u: setuid
		fmt.Println("	ModeSetgid", file.Mode()&fs.ModeSetgid == fs.ModeSetgid)             // g: setgid
		fmt.Println("	ModeCharDevice", file.Mode()&fs.ModeCharDevice == fs.ModeCharDevice) // c: Unix character device, when ModeDevice is set
		fmt.Println("	ModeSticky", file.Mode()&fs.ModeSticky == fs.ModeSticky)             // t: sticky
		fmt.Println("	ModeIrregular", file.Mode()&fs.ModeIrregular == fs.ModeIrregular)    // ?: non-regular file; nothing else is known about this file
	}
}
