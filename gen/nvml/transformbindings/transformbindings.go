/**
# Copyright 2025 NVIDIA CORPORATION
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
**/

package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

var (
	reAutogenComment  = regexp.MustCompile(`(?m)// WARNING: This file has automatically been generated on.*$`)
	reNvmlHLineNumber = regexp.MustCompile(`(?m)(// .* nvml/nvml\.h):[0-9]+$`)
)

func main() {
	sourceDir := flag.String("sourceDir", "", "Path to directory containing generated bindings")
	flag.Parse()

	if *sourceDir == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := transformFiles(*sourceDir); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func transformFiles(sourceDir string) error {
	return filepath.WalkDir(sourceDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		ext := filepath.Ext(path)
		if d.IsDir() || (ext != ".go" && ext != ".h") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		transformed := reAutogenComment.ReplaceAll(
			content,
			[]byte("// WARNING: THIS FILE WAS AUTOMATICALLY GENERATED."),
		)
		transformed = reNvmlHLineNumber.ReplaceAll(transformed, []byte("$1"))

		if string(transformed) == string(content) {
			return nil
		}
		return os.WriteFile(path, transformed, 0644)
	})
}
