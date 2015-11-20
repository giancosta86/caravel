/*§
  ===========================================================================
  Caravel
  ===========================================================================
  Copyright (C) 2015 Gianluca Costa
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
	"bytes"
	"os"
	"unicode"
	"unicode/utf8"
)

/*
FileExists returns true if "filePath" exists and is a file.
*/
func FileExists(filePath string) bool {
	stat, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	return !stat.IsDir()
}

/*
DirectoryExists returns true if "dirPath" exists and is a directory.
*/
func DirectoryExists(dirPath string) bool {
	stat, err := os.Stat(dirPath)
	if err != nil {
		return false
	}

	return stat.IsDir()
}

/*
FormatFileName receives a file name and returns it with only letters, digits, '.', ' ' and '-',
replacing every other rune with an underscore ('_'). Furthermore, each sequence of underscores
is contracted into a single underscore, and trailing underscores are suppressed.
*/
func FormatFileName(fileName string) string {
	const defaultRune rune = '_'

	var result bytes.Buffer
	var latestRune rune

	for _, currentRune := range fileName {
		if isAcceptableForFileName(currentRune) {
			result.WriteRune(currentRune)
			latestRune = currentRune
		} else {
			if latestRune != defaultRune {
				result.WriteRune(defaultRune)
				latestRune = defaultRune
			}
		}
	}

	if latestRune == defaultRune {
		result.Truncate(result.Len() - utf8.RuneLen(defaultRune))
	}

	return result.String()
}

func isAcceptableForFileName(runeToTest rune) bool {
	return unicode.IsLetter(runeToTest) ||
		unicode.IsDigit(runeToTest) ||
		runeToTest == '-' ||
		runeToTest == '.' ||
		runeToTest == ' '
}
