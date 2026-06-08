package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	teks, _ := reader.ReadString('\n')
	return strings.TrimSpace(teks)
}
