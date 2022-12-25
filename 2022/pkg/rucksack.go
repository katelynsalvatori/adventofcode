package pkg

import (
    "fmt"
    "os"
    "strings"
    "unicode"
)

type RucksackItem rune

type Rucksack struct {
    AllItems             []RucksackItem
    Compartment1         []RucksackItem
    Compartment2         []RucksackItem
    RepeatedItem         RucksackItem
    RepeatedItemPriority int
}

type RucksackCollection struct {
    Rucksacks                 []Rucksack
    TotalRepeatedItemPriority int
    BadgePriorityTotal        int
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
    badgePriorityTotal := 0
    var currentBadgeGroup []Rucksack

    for index, line := range lines {
        rucksack := ParseStringToRucksack(line)
        currentBadgeGroup = append(currentBadgeGroup, rucksack)
        rucksacks = append(rucksacks, rucksack)
        totalRepeatedItemPriority += rucksack.RepeatedItemPriority

        if (index+1)%3 == 0 {
            badge := FindBadge(currentBadgeGroup)
            currentBadgeGroup = []Rucksack{}
            badgePriorityTotal += badge.Priority()
        }

    }

    return &RucksackCollection{
        Rucksacks:                 rucksacks,
        TotalRepeatedItemPriority: totalRepeatedItemPriority,
        BadgePriorityTotal:        badgePriorityTotal,
    }
}

// FindBadge is given 3 rucksacks and finds the badge, which is the item that appears in all three
func FindBadge(badgeGroup []Rucksack) RucksackItem {
    var firstIntersection []RucksackItem

    for _, item := range badgeGroup[0].AllItems {
        if CompartmentContains(badgeGroup[1].AllItems, item) {
            firstIntersection = append(firstIntersection, item)
        }
    }

    for _, item := range firstIntersection {
        if CompartmentContains(badgeGroup[2].AllItems, item) {
            return item
        }
    }

    return 0
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
        AllItems:             rucksackItems,
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

func (rc RucksackCollection) GetFormattedBadgePriorities() string {
    return fmt.Sprintf("The total of the priorities of the badges is %d", rc.BadgePriorityTotal)
}
