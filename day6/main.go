package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ditta1337/AdventOfCode2025/util"
)

const Mul = "*"
const Add = "+"
const Space = " "

func main() {
	filePath, err := util.InputFilePath()
	util.Check(err)

	lines := util.ReadFile(filePath)
	fmt.Printf("Day 6 part 1: %d\n", part1(lines))
	fmt.Printf("Day 6 part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	result := 0
	numbers, operators, err := parseLines(lines)
	util.Check(err)

	numbersLength := len(numbers)

	for idx, operator := range operators {
		numbersInRow := make([]int, numbersLength)
		for rowIdx := range numbersInRow {
			numbersInRow[rowIdx] = numbers[rowIdx][idx]
		}
		switch operator {
		case Add:
			result += applyAdd(numbersInRow)
		case Mul:
			result += applyMul(numbersInRow)
		}
	}

	return result
}

func part2(lines []string) int {
	result := 0
	numbers, operators, err := parseLinesCephalopods(lines)
	util.Check(err)

	for idx, operator := range operators {
		numbersInCol := numbers[idx]
		switch operator {
		case Add:
			result += applyAdd(numbersInCol)
		case Mul:
			result += applyMul(numbersInCol)
		}
	}

	return result
}

func parseLinesCephalopods(lines []string) ([][]int, []string, error) {
	rowsLength := len(lines)
	columnsLength := len(lines[rowsLength-1])
	var numbers [][]int
	var operators []string

	var numberBuffer []int
	for _, col := range util.IntRange(columnsLength-1, 0) {
		currentNumber := 0
		possibleOperator := string(lines[rowsLength-1][col])
		var possibleDigits []int
		didFindDigit := false
		for row := range rowsLength - 1 {
			char := string(lines[row][col])
			if char != Space {
				didFindDigit = true
				digit, err := strconv.Atoi(string(lines[row][col]))
				if err != nil {
					return nil, nil, fmt.Errorf("error converting %s to int", lines[col][row])
				}
				possibleDigits = append(possibleDigits, digit)
			}
		}
		possibleDigitsLength := len(possibleDigits)
		for idx, digit := range possibleDigits {
			currentNumber += digit * util.IntPow(10, possibleDigitsLength-1-idx)
		}

		if didFindDigit {
			numberBuffer = append(numberBuffer, currentNumber)
		}

		if possibleOperator != Space {
			numbers = append(numbers, numberBuffer)
			operators = append(operators, possibleOperator)
			numberBuffer = []int{}
			didFindDigit = false
		}
	}
	return numbers, operators, nil
}

func parseLines(lines []string) ([][]int, []string, error) {
	length := len(lines)
	var numbers [][]int
	var operators = strings.Fields(lines[length-1])

	for idx := range length - 1 {
		numbersString := strings.Fields(lines[idx])
		parsedNumbers, err := parseNumbers(numbersString)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing numbers in line %d: %v", idx, err)
		}
		numbers = append(numbers, parsedNumbers)
	}

	return numbers, operators, nil
}

func parseNumbers(numbersString []string) ([]int, error) {
	var numbers []int
	for _, numberString := range numbersString {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			return nil, fmt.Errorf("error converting %s to int", numberString)
		}
		numbers = append(numbers, number)
	}

	return numbers, nil
}

func applyAdd(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}

	return result
}

func applyMul(numbers []int) int {
	result := 1
	for _, number := range numbers {
		result = result * number
	}

	return result
}
