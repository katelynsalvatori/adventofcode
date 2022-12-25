package pkg

import (
    "reflect"
    "testing"
)

func TestParseStringToAssignment(t *testing.T) {
    tests := []struct {
        name             string
        assignmentString string
        want             *Assignment
    }{
        {
            name:             "test case",
            assignmentString: "2-4",
            want: &Assignment{
                MinSection: 2,
                MaxSection: 4,
            },
        },
        {
            name:             "test case 2",
            assignmentString: "6-6",
            want: &Assignment{
                MinSection: 6,
                MaxSection: 6,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ParseStringToAssignment(tt.assignmentString); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ParseStringToAssignment() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestAssignmentPair_PairFullyOverlaps(t *testing.T) {
    var tests = []struct {
        name        string
        Assignment1 Assignment
        Assignment2 Assignment
        want        bool
    }{
        {
            name:        "does overlap",
            Assignment1: *ParseStringToAssignment("2-8"),
            Assignment2: *ParseStringToAssignment("3-7"),
            want:        true,
        },
        {
            name:        "does not overlap",
            Assignment1: *ParseStringToAssignment("1-3"),
            Assignment2: *ParseStringToAssignment("4-6"),
            want:        false,
        },
        {
            name:        "does also overlap",
            Assignment1: *ParseStringToAssignment("35-50"),
            Assignment2: *ParseStringToAssignment("49-50"),
            want:        true,
        },
        {
            name:        "partially overlaps",
            Assignment1: *ParseStringToAssignment("1-4"),
            Assignment2: *ParseStringToAssignment("3-6"),
            want:        false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ap := AssignmentPair{
                Assignment1: tt.Assignment1,
                Assignment2: tt.Assignment2,
            }
            if got := ap.PairFullyOverlaps(); got != tt.want {
                t.Errorf("PairFullyOverlaps() = %v, want %v", got, tt.want)
            }
        })
    }
}
