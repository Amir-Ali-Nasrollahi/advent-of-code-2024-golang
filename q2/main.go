package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Slice_Atoi(strArr []string) ([]int, error) {
    var str string 
    var i int
    var err error 

    iArr := make([]int, 0, len(strArr))
    for _, str = range strArr {
        i, err = strconv.Atoi(str)
        if err != nil {
            return nil, err
		}
        iArr = append(iArr, i)
    }
    return iArr, nil
}


func main() {
	read, _ := os.ReadFile("./test.txt")
	str := string(read)
	strings_array := strings.Split(str, "\n")
	unsafe := 0
	safe := 0

	for _,v := range strings_array {

		iArr, err := Slice_Atoi(strings.Split(v, " "))

		if err != nil {
            panic(err)
        }
		
		count_of_safe := 0
		
		for key, value := range iArr {
			if key > 0 {
				match := value - iArr[key-1]

				if match < 0 && match >= -3{
					count_of_safe--
				}else if match > 0 && match <= 3{
					count_of_safe++
				}
			}
		}

		// we start from second item of the list, so our count should +1 or -1 
		if count_of_safe+1 == len(iArr) || (count_of_safe-1)*(-1) == len(iArr) {
			safe ++
			continue
		} 
		unsafe++
		
	}
	
	fmt.Println(safe, unsafe)
	

}
