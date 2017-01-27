package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	textToPrint := os.Args[1]

	fileHandle, _ := os.Create("output.txt")
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprintln(writer, textToPrint)
	writer.Flush()
}
