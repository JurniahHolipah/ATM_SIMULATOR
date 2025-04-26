package utils

import (
    "fmt"
)

func GetInput(prompt string) string {
    var input string
    fmt.Print(prompt)
    fmt.Scanln(&input)
    return input
}

func GetInputInt(prompt string) int {
    var input int
    fmt.Print(prompt)
    fmt.Scanln(&input)
    return input
}