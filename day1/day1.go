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

    value := 50
    count := 0

    for scanner.Scan() {
        line := scanner.Text()
	char := line[0]
	b := []byte(line)
	b[0] = ' '
	line = string(b)
	num, err := strconv.ParseInt(strings.TrimSpace(line), 10, 64)
        if err != nil {
            panic(err)
        }
	if(char == 'L'){
	    num = num * -1
	}
	value = value + int(num)
	incr := 0
	if(char == 'R'){
	    incr = value / 100
	}else{
	    if(value <= 0){
		incr = -1 * value / 100 + 1
		if(value - int(num) == 0){
		    incr = incr - 1
		}
	    }
	}
	count = count + incr
	value = ((value % 100) + 100) % 100
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    fmt.Println(count)
}
