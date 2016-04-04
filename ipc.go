/*§
  ===========================================================================
  Caravel
  ===========================================================================
  Copyright (C) 2015-2016 Gianluca Costa
  ===========================================================================
  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
  ===========================================================================
*/

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
