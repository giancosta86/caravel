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
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

/*
GoInstall compiles and installs the given Go program - located at a
relative path under GOPATH - into $GOPATH/bin by executing "go install".

Of course, the GOPATH environment variable must be set, and "go" must be
accessible via the PATH environment variable.

It returns the path of the compiled program, or an error on failure.
*/
func GoInstall(relativePathInGoPath string) (programPath string, err error) {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		return "", fmt.Errorf("The GOPATH environment variable must be set")
	}

	installCommand := exec.Command(
		"go",

		"install",
		relativePathInGoPath)

	err = installCommand.Run()
	if err != nil {
		return "", err
	}

	programBaseName := filepath.Base(relativePathInGoPath)

	return filepath.Join(goPath, "bin", programBaseName), nil
}
