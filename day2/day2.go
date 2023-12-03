package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stuffInTheBagPt1("fullInput.txt")
}

func stuffInTheBagPt1(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()

	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	lineIndexSum := 0
	lineIndex := 1
	validLine := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		_, line, _ = strings.Cut(line, ": ")
		sets := strings.Split(line, "; ")
		fmt.Println(sets)
		for _, x := range sets {
			vals := strings.Split(x, ", ")
			for _, val := range vals {
				var count int
				if val[1] == ' ' {
					count, _ = strconv.Atoi(val[:1])
				} else {
					count, _ = strconv.Atoi(val[:2])
				}

				fmt.Println(count)
				if strings.Contains(val, "red") {
					if count > maxRed {
						validLine = false
					}
				}
				if strings.Contains(val, "green") {
					if count > maxGreen {
						validLine = false
					}
				}
				if strings.Contains(val, "blue") {
					if count > maxBlue {
						validLine = false
					}
				}
				if !validLine {
					break
				}
			}
			if !validLine {
				break
			}
		}

		if validLine {
			lineIndexSum += lineIndex
		}
		lineIndex++
		validLine = true
	}

	fmt.Println(lineIndexSum)
}
