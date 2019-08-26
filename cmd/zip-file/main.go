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
			Content : "Troi hom nay dep qua",
		},
		{
			Name : "file2.text",
			Content : "Ngay mai troi co mua",
		},
	}

	err := zipFile.ZipFolder(files, "./dist/output.zip")
	if err != nil {
		fmt.Println("Error when zip file : ", err)
	}
}