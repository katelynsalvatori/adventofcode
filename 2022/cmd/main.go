package main

import (
    "fmt"
    "github.com/katelynsalvatori/advent2022/pkg"
)

func main() {
    fileLines, err := pkg.ReadFileLinesForDay(4)

    if err != nil {
        fmt.Errorf(err.Error())
    }
    ap := pkg.ParseLinesToAssignmentPairs(fileLines)
    fmt.Println(pkg.GetFormattedPartialOverlapCounts(ap))
}
