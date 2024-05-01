package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) bool {
	return e != nil
}

func getFileInfo(fileHandler *os.File, functionType bufio.SplitFunc) int {
	fileScanner := bufio.NewScanner(fileHandler)
	if err := fileScanner.Err(); err != nil {
		println(err)
	}
	fileScanner.Split(functionType)
	count := 0

	for fileScanner.Scan() {
		count++
	}

	defer fileHandler.Close()

	return count
}

func getFile(filepath string) *os.File {
	fileHandler, err := os.Open(filepath)
	if check(err) {
		fmt.Println(err)
		panic(err)
	}

	return fileHandler
}

func main() {
	args := os.Args[1:]
	bufio_function_map := map[string]bufio.SplitFunc{
		"-c": bufio.ScanBytes,
		"-l": bufio.ScanLines,
		"-w": bufio.ScanWords,
		"-m": bufio.ScanRunes,
	}
	if len(args) > 1 {
		commandType := args[0]
		filepath := args[1]

		functionType, ok := bufio_function_map[commandType]
		if !ok {
			fmt.Println("Invalid command type:", commandType)
			return
		}
		file := getFile(filepath)
		fileInfo := getFileInfo(file, functionType)
		fmt.Println(fileInfo)
	} else {
		filepath := args[0]
		function, ok := bufio_function_map[filepath]
		stat, _ := os.Stdin.Stat()
		if ok && stat.Size() == 0 {
			fmt.Println("Invalid File path:", filepath)
		} else if stat.Size() == 0 {
			for _, function := range bufio_function_map {
				file := getFile(filepath)
				fileInfo := getFileInfo(file, function)
				fmt.Println(fileInfo)
			}
		} else {
			fileInfo := getFileInfo(os.Stdin, function)
			println(fileInfo)
		}

	}
}

/*
go run wc_tool.go -l test.txt
7145

go run wc_tool.go -c test.txt
342190

go run wc_tool.go -w test.txt
58164

go run wc_tool.go -m test.txt
339292

go run wc_tool.go test.txt
342190
7145
58164
339292
*/
