package main

import ( 
	"fmt"
	"bufio"
	"os"
    "strings"
    "strconv"
    "sort"
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

type box struct {
    x int64
    y int64
    z int64
}

type circuit struct {
    b []int
}

type pair struct {
    b [2]int
}

func connect(p pair, circuits []circuit, boxes []box) int64 {
    for i := 0; i < len(circuits[p.b[0]].b); i++ {
        if(circuits[p.b[0]].b[i] == p.b[1]){
            return 0
        }
    }
    c := append(circuits[p.b[0]].b, circuits[p.b[1]].b...)
    for i := 0; i < len(circuits[p.b[0]].b); i++ {
        circuits[circuits[p.b[0]].b[i]].b = c
    }
    return boxes[p.b[0]].x * boxes[p.b[1]].x
}

func main() {
    lines, err := readLines("input.txt")
    if err != nil {
        panic(err)
    }

    var boxes []box
    var circuits []circuit
    var pairs []pair

    // read boxes
    for i := 0; i < len(lines); i++ {
        coords := strings.Split(lines[i], ",")
        x, err := strconv.ParseInt(coords[0], 10,64)
        if err != nil {
            panic(err)
        }
        y, err := strconv.ParseInt(coords[1], 10,64)
        if err != nil {
            panic(err)
        }
        z, err := strconv.ParseInt(coords[2], 10,64)
        if err != nil {
            panic(err)
        }
        var b box
        b.x = x
        b.y = y
        b.z = z
        boxes = append(boxes, b)
        var c circuit
        c.b = append(c.b, i)
        circuits = append(circuits, c)
        var p pair
        p.b[1] = i
        for j := 0; j < i; j++ {
            p.b[0] = j
            pairs = append(pairs, p)
        }
    }

    // sort
    sort.Slice(pairs, func(i, j int) bool {
        xi := boxes[pairs[i].b[1]].x - boxes[pairs[i].b[0]].x
        yi := boxes[pairs[i].b[1]].y - boxes[pairs[i].b[0]].y
        zi := boxes[pairs[i].b[1]].z - boxes[pairs[i].b[0]].z
        ri := xi * xi + yi * yi + zi * zi
        xj := boxes[pairs[j].b[1]].x - boxes[pairs[j].b[0]].x
        yj := boxes[pairs[j].b[1]].y - boxes[pairs[j].b[0]].y
        zj := boxes[pairs[j].b[1]].z - boxes[pairs[j].b[0]].z
        rj := xj * xj + yj * yj + zj * zj
        return ri < rj
    })

    var product int64

    // connect
    for i := 0; len(circuits[0].b) < len(circuits); i++ {
        tmp := connect(pairs[i], circuits, boxes)
        if(tmp != 0){
            product = tmp
        }
    }

    fmt.Println(product)

}
