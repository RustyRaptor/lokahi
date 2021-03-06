// +build mage

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

const (
	// the version of Go to use in docker images.
	goVersion = "1.9.3"
)

// Generate runs all relevant code generation tasks.
func Generate() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shouldWork(ctx, nil, wd, "statik", "-src", "./public", "-f")
	shouldWork(ctx, nil, filepath.Join(wd, "rpc", "lokahi"), "sh", "./regen.sh")
	shouldWork(ctx, nil, filepath.Join(wd, "rpc", "lokahiadmin"), "sh", "./regen.sh")

	fmt.Println("reran code generation")
}

// Build builds the command code into binaries in ./bin.
func Build() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	os.Mkdir("bin", 0777)

	outd := filepath.Join(wd, "bin")
	cmds := []string{"lokahid", "lokahictl", "sample_hook"}

	for _, c := range cmds {
		shouldWork(ctx, nil, outd, "go", "build", "../cmd/"+c)
		fmt.Println("built ./bin/" + c)
	}
}

// Dep reruns `dep`.
func Dep() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shouldWork(ctx, nil, wd, "dep", "ensure", "-update")
	shouldWork(ctx, nil, wd, "dep", "prune")
}

// Docker creates the docker image xena/lokahi with the lokahi server.
func Docker() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shouldWork(ctx, nil, wd, "docker", "pull", "xena/go-mini:"+goVersion)
	shouldWork(ctx, nil, wd, "docker", "build", "-t", "xena/lokahi", ".")
}
