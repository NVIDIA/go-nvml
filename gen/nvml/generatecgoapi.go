package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

// Constants for type checking
const (
	nvmlPrefix    = "nvml"
	pointerPrefix = "*"
	arrayPrefix   = "[]"
)

// Template definitions
const fileTemplate = `
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

// Generated Code; DO NOT EDIT.

package nvml

// CgoAPI is a global variable providing direct calls to the cgo API rather than the standard wrappers
var CgoAPI = cgoapi{}

type cgoapi struct{}

{{range $i, $method := .Methods}}
{{- if $i }}

{{- end }}
{{$method}}
{{end}}`

const methodTemplate = `
{{- $params := "" -}}
{{- range $i, $param := .Params -}}
{{- if $i -}}
{{- $params = printf "%s, %s %s" $params .Name .Type -}}
{{- else -}}
{{- $params = printf "%s %s" .Name .Type -}}
{{- end -}}
{{- end -}}

{{- $returns := "" -}}
{{- range $i, $ret := .Returns -}}
{{- if $i -}}
{{- $returns = printf "%s, %s" $returns .Type -}}
{{- else -}}
{{- $returns = printf "%s" .Type -}}
{{- end -}}
{{- end -}}

{{- $callParams := "" -}}
{{- range $i, $param := .CallParams -}}
{{- if $i -}}
{{- $callParams = printf "%s, %s" $callParams . -}}
{{- else -}}
{{- $callParams = printf "%s" . -}}
{{- end -}}
{{- end -}}

func (cgoapi) {{.Name}}({{$params}}) {{if .Returns}}({{$returns}}){{end}} {
{{- range .Conversions}}
	{{.}}
{{- end}}
	ret := {{.CallName}}({{$callParams}})
{{- range .Assignments}}
	{{.}}
{{- end}}
	return ret
}`

// Dynamic type caches - populated during generation
var interfaceTypes = make(map[string]bool)
var structTypesWithEmbeddedNvml = make(map[string]string)

// Template data structures
type FileTemplateData struct {
	Methods []string
}

type MethodTemplateData struct {
	Name         string
	Params       []ParamData
	Returns      []ReturnData
	Conversions  []string
	CallName     string
	CallParams   []string
	Assignments  []string
}

type ParamData struct {
	Name string
	Type string
}

type ReturnData struct {
	Type string
}



// ParameterInfo holds processed parameter information
type ParameterInfo struct {
	name            string
	convertedType   string
	convertedName   string
	needsConversion bool
	isPointer       bool
	baseType        string
}

func main() {
	// Parse command line flags
	sourceDir := flag.String("sourceDir", "", "Path to the source directory for all go files")
	output := flag.String("output", "", "Path to the output file (default: stdout)")
	flag.Parse()

	if *sourceDir == "" {
		flag.Usage()
		return
	}

	// Discover interface types and struct types with embedded nvml* types
	if err := discoverTypes(*sourceDir); err != nil {
		fmt.Printf("Error discovering types: %v", err)
		return
	}

	// Set up output writer
	writer, closer, err := getWriter(*output)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer closer()

	// Generate the cgoapi file content
	fileContent, err := generateCgoAPIFile(*sourceDir)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Fprint(writer, fileContent)
}

// getWriter returns an appropriate writer for the output file
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

// generateCgoAPIFile generates the full cgoapi.go file content using templates
func generateCgoAPIFile(sourceDir string) (string, error) {
	// Parse the main file template
	tmpl, err := template.New("file").Parse(fileTemplate)
	if err != nil {
		return "", fmt.Errorf("parsing file template: %w", err)
	}

	// Collect methods
	methods, err := collectCgoApiMethods(sourceDir)
	if err != nil {
		return "", err
	}

	// Prepare file data
	fileData := FileTemplateData{
		Methods: methods,
	}

	// Execute template
	var sb strings.Builder
	if err := tmpl.Execute(&sb, fileData); err != nil {
		return "", fmt.Errorf("executing file template: %w", err)
	}

	return sb.String(), nil
}



// collectCgoApiMethods collects and returns all cgoapi methods, sorted alphabetically
func collectCgoApiMethods(sourceDir string) ([]string, error) {
	// Extract all functions from nvml.go
	funcs, err := extractFunctionsFromNvmlFile(sourceDir)
	if err != nil {
		return nil, err
	}

	// Create a map for quick function lookup
	funcMap := make(map[string]*ast.FuncDecl)
	for _, f := range funcs {
		funcMap[f.Name.Name] = f
	}

	var methods []string

	// Generate methods for all functions from nvml.go
	for _, function := range funcs {
		methodName := removeNvmlPrefix(function.Name.Name)
		if methodName == function.Name.Name {
			continue // Skip functions that don't start with "nvml"
		}
		methodCode, err := generateMethodData(methodName, function, function.Name.Name)
		if err != nil {
			return nil, err
		}
		methods = append(methods, methodCode)
	}

	// Generate methods for unversioned functions from lib.go
	unversionedFuncs, err := extractUnversionedNvmlFunctionsFromLibFile(sourceDir)
	if err != nil {
		return nil, err
	}
	for _, pair := range unversionedFuncs {
		nvmlFunc := pair[0]
		aliasedFunc := pair[1]
		methodName := removeNvmlPrefix(nvmlFunc)
		targetFunc := funcMap[aliasedFunc]
		if targetFunc == nil {
			continue // Skip if the aliased function doesn't exist
		}
		methodCode, err := generateMethodData(methodName, targetFunc, nvmlFunc)
		if err != nil {
			return nil, err
		}
		methods = append(methods, methodCode)
	}

	// Sort methods alphabetically
	sort.Strings(methods)
	return methods, nil
}

// generateMethodData generates method code for a given function declaration
func generateMethodData(methodName string, decl *ast.FuncDecl, callName string) (string, error) {
	// Parse the method template
	tmpl, err := template.New("method").Parse(methodTemplate)
	if err != nil {
		return "", fmt.Errorf("parsing method template: %w", err)
	}

	// Process parameters
	params := processParameters(decl.Type.Params.List)
	paramData := make([]ParamData, len(params))
	for i, param := range params {
		paramData[i] = ParamData{
			Name: param.name,
			Type: param.convertedType,
		}
	}

	// Process return types
	var returnData []ReturnData
	if decl.Type.Results != nil {
		for _, result := range decl.Type.Results.List {
			returnData = append(returnData, ReturnData{
				Type: formatFieldList(result),
			})
		}
	}

	// Generate conversions
	var conversions []string
	for _, param := range params {
		if param.needsConversion && param.isPointer {
			conversions = append(conversions, fmt.Sprintf("var nvml%s %s", strings.Title(param.name), param.baseType))
		}
	}

	// Generate call parameters
	callParams := make([]string, len(params))
	for i, param := range params {
		callParams[i] = param.convertedName
	}

	// Generate assignments
	var assignments []string
	for _, param := range params {
		if !param.needsConversion || !param.isPointer {
			continue
		}

		// Skip assignment for structs with embedded nvml* types
		if isStructWithEmbeddedNvmlType(convertTypeName(param.baseType)) {
			continue
		}

		// Only convert if the target type (without nvml prefix) is an interface
		if strings.HasPrefix(param.baseType, nvmlPrefix) {
			targetType := convertTypeName(param.baseType)
			if isInterfaceType(targetType) {
				assignment := generateAssignmentStatement(param, targetType)
				assignments = append(assignments, assignment)
			}
		}
	}

	methodData := MethodTemplateData{
		Name:        methodName,
		Params:      paramData,
		Returns:     returnData,
		Conversions: conversions,
		CallName:    callName,
		CallParams:  callParams,
		Assignments: assignments,
	}

	// Generate the method code using template
	var sb strings.Builder
	if err := tmpl.Execute(&sb, methodData); err != nil {
		return "", fmt.Errorf("executing method template: %w", err)
	}

	return sb.String(), nil
}



// processParameters processes all parameters and returns their conversion info
func processParameters(params []*ast.Field) []ParameterInfo {
	var results []ParameterInfo

	for _, param := range params {
		paramType := getTypeString(param.Type)
		info := ParameterInfo{
			convertedType: paramType,
			isPointer:     isPointerType(paramType),
		}

		if len(param.Names) > 0 {
			paramName := param.Names[0].Name
			info.name = convertParamName(paramName)
			info.convertedName = info.name
			info.baseType = paramType

			// Apply type-specific conversion logic
			if info.isPointer {
				info.baseType = paramType[len(pointerPrefix):] // Remove * prefix
				info = processPointerParameter(info)
			} else if shouldConvertType(paramType) {
				info = processNonPointerParameter(info)
			}
		}

		results = append(results, info)
	}

	return results
}

// processPointerParameter handles pointer parameter conversion logic
func processPointerParameter(info ParameterInfo) ParameterInfo {
	// Check if this is a pointer to a struct that embeds an nvml* type
	if isStructWithEmbeddedNvmlType(convertTypeName(info.baseType)) {
		embeddedType := getEmbeddedNvmlType(convertTypeName(info.baseType))
		info.convertedType = "*" + convertTypeName(info.baseType)
		info.convertedName = fmt.Sprintf("&%s.%s", info.name, embeddedType)
		info.needsConversion = false // No conversion needed - pass embedded field reference
	} else {
		// Regular pointer to nvml* type - create local variable for conversion
		info.convertedType = "*" + convertTypeName(info.baseType)
		info.convertedName = "&nvml" + strings.Title(info.name)
		info.needsConversion = true
	}

	return info
}

// processNonPointerParameter handles non-pointer parameter conversion logic
func processNonPointerParameter(info ParameterInfo) ParameterInfo {
	targetType := convertTypeName(info.baseType)
	if isInterfaceType(targetType) {
		// Convert to interface type using Handle() method
		info.convertedType = targetType
		info.convertedName = fmt.Sprintf("%sHandle(%s)", info.baseType, info.name)
		info.needsConversion = true
	} else {
		// Pass the nvml* type directly - no conversion needed
		info.convertedType = info.baseType
		info.convertedName = info.name
		info.needsConversion = false
	}

	return info
}



// generateAssignmentStatement generates a single assignment statement
func generateAssignmentStatement(param ParameterInfo, targetType string) string {
	// Check if it's a struct type that has a convert() method
	if strings.HasSuffix(targetType, "Info") || strings.HasSuffix(targetType, "Stats") || strings.HasSuffix(targetType, "Settings") {
		return fmt.Sprintf("\t*%s = nvml%s.convert()", param.name, strings.Title(param.name))
	} else {
		// Use type conversion for other interface types
		return fmt.Sprintf("\t*%s = %s(nvml%s)", param.name, targetType, strings.Title(param.name))
	}
}

// discoverTypes discovers interface types and struct types with embedded nvml* types
func discoverTypes(sourceDir string) error {
	// Parse all Go files in the source directory
	files, err := filepath.Glob(filepath.Join(sourceDir, "*.go"))
	if err != nil {
		return fmt.Errorf("finding Go files: %w", err)
	}

	for _, file := range files {
		if err := parseFileForTypes(file); err != nil {
			return fmt.Errorf("parsing %s: %w", file, err)
		}
	}

	return nil
}

// parseFileForTypes parses a single file to discover interface and struct types
func parseFileForTypes(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, content, parser.ParseComments)
	if err != nil {
		return err
	}

	// Look for interface and struct type declarations
	for _, decl := range node.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					analyzeTypeDeclaration(typeSpec)
				}
			}
		}
	}

	return nil
}

// analyzeTypeDeclaration analyzes a type declaration to determine if it's an interface or struct with embedded nvml* type
func analyzeTypeDeclaration(typeSpec *ast.TypeSpec) {
	typeName := typeSpec.Name.Name

	switch t := typeSpec.Type.(type) {
	case *ast.InterfaceType:
		// This is an interface type
		interfaceTypes[typeName] = true
	case *ast.StructType:
		// Check if this struct embeds an nvml* type
		for _, field := range t.Fields.List {
			if len(field.Names) == 0 && field.Type != nil {
				// This is an embedded field
				if embeddedType := getTypeString(field.Type); strings.HasPrefix(embeddedType, nvmlPrefix) {
					structTypesWithEmbeddedNvml[typeName] = embeddedType
					break
				}
			}
		}
	}
}

// extractFunctionsFromNvmlFile parses nvml.go and returns all function declarations
func extractFunctionsFromNvmlFile(sourceDir string) ([]*ast.FuncDecl, error) {
	nvmlFilePath := filepath.Join(sourceDir, "nvml.go")
	content, err := os.ReadFile(nvmlFilePath)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", nvmlFilePath, err)
	}
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, nvmlFilePath, content, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parsing %s: %w", nvmlFilePath, err)
	}
	var functions []*ast.FuncDecl
	for _, decl := range node.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			functions = append(functions, funcDecl)
		}
	}
	return functions, nil
}

// removeNvmlPrefix removes the nvml prefix and capitalizes the first letter
func removeNvmlPrefix(name string) string {
	if strings.HasPrefix(name, nvmlPrefix) {
		name = name[len(nvmlPrefix):]
	}
	if len(name) > 0 {
		return strings.ToUpper(name[:1]) + name[1:]
	}
	return name
}

// extractUnversionedNvmlFunctionsFromLibFile extracts nvml* function aliases from the marked block in lib.go
func extractUnversionedNvmlFunctionsFromLibFile(sourceDir string) ([][2]string, error) {
	libFilePath := filepath.Join(sourceDir, "lib.go")
	content, err := os.ReadFile(libFilePath)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", libFilePath, err)
	}
	lines := strings.Split(string(content), "\n")
	var results [][2]string
	inBlock := false
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "// BEGIN_UNVERSIONED_FUNCTIONS" {
			inBlock = true
			continue
		}
		if line == "// END_UNVERSIONED_FUNCTIONS" {
			inBlock = false
			continue
		}
		if !inBlock {
			continue
		}
		if strings.HasPrefix(line, "var nvml") && strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				nvmlFunc := strings.TrimSpace(parts[0][4:])
				aliasedFunc := strings.TrimSpace(parts[1])
				results = append(results, [2]string{nvmlFunc, aliasedFunc})
			}
		}
	}
	return results, nil
}

// getTypeString returns the string representation of a type
func getTypeString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + getTypeString(t.X)
	case *ast.ArrayType:
		return "[]" + getTypeString(t.Elt)
	case *ast.SelectorExpr:
		if x, ok := t.X.(*ast.Ident); ok {
			return x.Name + "." + t.Sel.Name
		}
		return t.Sel.Name
	default:
		return "interface{}"
	}
}

// shouldConvertType checks if a type should be converted (non-exported nvml* type that's not a pointer)
func shouldConvertType(typeName string) bool {
	// Check if it's a non-exported type (starts with lowercase)
	if len(typeName) == 0 || typeName[0] < 'a' || typeName[0] > 'z' {
		return false
	}

	// Check if it starts with "nvml"
	if !strings.HasPrefix(typeName, nvmlPrefix) {
		return false
	}

	// Check if it's not an array
	if strings.HasPrefix(typeName, arrayPrefix) {
		return false
	}

	return true
}

// isInterfaceType checks if a type (without nvml prefix) is an interface type
func isInterfaceType(typeName string) bool {
	return interfaceTypes[typeName]
}

// isStructWithEmbeddedNvmlType checks if a type is a struct that embeds an nvml* type
func isStructWithEmbeddedNvmlType(typeName string) bool {
	return structTypesWithEmbeddedNvml[typeName] != ""
}

// getEmbeddedNvmlType returns the embedded nvml* type for a struct type
func getEmbeddedNvmlType(typeName string) string {
	return structTypesWithEmbeddedNvml[typeName]
}

// isPointerType checks if a type is a pointer to an nvml* type that should be converted
func isPointerType(typeName string) bool {
	if !strings.HasPrefix(typeName, pointerPrefix) {
		return false
	}

	// Extract the base type (remove the * prefix)
	baseType := typeName[len(pointerPrefix):]

	// Check if it's a non-exported type (starts with lowercase)
	if len(baseType) == 0 || baseType[0] < 'a' || baseType[0] > 'z' {
		return false
	}

	// Check if it starts with "nvml"
	if !strings.HasPrefix(baseType, nvmlPrefix) {
		return false
	}

	return true
}

// convertTypeName removes the "nvml" prefix and capitalizes the first letter
func convertTypeName(typeName string) string {
	if strings.HasPrefix(typeName, nvmlPrefix) {
		typeName = typeName[len(nvmlPrefix):]
	}
	if len(typeName) > 0 {
		return strings.ToUpper(typeName[:1]) + typeName[1:]
	}
	return typeName
}

// convertParamName removes the "nvml" prefix and converts to lowercase
func convertParamName(paramName string) string {
	if strings.HasPrefix(paramName, nvmlPrefix) {
		return strings.ToLower(paramName[len(nvmlPrefix):])
	}
	return strings.ToLower(paramName)
}

// formatFieldList returns a string for a parameter or result field
func formatFieldList(field *ast.Field) string {
	var builder strings.Builder
	if len(field.Names) > 0 {
		for i, name := range field.Names {
			if i > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(name.Name)
		}
		builder.WriteString(" ")
	}
	switch t := field.Type.(type) {
	case *ast.Ident:
		builder.WriteString(t.Name)
	case *ast.StarExpr:
		builder.WriteString(pointerPrefix)
		if ident, ok := t.X.(*ast.Ident); ok {
			builder.WriteString(ident.Name)
		}
	case *ast.ArrayType:
		builder.WriteString(arrayPrefix)
		if ident, ok := t.Elt.(*ast.Ident); ok {
			builder.WriteString(ident.Name)
		}
	case *ast.SelectorExpr:
		if x, ok := t.X.(*ast.Ident); ok {
			builder.WriteString(x.Name)
			builder.WriteString(".")
		}
		builder.WriteString(t.Sel.Name)
	default:
		builder.WriteString("interface{}")
	}
	return builder.String()
}
