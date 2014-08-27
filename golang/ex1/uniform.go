package main

import (
    "fmt"
    "math/rand"
    "time"
    "errors"
)

func main() {
    rs, err := RandInts(10, 20)
    if err != nil {
            panic(err)
    }
    fmt.Println(rs)
}

func RandInts(min, max int) ([]int, error) {
    if max <= min {
        return nil, errors.New("RandInts: max must be greater than min")
    }
    // Make a slice of ints sized for inclusive of min and max
    rs := make([]int, max - min + 1)
    // Initialise with all the numbers between min and max
    for i, _ := range rs {
        rs[i] = min + i
    }
    rand.Seed(time.Now().UnixNano())
    // Now shuffle them
    for i, _ := range rs {
        j := rand.Intn(len(rs) - 1)
        rs[i], rs[j] = rs[j], rs[i]
    }
    return rs, nil
}