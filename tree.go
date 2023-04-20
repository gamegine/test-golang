package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {

	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// skip patern
			{
				var reg = regexp.MustCompile(`\.git|a.txt|\.go`)
				if reg.MatchString(info.Name()) {
					if info.IsDir() {
						return filepath.SkipDir
					}
					return nil
				}
			}
			fmt.Println(path, info.Size()/1024, "Ko")
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
