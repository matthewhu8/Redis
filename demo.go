package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := "$5\r\nMatt\r\n"

	reader := bufio.NewReader(strings.NewReader(input))

	b, _ := reader.ReadByte()

	if b != '$' {
		fmt.Println("Not a $ so invalid type, expecting bulk string only")
		os.Exit(1)
	}

	size, _ := reader.ReadByte()
	strSize, _ := strconv.ParseInt(string(size), 10, 64)

	name := make([]byte, strSize)
	reader.Read(name)

	fmt.Println(string(name))

}
