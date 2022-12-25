package pkg

import (
    "fmt"
    "os"
    "strings"
    "unicode"
)

type RucksackItem rune

type Rucksack struct {
    Compartment1         []RucksackItem
    Compartment2         []RucksackItem
    RepeatedItem         RucksackItem
    RepeatedItemPriority int
}

type RucksackCollection struct {
    Rucksacks                 []Rucksack
    TotalRepeatedItemPriority int
}

func ParseFileToRucksacks(filename string) (*RucksackCollection, error) {
    rucksackBytes, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    rucksackLines := strings.Split(string(rucksackBytes), "\n")
    return ParseLinesToRucksacks(rucksackLines), nil
}

func ParseLinesToRucksacks(lines []string) *RucksackCollection {
    var rucksacks []Rucksack
    totalRepeatedItemPriority := 0

    for _, line := range lines {
        rucksack := ParseStringToRucksack(line)
        rucksacks = append(rucksacks, rucksack)
        totalRepeatedItemPriority += rucksack.RepeatedItemPriority
    }

    return &RucksackCollection{
        Rucksacks:                 rucksacks,
        TotalRepeatedItemPriority: totalRepeatedItemPriority,
    }
}

func ParseStringToRucksack(rucksackString string) Rucksack {
    rucksackItems := []RucksackItem(rucksackString)
    compartment1 := rucksackItems[:len(rucksackItems)/2]
    compartment2 := rucksackItems[len(rucksackItems)/2:]
    var repeatedItem RucksackItem

    for _, item := range compartment1 {
        if CompartmentContains(compartment2, item) {
            repeatedItem = item
            break
        }
    }

    return Rucksack{
        Compartment1:         compartment1,
        Compartment2:         compartment2,
        RepeatedItem:         repeatedItem,
        RepeatedItemPriority: repeatedItem.Priority(),
    }
}

func SplitRucksackItems(items []RucksackItem) ([]RucksackItem, []RucksackItem) {
    return items[:len(items)/2], items[len(items)/2:]
}

func CompartmentContains(compartment []RucksackItem, item RucksackItem) bool {
    for _, c := range compartment {
        if c == item {
            return true
        }
    }

    return false
}

func (item RucksackItem) Priority() int {
    if unicode.IsUpper(rune(item)) {
        // Uppercase letters' priorities start at 27
        return int(item) - int('A') + 27
    }

    // Lowercase letters' priorities start at 1
    return int(item) - int('a') + 1
}

func (rc RucksackCollection) GetFormattedTotalPriorities() string {
    return fmt.Sprintf("The total of the priorities of the repeated items is %d", rc.TotalRepeatedItemPriority)
}
