package main

import (
    "fmt"
    "log"
    "os"
    "regexp"

    "github.com/urfave/cli/v2"
)

var modes = {
    "1: Basic Arithmetics",
    "2: Comparisons",
    "3: Fractions, Decimals, and Percentages",
    "4: Trigonometry",
}

func main() {
    var app = &cli.App{
        Name: "cli-calculator",
        Description: "A CLI Calculator.",
        Action: func(c *cli.Context) error {
            fmt.Println("Welcome to CLI-Calculator.")
            fmt.Println("Please select a mode to continue:")
            for _, mode := range modes {
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
        },
        Commands: []*cli.Command{
			{
				Name:    "arithmics",
				Aliases: []string{"a", "1"},
				Usage:   "Quickly compute arithmic problems without using the interactive mode.",
				Action: func(c *cli.Context) error {
					arithProcess(c.Args())
					return nil
				},
			},
			{
				Name:    "comparisons",
				Aliases: []string{"c", "2"},
				Usage:   "Quickly compare values without using the interactive mode.",
				Action: func(c *cli.Context) error {
					compProcess(c.Args())
					return nil
				},
			},
			{
				Name:    "fdp",
				Aliases: []string{"2"},
				Usage:   "Quickly compute fractions, decimals or percentages without using the interactive mode.",
				Action: func(c *cli.Context) error {
					fdpProcess(c.Args())
					return nil
				},
			},
			{
				Name:    "trigonometry",
				Aliases: []string{"t", "4"},
				Usage:   "Quickly compute trigonometry-related problems without using the interactive mode.",
				Action: func(c *cli.Context) error {
					trigProcess(c.Args())
					return nil
				},
			},
		},
    }
}

func arithProcess(args cli.Args) {
    arithRegexVerify = regexp.MustCompile(`(?<!.)\d+ ?(?:(?:[\+\-]|(?:[\*][\*]?))|(?:[\/][\/]?)) ?\d+(?: ?(?:(?:[\+\-]|(?:[\*][\*]?))|(?:[\/][\/]?)) ?\d+)*(?!.)`)
    var matches = arithRegexVerify.FindString(args[0])
    
    fmt.Println("something")
    for idx, match := range matches {
        fmt.Printf("%v == %s\n", idx, match)
    }
}

func compProcess(args cli.Args) {

} 

func fdpProcess(args cli.Args) {

}

func trigProcess(args cli.Args) {

}