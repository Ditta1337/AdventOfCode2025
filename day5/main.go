package main

import (
	"fmt"
	"github.com/ditta1337/AdventOfCode2025/util"
	"sort"
	"strconv"
	"strings"
)

const RangeSeparator = "-"
const RangesIdxSeparator = ""

type value struct {
	val     int
	rangeId int
}

func main() {
	filePath, err := util.InputFilePath()
	util.Check(err)

	lines := util.ReadFile(filePath)
	fmt.Printf("Day 5 part 1: %d\n", part1(lines))
	fmt.Printf("Day 5 part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	result := 0
	rangesMap, values, ids, err := parseLines(lines)
	util.Check(err)
	sortValues(values)

	for _, id := range ids {
		foundIdx := findValueIndexLeftToId(id, values)
		if foundIdx != -1 {
			for _, idx := range util.IntRange(foundIdx, 0) {
				val := values[idx]
				rangeVals := rangesMap[val.rangeId]
				if rangeVals[0] == val && id <= rangeVals[1].val {
					result++
					break
				}
			}
		}
	}

	return result
}

func part2(lines []string) int {
	result := 0
	rangesMap, _, _, err := parseLines(lines)
	util.Check(err)

	merged := mergeRanges(rangesMap)
	for _, r := range merged {
		result += r[1] - r[0] + 1
	}

	return result
}

func findValueIndexLeftToId(id int, values []value) int {
	for idx, val := range values[1:] {
		if id > values[idx].val && id <= val.val {
			return idx
		}
	}

	return -1
}

func sortValues(values []value) {
	sort.Slice(values, func(i, j int) bool {
		return values[i].val < values[j].val
	})
}

func parseLines(lines []string) (map[int][]value, []value, []int, error) {
	rangesMap := make(map[int][]value)
	var values []value
	var ids []int
	currentId := 0
	parsingIds := false

	for idx, line := range lines {
		if line == RangesIdxSeparator {
			parsingIds = true
			continue
		}
		if parsingIds {
			err := parseId(line, &ids)
			if err != nil {
				return nil, nil, nil, fmt.Errorf("error in line %d: %s", idx, err.Error())
			}

		} else {
			err := parseRange(line, &currentId, &values, rangesMap)
			if err != nil {
				return nil, nil, nil, fmt.Errorf("error in line %d: %s", idx, err.Error())
			}
		}
	}

	return rangesMap, values, ids, nil
}

func parseRange(line string, currentId *int, values *[]value, rangesMap map[int][]value) error {
	parts := strings.Split(line, RangeSeparator)
	if len(parts) != 2 {
		return fmt.Errorf("number of parts not equal to 2")
	}
	left, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("cannot parse range left value (%s)", parts[0])
	}

	right, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("cannot parse range right value (%s)", parts[1])
	}

	leftVal := value{left, *currentId}
	rightVal := value{right, *currentId}

	*values = append(*values, leftVal, rightVal)
	rangesMap[*currentId] = []value{leftVal, rightVal}
	*currentId++
	return nil
}

func parseId(line string, ids *[]int) error {
	id, err := strconv.Atoi(line)
	if err != nil {
		return fmt.Errorf("cannot parse ingredient id (%s)", line)
	}
	*ids = append(*ids, id)
	return nil
}

func mergeRanges(rangesMap map[int][]value) [][2]int {
	var ranges [][2]int
	for _, r := range rangesMap {
		ranges = append(ranges, [2]int{r[0].val, r[1].val})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	merged := make([][2]int, 0)
	current := ranges[0]

	for _, r := range ranges[1:] {
		if r[0] <= current[1] {
			if r[1] > current[1] {
				current[1] = r[1]
			}
		} else {
			merged = append(merged, current)
			current = r
		}
	}
	merged = append(merged, current)

	return merged
}
