package main

import (
	"fmt"

	zipFile "github.com/duynhanf/zip-file"
)

func main() {
	fmt.Println("zip file run")

	files := []zipFile.FileContent{
		{
			Name : "file1.text",
			Content : "今日、いい天気ですね。",
		},
		{
			Name : "file2.text",
			Content : "明日、雨が降りそうですね。",
		},
	}

	err := zipFile.ZipFolder(files, "output.zip")
	if err != nil {
		fmt.Println("Error when zip file : ", err)
	}
}