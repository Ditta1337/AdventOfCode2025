package main

import (
	"fmt"
	"github.com/ditta1337/AdventOfCode2025/util"
)

const EmptyCell = "."
const PaperCell = "@"
const MaskSize = 3

func main() {
	util.Assert(MaskSize%2 == 1, "mask size not odd")

	filePath, err := util.InputFilePath()
	util.Check(err)

	lines := util.ReadFile(filePath)
	fmt.Printf("Day 4 part 1: %d\n", part1(lines))
	fmt.Printf("Day 4 part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	result := 0
	cafeMap, err := parseLinesToMap(lines)
	util.Check(err)

	length, height := len(cafeMap[0]), len(cafeMap)
	shift := (MaskSize - 1) / 2

	for row := range height {
		for col := range length {
			if cafeMap[row][col] == PaperCell {
				paperCount := checkMask(cafeMap, row, col, shift, height, length)
				if paperCount < 4 {
					result++
				}
			}
		}
	}

	return result
}

func part2(lines []string) int {
	result := 0
	cafeMap, err := parseLinesToMap(lines)
	util.Check(err)

	length, height := len(cafeMap[0]), len(cafeMap)
	shift := (MaskSize - 1) / 2
	removedPapers := -1

	for removedPapers != 0 {
		removedPapers = 0
		for row := range height {
			for col := range length {
				if cafeMap[row][col] == PaperCell {
					paperCount := checkMask(cafeMap, row, col, shift, height, length)
					if paperCount < 4 {
						cafeMap[row][col] = "x"
						removedPapers++
					}
				}
			}
		}
		result += removedPapers
	}
	return result
}

func checkMask(cafeMap [][]string, row, col, shift, height, length int) int {
	paperCount := 0
	originRow, originCol := row-shift, col-shift
	for rowShift := range MaskSize {
		for colShift := range MaskSize {
			actualRow, actualCol := originRow+rowShift, originCol+colShift
			if isNotOutOfTheBounds(actualRow, actualCol, height, length) && isNotCenter(actualRow, actualCol, row, col) {
				if cafeMap[actualRow][actualCol] == PaperCell {
					paperCount++
				}
			}
		}
	}

	return paperCount
}

func isNotCenter(actualRow, actualCol, row, col int) bool {
	return actualRow != row || actualCol != col
}

func isNotOutOfTheBounds(actualRow, actualCol, height, length int) bool {
	return actualRow >= 0 && actualRow < height && actualCol >= 0 && actualCol < length
}

func parseLinesToMap(lines []string) ([][]string, error) {
	var cafeMap [][]string
	for row, line := range lines {
		var cafeMapRow []string
		for col, cellASCII := range line {
			cell := string(cellASCII)
			if cell != EmptyCell && cell != PaperCell {
				return [][]string{}, fmt.Errorf("character in row: %d, column: %d is not \"%s\" or \"%s\"", row, col, EmptyCell, PaperCell)
			}
			cafeMapRow = append(cafeMapRow, cell)
		}
		cafeMap = append(cafeMap, cafeMapRow)
	}

	return cafeMap, nil
}
