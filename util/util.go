package util

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
)

func Check(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Fatalf("%s:%d: %v", filepath.Base(file), line, err)
	}
}

func InputFilePath() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("provide path to input .txt file")
	}

	p := os.Args[1]
	if filepath.Ext(p) != ".txt" {
		return "", fmt.Errorf("file must be .txt")
	}

	return p, nil
}

func ReadFile(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func IntRange(x, y int) []int {
	if x > y {
		return []int{}
	}

	r := make([]int, y-x+1)
	for i := range r {
		r[i] = x + i
	}
	return r
}

func IntPow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
