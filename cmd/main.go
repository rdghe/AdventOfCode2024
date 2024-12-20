package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"

	"AdventOfCode2024/internal/day1"
	"AdventOfCode2024/internal/input"
)

type inputData struct {
	session string
	day     string
}

type day struct {
	part    string
	example bool
}

type days struct {
	day1  day
	day2  day
	day3  day
	day4  day
	day5  day
	day6  day
	day7  day
	day8  day
	day9  day
	day10 day
	day11 day
	day12 day
	day13 day
	day14 day
	day15 day
	day16 day
	day17 day
	day18 day
	day19 day
	day20 day
	day21 day
	day22 day
	day23 day
	day24 day
	day25 day
}

func main() {
	var i inputData
	var d days

	cmd := &cli.Command{
		Name:  "aoc2024",
		Usage: "Solve Advent of Code 2024 problems",
		Flags: []cli.Flag{},
		Commands: []*cli.Command{
			{
				Name:  "test-command",
				Usage: "Run a test command",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "test-arg",
						Usage: "Test your argument here",
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					log.Println("successfully tested your command", cmd.Args().First())
					return nil
				},
			},
			{
				Name:  "get-input",
				Usage: "Download an exercise input file for the given day",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "session",
						Aliases:     []string{"s"},
						Usage:       "specify the session such that the user can obtain its tailored input",
						Sources:     cli.EnvVars("SESSION"),
						Destination: &i.session,
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "day",
						Destination: &i.day,
						Value:       "1",
						Usage:       "Specify the day for which to download the input file. Defaults to 1",
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					err := input.Download(i.session, i.day)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:  "day1",
				Usage: "solve day1 exercise",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "part",
						Value:       "1",
						Destination: &d.day1.part,
						Usage:       "Specify which part of the problem should be solved",
					},
					&cli.BoolFlag{
						Name:        "example",
						Value:       false,
						Destination: &d.day1.example,
						Usage:       "Whether to use the example input or not",
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					err := day1.Solve(d.day1.part, d.day1.example)
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
