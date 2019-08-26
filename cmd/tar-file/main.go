package main

import (
	"archive/tar"
	"fmt"
	"os"
)

func main() {
	tf, err := os.Create("output.tar")
	if err != nil {
		fmt.Println("Create file error : ", err)
	}
	defer tf.Close()

	of, err := os.Open("file2.txt")
	if err != nil {
		fmt.Println("Create file error : ", err)
	}
	defer of.Close()

	tarReader := tar.NewWriter(tf)

	tarReader.WriteHeader(&tar.Header{
		Name : "file1.text",
		Mode: 777,
		Size: 22,
	})

	// c, _ := tarReader.Write([]byte("hom nay toi lam gi day"))
	// fmt.Println(c)

	tarReader.Close()

}