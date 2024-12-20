# Advent of Code 2024

My solutions for [2024's edition of Advent of Code](https://adventofcode.com/2024).

The application is build using Golang and runs as a CLI tool.

### To run `help` the command, use:

```shell
go run ./cmd/... help
```

## Retrieving inputs

To download the inputs for each problem/day, you need to do the following:

- Find your `session` token in the `Cookies` of your browser of choice. This is required to authenticate you as a user and retrieve your personalised inputs. Then, run: 
```shell
export SESSION=<your-session-string>
```
- Following that, you can retrieve the input using the following command:
```shell
go run ./cmd/... get-input --day=1
```
Where the optional parameter `--day` specified for which day to fetch the input.
The session string may also be provided using the `--session` argument, instead of the environment variable.

## Solving the exercises

To solve the exercise for a given day, using your previously fetched personalised inputs, run the following:
```shell
go run ./cmd/... day1
```
**_Replace the name of the subcommand with the day you are looking to solve (1 to 25)_**

This will print the solution to the problem, calculated based on your personalised inputs!
