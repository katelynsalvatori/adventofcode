package pkg

import (
    "reflect"
    "testing"
)

func TestItemToPriority(t *testing.T) {
    tests := []struct {
        name string
        item RucksackItem
        want int
    }{
        {
            name: "uppercase letter should be greater than 26",
            item: 'B',
            want: 28,
        },
        {
            name: "lowercase a is 1",
            item: 'a',
            want: 1,
        },
        {
            name: "lowercase z is 26",
            item: 'z',
            want: 26,
        },
        {
            name: "uppercase A is 27",
            item: 'A',
            want: 27,
        },
        {
            name: "uppercase Z is 52",
            item: 'Z',
            want: 52,
        },
        {
            name: "lowercase letter is between 1 and 26",
            item: 'k',
            want: 11,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := tt.item.Priority(); got != tt.want {
                t.Errorf("Priority() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSplitRucksackItems(t *testing.T) {
    tests := []struct {
        name  string
        items []RucksackItem
        want1 []RucksackItem
        want2 []RucksackItem
    }{
        {
            name:  "empty",
            items: []RucksackItem{},
            want1: []RucksackItem{},
            want2: []RucksackItem{},
        },
        {
            name:  "full",
            items: []RucksackItem{'A', 'a', 'b', 'R', 'k', 'Z'},
            want1: []RucksackItem{'A', 'a', 'b'},
            want2: []RucksackItem{'R', 'k', 'Z'},
        },
        {
            name:  "from string",
            items: []RucksackItem("AabRkZ"),
            want1: []RucksackItem{'A', 'a', 'b'},
            want2: []RucksackItem{'R', 'k', 'Z'},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, got1 := SplitRucksackItems(tt.items)
            if !reflect.DeepEqual(got, tt.want1) {
                t.Errorf("SplitRucksackItems() got = %v, want %v", got, tt.want1)
            }
            if !reflect.DeepEqual(got1, tt.want2) {
                t.Errorf("SplitRucksackItems() got1 = %v, want %v", got1, tt.want2)
            }
        })
    }
}

func TestFindBadge(t *testing.T) {
    tests := []struct {
        name       string
        badgeGroup []Rucksack
        want       RucksackItem
    }{
        {
            name: "test case",
            badgeGroup: []Rucksack{
                {
                    AllItems: []RucksackItem("vJrwpWtwJgWrhcsFMMfFFhFp"),
                },
                {
                    AllItems: []RucksackItem("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"),
                },
                {
                    AllItems: []RucksackItem("PmmdzqPrVvPwwTWBwg"),
                },
            },
            want: 'r',
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := FindBadge(tt.badgeGroup); got != tt.want {
                t.Errorf("FindBadge() = %v, want %v", got, tt.want)
            }
        })
    }
}
