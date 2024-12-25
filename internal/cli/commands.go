package cli

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"AdventOfCode2024/internal/day1"
	"AdventOfCode2024/internal/day2"
	"AdventOfCode2024/internal/day3"
	"AdventOfCode2024/internal/input"
	"AdventOfCode2024/internal/xerrors"
)

type inputData struct {
	session string
	day     string
}

type day struct {
	part2   bool
	example bool
}

var days = make(map[int]*day)

func init() {
	// Populate the 25 days
	for i := 1; i <= 25; i++ {
		days[i] = &day{}
	}
}

var solvers = map[int]func(bool, bool) error{
	1: day1.Solve,
	2: day2.Solve,
	3: day3.Solve,
	// Add more solvers as needed
}

func solveDay(dayNumber int, part2, example bool) error {
	if solver, exists := solvers[dayNumber]; exists {
		return solver(part2, example)
	}
	return xerrors.NotImplementedError
}

func RegisterCommands() []*cli.Command {
	var commands []*cli.Command
	var iData inputData

	for dayNumber := 1; dayNumber <= 25; dayNumber++ {
		commands = append(commands, &cli.Command{
			Name:  fmt.Sprintf("day%d", dayNumber),
			Usage: fmt.Sprintf("Solve day %d exercise", dayNumber),
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:        "part2",
					Usage:       "Solve part 2 of the problem",
					Destination: &days[dayNumber].part2,
				},
				&cli.BoolFlag{
					Name:        "example",
					Usage:       "Use the example input",
					Destination: &days[dayNumber].example,
				},
			},
			Action: func(ctx context.Context, cmd *cli.Command) error {
				return solveDay(dayNumber, days[dayNumber].part2, days[dayNumber].example)
			},
		})
	}

	commands = append(commands, &cli.Command{
		Name:  "get-input",
		Usage: "Download an exercise input file for the given day",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "session",
				Aliases:     []string{"s"},
				Sources:     cli.EnvVars("SESSION"),
				Required:    true,
				Usage:       "Specify the session such that the user can obtain its tailored input",
				Destination: &iData.session,
			},
			&cli.StringFlag{
				Name:        "day",
				Aliases:     []string{"d"},
				Usage:       "Specify the day for which to download the input file",
				Value:       "1",
				Destination: &iData.day,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return input.Download(iData.session, iData.day)
		},
	})

	return commands
}
