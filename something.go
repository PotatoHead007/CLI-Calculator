package main

import (
    "fmt"
    "log"
    "os"
    "bufio"
    "strings"
    "sort"

    "github.com/urfave/cli/v2"
    regexp "github.com/dlclark/regexp2"
)


const DEBUG = true


// function for printing debug info if DEBUG is true.
func printDebug(v ...interface{}) {
    if DEBUG {
        fmt.Println(v...)
    }
}
func printfDebug(format string, v ...interface{}) {
    if DEBUG {
        fmt.Printf(format, v...)
    }
}

var modes = []string{
    "\033[1;34m1:\033[0m Basic Arithmetics",
    "\033[1;34m2:\033[0m Comparisons",
    "\033[1;34m3:\033[0m Fractions, Decimals, and Percentages",
    "\033[1;34m4:\033[0m Trigonometry",
}

// custom type for type sort.Interface
type captureSort []regexp.Capture
func (capt captureSort) Len() int {
    return len(capt)
}
func (capt captureSort) Swap(one, two int) {
    capt[one], capt[two] = capt[two], capt[one]
}
func (capt captureSort) Less(one, two int) bool {
    return capt[one].Index < capt[two].Index
}

func main() {
    var app = &cli.App{
        Name: "cli-calculator",
        Description: "A CLI Calculator.",
        Action: func(c *cli.Context) error {
            fmt.Println("\033[1;34mWelcome to CLI-Calculator.\033[0m")
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
				Name:    "arithmetic",
				Aliases: []string{"a", "1"},
				Usage:   "Quickly compute arithmetic problems without using the interactive mode.",
				Action: func(c *cli.Context) error {
					arithProcess(c.Args(), nil)
					return nil
				},
                SkipFlagParsing: true,
			},
			{
				Name:    "comparisons",
				Aliases: []string{"c", "2"},
				Usage:   "Quickly compare values without using the interactive mode.",
				Action: func(c *cli.Context) error {
					compProcess(c.Args())
					return nil
				},
                SkipFlagParsing: true,
			},
			{
				Name:    "fdp",
				Aliases: []string{"2"},
				Usage:   "Quickly compute fractions, decimals or percentages without using the interactive mode.",
				Action: func(c *cli.Context) error {
					fdpProcess(c.Args())
					return nil
				},
                SkipFlagParsing: true,
			},
			{
				Name:    "trigonometry",
				Aliases: []string{"t", "4"},
				Usage:   "Quickly compute trigonometry-related problems without using the interactive mode.",
				Action: func(c *cli.Context) error {
					trigProcess(c.Args())
					return nil
				},
                SkipFlagParsing: true,
			},
		},
    }

    var err = app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}

// function for processing arithmetic mode input
// quickArgs is for input from command line arguments,
// and interactiveArgs is for input in interactive mode.
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
    printDebug(input)
    
    var arithRegexVerify = regexp.MustCompile(`(?<!.|\n)(\(*)*(-?\d+)(\)*)* ?((?:[\+\-]|(?:[\*][\*]?))|(?:[\/][\/]?)) ?(\(*)*(-?\d+)(\)*)*(?: ?((?:[\+\-]|(?:[\*][\*]?))|(?:[\/][\/]?)) ?(\(*)*(-?\d+)(\)*)*)*(?!.|\n)`, 0)
    var match, err = arithRegexVerify.FindStringMatch(input)
    if match == nil {
        log.Fatalln("Invalid input in arithmetic module: input did not match defined regexp.")
    }
    var matchGroupCaptues = match.Groups()
    if err != nil {
        log.Fatal(err)
    }

    var validCaptures []regexp.Capture
    for gIdx, group := range matchGroupCaptues {
        for Idx, capture := range group.Captures {
            printfDebug("%v is '%s' (%v|%v)\n", capture.Index, capture.String(), gIdx, Idx)
            if gIdx == 0 && Idx == 0 {
                printDebug("^ not appended (first capture)")
                continue
            } else if capture.Length <= 0 {
                printDebug("^ not appended (length is 0)")
            } else {
                validCaptures = append(validCaptures, capture)
            }
        }
    }
    printDebug("\n")
    for _, capt := range validCaptures {
        printfDebug("%s : %v : %v\n", capt.String(), capt.Index, capt.Length)
    }
    sort.Sort(captureSort(validCaptures))
    printDebug("\n")
    for _, capt := range validCaptures {
        printfDebug("%s : %v : %v\n", capt.String(), capt.Index, capt.Length)
    }
    
    printDebug(input)
}

func compProcess(args cli.Args) {
    fmt.Println("incomplete")
} 

func fdpProcess(args cli.Args) {
    fmt.Println("incomplete")
}

func trigProcess(args cli.Args) {
    fmt.Println("incomplete")
}