# Go CLI Net Tool

In this application, we will see the creation of a basic CLI using GoLang to exercise key concepts of the language.

## Requirements

- GoLang installed on your machine
- A code editor

## How to run

1. Clone the repository
2. Open the terminal and navigate to the project folder
3. Run the command `go run main.go` to start the application

## Commands

- `ip` - Get the IP address of a domain
- `ns` - Get the Name Servers of a domain

## Flags

- `--host` - The domain to be analyzed

## Example

```bash
$ go run main.go ip --host google.com
$ go run main.go ns --host google.com
```
