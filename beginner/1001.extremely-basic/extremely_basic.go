package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers := readInput()
	sum := numbers[0] + numbers[1]
	fmt.Println("X =", sum)
}

func readInput() []int {
	numbers := make([]int, 0)
	for _, line := range readLines() {
		number, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func readLines() []string {
	stdin := readStdin()
	stdin = strings.TrimSpace(stdin)
	return strings.Split(stdin, "\n")
}

func readStdin() string {
	reader := bufio.NewReader(os.Stdin)
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
