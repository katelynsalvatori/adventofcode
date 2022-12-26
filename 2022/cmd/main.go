package main

import (
    "fmt"
    "github.com/katelynsalvatori/advent2022/pkg"
)

func main() {
    fileString, err := pkg.ReadFileForDay(6)

    if err != nil {
        fmt.Errorf(err.Error())
    }

    fmt.Printf("%d", pkg.FindStartOfMessageMarker(fileString))
}
