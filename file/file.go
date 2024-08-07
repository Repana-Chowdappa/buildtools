/*
Copyright 2016 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package file provides utility file operations.
package file

import (
	"fmt"
	"io"
	"os"
)

// ReadFile can be updated from the caller to change the API
// for reading a file.
var ReadFile = readFile

// WriteFile can be updated from the caller to change the API
// for writing a file.
var WriteFile = writeFile

// OpenReadFile can be updated from the caller to change the API
// for opening a file.
var OpenReadFile = openReadFile

// readFile is like os.ReadFile.
func readFile(name string) ([]byte, os.FileInfo, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return nil, nil, err
	}

	data, err := os.ReadFile(name)
	return data, fi, err
}

// writeFile is like os.WriteFile.
func writeFile(name string, data []byte) error {
	return os.WriteFile(name, data, 0644)
}

// openReadFile is like os.Open.
func openReadFile(name string) io.ReadCloser {
	f, err := os.Open(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open %s\n", name)
		os.Exit(1)
	}
	return f
}
