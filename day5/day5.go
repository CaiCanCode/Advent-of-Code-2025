package main

import ( 
	"fmt"
	"bufio"
	"os"
    "strconv"
    "strings"
    "sort"
    "math"
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

    type pair struct {
        values [2]int64
    }

    var ranges []pair
    idx := 0

    for len(lines[idx]) != 0 {
        line := strings.Split(lines[idx], "-")
        var err error
        var p pair
        p.values[0], err = strconv.ParseInt(line[0], 10, 64)
        if err != nil {
            panic(err)
        }
        p.values[1], err = strconv.ParseInt(line[1], 10, 64)
        if err != nil {
            panic(err)
        }
        ranges = append(ranges, p)
        idx++
    }

    // part 1 counting variable 
    var fresh int64 = 0

    for i := idx + 1; i < len(lines); i++ {
        id, err := strconv.ParseInt(lines[i], 10, 64)
        if err != nil {
            panic(err)
        }
        for j := 0; j < idx; j++ {
            if(id >= ranges[j].values[0]){
                if(id <= ranges[j].values[1]){
                    fresh++
                    break
                }
            }
        }
    }

    // part 2 counting variable
    var count int64 = 0

    // sort
    sort.Slice(ranges, func(i, j int) bool {
        return (ranges[i].values[0] - ranges[j].values[0]) < 0
    })

    // process
    for i := 0; i < idx - 1; i++ {
        // check if ranges overlap
        offset := 1
        for ranges[i].values[1] >= ranges[i + offset].values[0] {
            current := float64(ranges[i].values[1])
            candidate := float64(ranges[i + offset].values[1])
            max := int64(math.Max(current, candidate))
            ranges[i].values[1] = max
            ranges[i + offset].values[1] = max
            offset++
            if(i + offset >= idx){
                break;
            }
        }
    }

    // reduce
    index := 0
    for i := 1; i < idx; i++ {
        if(ranges[i].values[1] != ranges[index].values[1]){
            index++
            ranges[index] = ranges[i]
        }
    }
    index++

    // count
    for i := 0; i < index; i++ {
        count += ranges[i].values[1] - ranges[i].values[0] + 1
    }

    fmt.Println(fresh)
    fmt.Println(count)
}
