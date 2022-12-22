package pkg

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CalorieCollection struct {
	calorieCollection []int
}

// ParseCollectionFromString expects a string that looks like this:
// 1000
// 2000
// 3000
//
// 4000
//
// 5000
// 6000
// where each line represents the calorie count of an item and each individual elves' calories are separated by blank lines
func ParseCollectionFromString(collectionString string) (*CalorieCollection, error) {
	splitString := strings.Split(collectionString, "\n")
	index := 0

	collection := []int{0}

	for _, s := range splitString {
		if s == "" {
			index += 1
			collection = append(collection, 0)
		} else {
			calorieCount, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			collection[index] = collection[index] + calorieCount
		}
	}

	return &CalorieCollection{
		calorieCollection: collection,
	}, nil
}

func ParseCollectionFromFile(collectionFileName string) (*CalorieCollection, error) {
	collectionBytes, err := os.ReadFile(collectionFileName)

	if err != nil {
		return nil, err
	}

	return ParseCollectionFromString(string(collectionBytes))
}

func (cc CalorieCollection) GetMaxCalories() int {
	sort.Ints(cc.calorieCollection)
	return cc.calorieCollection[len(cc.calorieCollection)-1]
}

func (cc CalorieCollection) GetTop3Calories() []int {
	sort.Ints(cc.calorieCollection)
	maxIndex := len(cc.calorieCollection)
	return cc.calorieCollection[maxIndex-3 : maxIndex]
}

func (cc CalorieCollection) GetTop3CaloriesSum() int {
	top3 := cc.GetTop3Calories()
	sum := 0
	for _, c := range top3 {
		sum += c
	}

	return sum
}

// GetFormattedMaxCalories is the answer to Day 1, part 1
func (cc CalorieCollection) GetFormattedMaxCalories() string {
	return fmt.Sprintf("The max amount of calories is %d\n", cc.GetMaxCalories())
}

// GetFormattedTop3CaloriesSum is the answer to Day 1, part 2
func (cc CalorieCollection) GetFormattedTop3CaloriesSum() string {
	return fmt.Sprintf("The sum of the top 3 calories is %d\n", cc.GetTop3CaloriesSum())
}
