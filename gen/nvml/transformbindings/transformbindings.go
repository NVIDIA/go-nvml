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
	"strings"
)

var (
	reAutogenComment  = regexp.MustCompile(`(?m)// WARNING: This file has automatically been generated on.*$`)
	reNvmlHLineNumber = regexp.MustCompile(`(?m)(// .* nvml/nvml\.h):[0-9]+$`)
	reNvmlCall        = regexp.MustCompile(`C\.(nvml[A-Za-z0-9_]+)\(`)
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

		transformed := applyGeneralTransforms(content)
		if filepath.Base(path) == "nvml.go" {
			transformed = []byte(applyNvmlGoTransforms(string(transformed)))
		}

		if string(transformed) == string(content) {
			return nil
		}
		return os.WriteFile(path, transformed, 0644)
	})
}

func applyGeneralTransforms(content []byte) []byte {
	content = reAutogenComment.ReplaceAll(content, []byte("// WARNING: THIS FILE WAS AUTOMATICALLY GENERATED."))
	content = reNvmlHLineNumber.ReplaceAll(content, []byte("$1"))
	return content
}

func applyNvmlGoTransforms(src string) string {
	src = stripLDFlags(src)
	src = rewriteCalls(src)
	src = insertDispatchInclude(src)
	return src
}

func stripLDFlags(src string) string {
	var out strings.Builder
	for _, line := range strings.SplitAfter(src, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "#cgo") && strings.Contains(trimmed, "LDFLAGS:") {
			continue
		}
		out.WriteString(line)
	}
	return out.String()
}

func rewriteCalls(src string) string {
	return reNvmlCall.ReplaceAllStringFunc(src, func(match string) string {
		return "C.dispatch_" + match[len("C."):]
	})
}

func insertDispatchInclude(src string) string {
	const dispatchInclude = `#include "nvml_dispatch.h"`
	if strings.Contains(src, dispatchInclude) {
		return src
	}

	importC := `import "C"`
	importIdx := strings.Index(src, importC)
	if importIdx < 0 {
		return src
	}
	preamble := src[:importIdx]
	lastInclude := strings.LastIndex(preamble, "#include")
	if lastInclude < 0 {
		return src
	}
	lineEnd := strings.Index(preamble[lastInclude:], "\n")
	if lineEnd < 0 {
		return src
	}
	insertAt := lastInclude + lineEnd + 1

	return preamble[:insertAt] + dispatchInclude + "\n" + src[insertAt:]
}
