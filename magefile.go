//+build mage

package main

import (
	"github.com/brittonhayes/dnd/internal/gen"
	"github.com/magefile/mage/sh"
	"github.com/princjef/mageutil/bintool"
	"github.com/princjef/mageutil/shellcmd"
	"os"
)

var linter = bintool.Must(bintool.New(
	"golangci-lint{{.BinExt}}",
	"1.23.6",
	"https://github.com/golangci/golangci-lint/releases/download/v{{.Version}}/golangci-lint-{{.Version}}-{{.GOOS}}-{{.GOARCH}}{{.ArchiveExt}}",
))

// Lint lints the library with golangci-lint
func Lint() error {
	if err := linter.Ensure(); err != nil {
		return err
	}
	return linter.Command(`run`).Run()
}

// Test tests packages
func Test() error {
	return shellcmd.Command(`go test -v ./...`).Run()
}

// Test tests the packages with coverage
func Cover() error {
	return shellcmd.Command(`courtney -t="-race" -v github.com/brittonhayes/dnd github.com/brittonhayes/dnd/models`).Run()
}

// Download downloads the library dependencies
func Download() error {
	return sh.Run("go", "mod", "download")
}

// Generate runs all generation tasks
func Generate() error {
	return shellcmd.Command(`go generate .`).Run()
}

func Mocks() error {
	return gen.GenerateMocks()
}

// Create a new service with mage service [resource] [model]
func Service() error {
	args := os.Args
	return gen.GenerateService(args[2], args[3])
}

// Create a new model with mage model [name]
func Model() error {
	args := os.Args
	return gen.GenerateModel(args[2])
}
