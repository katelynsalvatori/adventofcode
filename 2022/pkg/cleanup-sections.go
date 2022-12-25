package pkg

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Assignment struct {
    MinSection int // inclusive
    MaxSection int // inclusive
}

type AssignmentPair struct {
    Assignment1 Assignment
    Assignment2 Assignment
}

func (ap AssignmentPair) PairFullyOverlaps() bool {
    if ap.Assignment1.MinSection >= ap.Assignment2.MinSection &&
        ap.Assignment1.MaxSection <= ap.Assignment2.MaxSection {
        return true
    }

    if ap.Assignment2.MinSection >= ap.Assignment1.MinSection &&
        ap.Assignment2.MaxSection <= ap.Assignment1.MaxSection {
        return true
    }

    return false
}

func (ap AssignmentPair) PairPartiallyOverlaps() bool {
    if ap.PairFullyOverlaps() {
        return true
    }

    return ap.Assignment1.Contains(ap.Assignment2.MinSection) ||
        ap.Assignment1.Contains(ap.Assignment2.MaxSection) ||
        ap.Assignment2.Contains(ap.Assignment1.MinSection) ||
        ap.Assignment2.Contains(ap.Assignment1.MaxSection)
}

func (a Assignment) Contains(section int) bool {
    return section >= a.MinSection && section <= a.MaxSection
}

func ParseFileToAssignmentPairs(filename string) ([]AssignmentPair, error) {
    fileBytes, err := os.ReadFile(filename)

    if err != nil {
        return nil, err
    }

    return ParseLinesToAssignmentPairs(strings.Split(string(fileBytes), "\n")), nil
}

func ParseLinesToAssignmentPairs(lines []string) []AssignmentPair {
    var assignmentPairs []AssignmentPair

    for _, line := range lines {
        assignmentPairs = append(assignmentPairs, *ParseLineToAssignmentPair(line))
    }

    return assignmentPairs
}

// ParseLineToAssignmentPair expects a string of the following form: 51-88,52-87
// where each member of the pair is delimited by a comma
func ParseLineToAssignmentPair(line string) *AssignmentPair {
    splitLine := strings.Split(line, ",")
    return &AssignmentPair{
        Assignment1: *ParseStringToAssignment(splitLine[0]),
        Assignment2: *ParseStringToAssignment(splitLine[1]),
    }
}

// ParseStringToAssignment expects a string of the following form: 2-4
// where 2 is the inclusive minimum and 4 is the inclusive maximum
func ParseStringToAssignment(assignmentString string) *Assignment {
    splitString := strings.Split(assignmentString, "-")
    // Normally I'd handle the error here but the challenge input should be valid
    min, _ := strconv.Atoi(splitString[0])
    max, _ := strconv.Atoi(splitString[1])
    return &Assignment{
        MinSection: min,
        MaxSection: max,
    }
}

func CountFullOverlaps(assignmentPairs []AssignmentPair) int {
    count := 0

    for _, assignmentPair := range assignmentPairs {
        if assignmentPair.PairFullyOverlaps() {
            count += 1
        }
    }

    return count
}

func CountPartialOverlaps(assignmentPairs []AssignmentPair) int {
    count := 0

    for _, assignmentPair := range assignmentPairs {
        if assignmentPair.PairPartiallyOverlaps() {
            count += 1
        }
    }

    return count
}

func GetFormattedOverlapCounts(assignmentPairs []AssignmentPair) string {
    return fmt.Sprintf("The number of overlapping assignments is %d", CountFullOverlaps(assignmentPairs))
}

func GetFormattedPartialOverlapCounts(assignmentPairs []AssignmentPair) string {
    return fmt.Sprintf("The number of partial overlapping assignments is %d", CountPartialOverlaps(assignmentPairs))
}
