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
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

/*
ExtractZip extracts the zip file located at the given path to the given
target directory, restoring the file-system permission bits for every entry.

Absolute entry paths in the zip file are not allowed.

The function returns nil on success, an error on failure.
*/
func ExtractZip(zipPath string, targetDir string) (err error) {
	return ExtractZipSkipLevels(zipPath, targetDir, 0)
}

/*
ExtractZipSkipLevels extracts the zip as ExtractZip does, but ignoring the first
"skippedLevels" directory nodes in the tree. "skippedLevels" must be >= 0.
*/
func ExtractZipSkipLevels(zipPath string, targetDir string, skippedLevels int) (err error) {
	if skippedLevels < 0 {
		return fmt.Errorf("The number of skipped levels must be >= 0")
	}

	zipReader, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, entryFile := range zipReader.File {
		entryRelativePath := entryFile.Name

		if path.IsAbs(entryRelativePath) {
			return fmt.Errorf("Absolute entry paths are not allowed: '%v'", entryRelativePath)
		}

		if skippedLevels > 0 {
			entryRelativePathComponents := strings.Split(entryRelativePath, "/")

			if len(entryRelativePathComponents) < skippedLevels+1 {
				continue
			}

			newEntryRelativePathComponents := entryRelativePathComponents[skippedLevels:]

			entryRelativePath = path.Join(newEntryRelativePathComponents...)
		}

		osSpecificEntryRelativePath := filepath.FromSlash(entryRelativePath)
		targetFilePath := filepath.Join(targetDir, osSpecificEntryRelativePath)

		entryFileMode := entryFile.FileInfo().Mode()

		if entryFile.FileInfo().IsDir() {
			err = os.MkdirAll(targetFilePath, entryFileMode)
			if err != nil {
				return err
			}
		} else {
			targetFile, err := os.OpenFile(targetFilePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, entryFileMode)
			if err != nil {
				return err
			}
			defer targetFile.Close()

			entryReader, err := entryFile.Open()
			if err != nil {
				return err
			}
			defer entryReader.Close()

			_, err = io.Copy(targetFile, entryReader)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
