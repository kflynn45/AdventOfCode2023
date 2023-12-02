package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//decodePt2("sampleInput2.txt")
	decodePt2("fullInput.txt")
}

func decodePt1(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var codeSum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		first := '0'
		last := '0'
		for _, x := range scanner.Text() {
			if x >= '0' && x <= '9' {
				if first == '0' {
					first = x
				}
				last = x
			}
		}
		sum, err := strconv.Atoi(string(first) + string(last))
		if err != nil {
			log.Fatal(err)
		}
		codeSum += sum
	}

	fmt.Println(codeSum)
}

func decodePt2(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var codeSum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var vals []string
		line := scanner.Text()
		for i, x := range line {
			if unicode.IsDigit(x) {
				vals = append(vals, string(x))
			}
			for j, y := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
				if strings.HasPrefix(line[i:], y) {
					vals = append(vals, strconv.Itoa(j+1))
				}
			}
		}
		// for _, a := range vals {
		// 	fmt.Print(a)
		// 	fmt.Print(",")
		// }
		sum, _ := strconv.Atoi(vals[0] + vals[len(vals)-1])
		//fmt.Println(sum)
		codeSum += sum
	}

	fmt.Println(codeSum)
}
