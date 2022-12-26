package pkg

import (
    "fmt"
    "github.com/golang-collections/collections/stack"
    "strconv"
    "strings"
    "unicode"
)

type Instruction struct {
    NumberToMove int
    Source       int
    Destination  int
}

type CrateMovement struct {
    CrateStacks  []stack.Stack
    Instructions []Instruction
}

func ParseLinesToCrateMovement(lines []string) *CrateMovement {
    var delimiterIndex, numberOfStacks int

    for index, line := range lines {
        // A blank line separates the crate stacks from the instructions
        if lines[index+1] == "" {
            delimiterIndex = index
            numberOfStacks = getNumberOfStacks(line)
            break
        }
    }

    instructions := ParseLinesToInstructions(lines[delimiterIndex+1:])
    crateStacks := ParseLinesToStacks(lines[:delimiterIndex], numberOfStacks)

    return &CrateMovement{
        Instructions: instructions,
        CrateStacks:  crateStacks,
    }
}

func getNumberOfStacks(line string) int {
    stackNum, _ := strconv.Atoi(line[len(line)-1:])
    return stackNum
}

// ParseLinesToInstructions expects strings in the form of "move 2 from 8 to 7"
func ParseLinesToInstructions(lines []string) []Instruction {
    var instructions []Instruction

    for _, line := range lines {
        if line != "" {
            instructions = append(instructions, ParseLineToInstruction(line))
        }
    }

    return instructions
}

func ParseLineToInstruction(line string) Instruction {
    splitLine := strings.Split(line, " ")

    numberToMove, _ := strconv.Atoi(splitLine[1])
    source, _ := strconv.Atoi(splitLine[3])
    destination, _ := strconv.Atoi(splitLine[5])

    return Instruction{
        NumberToMove: numberToMove,
        Source:       source,
        Destination:  destination,
    }
}

func ParseLinesToStacks(lines []string, numberOfStacks int) []stack.Stack {
    lines = reverse(lines)
    crateStacks := make([]stack.Stack, numberOfStacks)
    for _, line := range lines {
        for index, character := range line {
            if unicode.IsUpper(character) {
                crateStacks[index/4].Push(string(character))
            }
        }
    }

    return crateStacks
}

func reverse(s []string) []string {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }

    return s
}

// ExecuteInstructions updates the crate stacks with the existing instructions
func (cm CrateMovement) ExecuteInstructions() *CrateMovement {
    for len(cm.Instructions) > 0 {
        instruction := cm.Instructions[0]
        cm.Instructions = cm.Instructions[1:]

        for i := 0; i < instruction.NumberToMove; i++ {
            crate := cm.CrateStacks[instruction.Source-1].Pop()
            cm.CrateStacks[instruction.Destination-1].Push(crate)
        }
    }

    return &CrateMovement{
        Instructions: []Instruction{},
        CrateStacks:  cm.CrateStacks,
    }
}

func (cm CrateMovement) GetTopCrates() string {
    var topCrates string

    for _, crateStack := range cm.CrateStacks {
        crate := crateStack.Peek()
        if crate != nil {
            topCrates += fmt.Sprintf("%v", crate)
        }
    }

    return topCrates
}
