package main

import ( 
	"fmt"
	"bufio"
	"os"
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

type call struct {
    result int64
    params struct {
        row int
        col int
    }
}

// memoization as suggested by Reddit
// pointer-to-array due to size/capacity changes
// and unstable references during re-allocation
func qu_rec(lines []string, table* []call, row int, col int) int64 {
    var ret int64
    next := row
    for i := row + 1; i < len(lines); i++ {
        if(lines[i][col] == '^'){
            next = i
            break
        }
    }
    if(next == row){
        ret = 1
    }else{
        var left int64 = 0
        var right int64 = 0
        for _, val := range *table {
            if(val.params.row == next){
                if(val.params.col == col - 1){
                    left = val.result
                }
                if(val.params.col == col + 1){
                    right = val.result
                }
            }    
        }
        if(left == 0){
            left = qu_rec(lines, table, next, col - 1)
            var left_call call
            left_call.result = left
            left_call.params.row = next
            left_call.params.col = col - 1
            *table = append(*table, left_call)
        }
        if(right == 0){
            right = qu_rec(lines, table, next, col + 1)
            var right_call call
            right_call.result = right
            right_call.params.row = next
            right_call.params.col = col + 1
            *table = append(*table, right_call)
        }
        ret = left + right
    }
    var cur_call call
    cur_call.result = ret
    cur_call.params.row = row
    cur_call.params.col = col
    *table = append(*table, cur_call)
    return ret
}

func main() {
    lines, err := readLines("input.txt")
    if err != nil {
        panic(err)
    }

    var count int64 = 0
    var col int

    for i := 1; i < len(lines); i++ {
        line := []byte(lines[i])
        for j := 0; j < len(line); j++ {
            if(lines[i - 1][j] == 'S'){
                if(i == 1){
                    col = j
                }
                if(line[j] == byte('^')){
                    line[j - 1] = byte('S')
                    line[j + 1] = byte('S')
                    count++
                }else{
                    line[j] = byte('S')
                }
            }
        }
        lines[i] = string(line)
    }

    var table []call
    timelines := qu_rec(lines, &table, 0, col)

    fmt.Println(count)
    fmt.Println(timelines)
}
