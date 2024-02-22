package test_integration

import (
	"context"
	"os"

	"dagger.io/dagger"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Health", func() {
	var (
		ctx    context.Context
		client *dagger.Client
	)
	BeforeEach(func() {
		var err error
		ctx = context.Background()
		client, err = dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
		Expect(err).Should(BeNil())
	})

	Context("Register", func() {

		FIt("Success", func() {
			var err error
			src := client.Host().Directory(".")

			// get `golang` image
			golang := client.Container().From("golang:latest")

			// mount cloned repository into `golang` image
			golang = golang.WithDirectory("/apps/backend", src).WithWorkdir("/apps/backend")

			// define the application build command
			outputPath := "build/"
			golang = golang.WithExec([]string{"go", "build", "./cmd/adapticc", "-o", outputPath})

			// get reference to build output directory in container
			output := golang.Directory(outputPath)

			// write contents of container build/ directory to the host
			_, err = output.Export(ctx, outputPath)
			Expect(err).Should(BeNil())
		})
	})

	AfterEach(func() {
		client.Close()
	})

})
