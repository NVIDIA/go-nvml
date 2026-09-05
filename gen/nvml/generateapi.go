/**
# Copyright 2024 NVIDIA CORPORATION
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
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"unicode"
)

type GeneratableInterfacePoperties struct {
	Type                      string
	Interface                 string
	Exclude                   []string
	PackageMethodsAliasedFrom string
}

var GeneratableInterfaces = []GeneratableInterfacePoperties{
	{
		Type:                      "library",
		Interface:                 "Interface",
		Exclude:                   []string{"LookupSymbol"},
		PackageMethodsAliasedFrom: "libnvml",
	},
	{
		Type:      "nvmlDevice",
		Interface: "Device",
	},
	{
		Type:      "nvmlGpuInstance",
		Interface: "GpuInstance",
	},
	{
		Type:      "nvmlComputeInstance",
		Interface: "ComputeInstance",
	},
	{
		Type:      "nvmlEventSet",
		Interface: "EventSet",
	},
	{
		Type:      "nvmlGpmSample",
		Interface: "GpmSample",
	},
	{
		Type:      "nvmlUnit",
		Interface: "Unit",
	},
	{
		Type:      "nvmlVgpuInstance",
		Interface: "VgpuInstance",
	},
	{
		Type:      "nvmlVgpuTypeId",
		Interface: "VgpuTypeId",
	},
}

func main() {
	sourceDir := flag.String("sourceDir", "", "Path to the source directory for all go files")
	output := flag.String("output", "", "Path to the output file (default: stdout)")
	flag.Parse()

	// Check if required flags are provided
	if *sourceDir == "" {
		flag.Usage()
		return
	}

	var buf strings.Builder

	imports, err := extractImports(*sourceDir)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	header, err := generateHeader()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Fprint(&buf, header)

	if len(imports) > 0 {
		importBlock, err := generateImports(imports)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		fmt.Fprint(&buf, importBlock)
	}

	for i, p := range GeneratableInterfaces {
		if p.PackageMethodsAliasedFrom != "" {
			comment, err := generatePackageMethodsComment(p)
			if err != nil {
				fmt.Printf("Error: %v", err)
				return
			}
			fmt.Fprint(&buf, comment)

			out, err := generatePackageMethods(*sourceDir, p)
			if err != nil {
				fmt.Printf("Error: %v", err)
				return
			}
			fmt.Fprintf(&buf, "%s\n", out)
		}

		comment, err := generateInterfaceComment(p)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		fmt.Fprint(&buf, comment)

		out, err := generateInterface(*sourceDir, p)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		fmt.Fprint(&buf, out)

		if i < (len(GeneratableInterfaces) - 1) {
			fmt.Fprint(&buf, "\n")
		}
	}

	formatted, err := format.Source([]byte(buf.String()))
	if err != nil {
		fmt.Printf("Error: formatting output: %v", err)
		return
	}

	writer, closer, err := getWriter(*output)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer closer()

	_, err = writer.Write(formatted)
	if err != nil {
		fmt.Printf("Error: writing output: %v", err)
	}
}

func getWriter(outputFile string) (io.Writer, func() error, error) {
	if outputFile == "" {
		return os.Stdout, func() error { return nil }, nil
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return nil, nil, err
	}

	return file, file.Close, nil
}

func generateHeader() (string, error) {
	lines := []string{
		"/**",
		"# Copyright 2024 NVIDIA CORPORATION",
		"#",
		"# Licensed under the Apache License, Version 2.0 (the \"License\");",
		"# you may not use this file except in compliance with the License.",
		"# You may obtain a copy of the License at",
		"#",
		"#     http://www.apache.org/licenses/LICENSE-2.0",
		"#",
		"# Unless required by applicable law or agreed to in writing, software",
		"# distributed under the License is distributed on an \"AS IS\" BASIS,",
		"# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.",
		"# See the License for the specific language governing permissions and",
		"# limitations under the License.",
		"**/",
		"",
		"// Generated Code; DO NOT EDIT.",
		"",
		"package nvml",
		"",
		"",
	}
	return strings.Join(lines, "\n"), nil
}

func generateImports(imports []string) (string, error) {
	lines := []string{"import ("}
	for _, pkg := range imports {
		lines = append(lines, "\t\""+pkg+"\"")
	}
	lines = append(lines, ")", "")
	return strings.Join(lines, "\n"), nil
}

func generatePackageMethodsComment(input GeneratableInterfacePoperties) (string, error) {
	commentFmt := []string{
		"// The variables below represent package level methods from the %s type.",
	}

	var signature strings.Builder
	comment := strings.Join(commentFmt, "\n")
	comment = fmt.Sprintf(comment, input.Type)
	fmt.Fprintf(&signature, "%s\n", comment)
	return signature.String(), nil
}

func generateInterfaceComment(input GeneratableInterfacePoperties) (string, error) {
	commentFmt := []string{
		"// %s represents the interface for the %s type.",
		"//",
		"//go:generate moq -out mock/%s.go -pkg mock . %s:%s",
	}

	var signature strings.Builder
	comment := strings.Join(commentFmt, "\n")
	comment = fmt.Sprintf(comment, input.Interface, input.Type, strings.ToLower(input.Interface), input.Interface, input.Interface)
	fmt.Fprintf(&signature, "%s\n", comment)
	return signature.String(), nil
}

func generatePackageMethods(sourceDir string, input GeneratableInterfacePoperties) (string, error) {
	var signature strings.Builder

	signature.WriteString("var (\n")

	methods, err := extractMethodsFromPackage(sourceDir, input)
	if err != nil {
		return "", err
	}

	for _, method := range methods {
		name := method.Name.Name
		formatted := fmt.Sprintf("\t%s = %s.%s\n", name, input.PackageMethodsAliasedFrom, name)
		signature.WriteString(formatted)
	}

	signature.WriteString(")\n")

	return signature.String(), nil
}

func generateInterface(sourceDir string, input GeneratableInterfacePoperties) (string, error) {
	var signature strings.Builder

	fmt.Fprintf(&signature, "type %s interface {\n", input.Interface)

	methods, err := extractMethodsFromPackage(sourceDir, input)
	if err != nil {
		return "", err
	}

	for _, method := range methods {
		formatted := fmt.Sprintf("\t%s\n", formatMethodSignature(method))
		signature.WriteString(formatted)
	}

	signature.WriteString("}\n")

	return signature.String(), nil
}

func getGoFiles(sourceDir string) (map[string][]byte, error) {
	gofiles := make(map[string][]byte)

	err := filepath.WalkDir(sourceDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		gofiles[path] = content

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walking %s: %w", sourceDir, err)
	}

	return gofiles, nil
}

func extractMethodsFromPackage(sourceDir string, input GeneratableInterfacePoperties) ([]*ast.FuncDecl, error) {
	gofiles, err := getGoFiles(sourceDir)
	if err != nil {
		return nil, err
	}

	var methods []*ast.FuncDecl
	for file, content := range gofiles {
		m, err := extractMethods(file, content, input)
		if err != nil {
			return nil, err
		}
		methods = append(methods, m...)
	}

	sort.Slice(methods, func(i, j int) bool {
		return methods[i].Name.Name < methods[j].Name.Name
	})

	return methods, nil
}

func extractMethods(sourceFile string, sourceContent []byte, input GeneratableInterfacePoperties) ([]*ast.FuncDecl, error) {
	// Parse source file
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, sourceFile, sourceContent, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// Traverse AST to find type declarations and associated methods
	var methods []*ast.FuncDecl
	for _, decl := range node.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		// Check if the function is a method associated with the specified type
		if receiverType := funcDecl.Recv; receiverType != nil {
			var ident *ast.Ident

			for _, field := range receiverType.List {
				switch fieldType := field.Type.(type) {
				case *ast.Ident:
					ident = fieldType
				case *ast.StarExpr:
					// Update ident if it's a *ast.StarExpr
					if newIdent, ok := fieldType.X.(*ast.Ident); ok {
						// If the inner type is an *ast.Ident, update ident
						ident = newIdent
					}
				}

				// No identifier found
				if ident == nil {
					continue
				}

				// Identifier is not the one we are looking for
				if ident.Name != input.Type {
					continue
				}

				// Ignore non-public methods
				if !isPublic(funcDecl.Name.Name) {
					continue
				}

				// Ignore method in the exclude list
				if slices.Contains(input.Exclude, funcDecl.Name.Name) {
					continue
				}

				methods = append(methods, funcDecl)
			}
		}
	}

	return methods, nil
}

func formatMethodSignature(decl *ast.FuncDecl) string {
	var signature strings.Builder

	// Write method name
	signature.WriteString(decl.Name.Name)
	signature.WriteString("(")

	// Write parameters
	if decl.Type.Params != nil {
		for i, param := range decl.Type.Params.List {
			if i > 0 {
				signature.WriteString(", ")
			}
			signature.WriteString(formatFieldList(param))
		}
	}

	signature.WriteString(")")

	// Write return types
	if decl.Type.Results != nil {
		signature.WriteString(" ")
		if len(decl.Type.Results.List) > 1 {
			signature.WriteString("(")
		}
		for i, result := range decl.Type.Results.List {
			if i > 0 {
				signature.WriteString(", ")
			}
			signature.WriteString(formatFieldList(result))
		}
		if len(decl.Type.Results.List) > 1 {
			signature.WriteString(")")
		}
	}

	return signature.String()
}

func formatFieldList(field *ast.Field) string {
	var builder strings.Builder
	switch fieldType := field.Type.(type) {
	case *ast.Ident:
		builder.WriteString(fieldType.Name)
	case *ast.ArrayType:
		builder.WriteString("[]")
		builder.WriteString(formatFieldList(&ast.Field{Type: fieldType.Elt}))
	case *ast.StarExpr:
		builder.WriteString("*")
		builder.WriteString(formatFieldList(&ast.Field{Type: fieldType.X}))
	case *ast.SelectorExpr:
		if pkg, ok := fieldType.X.(*ast.Ident); ok {
			builder.WriteString(pkg.Name)
			builder.WriteString(".")
		}
		builder.WriteString(fieldType.Sel.Name)
	}
	return builder.String()
}

func extractImports(sourceDir string) ([]string, error) {
	importSet := make(map[string]bool)
	for _, p := range GeneratableInterfaces {
		methods, err := extractMethodsFromPackage(sourceDir, p)
		if err != nil {
			return nil, err
		}
		getQualifiedPackages(methods, importSet)
	}
	var imports []string
	for pkg := range importSet {
		imports = append(imports, pkg)
	}
	sort.Strings(imports)
	return imports, nil
}

func getQualifiedPackages(methods []*ast.FuncDecl, seen map[string]bool) {
	for _, method := range methods {
		getPackagesFromFieldList(method.Type.Params, seen)
		getPackagesFromFieldList(method.Type.Results, seen)
	}
}

func getPackagesFromFieldList(fieldList *ast.FieldList, seen map[string]bool) {
	if fieldList == nil {
		return
	}
	for _, field := range fieldList.List {
		collectPackagesFromExpr(field.Type, seen)
	}
}

func collectPackagesFromExpr(expr ast.Expr, seen map[string]bool) {
	switch e := expr.(type) {
	case *ast.SelectorExpr:
		if pkg, ok := e.X.(*ast.Ident); ok {
			seen[pkg.Name] = true
		}
	case *ast.StarExpr:
		collectPackagesFromExpr(e.X, seen)
	case *ast.ArrayType:
		collectPackagesFromExpr(e.Elt, seen)
	}
}

func isPublic(name string) bool {
	if len(name) == 0 {
		return false
	}
	return unicode.IsUpper([]rune(name)[0])
}
