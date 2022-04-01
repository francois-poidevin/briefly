//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	Default = Build
	goFiles = getGoFiles()
)

// Runs go mod download and then installs the binary.
func Build() {
	banner := figure.NewFigure("Briefly", "", true)
	banner.Print()

	fmt.Println("")

	color.Red("# Pre Build ------------------------------------------------------------")
	mg.SerialDeps(Go.Generate, Go.Tidy, Go.Deps, Go.Format)

	fmt.Println("")
	color.Red("# Build Artefact ----------------------------------------------------------------")
	mg.Deps(Bin.Briefly)
}

type Go mg.Namespace

// Generate go code
func (Go) Generate() error {
	color.Cyan("## Generate code")
	mg.SerialDeps(Gen.Protobuf)
	return nil
}

// Tidy add/remove depenedencies.
func (Go) Tidy() error {
	color.Cyan("## Cleaning go modules")
	return sh.RunV("go", "mod", "tidy", "-v")
}

// Deps install dependency tools.
func (Go) Deps() error {
	color.Cyan("## Vendoring dependencies")
	return sh.RunV("go", "mod", "vendor")
}

// Format runs gofmt on everything
func (Go) Format() error {
	color.Cyan("## Format everything")
	args := []string{"-w"}
	args = append(args, goFiles...)
	return sh.RunV("gofumpt", args...)
}

type Gen mg.Namespace

// Generate protobuf
func (Gen) Protobuf() error {
	color.Blue("### Protobuf")
	return sh.RunV("prototool", "all", "pkg/protocol")
}

type Bin mg.Namespace

func (Bin) Briefly() error {
	return goBuild("briefly")
}

func goBuild(packageName string) error {
	color.Green(" > Building %s\n", packageName)

	return sh.RunV("go", "build", "-o", fmt.Sprintf("bin/%s", packageName))
}

func getGoFiles() []string {
	var goFiles []string

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, "vendor/") {
			return filepath.SkipDir
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		goFiles = append(goFiles, path)
		return nil
	})

	return goFiles
}
