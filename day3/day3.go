package main

import ( 
	"fmt"
	"bufio"
	"os"
    "strconv"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var total int64 = 0

    for scanner.Scan() {
        line := scanner.Text()
        cur_str := []byte(line[len(line) - 12 : len(line)])
        for i := len(line) - 13; i >= 0; i-- {
            if(line[i] >= cur_str[0]){
                // Reddit user u/IsatisCrucifer
                // helped deciding which digit to remove
                var idx int
                for idx = 0; idx < 11; idx++ {
                    if(cur_str[idx] < cur_str[idx + 1]){
                        break
                    }
                }
                // remove idx (right shift)
                for j := idx; j > 0; j-- {
                    cur_str[j] = cur_str[j - 1]
                }
                // pre-pend line[i]
                cur_str[0] = line[i]
            }
        }
        current, err := strconv.ParseInt(string(cur_str), 10, 64)
        if err != nil {
            panic(err)
        }
        total += current
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    fmt.Println(total)
}
