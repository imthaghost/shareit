package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
)

func shareit(args []string) {
	filepath := args[0]
	fmt.Println(filepath)

	clipboard.WriteAll(filepath)
}
