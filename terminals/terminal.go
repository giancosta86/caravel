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

/*
Terminal is a text-based user interface
*/
type Terminal interface {
	/*
	  GetRows returns the number of rows (as characters) in the terminal
	*/
	GetRows() int

	/*
	  GetColumns returns the number of columns (as characters) in the terminal
	*/
	GetColumns() int

	/*
		  Clear empties the terminal, ensuring that the background color
			is applied to the whole screen, then moves the cursor to its
			upper-left corner
	*/
	Clear()

	/*
		  MoveCursor moves the cursor to the given row and column -
			both values are 1-based
	*/
	MoveCursor(row int, column int)

	/*
		ResetStyle resets the style (including colors) of the terminal
		to its default parameters
	*/
	ResetStyle()

	/*
		EnableTextHidden enables the Hidden text style - also hiding user input
	*/
	EnableTextHidden()

	/*
		DisableTextHidden disables the Hidden text style - also showing user input
	*/
	DisableTextHidden()

	/*
	  EnableTextBold enables the Bold text style
	*/
	EnableTextBold()

	/*
	  DisableTextBold disables the Bold text style
	*/
	DisableTextBold()

	/*
		EnableTextUnderlined enables the Underlined text style
	*/
	EnableTextUnderlined()

	/*
		DisableTextUnderlined enables the Underlined text style
	*/
	DisableTextUnderlined()

	/*
	  ShowCursor shows the input cursor
	*/
	ShowCursor()

	/*
	  HideCursor hides the input cursor
	*/
	HideCursor()

	/*
	  SetBackgroundColor sets the color code for the text background
	*/
	SetBackgroundColor(colorCode int)

	/*
	  SetForegroundColor sets the color code for the text
	*/
	SetForegroundColor(colorCode int)

	/*
		PrintCenteredInRow prints the given text on the given row,
		horizontally centering it
	*/
	PrintCenteredInRow(row int, text string)

	/*
		DrawHorizontalProgressBar draws a horizontal progress bar starting
		at the given row and column and having an internal line long at most
		"maxLineLength" characters.
		The given fractional value must be in the [0; 1] range
	*/
	DrawHorizontalProgressBar(
		row int,
		column int,
		maxLineLength int,
		fractionalValue float64)

	/*
		SupportsANSI returns true if the terminal supports the ANSI escape codes
	*/
	SupportsANSI() bool

	/*
		Escape prints an escape sequence according to the terminal's specific format,
		employing the escapeFormatString - which gets formatted via the
		provided optional args
	*/
	Escape(escapeFormatString string, args ...interface{})
}
