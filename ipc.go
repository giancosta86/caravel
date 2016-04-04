package caravel

import (
	"os/exec"
)

/*
GetProcessOutputString runs the given app with the given arguments,
returning its output as a string or an error on failure.
*/
func GetProcessOutputString(name string, arg ...string) (outputString string, err error) {
	bytes, err := exec.Command(name, arg...).Output()

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
