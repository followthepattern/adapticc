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

		ctx     context.Context
		client  *dagger.Client
		backend *dagger.Service
	)

	BeforeEach(func() {
		var err error
		ctx = context.Background()
		client, err = dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
		backendDirectory := client.Host().Directory(backendAbsolutePath)

		// Database service used for application tests
		backend = client.Container().From("golang:1.21").
			WithDirectory("/backend", backendDirectory).
			WithWorkdir("/backend").
			WithExec([]string{"go", "run", "./cmd/adapticc"}).
			WithExposedPort(8080).
			AsService()

		Expect(err).Should(BeNil())
	})

	When("backend service successfully runs", func() {
		Context("http tester client", func() {
			FIt("calles healthcheck endpoint", func() {

				testFolder := client.Host().Directory(".")

				// Run application tests
				out, err := client.Container().From("golang:1.21").
					WithServiceBinding("backend", backend).
					WithDirectory("/httpClient", testFolder).
					WithWorkdir("/httpClient").
					WithExec([]string{"go", "run", "./tester_client/client.go"}). // execute go test
					Stdout(ctx)

				Expect(err).Should(BeNil())

				fmt.Println(out)
			})
			FIt("just test", func() {
				Expect(1).To(Equal(1))
			})
		})
	})

	AfterEach(func() {
		client.Close()
	})

})
