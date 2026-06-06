package main

// ex file /home/ruby/Downloads/sample_setup.exe

import (
	"fmt"
	// "github.com/emgeorrk/go-hashlib"
	"os"
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
	fileSizeB := getFileSize(file)
	fileSizeKB := float64(fileSizeB) / 1024
	fileSizeMB := fileSizeKB / 1024
	fmt.Printf("%.2f MB in size", fileSizeMB)
}
// func getHash() {
//
//
// }
