package caravel

import (
	"os/exec"
)

/*
GetProcessOutputString runs the given program with the given arguments,
returning its output as a string - or an error on failure.
*/
func GetProcessOutputString(name string, arg ...string) (outputString string, err error) {
	command := exec.Command(name, arg...)

	return GetCommandOutputString(command)
}

/*
GetCommandOutputString runs the given command, returning its output
as a string - or an error on failure.
*/
func GetCommandOutputString(command *exec.Cmd) (outputString string, err error) {
	bytes, err := command.Output()
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
