package main

// TODO fix string dumpStrings
// TODO create a func to find c2's and imports
// TODO add a func to detect malware using prexesting hash's and or 

import (
	"fmt"
	"os"
	"crypto/sha256"
	"crypto/md5"
	"encoding/hex"	
	"bufio"
	"io"
	"strings"
	"unicode"
)

func getFileSize(file string) int64 {
    fileInfo, err := os.Stat(file)
    if err != nil {
        fmt.Println(err)
        return 0
    }
    return fileInfo.Size()
}

func GetHashes(file string) (string, string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", "", err
	}
	defer f.Close()

	md5Hash := md5.New()
	sha256Hash := sha256.New()

	mw := io.MultiWriter(md5Hash, sha256Hash)

	if _, err := io.Copy(mw, f); err != nil {
		return "", "", err
	}

	md5Sum := hex.EncodeToString(md5Hash.Sum(nil))
	sha256Sum := hex.EncodeToString(sha256Hash.Sum(nil))

	return md5Sum, sha256Sum, nil
}


type StringEntry struct {
	Offset int64
	Value  string
}

func dumpStrings(file string, minLen int) ([]StringEntry, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var results []StringEntry
	var current strings.Builder
	var offset int64
	var startOffset int64
	reader := bufio.NewReader(f)
	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if b >= 32 && b < 127 && unicode.IsPrint(rune(b)) {
			if current.Len() == 0 {
				startOffset = offset
			}
			current.WriteByte(b)
		} else {
			if current.Len() >= minLen {
				results = append(results, StringEntry{Offset: startOffset, Value: current.String()})
			}
			current.Reset()
		}
		offset++
	}
	if current.Len() >= minLen {
		results = append(results, StringEntry{Offset: startOffset, Value: current.String()})
	}
	return results, nil
}

func main() {
	var file string

	fmt.Println("Enter File Path: ")
	fmt.Scanf("%s", &file)
	// file size shit
	fileSizeB := getFileSize(file)
	fileSizeKB := float64(fileSizeB) / 1024
	fileSizeMB := fileSizeKB / 1024

	// Hash Shit

	md5sum, sha256sum, err := GetHashes(file)
	if err != nil {
		fmt.Println("Error hashing file:", err)
		return
 	}

	fmt.Println("MD5:", md5sum)
	fmt.Println("SHA256:", sha256sum)
	fmt.Printf("%.2f MB in size", fileSizeMB)
	
	// String Extraction
	
	strs, err := dumpStrings(file, 4)
	if err != nil {
		fmt.Println("Error extracting strings:", err)
		return
	}
	out, err := os.Create(file + "_strings.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer out.Close()
	fmt.Printf("\n%d strings written to %s_strings.txt\n", len(strs), file)
}






