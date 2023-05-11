package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Process struct {
	Pid         int
	Executable  string
	CommandLine string
	Tty         string
}

func GetProcesses() ([]Process, error) {
	cmd := exec.Command("ps", "-eo", "pid,comm,args,tty")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		return nil, err
	}

	processes := []Process{}
	lines := strings.Split(out.String(), "\n")
	for _, line := range lines[1:] {
		fields := strings.Fields(line)
		if len(fields) >= 4 {
			pid, err := strconv.Atoi(fields[0])
			if err == nil {
				process := Process{
					Pid:         pid,
					Executable:  fields[1],
					CommandLine: fields[2],
					Tty:         fields[3],
				}
				processes = append(processes, process)
			}
		}
	}

	return processes, nil
}

func FindProcessByPid(pid int) (*Process, error) {
	processes, err := GetProcesses()
	if err != nil {
		return nil, err
	}

	for _, process := range processes {
		if process.Pid == pid {
			return &process, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Process with ID %d not found", pid))
}

func (p *Process) Kill() error {
	process, err := os.FindProcess(p.Pid)
	if err != nil {
		return err
	}
	return process.Kill()
}
