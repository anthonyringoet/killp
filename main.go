package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	listFlag := flag.Bool("list", false, "List active processes")
	flag.Parse()

	var filter string
	if flag.NArg() > 0 {
		filter = flag.Arg(0)
	}

	if *listFlag {
		processes, err := GetProcesses()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Add headers to the output
		fmt.Println("PID\tExecutable\tCommandLine\tTTY")

		for _, process := range processes {
			if filter == "" || strings.Contains(process.Executable, filter) {
				fmt.Printf("%d\t%s\t%s\t%s\n", process.Pid, process.Executable, process.CommandLine, process.Tty)
			}
		}
	} else {
		if len(flag.Args()) == 0 {
			fmt.Println("Please provide a process ID or use -list flag")
			os.Exit(1)
		}

		pid, err := strconv.Atoi(flag.Arg(0))
		if err != nil {
			fmt.Printf("Invalid process ID: %s\n", flag.Arg(0))
			os.Exit(1)
		}

		process, err := FindProcessByPid(pid)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if err := process.Kill(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Process %d (%s) killed.\n", process.Pid, process.Executable)
	}
}
