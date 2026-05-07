// Copyright (c) 2025, NVIDIA CORPORATION.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// gen-dispatch generates a C dispatch table for NVML functions.
//
// It reads nvml.h and emits:
//   - nvml_dispatch.h: struct definition + static inline call-through wrappers
//   - zz_generated_nvml_dispatch.go: Go functions populateDispatch / clearDispatch
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	headerInput = flag.String("header", "", "path to nvml.h (required)")
	outputH     = flag.String("output-h", "", "output C header path (required)")
	outputGo    = flag.String("output-go", "", "output Go source path (required)")
	goPackage   = flag.String("package", "nvml", "Go package name for the generated file")
)

// funcDecl holds a parsed C function declaration.
type funcDecl struct {
	ReturnType string
	Name       string
	rawParams  string
	params     []param
}

type param struct {
	typ  string
	name string
}

var declDirRe = regexp.MustCompile(`\bDECLDIR\b`)
var deprecatedRe = regexp.MustCompile(`^DEPRECATED\([^)]*\)\s*`)

func main() {
	flag.Parse()
	if *headerInput == "" || *outputH == "" || *outputGo == "" {
		fmt.Fprintln(os.Stderr, "usage: gen-dispatch -header <nvml.h> -output-h <out.h> -output-go <out.go>")
		os.Exit(1)
	}

	decls, err := parseHeader(*headerInput)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing header: %v\n", err)
		os.Exit(1)
	}

	if err := writeHeader(*outputH, decls); err != nil {
		fmt.Fprintf(os.Stderr, "error writing %s: %v\n", *outputH, err)
		os.Exit(1)
	}
	if err := writeGoDispatch(*outputGo, decls); err != nil {
		fmt.Fprintf(os.Stderr, "error writing %s: %v\n", *outputGo, err)
		os.Exit(1)
	}

	fmt.Printf("Generated %s, %s (%d functions)\n", *outputH, *outputGo, len(decls))
}

// parseHeader reads nvml.h and extracts all DECLDIR function declarations.
func parseHeader(path string) ([]funcDecl, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	var buf string
	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), " \t")
		if buf != "" {
			buf += " " + strings.TrimSpace(line)
			if strings.Contains(line, ";") {
				lines = append(lines, buf)
				buf = ""
			}
			continue
		}
		if declDirRe.MatchString(line) && !strings.HasPrefix(line, "#") {
			if strings.Contains(line, ";") {
				lines = append(lines, line)
			} else {
				buf = line
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	var decls []funcDecl
	for _, line := range lines {
		d, ok := parseLine(line)
		if !ok {
			continue
		}
		decls = append(decls, d)
	}
	return decls, nil
}

func parseLine(line string) (funcDecl, bool) {
	line = deprecatedRe.ReplaceAllString(line, "")
	line = strings.ReplaceAll(line, "DECLDIR", "")
	line = strings.TrimSuffix(strings.TrimSpace(line), ";")
	line = collapseSpaces(line)

	parenOpen := strings.Index(line, "(")
	if parenOpen < 0 {
		return funcDecl{}, false
	}
	parenClose := strings.LastIndex(line, ")")
	if parenClose < parenOpen {
		return funcDecl{}, false
	}

	beforeParen := strings.TrimSpace(line[:parenOpen])
	rawParams := strings.TrimSpace(line[parenOpen+1 : parenClose])

	lastSpace := strings.LastIndex(beforeParen, " ")
	if lastSpace < 0 {
		return funcDecl{}, false
	}
	retType := strings.TrimSpace(beforeParen[:lastSpace])
	name := strings.TrimSpace(beforeParen[lastSpace+1:])

	if !strings.HasPrefix(name, "nvml") {
		return funcDecl{}, false
	}

	var params []param
	if rawParams != "" && rawParams != "void" {
		for _, p := range splitParams(rawParams) {
			p = strings.TrimSpace(p)
			if p == "" {
				continue
			}
			pt, pn := parseParam(p)
			params = append(params, param{typ: pt, name: pn})
		}
	}

	return funcDecl{
		ReturnType: retType,
		Name:       name,
		params:     params,
	}, true
}

func splitParams(s string) []string {
	var result []string
	depth := 0
	start := 0
	for i, ch := range s {
		switch ch {
		case '(':
			depth++
		case ')':
			depth--
		case ',':
			if depth == 0 {
				result = append(result, s[start:i])
				start = i + 1
			}
		}
	}
	return append(result, s[start:])
}

func parseParam(p string) (typ, name string) {
	tokens := strings.Fields(p)
	if len(tokens) == 0 {
		return p, ""
	}
	last := tokens[len(tokens)-1]
	name = strings.TrimLeft(last, "*")
	stars := last[:len(last)-len(name)]
	typeTokens := tokens[:len(tokens)-1]
	if stars != "" {
		typeTokens = append(typeTokens, stars)
	}
	return strings.Join(typeTokens, " "), name
}

func collapseSpaces(s string) string {
	return regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")
}

// ---- Header ----------------------------------------------------------------

const headerTemplate = `// Copyright (c) 2025, NVIDIA CORPORATION.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// WARNING: THIS FILE WAS AUTOMATICALLY GENERATED.
// Code generated by gen/nvml/gen-dispatch. DO NOT EDIT.

#ifndef NVML_DISPATCH_H
#define NVML_DISPATCH_H

#include "nvml.h"

// nvml_dispatch_t holds function pointers for all NVML API functions.
// Populated at library load time via populateDispatch() in Go.
typedef struct {
%s
} nvml_dispatch_t;

extern nvml_dispatch_t nvml_dispatch;

// Static inline wrappers — cgo calls these instead of the NVML symbols directly,
// routing through the dispatch table rather than the PLT.
%s
#endif /* NVML_DISPATCH_H */
`

func writeHeader(path string, decls []funcDecl) error {
	var structFields strings.Builder
	var inlineWrappers strings.Builder

	for _, d := range decls {
		fmt.Fprintf(&structFields, "    %s (*%s)(%s);\n",
			d.ReturnType, d.Name, paramList(d, true))

		fmt.Fprintf(&inlineWrappers, "static inline %s dispatch_%s(%s) {\n",
			d.ReturnType, d.Name, paramList(d, true))
		switch d.ReturnType {
		case "void":
			fmt.Fprintf(&inlineWrappers, "    if (!nvml_dispatch.%s) return;\n", d.Name)
			fmt.Fprintf(&inlineWrappers, "    nvml_dispatch.%s(%s);\n", d.Name, argList(d))
		case "nvmlReturn_t":
			fmt.Fprintf(&inlineWrappers, "    if (!nvml_dispatch.%s) return NVML_ERROR_FUNCTION_NOT_FOUND;\n", d.Name)
			fmt.Fprintf(&inlineWrappers, "    return nvml_dispatch.%s(%s);\n", d.Name, argList(d))
		default:
			fmt.Fprintf(&inlineWrappers, "    if (!nvml_dispatch.%s) return (%s)0;\n", d.Name, d.ReturnType)
			fmt.Fprintf(&inlineWrappers, "    return nvml_dispatch.%s(%s);\n", d.Name, argList(d))
		}
		inlineWrappers.WriteString("}\n\n")
	}

	return os.WriteFile(path, []byte(fmt.Sprintf(headerTemplate,
		structFields.String(), inlineWrappers.String())), 0644)
}

// ---- Go dispatch file ------------------------------------------------------

const goTemplate = `// Copyright (c) 2025, NVIDIA CORPORATION.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// WARNING: THIS FILE WAS AUTOMATICALLY GENERATED.
// Code generated by gen/nvml/gen-dispatch. DO NOT EDIT.

package %s

/*
#include "nvml_dispatch.h"

// nvml_dispatch is the single global instance of the dispatch table.
// It is populated by populateDispatch() and zeroed by clearDispatch() in Go.
nvml_dispatch_t nvml_dispatch;
*/
import "C"
import (
	"unsafe"

	"github.com/NVIDIA/go-nvml/pkg/dl"
)

// populateDispatch resolves all NVML function pointers via the given library
// and stores them in the C dispatch table.
func populateDispatch(lib *dl.DynamicLibrary) error {
%s	return nil
}

// clearDispatch zeros the dispatch table, preventing calls into an unloaded library.
func clearDispatch(_ *dl.DynamicLibrary) error {
	C.nvml_dispatch = C.nvml_dispatch_t{}
	return nil
}
`

func writeGoDispatch(path string, decls []funcDecl) error {
	var populate strings.Builder
	for _, d := range decls {
		fmt.Fprintf(&populate,
			"\t*(*unsafe.Pointer)(unsafe.Pointer(&C.nvml_dispatch.%s)) = lib.Resolve(%q)\n",
			d.Name, d.Name)
	}
	return os.WriteFile(path, []byte(fmt.Sprintf(goTemplate, *goPackage, populate.String())), 0644)
}

// ---- helpers ---------------------------------------------------------------

func paramList(d funcDecl, withNames bool) string {
	if len(d.params) == 0 {
		return "void"
	}
	var parts []string
	for _, p := range d.params {
		if withNames && p.name != "" {
			parts = append(parts, p.typ+" "+p.name)
		} else {
			parts = append(parts, p.typ)
		}
	}
	return strings.Join(parts, ", ")
}

func argList(d funcDecl) string {
	var names []string
	for _, p := range d.params {
		names = append(names, p.name)
	}
	return strings.Join(names, ", ")
}
