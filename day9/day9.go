package main

import ( 
	"fmt"
	"bufio"
	"os"
    "strings"
    "strconv"
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

type tile struct {
    x int64
    y int64
}

func mul(a tile, b tile) float64 {
    x := math.Abs(float64(b.x - a.x))
    y := math.Abs(float64(b.y - a.y))
    return (x + float64(1)) * (y + float64(1))
}

// a point is in a polygon
// if it crosses polygon edges an odd number of times
func in_polygon(t tile, tiles []tile) bool {
    crossings := 0
    for i := 0; i < len(tiles); i++ {
        next := tiles[(i + 1) % len(tiles)]
        // Vertical edge
        if tiles[i].x == next.x {
            x := next.x
            y_ := math.Min(float64(tiles[i].y), float64(next.y))
            y := math.Max(float64(tiles[i].y), float64(next.y))
            if t.x < x && t.y > int64(y_) && t.y < int64(y) {
                crossings++
            }
        }
        // Horizontal edge ignored
    }
    return crossings % 2 == 1
}

// legal rectangle if no edge is strictly inside (a,c)
func area(a tile, c tile, tiles []tile) float64 {
    x := int64(math.Max(float64(a.x), float64(c.x)))
    x_ := int64(math.Min(float64(a.x), float64(c.x)))
    y := int64(math.Max(float64(a.y), float64(c.y)))
    y_ := int64(math.Min(float64(a.y), float64(c.y)))
    edge := false
    for i := 0; i < len(tiles); i++ {
        next := tiles[(i + 1) % len(tiles)]
        // right edge
        if(tiles[i].x >= x && tiles[i].y > y_ && tiles[i].y < y){
            if(next.x < x){
                edge = true
                break
            }
        }
        // left edge
        if(tiles[i].x <= x_ && tiles[i].y > y_ && tiles[i].y < y){
            if(next.x > x_){
                edge = true
                break
            }
        }
        // bottom edge
        if(tiles[i].y >= y && tiles[i].x > x_ && tiles[i].x < x){
            if(next.y < y){
                edge = true
                break
            }
        }
        // top edge
        if(tiles[i].y <= y_ && tiles[i].x > x_ && tiles[i].x < x){
            if(next.y > y_){
                edge = true
                break
            }
        }
    }
    if(!edge){
        var t tile
        t.x = x_ + 1
        t.y = y_ + 1
        if(in_polygon(t, tiles)){
            return mul(a, c)
        }
    }
    return float64(1)
}

func main() {
    lines, err := readLines("input.txt")
    if err != nil {
        panic(err)
    }

    var tiles []tile

    // read tiles
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
        var t tile
        t.x = x
        t.y = y
        tiles = append(tiles, t)
    }

    product := float64(1)

    // I tried to be clever and not check every pair
    // but it didn't work
    for i := 0; i < len(tiles); i++ {
        for j := 0; j < i; j++ {
            rect := area(tiles[i], tiles[j], tiles)
            product = math.Max(product, rect)
        }
    }

    fmt.Println(int64(product))

}
