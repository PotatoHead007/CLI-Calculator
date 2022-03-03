package main

import (
    "fmt"
    "log"
    "os"
    "bufio"
    "strings"

    "github.com/urfave/cli/v2"
    regexp "github.com/dlclark/regexp2"
)

var modes = []string{
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
            var _, err = fmt.Scanln(&chosenMode)
            if err != nil {
                log.Println("Failiure to read from user input. Input may be of incorrect type.")
                return err
            }
            switch chosenMode {
            case 1:
                var reader = bufio.NewReader(os.Stdin)
                fmt.Println("Basic Arithmetics selected.")
                fmt.Println("What equation do you want to solve?")
                var input string
                input, err := reader.ReadString('\n')
                if err != nil {
                    return err
                }
                arithProcess(nil, []string{strings.TrimSpace(input)})
            case 2:
                fmt.Println("Comparisons selected.")
                compProcess(c.Args())
            case 3: 
                fmt.Println("Fractions, Decimals, & Percentages selected.")
                fdpProcess(c.Args())
            case 4:
                fmt.Println("Trigonometry selected")
                trigProcess(c.Args())
            default:
                log.Fatalf("%v is not a valid choice of mode currently.\n", chosenMode)
            }
            return nil
        },
        Commands: []*cli.Command{
			{
				Name:    "arithmics",
				Aliases: []string{"a", "1"},
				Usage:   "Quickly compute arithmic problems without using the interactive mode.",
				Action: func(c *cli.Context) error {
					arithProcess(c.Args(), nil)
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

    var err = app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}

func arithProcess(quickArgs cli.Args, interactiveArgs []string) {
    var input string
    if quickArgs == nil && interactiveArgs == nil {
        log.Fatalln("Internal error in arithmetic module: both quickArgs and interactiveArgs were nil.")
    } else if quickArgs != nil && interactiveArgs != nil {
        log.Fatalln("Internal error in arithmetic module: both quickArgs and interactiveArgs were NOT nil.")
    } else if quickArgs == nil {
        input = interactiveArgs[0]
    } else {
        input = quickArgs.Get(0)
    }
    fmt.Println(input)
    
    fmt.Println("numbers beep boop")
    var arithRegexVerify = regexp.MustCompile(`(?<!.|\n)(\d+) ?((?:[\+\-]|(?:[\*][\*]?))|(?:[\/][\/]?)) ?(\d+)(?: ?((?:[\+\-]|(?:[\*][\*]?))|(?:[\/][\/]?)) ?(\d+))*(?!.|\n)`, 0)
    var match, err = arithRegexVerify.FindStringMatch(input)
    if match == nil {
        log.Fatalln("Invalid input in arithmetic module: input did not match defined regexp.")
    }
    var matchGroupCaptues = match.Groups()
    if err != nil {
        log.Fatal(err)
    }
    for _, group := range matchGroupCaptues {
        for _, capture := range group.Captures {
            fmt.Printf("%v is '%s'\n", capture.Index, capture.String())
        }
    }
    
    fmt.Println("something")
    /*log.Println(isMatch)
    log.Println(interactiveArgs[0])*/
    fmt.Println(input)
}

func compProcess(args cli.Args) {
    fmt.Println("comparing beep boop")
} 

func fdpProcess(args cli.Args) {
    fmt.Println("fractions beep boop")
}

func trigProcess(args cli.Args) {
    fmt.Println("triangles beep boop")
}