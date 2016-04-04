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
	"os"
	"os/user"
	"path/filepath"
)

/*
GetUserDirectory returns the user's directory, or an error on failure.
First it looks for a "HOME" environment variable; then it employs user.Current()
*/
func GetUserDirectory() (userDir string, err error) {
	environmentHome := os.Getenv("HOME")
	if environmentHome != "" {
		return environmentHome, nil
	}

	user, err := user.Current()
	if err != nil {
		return "", err
	}

	return user.HomeDir, nil
}

/*
GetUserDesktop returns the user's "Desktop" directory, or an error on failure.
*/
func GetUserDesktop() (desktopDir string, err error) {
	userDir, err := GetUserDirectory()
	if err != nil {
		return "", err
	}

	desktopDir = filepath.Join(userDir, "Desktop")
	return desktopDir, nil
}
