package cstore

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func GetLastLineWithSeek(filepath string) string {
	fileHandle, err := os.Open(filepath)

	if err != nil {
		panic("Cannot open file")
		os.Exit(1)
	}
	defer fileHandle.Close()

	line := ""
	var lastLine = true
	var cursor int64 = 0
	stat, _ := fileHandle.Stat()
	filesize := stat.Size()
	for {
		cursor -= 1
		fileHandle.Seek(cursor, io.SeekEnd)

		char := make([]byte, 1)
		fileHandle.Read(char)

		if cursor != -1 && (char[0] == 10 || char[0] == 13) { // stop if we find a line
			if lastLine {
				lastLine = false
			} else {
				break
			}
		}
		if !lastLine {
			line = fmt.Sprintf("%s%s", string(char), line) // there is more efficient way
		}

		if cursor == -filesize { // stop if we are at the begining
			break
		}
	}

	return line
}

func GetNthLine(filepath string, lineNum int) string {
	lastLine := 0
	fileHandle, err := os.Open(filepath)
	if err != nil {
		panic("Cannot open file")
		os.Exit(1)
	}
	defer fileHandle.Close()

	sc := bufio.NewScanner(fileHandle)
	for sc.Scan() {
		lastLine++
		if lastLine == lineNum {
			// you can return sc.Bytes() if you need output in []bytes
			return sc.Text()
		}
	}
	return sc.Text()
}
