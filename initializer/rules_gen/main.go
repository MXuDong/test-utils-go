package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"
)

type FunctionValue struct {
	Type          string
	FunctionType  string
	InnerFunction string
}

var TypeMapper map[string]FunctionValue = map[string]FunctionValue{
	"Uint":    {Type: "Uint", InnerFunction: "FixedUintBaseRule", FunctionType: "uint64"},
	"Uint8":   {Type: "Uint8", InnerFunction: "FixedUintBaseRule", FunctionType: "uint64"},
	"Uint16":  {Type: "Uint16", InnerFunction: "FixedUintBaseRule", FunctionType: "uint64"},
	"Uint32":  {Type: "Uint32", InnerFunction: "FixedUintBaseRule", FunctionType: "uint64"},
	"Uint64":  {Type: "Uint64", InnerFunction: "FixedUintBaseRule", FunctionType: "uint64"},
	"Int":     {Type: "Int", InnerFunction: "FixedIntBaseRule", FunctionType: "int64"},
	"Int8":    {Type: "Int8", InnerFunction: "FixedIntBaseRule", FunctionType: "int64"},
	"Int16":   {Type: "Int16", InnerFunction: "FixedIntBaseRule", FunctionType: "int64"},
	"Int32":   {Type: "Int32", InnerFunction: "FixedIntBaseRule", FunctionType: "int64"},
	"Int64":   {Type: "Int64", InnerFunction: "FixedIntBaseRule", FunctionType: "int64"},
	"Float32": {Type: "Float32", InnerFunction: "FixedFloatBaseRule", FunctionType: "float64"},
	"Float64": {Type: "Float64", InnerFunction: "FixedFloatBaseRule", FunctionType: "float64"},
}

// run with: go generate github.com/MXuDong/test-utils-go/initializer/rules_gen
//go:generate go run $PWD/initializer/rules_gen/main.go --template $PWD/initializer/rules_gen/rules.go.temp --output $PWD/initializer/rules.gen.go
func main() {
	templateFile := flag.String("template", "", "Template file")

	outputFile := flag.String("output", "", "Output file. Leave blank to go to stdout")

	flag.Parse()

	if len(*templateFile) == 0 {
		return
	}
	if len(*outputFile) == 0 {
		return
	}

	tmpl := template.Must(template.ParseFiles(*templateFile))

	var buffer bytes.Buffer

	if err := tmpl.Execute(&buffer, TypeMapper); err != nil {
		log.Fatal(fmt.Errorf("template: %v", err))
	}

	// Format source code.
	out, err := format.Source(buffer.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	// Output
	if outputFile == nil || *outputFile == "" {
		fmt.Println(string(out))
	} else if err := os.WriteFile(*outputFile, out, 0o644); err != nil {
		panic(err)
	}
}
