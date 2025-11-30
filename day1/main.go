package main

import (
	"fmt"
	"github.com/ditta1337/AdventOfCode2025/util"
)

func main() {
	filePath, err := util.InputFilePath()
	util.Check(err)
	fmt.Println("Using file:", filePath)

	lines := util.ReadFile(filePath)
	for index, line := range lines {
		fmt.Printf("%d: %s\n", index, line)
	}
}
