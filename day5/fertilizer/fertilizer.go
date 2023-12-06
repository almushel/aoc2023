package fertilizer

import (
	"strings"

	"github.com/almushel/aoc2023/utility"
)

func ParseAlmanac(input string) (seedNums []int, seedMaps [][][3]int) {
	sections := strings.Split(input, "\n\n")
	seedMaps = make([][][3]int, 0)
	for _, section := range sections {
		if strings.HasPrefix(section, "seeds: ") {
			seeds := strings.Split(strings.TrimSpace(section[len("seeds: "):]), " ")
			for _, num := range seeds {
				inum, _ := utility.StringToInt(num)
				seedNums = append(seedNums, inum)
			}
		} else {
			var rangeList [][3]int
			_, sectionRanges, _ := strings.Cut(section, ":")
			for _, line := range strings.Split(sectionRanges, "\n") {
				line = strings.TrimSpace(line)
				if len(line) == 0 {
					continue
				}
				nums := strings.Split(line, " ")
				r1, _ := utility.StringToInt(nums[0])
				r2, _ := utility.StringToInt(nums[1])
				r3, _ := utility.StringToInt(nums[2])
				newRange := [3]int{r1, r2, r3}
				rangeList = append(rangeList, newRange)
			}
			seedMaps = append(seedMaps, rangeList)
		}
	}
	return
}

func mapNum(num int, conversionRange [3]int) int {
	if num < conversionRange[1] || num >= conversionRange[1]+conversionRange[2] {
		return num
	}

	return conversionRange[0] + (num - conversionRange[1])
}

func seedToLocation(seed int, almanac [][][3]int) (result int) {
	result = seed
	for _, conversionMap := range almanac {
		for _, conversionRange := range conversionMap {
			conversion := mapNum(result, conversionRange)
			if conversion != result {
				result = conversion
				break
			}
		}
	}
	return
}

func SolvePart1(seeds []int, almanac [][][3]int) (result int) {
	result = utility.INT_MAX
	for _, seed := range seeds {
		location := seedToLocation(seed, almanac)
		if location < result {
			result = location
		}
	}
	return
}

// bad and dumb brute force method
func PoorlySolvePart2(seedRanges []int, almanac [][][3]int) (result int) {
	result = utility.INT_MAX
	for s := 0; s+1 < len(seedRanges); s += 2 {
		for seed := seedRanges[s]; seed < seedRanges[s]+seedRanges[s+1]; seed++ {
			location := seedToLocation(seed, almanac)
			if location < result {
				result = location
			}
		}
	}
	return
}
