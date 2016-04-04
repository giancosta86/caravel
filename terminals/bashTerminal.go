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

package terminals

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/giancosta86/caravel"
)

/*
BashTerminal is a terminal provided by the BASH-shell
*/
type BashTerminal struct {
	AnsiTerminal
}

/*
NewBashTerminal creates an ANSI terminal implemented by the BASH shell running the program
*/
func NewBashTerminal() *BashTerminal {
	rows, columns := getBashTerminalSize()

	return &BashTerminal{
		AnsiTerminal{
			rows:    rows,
			columns: columns,
		},
	}
}

func getBashTerminalSize() (rows int, columns int) {
	const defaultRows = 25
	const defaultColumns = 80

	sizeCommand := exec.Command("stty", "size")
	sizeCommand.Stdin = os.Stdin

	sizeOutputString, err := caravel.GetCommandOutputString(sizeCommand)
	if err != nil {
		return defaultRows, defaultColumns
	}

	sizeOutputString = strings.TrimSpace(sizeOutputString)
	_, err = fmt.Sscanf(sizeOutputString, "%d %d", &rows, &columns)
	if err != nil {
		return defaultRows, defaultColumns
	}

	return rows, columns
}
