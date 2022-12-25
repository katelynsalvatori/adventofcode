package pkg

import (
    "fmt"
    "os"
    "strings"
)

func Filename(day int) string {
    return fmt.Sprintf("inputs/%d.txt", day)
}

func ReadFileLinesForDay(day int) ([]string, error) {
    fileBytes, err := os.ReadFile(Filename(day))
    if err != nil {
        return nil, err
    }
    return strings.Split(string(fileBytes), "\n"), nil
}
