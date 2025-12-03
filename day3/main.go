package main

import (
	"fmt"
	"github.com/ditta1337/AdventOfCode2025/util"
	"strconv"
)

type batteriesPack struct {
	batteries []int
	foundIdxs []int
}

func main() {
	filePath, err := util.InputFilePath()
	util.Check(err)

	lines := util.ReadFile(filePath)
	fmt.Printf("Day 3 part 1: %d\n", part1(lines))
	fmt.Printf("Day 3 part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	result := 0
	bankSize := 2
	packs, err := parseLinesToBatteriesPacks(lines)
	util.Check(err)

	for _, pack := range packs {
		setFoundIdxs(bankSize, &pack)
		for idx := range pack.foundIdxs {
			result += util.IntPow(10, bankSize-idx-1) * pack.batteries[pack.foundIdxs[idx]]
		}
	}

	return result
}

func part2(lines []string) int {
	result := 0
	bankSize := 12
	packs, err := parseLinesToBatteriesPacks(lines)
	util.Check(err)

	for _, pack := range packs {
		setFoundIdxs(bankSize, &pack)
		for idx := range pack.foundIdxs {
			result += util.IntPow(10, bankSize-idx-1) * pack.batteries[pack.foundIdxs[idx]]
		}
	}

	return result
}

func setFoundIdxs(idxsToFind int, pack *batteriesPack) {
	var foundIdxs []int
	batteries := pack.batteries
	batteriesLength := len(batteries)
	lastFoundIdx := -1
	for idxToFind := range idxsToFind {
		topVoltage := 0
		topVoltageIdx := -1
		for _, batteryIdx := range util.IntRange(lastFoundIdx+1, batteriesLength+idxToFind-idxsToFind) {
			if batteries[batteryIdx] > topVoltage {
				topVoltage = batteries[batteryIdx]
				topVoltageIdx = batteryIdx
			}
		}
		foundIdxs = append(foundIdxs, topVoltageIdx)
		lastFoundIdx = topVoltageIdx
	}

	pack.foundIdxs = foundIdxs
}

func parseLinesToBatteriesPacks(lines []string) ([]batteriesPack, error) {
	var packs []batteriesPack
	for lineIdx, line := range lines {
		var batteries []int
		for batteryIdx, battery := range line {
			voltage, err := strconv.Atoi(string(battery))
			if err != nil {
				return []batteriesPack{}, fmt.Errorf("cannot parse battery \"%s\" (row %d column %d)", string(battery), lineIdx, batteryIdx)
			}
			batteries = append(batteries, voltage)
		}

		packs = append(packs, batteriesPack{batteries, []int{}})
	}

	return packs, nil
}
