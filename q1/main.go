package main

import (
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)


func main() {
	


	start := time.Now()

    r := new(big.Int)
    fmt.Println(r.Binomial(1000, 10))

    elapsed := time.Since(start)


	test, _ := os.ReadFile("./test.txt")
	str := string(test)
	str = strings.ReplaceAll(str, "   ", "\n")
	list := strings.Split(str, "\n")
	var right_list []int32	
	var left_list  []int32

	for i := 0; i < len(list); i++ {
		intValue, _ := strconv.ParseInt(list[i], 10, 32)
		if i%2 == 0 {
			right_list = append(right_list, int32(intValue))
		} else {
			left_list = append(left_list, int32(intValue))
		}
	}
	sort.Slice(left_list, func(i ,j int) bool {
		return left_list[i] < left_list[j]
	})

	sort.Slice(right_list, func(i ,j int) bool {
		return right_list[i] < right_list[j]
	})
	
	var sum int64 	
	for j := 0; j < len(right_list); j++ {

		if right_list[j] > left_list[j] {
			sum += int64(right_list[j] - left_list[j])
		} else {
			sum += int64(left_list[j] - right_list[j])
		}
	}

	fmt.Println(sum)
	fmt.Println(elapsed)

}
