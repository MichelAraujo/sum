package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	const characterHashAsByte byte = 35

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Print(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	total := 0
	for scanner.Scan() {
		if scanner.Bytes()[0] != characterHashAsByte {
			currentValue, _ := strconv.Atoi(scanner.Text())
			total = total + currentValue
		}
	}

	fmt.Println(total)
	file.Close()
}
