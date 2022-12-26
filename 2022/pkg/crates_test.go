package pkg

import (
    "github.com/golang-collections/collections/stack"
    "reflect"
    "testing"
)

func TestParseLineToInstructions(t *testing.T) {
    tests := []struct {
        name string
        line string
        want Instruction
    }{
        {
            name: "test case",
            line: "move 2 from 8 to 7",
            want: Instruction{
                NumberToMove: 2,
                Source:       8,
                Destination:  7,
            },
        },
        {
            name: "another test case",
            line: "move 36 from 8 to 6",
            want: Instruction{
                NumberToMove: 36,
                Source:       8,
                Destination:  6,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ParseLineToInstruction(tt.line); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ParseLineToInstruction() = %v, want %v", got, tt.want)
            }
        })
    }
}

func createCrateStack(crates []string) stack.Stack {
    var crateStack stack.Stack
    for _, crate := range crates {
        crateStack.Push(crate)
    }
    return crateStack
}

func createCrateStacks(cratesList [][]string) []stack.Stack {
    var crateStacks []stack.Stack
    for _, crates := range cratesList {
        crateStacks = append(crateStacks, createCrateStack(crates))
    }
    return crateStacks
}

func TestCrateMovement_ExecuteInstructions(t *testing.T) {
    type fields struct {
        CrateStacks  []stack.Stack
        Instructions []Instruction
    }
    tests := []struct {
        name   string
        fields fields
        want   string
    }{
        {
            name: "test1",
            fields: fields{
                CrateStacks: createCrateStacks([][]string{
                    {"A", "B", "C"},
                    {"D", "E"},
                }),
                Instructions: []Instruction{
                    {
                        NumberToMove: 2,
                        Source:       1,
                        Destination:  2,
                    },
                },
            },
            want: "AB",
        },
        {
            name: "test2",
            fields: fields{
                CrateStacks: createCrateStacks([][]string{
                    {"A", "B", "C"},
                    {"D", "E"},
                    {"F", "G", "H"},
                }),
                Instructions: ParseLinesToInstructions([]string{
                    "move 1 from 2 to 1",
                    "move 2 from 3 to 2",
                    "move 1 from 2 to 1",
                }),
            },
            want: "GHF",
        },
        {
            name: "test3",
            fields: fields{
                CrateStacks: createCrateStacks([][]string{
                    {"A", "B"},
                    {"C"},
                }),
                Instructions: ParseLinesToInstructions([]string{
                    "move 1 from 2 to 1",
                }),
            },
            want: "C",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            cm := CrateMovement{
                CrateStacks:  tt.fields.CrateStacks,
                Instructions: tt.fields.Instructions,
            }

            newCm := cm.ExecuteInstructions()
            got := newCm.GetTopCrates()

            if got != tt.want {
                t.Errorf("ExecuteInstructions() = %v, want %v", got, tt.want)
            }

            if len(newCm.Instructions) != 0 {
                t.Errorf("Instructions length should be 0")
            }
        })
    }
}

func Test_getNumberOfStacks(t *testing.T) {
    tests := []struct {
        name string
        line string
        want int
    }{
        {
            name: "test case",
            line: " 1   2   3",
            want: 3,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := getNumberOfStacks(tt.line); got != tt.want {
                t.Errorf("getNumberOfStacks() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestParseLinesToStacks(t *testing.T) {
    type args struct {
        lines          []string
        numberOfStacks int
    }
    tests := []struct {
        name string
        args args
        want []stack.Stack
    }{
        {
            name: "test case",
            args: args{
                lines: []string{
                    "[A]",
                    "[B]     [C]",
                    "[D] [E] [F]",
                },
                numberOfStacks: 3,
            },
            want: createCrateStacks([][]string{
                {"D", "B", "A"},
                {"E"},
                {"F", "C"},
            }),
        },
        {
            name: "test 2",
            args: args{
                lines: []string{
                    "    [D]",
                    "[N] [C]",
                    "[Z] [M] [P]",
                },
                numberOfStacks: 3,
            },
            want: createCrateStacks([][]string{
                {"Z", "N"},
                {"M", "C", "D"},
                {"P"},
            }),
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ParseLinesToStacks(tt.args.lines, tt.args.numberOfStacks); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ParseLinesToStacks() = %v, want %v", got, tt.want)
            }
        })
    }
}
