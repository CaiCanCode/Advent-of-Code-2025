package main

import ( 
	"fmt"
	"bufio"
	"os"
    "strings"
    "strconv"
    "math"
    "github.com/draffensperger/golp"
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

func min_ops(idx int, b []int, costs []int) {
    cost := costs[idx]
    var states []int
    for i := 0; i < len(b); i++ {
        state := idx ^ b[i]
        if(costs[state] > costs[idx] + 1){
            costs[state] = cost + 1
            states = append(states, state)
        }
    }
    for i := 0; i < len(states); i++ {
       min_ops(states[i], b, costs) 
    }
}

func main() {
    lines, err := readLines("input.txt")
    if err != nil {
        panic(err)
    }

    var total int64 = 0
    var count int64 = 0

    for i := 0; i < len(lines); i++ {
        // part 1
        feilds := strings.Split(lines[i], " ")
        var lights []bool
        for j := 1; j < len(feilds[0]) - 1; j++ {
            lights = append(lights, feilds[0][j] == '#')
        }
        idx := 0
        shift := len(lights) - 1
        for j := 0; j < len(lights); j++ {
            if(lights[j]){
                idx += (1 << shift)
            }
            shift--
        }
        var states []int
        for j := 0; j < (1 << len(lights)); j++ {
            states = append(states, math.MaxInt - 1)
        }
        states[0] = 0
        var buttons []int
        for j := 1; j < len(feilds) - 1; j++ {
            input := feilds[j][1 : len(feilds[j]) - 1]
            inputs := strings.Split(input, ",")
            b := 0
            for k := 0; k < len(inputs); k++ {
                index, err := strconv.ParseInt(inputs[k], 10, 64)
                if err != nil {
                    panic(err)
                }
                b += (1 << (len(lights) - int(index) - 1))
            }
            buttons = append(buttons, b)
        }
        min_ops(0, buttons, states)
        total += int64(states[idx])
        // part 2
        var equals []float64 // size len(lights)
        target_str := feilds[len(feilds) - 1]
        target_str = target_str[1 : len(target_str) - 1]
        targets := strings.Split(target_str, ",")
        for j := 0; j < len(lights); j++ {
            // should be integer values
            num, err := strconv.ParseInt(targets[j], 10, 64)
            if err != nil {
                panic(err)
            }
            // convert to float
            equals = append(equals, float64(num))
        }
        var sum []float64 // size len(buttons)
        for j := 0; j < len(buttons); j++ {
            sum = append(sum, float64(1))
        }
        var matrix []float64 // size len(lights) x len(buttons)
        for j := 0; j < len(lights); j++ {
            mask := (1 << (len(lights) - 1 - j))
            for k := 0; k < len(buttons); k++ {
                num := 0
                if((buttons[k] & mask) != 0){
                    num = 1
                }
                matrix = append(matrix, float64(num))
            }
        }
        /*
         *  minimize 
         *      sum^T * x 
         *  subject to
         *      A * x == equals
         *      x >= 0
         */
        ilp := golp.NewLP(0, len(buttons))
        ilp.SetObjFn(sum)
        for j := 0; j < len(lights); j++ {
            Aj := matrix[j * len(buttons) : (j + 1) * len(buttons)]
            ilp.AddConstraint(Aj, golp.EQ, equals[j])
        }
        for j := 0; j < len(buttons); j++ {
            ilp.SetInt(j, true)
        }
        ilp.Solve()
        count += int64(math.Round(ilp.Objective()))
    }

    fmt.Println(total)
    fmt.Println(count)

}
