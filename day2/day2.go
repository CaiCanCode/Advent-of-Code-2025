package main

import ( 
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan();
    line := scanner.Text()
    if err := scanner.Err(); err != nil {
        panic(err)
    }
    
    ranges := strings.SplitAfter(line, ",")

    var count int64 = 0

    for _, values := range ranges {
        endpoints := strings.SplitAfter(values, "-")
	start, err := strconv.ParseInt(strings.TrimSuffix(endpoints[0], "-"), 10, 64)
    	if err != nil {
            panic(err)
    	}
	end, err := strconv.ParseInt(strings.TrimSuffix(endpoints[1], ","), 10, 64)
    	if err != nil {
            panic(err)
    	}
	for i := start; i <= end; i++ {
	    str := strconv.FormatInt(i, 10)
	    digits := len(str)
    	    for k := digits / 2; k > 0; k-- {
	        if(digits % k == 0){
	   	    flag := true
		    for j := 0; j < k; j++ {
		  	symbol := str[j]
			for m := k + j; m < digits; m = m + k {
			    if(str[m] != symbol){
			    	flag = false;
				break
			    }
			} // m
			if(flag == false){
			    break
			}
		    } // j
		    if(flag){
		        count = count + i
			break
		    }
 	        }
    	    } // k
	} // i
    }

    fmt.Println(count)
}
