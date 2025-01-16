package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func grep(pattern string, input []byte ) []string {
	r, _ := regexp.Compile(pattern)

	match := r.FindAllString(string(input), len(string(input)))
	return match

}
func main() {
	read, _ := os.ReadFile("./test.txt")
	// fmt.Println(string(read))
	match := grep(`mul\(\d+,\d+\)`, read)
	var sum int 
	var new_pattern string = `\d+`
	for _, value := range match {

		new_match := grep(new_pattern, []byte(value))
		// fmt.Println(new_match[1])
		firstItem, _ :=  strconv.Atoi(new_match[0])
		secondItem, _ := strconv.Atoi(new_match[1])
		
		sum += firstItem * secondItem
	}
	fmt.Println(sum)
}
