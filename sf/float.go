package main

import (
	"bufio"
	"fmt"
	"os"
)

func Float(prompt string) float64 {
	var angka float64
	fmt.Print(prompt)
	fmt.Scan(&angka)
	bufio.NewReader(os.Stdin).ReadString('\n')
	return angka
}
