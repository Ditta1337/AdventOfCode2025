package main

import (
	"fmt"
	"github.com/ditta1337/AdventOfCode2025/util"
	"strconv"
)

const InitialPos = 50
const DialSize = 100

func main() {
	filePath, err := util.InputFilePath()
	util.Check(err)

	lines := util.ReadFile(filePath)
	fmt.Printf("Day 1 part 1: %d\n", part1(lines))
	fmt.Printf("Day 1 part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	result := 0
	currentPos := InitialPos

	for _, command := range lines {
		operator, amount, err := parseCommand(command)
		util.Check(err)

		currentPos += operator * amount

		if modulo(currentPos, DialSize) == 0 {
			result++
		}
	}

	return result
}

func part2(lines []string) int {
	result := 0
	currentPos := InitialPos

	for _, command := range lines {
		operator, amount, err := parseCommand(command)
		util.Check(err)

		newPos := currentPos + operator*amount
		start, end := currentPos, newPos

		if newPos < currentPos {
			start, end = newPos-1, currentPos-1
		}
		startSector := (start - modulo(start, DialSize)) / DialSize
		endSector := (end - modulo(end, DialSize)) / DialSize

		result += endSector - startSector
		currentPos = newPos
	}

	return result
}

func parseCommand(command string) (operator int, amount int, err error) {
	if len(command) < 2 {
		return 0, 0, fmt.Errorf("command is too short")
	}

	turn := command[:1]
	amountStr := command[1:]

	amount, err = strconv.Atoi(amountStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid amount: %w", err)
	}

	if turn == "R" {
		return 1, amount, nil
	}

	if turn == "L" {
		return -1, amount, nil
	}

	return 0, 0, fmt.Errorf("invalid turn: %s", turn)
}

func modulo(val int, modulus int) int {
	return ((val % modulus) + modulus) % modulus
}
