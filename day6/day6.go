package main

import ( 
	"fmt"
	"bufio"
	"os"
    "strconv"
)

// https://stackoverflow.com/a/18479916
// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func main() {
    lines, err := readLines("input.txt")
    if err != nil {
        panic(err)
    }

    var total int64 = 0

    current_op := lines[len(lines) - 1][0]
    var current_total int64 = 0
    if(current_op == '*'){
        current_total++
    }

    for i := 0; i < len(lines[0]); i++ {
        var num_str []byte
        for j := 0; j < len(lines) - 1; j++ {
            if(lines[j][i] != ' '){
                num_str = append(num_str, byte(lines[j][i]))
            }
        }
        // end of problem
        if(len(num_str) == 0){
            total += current_total;
            current_op = lines[len(lines) - 1][i + 1]
            current_total = 0
            if(current_op == '*'){
                current_total++
            }
        }else{
            num, err := strconv.ParseInt(string(num_str), 10, 64)
            if err != nil {
                panic(err)
            }
            if(current_op == '*'){
                current_total *= num
            }else{
                current_total += num
            }
        }
    }

    total += current_total

    fmt.Println(total)
}
