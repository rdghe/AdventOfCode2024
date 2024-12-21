package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"

	internalcli "AdventOfCode2024/internal/cli"
)

func main() {
	app := &cli.Command{
		Name:     "aoc2024",
		Usage:    "Solve Advent of Code 2024 problems",
		Commands: internalcli.RegisterCommands(),
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
