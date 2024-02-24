package test_integration

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"dagger.io/dagger"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HealthCheck", func() {
	var (
		backendAbsolutePath, _ = filepath.Abs("./../../")
		adapticcBuildPath      = "build/"

		ctx     context.Context
		client  *dagger.Client
		backend *dagger.Service
	)

	buildAdapticc := func(client *dagger.Client) error {
		backendDirectory := client.Host().Directory(backendAbsolutePath)

		golang := client.Container().From("golang:latest")
		golang = golang.WithDirectory("/backend", backendDirectory).WithWorkdir("/backend")

		golang = golang.WithExec([]string{"go", "build", "-o", adapticcBuildPath, "./cmd/adapticc"})

		output := golang.Directory(adapticcBuildPath)
		_, err := output.Export(ctx, adapticcBuildPath)
		return err
	}

	When("build adapticc", func() {
		var err error
		ctx = context.Background()
		client, err = dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
		Expect(err).Should(BeNil())

		err = buildAdapticc(client)
		Expect(err).Should(BeNil())

		testDir := client.Host().Directory(".")

		Context("http tester client", func() {
			FIt("calles healthcheck endpoint", func() {
				backend = client.Container().From("golang:1.21").
					WithDirectory("/backend", testDir).
					WithWorkdir("/backend").
					WithExec([]string{"./build/adapticc"}).
					WithExposedPort(8080).
					AsService()

				// Run application tests
				out, err := client.Container().From("golang:1.21").
					WithServiceBinding("backend", backend).
					WithDirectory("/httpClient", testDir).
					WithWorkdir("/httpClient").
					WithExec([]string{"go", "run", "./tester_client/client.go"}). // execute go test
					Stdout(ctx)

				Expect(err).Should(BeNil())

				fmt.Println(out)
			})
			It("just test", func() {

				// Run application tests
				out, err := client.Container().From("golang:1.21").
					WithServiceBinding("backend", backend).
					WithDirectory("/httpClient", testDir).
					WithWorkdir("/httpClient").
					WithExec([]string{"go", "run", "./tester_client/client.go"}). // execute go test
					Stdout(ctx)

				Expect(err).Should(BeNil())

				fmt.Println(out)
			})
		})
	})

	AfterEach(func() {
		client.Close()
	})

})
