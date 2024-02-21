package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	if err := build(context.Background()); err != nil {
		fmt.Println(err)
	}
}

func build(ctx context.Context) error {
	fmt.Println("Building with Dagger")

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		return err
	}
	defer client.Close()

	// get reference to the local project
	src := client.Host().Directory(".")

	// get `golang` image
	golang := client.Container().From("golang:latest")

	// mount cloned repository into `golang` image
	golang = golang.WithDirectory("/src", src).WithWorkdir("/src")

	// define the application build command
	path := "build/"
	golang = golang.WithExposedPort(8080)
	golang = golang.WithExec([]string{"go", "run", "./cmd/adapticc"})

	// get reference to build output directory in container
	// output := golang.Directory(path)

	// write contents of container build/ directory to the host
	_, err = golang.Export(ctx, path)

	return err
}
