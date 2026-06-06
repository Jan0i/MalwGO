package main

// ex file /home/ruby/Downloads/sample_setup.exe

import (
	"fmt"
	"os"
	"crypto/sha256"
	// "crypto/md5"
	"log"
	"io"
)

func getFileSize(file string) int64 {
    fileInfo, err := os.Stat(file)
    if err != nil {
        fmt.Println(err)
        return 0
    }
    return fileInfo.Size()
}

func main() {
	var file string

	fmt.Println("Enter File Path: ")
	fmt.Scanf("%s", &file)
	// file size shit
	fileSizeB := getFileSize(file)
	fileSizeKB := float64(fileSizeB) / 1024
	fileSizeMB := fileSizeKB / 1024
	fmt.Printf("%.2f MB in size", fileSizeMB)
	// TODO MAKE HASH'S WORK
// 	hash := sha256.New()
// 	sum := hash.Sum(nil)
// 	if _, err := io.Copy(hash, file); err != nil {
// 	log.Fatal(err)
// }
// 	fmt.Printf("%x", sum)
// }
}

