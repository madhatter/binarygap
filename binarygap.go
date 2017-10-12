package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readNum() int64 {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	text, _ := reader.ReadString('\n')
	i, err := strconv.ParseInt(strings.TrimRight(text, "\n"), 10, 64)
	if err == nil {
		return i
	} else {
		fmt.Printf("%v", err)
		return 0
	}
}

func findGap(zahl string) int {
	var start bool
	var count, gap int

	for _, char := range string(zahl) {
		if start == true && string(char) == "0" {
			count += 1
		} else if start == true && string(char) == "1" {
			if count > gap {
				gap = count
			}
			count = 0 // reset
		} else if string(char) == "1" && start == false {
			start = true
		}
	}
	return gap
}

func main() {
	zahl := readNum()
	zahlString := strconv.FormatInt(zahl, 2)

	fmt.Printf("Your number is %v in binary.\n", zahlString)

	gap := findGap(zahlString)

	if gap < 1 {
		fmt.Printf("There is no binary gap in your number.\n")
	} else {
		fmt.Printf("The largest binary gap is %v.\n", gap)
	}
}
