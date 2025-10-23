package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadLine reads input with spaces from the user
func ReadLine(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
