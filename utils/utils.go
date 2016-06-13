package utils

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/spf13/viper"
)

// predefined Config Errors
var (
	// ErrPidDirConfig error for :pid_dir missing in config
	ErrPidDirConfig error = errors.New(":pid_dir not defined in the config file")

	ErrProgramName error = errors.New(":program_namenot defined in the config file")
)

// GetPidFile retrieve pid file from the config file
func GetPidFile() (string, error) {
	pidDir := viper.GetString(":pid_dir")
	if pidDir == "" {
		return "", ErrPidDirConfig
	}

	programName := viper.GetString(":program_name")
	if programName == "" {
		return "", ErrProgramName
	}
	pidFile := pidDir + programName + ".pid"
	return pidFile, nil
}

// GetServiceProcess retrieve the ID of the service if exists
func GetServiceProcess() (*os.Process, error) {

	pidFile, err := GetPidFile()
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve pid file - %s\n", err)
	}

	pidID, err := GetPidIDFromFile(pidFile)

	if err != nil {
		return nil, fmt.Errorf("Unable to load %s - %s\n", pidFile, err)
	}
	return os.FindProcess(pidID)
}

// GetPidIDFromFile retrieve pidID from a file
func GetPidIDFromFile(file string) (int, error) {

	fi, err := os.Open(file)
	if err != nil {
		return 0, fmt.Errorf("Unable to open %s - %s\n", file, err)
	}
	return GetPidID(fi)
}

// SavePidIDToFile Save the pid to a file
func SavePidIDToFile(pifdile string, pid int) error {
	return ioutil.WriteFile(pifdile, []byte(strconv.Itoa(pid)+"\n"), 0644)
}

// GetPidID retrieve pidID from a file
func GetPidID(rd io.Reader) (int, error) {
	bs, err := ioutil.ReadAll(rd)
	if err != nil {
		return 0, fmt.Errorf("Can't read content - %s\n", err)
	}

	content := strings.Replace(string(bs), "\n", "", -1)
	pidID, err := strconv.Atoi(content)
	if err != nil {
		return 0, fmt.Errorf("Can't convert '%s' to int loaded - %s\n", content, err)
	}
	return pidID, nil
}

// IsProcessRunning find out if process is running
func IsProcessRunning(process *os.Process) error {
	return process.Signal(syscall.Signal(0))
}

// StopRunningProcess find out if process is running
func StopRunningProcess(process *os.Process) error {
	return process.Signal(syscall.Signal(syscall.SIGTERM))
}
