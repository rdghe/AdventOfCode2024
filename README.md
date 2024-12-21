# Advent Of Code 2024 CLI Tool
This is a CLI tool built to solve [2024's edition of Advent of Code](https://adventofcode.com/2024) challenges in Go. It provides subcommands for solving each day's exercises, downloading input files, and more.

## Features
- Solve any day’s challenge with a specific subcommand (e.g., day1, day2).
- Support for running both Part 1 and Part 2 of each problem.
- Ability to toggle between real and example inputs.
- Fetch puzzle inputs directly using your session token.

---

## How It's Built
- CLI Framework: urfave/cli for command-line interface management.
- **Modular Design**:
  - Each day’s logic is implemented in its own package (e.g., internal/day1, internal/day2).
  - Shared functionality (like input fetching) is implemented in reusable modules.
- **Dynamic Command Registration**: Subcommands for each day are generated programmatically to avoid repetitive boilerplate.
- **Custom Error Handling**: Uses a centralized error-handling package for consistent error messages.

## Project Structure

```shell
AdventOfCode2024/
├── cmd/
│   └── main.go         # CLI entry point
├── internal/
│   ├── day1/           # Day 1 logic
│   │   └── day1.go
│   ├── day2/           # Day 2 logic
│   │   └── day2.go
│   ├── ...
│   │
│   ├── input/          # Input downloading logic
│   │   └── downloadinput.go
│   ├── cli/            # Command registration logic
│   │   └── commands.go
│   └── xerrors/        # Custom error definitions
│       └── errors.go

```

---

## Installation
### Using go build
1. Clone the repository:

```shell
git clone https://github.com/your-username/AdventOfCode2024.git
cd AdventOfCode2024
```

2. Build the CLI tool:
```shell
go build -o aoc2024 cmd/main.go

```

3. Run the tool:
```shell
./aoc2024

```

### Using go run
Alternatively, you can run the tool without building it:
```shell
go run cmd/main.go
```

---

## Usage
### General Help
To see all available commands, run:

```shell
./aoc2024
```
Expected output:
```text
NAME:
   aoc2024 - Solve Advent of Code 2024 problems

USAGE:
   aoc2024 [global options] [command [command options]]

COMMANDS:
   day1       Solve day 1 exercise
   day2       Solve day 2 exercise
   ... 
   day25      Solve day 25 exercise
   get-input  Download an exercise input file for the given day
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

---

## Retrieving Puzzle Inputs

To download the inputs for each problem/day, you need to do the following:

1. Find your `session` token in the `Cookies` of your browser of choice. This is required to authenticate you as a user and retrieve your personalised inputs.
2. Then, run: 
```shell
export SESSION=<your-session-string>
```
3. Following that, you can retrieve the input using the following command:
```shell
./aoc2024 get-input --day=1
```
Where the optional parameter `--day` specifies for which day to fetch the input. The default value is "1".

The session string may also be directly provided using the `--session` argument, instead of the environment variable:
```shell
./aoc2024 get-input --session YOUR_SESSION_TOKEN --day 1
```

## Solving the exercises

To solve a specific day’s problem, use the corresponding subcommand:

Example Commands
- Solve Day 1, Part 1:
```shell
./aoc2024 day1
```

- Solve Day 1, Part 2:
```shell
./aoc2024 day1 --part2
```

- Solve Day 1 using the example input:
```shell
./aoc2024 day1 --example
```

- Solve Day 2, Part 2 with example input:
```shell
./aoc2024 day2 --part2 --example
```

This will print the solution to the problem, calculated based on your personalised inputs!
