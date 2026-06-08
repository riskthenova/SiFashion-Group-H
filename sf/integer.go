package main

import (
	"bufio"
	"fmt"
	"os"
)

func Integer(prompt string) int {
	var angka int
	fmt.Print(prompt)
	fmt.Scan(&angka)
	bufio.NewReader(os.Stdin).ReadString('\n')
	return angka
}
