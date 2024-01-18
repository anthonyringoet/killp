# killp

killp is a simple command-line tool to manage and kill processes. It allows you to view a list of active processes and kill them by providing their process ID.

## Building

To build the killp CLI tool, run the following command:

```sh
go build -o killp main.go process.go

# build for mac+win+linux in one go
./build.sh
```

This will generate an executable named `killp` in the current directory.

You can find the pre-built binaries on the [Releases overview](https://github.com/anthonyringoet/killp/releases).

## Running

To run the killp tool without building, use the following command:

```sh
go run main.go process.go [flags] [arguments]
```

Replace `[flags]` with any flags you want to use (e.g., -list) and `[arguments]` with any required arguments for the command.

## Available commands

- `killp <process_id>`: Kill a process by its process ID
- `killp -list`: List active processes
- `killp -list <filter>`: List active processes containing the filter in their name
