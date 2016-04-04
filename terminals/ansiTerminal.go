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
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

/*
AnsiTerminal is a 256-color terminal based on the ANSI escape codes
(https://en.wikipedia.org/wiki/ANSI_escape_code)
*/
type AnsiTerminal struct {
	rows    int
	columns int
}

func (terminal *AnsiTerminal) Escape(escapeFormatString string, args ...interface{}) {
	fmt.Printf("\033"+escapeFormatString, args...)
}

func (terminal *AnsiTerminal) GetRows() int {
	return terminal.rows
}

func (terminal *AnsiTerminal) GetColumns() int {
	return terminal.columns
}

func (terminal *AnsiTerminal) Clear() {
	emptyRow := strings.Repeat(" ", terminal.GetColumns())

	for row := 1; row <= terminal.GetRows(); row++ {
		fmt.Println(emptyRow)
	}

	terminal.MoveCursor(1, 1)
}

func (terminal *AnsiTerminal) MoveCursor(row int, column int) {
	terminal.checkRow(row)
	terminal.checkColumn(column)

	terminal.Escape("[%d;%dH", row, column)
}

func (terminal *AnsiTerminal) ResetStyle() {
	terminal.Escape("[0m")
}

func (terminal *AnsiTerminal) EnableTextHidden() {
	terminal.Escape("[8m")
}

func (terminal *AnsiTerminal) DisableTextHidden() {
	terminal.Escape("[28m")
}

func (terminal *AnsiTerminal) EnableTextBold() {
	terminal.Escape("[1m")
}

func (terminal *AnsiTerminal) DisableTextBold() {
	terminal.Escape("[21m")
}

func (terminal *AnsiTerminal) EnableTextUnderlined() {
	terminal.Escape("[4m")
}

func (terminal *AnsiTerminal) DisableTextUnderlined() {
	terminal.Escape("[24m")
}

func (terminal *AnsiTerminal) ShowCursor() {
	terminal.Escape("[?25h")
}

func (terminal *AnsiTerminal) HideCursor() {
	terminal.Escape("[?25l")
}

func (terminal *AnsiTerminal) SetBackgroundColor(colorCode int) {
	terminal.checkColorCode(colorCode)

	terminal.Escape("[48;5;%dm", colorCode)
}

func (terminal *AnsiTerminal) SetForegroundColor(colorCode int) {
	terminal.checkColorCode(colorCode)

	terminal.Escape("[38;5;%dm", colorCode)
}

func (terminal *AnsiTerminal) PrintCenteredInRow(row int, text string) {
	terminal.checkRow(row)

	textLength := utf8.RuneCountInString(text)
	textColumn := (terminal.GetColumns() - textLength) / 2

	terminal.MoveCursor(row, textColumn)
	fmt.Print(text)
}

func (terminal *AnsiTerminal) DrawHorizontalProgressBar(
	row int,
	column int,
	maxLineLength int,
	fractionalValue float64) {

	terminal.checkRow(row)
	terminal.checkColumn(column)

	if fractionalValue < 0 || fractionalValue > 1 {
		panic("Invalid fractional value")
	}

	const delimiter = "ǁ"
	const tick = "="
	const space = " "

	numberOfTicks := int(math.Ceil(float64(maxLineLength) * fractionalValue))
	percentage := 100 * fractionalValue

	terminal.MoveCursor(row, column)
	fmt.Printf("%v%v%v%v  %.2f%%",
		delimiter,
		strings.Repeat(tick, numberOfTicks),
		strings.Repeat(space, maxLineLength-numberOfTicks),
		delimiter,
		percentage)
}

func (terminal *AnsiTerminal) SupportsANSI() bool {
	return true
}

func (terminal *AnsiTerminal) checkRow(row int) {
	if row < 1 || row > terminal.GetRows() {
		panic("Invalid row: " + strconv.Itoa(row))
	}
}

func (terminal *AnsiTerminal) checkColumn(column int) {
	if column < 1 || column > terminal.GetColumns() {
		panic("Invalid column: " + strconv.Itoa(column))
	}
}

func (terminal *AnsiTerminal) checkColorCode(colorCode int) {
	if colorCode < 0 || colorCode > 255 {
		panic("Invalid color code: " + strconv.Itoa(colorCode))
	}
}
