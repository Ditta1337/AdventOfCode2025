package main

import (
	"fmt"
	"github.com/ditta1337/AdventOfCode2025/util"
	"math"
	"strconv"
	"strings"
)

const RangeSeparator = ","
const RangePairSeparator = "-"

type rangePair struct {
	first  int
	second int
}

func main() {
	filePath, err := util.InputFilePath()
	util.Check(err)

	lines := util.ReadFile(filePath)
	fmt.Printf("Day 2 part 1: %d\n", part1(lines))
	fmt.Printf("Day 2 part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	result := 0
	pairs, err := parseLinesToRangePairs(lines)
	util.Check(err)

	for _, pair := range pairs {
		first, second := pair.first, pair.second

		for _, id := range util.IntRange(first, second+1) {
			digitsCount := getDigitsCountInNumber(id)
			if digitsCount%2 == 1 {
				continue
			}

			firstHalf := getChunk(id, 2, digitsCount)
			if id == getNumberFromChunk(firstHalf, 2) {
				result += id
			}
		}
	}

	return result
}

func part2(lines []string) int {
	result := 0
	pairs, err := parseLinesToRangePairs(lines)
	util.Check(err)

	for _, pair := range pairs {
		first, second := pair.first, pair.second

		for _, id := range util.IntRange(first, second+1) {
			digitsCount := getDigitsCountInNumber(id)

			for _, numberOfParts := range getPossibleNumberOfParts(digitsCount) {
				chunk := getChunk(id, numberOfParts, digitsCount)
				if id == getNumberFromChunk(chunk, numberOfParts) {
					result += id
					break
				}
			}
		}
	}

	return result
}

func parseLinesToRangePairs(lines []string) ([]rangePair, error) {
	var pairs []rangePair
	for idx, line := range lines {
		err := parseLineToRangePairs(line, &pairs)
		if err != nil {
			return []rangePair{}, fmt.Errorf("error in input file in line %d: %s", idx, err.Error())
		}
	}

	return pairs, nil
}

func parseLineToRangePairs(line string, pairs *[]rangePair) error {
	ranges := strings.Split(line, RangeSeparator)
	for _, pair := range ranges {
		parts := strings.Split(pair, RangePairSeparator)
		if len(parts) != 2 {
			return fmt.Errorf("range pair \"%s\" has more parts than two", pair)
		}
		first, err := strconv.Atoi(parts[0])
		if err != nil {
			return fmt.Errorf("cannot parse first part of range \"%s\"", pair)
		}

		second, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("cannot parse second part of range \"%s\"", pair)
		}

		if first > second {
			return fmt.Errorf("second part of range (%d) cannot be bigger than first (%d)", first, second)
		}

		*pairs = append(*pairs, rangePair{first, second})
	}

	return nil
}

func getDigitsCountInNumber(num int) int {
	return int(math.Log10(float64(num))) + 1
}

func getChunk(num, numberOfParts, digitsCount int) int {
	return num / util.IntPow(10, digitsCount-digitsCount/numberOfParts)
}

func getNumberFromChunk(chunk, numberOfParts int) int {
	num := 0
	chunkSize := getDigitsCountInNumber(chunk)
	for i := range numberOfParts {
		num += chunk * util.IntPow(10, i*chunkSize)
	}

	return num
}

func getPossibleNumberOfParts(digitsCount int) []int {
	var numberOfParts []int
	for _, partSize := range util.IntRange(1, digitsCount-1) {
		if digitsCount%partSize == 0 {
			numberOfParts = append(numberOfParts, digitsCount/partSize)
		}
	}

	return numberOfParts
}
