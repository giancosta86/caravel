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
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

/*
RetrievalProgressCallback models a function receving 2 parameters: the number of bytes
retrieved up to now and the total number of bytes to be retrieved.
*/
type RetrievalProgressCallback func(retrievedSize int64, totalSize int64)

/*
RetrieveFromURL downloads and returns all the bytes from the given remote URL,
returning an error on failure.
*/
func RetrieveFromURL(remoteURL *url.URL) (bytes []byte, err error) {
	response, err := http.Get(remoteURL.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(response.Status)
	}

	return ioutil.ReadAll(response.Body)
}

/*
RetrieveChunksFromURL reads the bytes located at the given remote URL as subsequent
chunks having the passed size; whenever a chunk is retrieved, it is
written to the given Writer, and "progressCallback" is called.

The function returns nil on success, an error on failure.
*/
func RetrieveChunksFromURL(
	remoteURL *url.URL,
	outputWriter io.Writer,
	chunkSize int64,
	progressCallback RetrievalProgressCallback) (err error) {

	response, err := http.Get(remoteURL.String())
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf(response.Status)
	}

	retrievedSize := int64(0)
	totalSize := response.ContentLength

	if totalSize <= 0 {
		return fmt.Errorf("Invalid Content-Length for response: %v", totalSize)
	}

	for retrievedSize < totalSize {
		writtenBytes, err := io.CopyN(outputWriter, response.Body, chunkSize)
		if err != nil && err != io.EOF {
			return err
		}

		retrievedSize = retrievedSize + writtenBytes
		progressCallback(retrievedSize, totalSize)
	}

	return nil
}

/*
IsSecureURL returns true if the given url employs HTTPS.
*/
func IsSecureURL(urlToCheck *url.URL) bool {
	return urlToCheck.Scheme == "https"
}
