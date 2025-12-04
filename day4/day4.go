package main

import ( 
	"fmt"
	"bufio"
	"os"
    "math"
)

// assumes lines[row][col] == @
func checkForklift(lines []string, row int, col int) (bool) {
    count := 0
    row_start := math.Max(float64(0), float64(row - 1))
    row_end := math.Min(float64(len(lines)), float64(row + 2))
    for i := int(row_start); i < int(row_end); i++ {
        col_start := math.Max(float64(0), float64(col - 1))
        col_end := math.Min(float64(len(lines[i])), float64(col + 2))
        for j := int(col_start); j < int(col_end); j++ {
            if(lines[i][j] != '.'){
                count++;
            }
        }
    }
    return count <= 4
}

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

    var old int64 = -1
    var count int64 = 0

    for old != count {
        old = count
        for i := 0; i < len(lines); i++ {
            for j := 0; j < len(lines[i]); j++ {
                if(lines[i][j] == '@'){
                    if(checkForklift(lines, i, j)){
                        count++;
                        mut := []byte(lines[i])
                        mut[j] = 'x'
                        lines[i] = string(mut)
                    }
                }
            }
        }
        for i := 0; i < len(lines); i++ {
            for j := 0; j < len(lines[i]); j++ {
                if(lines[i][j] == 'x'){
                    mut := []byte(lines[i])
                    mut[j] = '.'
                    lines[i] = string(mut)
                }
            }
        }
    }

    fmt.Println(count)
}
