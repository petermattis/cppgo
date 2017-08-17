// Copyright 2017 Peter Mattis.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//go:generate go-bindata cppgo.go

var dryRun = flag.Bool("n", false, "dry run")

func usage() {
	fmt.Fprintf(os.Stderr, "usage: cppgo [type...]\n")
	fmt.Fprintf(os.Stderr, "options:\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if len(flag.Args()) == 0 {
		usage()
		return
	}

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	types := map[string]map[string]bool{}
	for _, arg := range flag.Args() {
		dir := filepath.Dir(arg)
		base := filepath.Base(arg)
		var path string
		var typeName string

		if parts := strings.Split(base, "."); len(parts) != 2 {
			path = "."
			typeName = base
		} else {
			path = filepath.Join(dir, parts[0])
			typeName = parts[1]
		}

		pkg, err := build.Import(path, pwd, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		pkgTypes := types[pkg.ImportPath]
		if pkgTypes == nil {
			pkgTypes = map[string]bool{}
			types[pkg.ImportPath] = pkgTypes
		}

		typeName = pkg.Name + "." + typeName
		pkgTypes[typeName] = true

		if *dryRun {
			fmt.Printf("%s.%s\n", pkg.ImportPath, typeName)
		}
	}

	var imports bytes.Buffer
	var walk bytes.Buffer
	for i, pkgTypes := range types {
		fmt.Fprintf(&imports, "\t\"%s\"\n", i)
		for t := range pkgTypes {
			fmt.Fprintf(&walk, "\twalk(reflect.TypeOf(&%s{}))\n", t)
		}
	}

	if *dryRun {
		return
	}

	dir, err := ioutil.TempDir("", "cppgo")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(dir)

	f, err := os.Create(filepath.Join(dir, "cppgo_gen.go"))
	if err != nil {
		fmt.Println(err)
		return
	}

	template, err := cppgoGoBytes()
	if err != nil {
		fmt.Println(err)
		return
	}
	template = bytes.Replace(template, []byte("\t// IMPORTS\n"), imports.Bytes(), 1)
	template = bytes.Replace(template, []byte("\t// WALK\n"), walk.Bytes(), 1)
	f.Write(template)

	cmd := exec.Command("go", "run", "cppgo_gen.go")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}
