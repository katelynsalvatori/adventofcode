package pkg

func FindMarker(datastream string, uniqueChars int) int {
    lastN := make([]rune, uniqueChars)

    for index, char := range datastream {
        lastN[index%uniqueChars] = char
        if index > 2 && allUnique(lastN) {
            return index + 1
        }
    }

    return 0

}

func FindStartOfPacketMarker(datastream string) int {
    return FindMarker(datastream, 4)
}

func FindStartOfMessageMarker(datastream string) int {
    return FindMarker(datastream, 14)
}

func contains(slice []rune, char rune) bool {
    for _, c := range slice {
        if c == char {
            return true
        }
    }

    return false
}

func allUnique(slice []rune) bool {
    for index, c := range slice {
        if contains(slice[index+1:], c) {
            return false
        }
    }
    return true
}
