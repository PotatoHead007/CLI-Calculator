package main

import "fmt"

var modes = {
    "1: Basic Arithmetics",
    "2: Comparisons",
    "3: Fractions, Decimals, and Percentages",
    "4: Trigonometry"
}

func main() {
    fmt.Println("Welcome to CLI-Calculator.")
    fmt.Println("Please select a mode to continue:")
    for _, mode = range modes {
        fmt.Println(mode)
    }

    var chosenMode int8
    var _, err = fmt.Scan(&chosenMode)
    if err != nil {
        fmt.Println("Failiure to read from user input. Input may be of incorrect type.")
        return
    }
    switch chosenMode {
    case 1:
        fmt.Println("Basic Arithmetics selected.")
        arithProcess()
    case 2:
        fmt.Println("Comparisons selected.")
        compProcess()
    case 3: 
        fmt.Println("Fractions, Decimals, & Percentages selected.")
        fdpProcess()
    case 4:
        fmt.Println("Trigonometry selected")
        trigProcess()
    default:
        fmt.Printf("%v is not a valid choice of mode currently.")
        return
    }
}
