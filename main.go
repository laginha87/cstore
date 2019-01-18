package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/urfave/cli.v1"
)

var filePath, _ = homedir.Expand("~/.command-storer")
var histPath, _ = homedir.Expand("~/.zsh_history")

func getLastLineWithSeek(filepath string) string {
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

func getNthLine(filepath string, lineNum int) string {
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

func add(c *cli.Context) error {
	var lastLine = getLastLineWithSeek(histPath)
	var options = strings.Split(lastLine, ";")
	var cmd = options[len(options)-1]

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(cmd); err != nil {
		panic(err)
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "command-storer"
	app.Usage = "make an explosive entrance"
	app.Commands = []cli.Command{
		{
			Action: add,
		},
		{
			Name: "init",
			Action: func(c *cli.Context) error {
				os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0666)
				return nil
			},
		},
		{
			Name: "add",

			Action: add,
		},
		{
			Name: "reset",
			Action: func(c *cli.Context) error {
				os.OpenFile(filePath, os.O_TRUNC, 0666)
				return nil
			},
			Aliases: []string{"r"},
		},
		{
			Name: "exec",
			Action: func(c *cli.Context) error {
				lineno, _ := strconv.Atoi(c.Args().Get(0))
				cmd := fmt.Sprintf("%s", getNthLine(filePath, lineno))
				syscall.Exec("/bin/zsh", []string{"zsh", "-c", "-l", cmd}, os.Environ())
				return nil
			},
			Aliases: []string{"e"},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
