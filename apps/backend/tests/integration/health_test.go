package test_integration

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"dagger.io/dagger"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HealthCheck", Ordered, func() {
	var (
		backendAbsolutePath, _ = filepath.Abs("./../../")
		testDir                *dagger.Directory

		ctx     context.Context
		client  *dagger.Client
		backend *dagger.Service
	)

	BeforeAll(func() {
		var err error
		ctx = context.Background()

		client, err = dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
		Expect(err).To(BeNil())

		backendDirectory := client.Host().Directory(backendAbsolutePath)

		builder := client.Container().From("golang:latest")
		builder = builder.WithDirectory("/backend", backendDirectory).WithWorkdir("/backend")

		outputPath := "out/"
		builder = builder.WithExec([]string{"go", "build", "-o", outputPath, "./cmd/adapticc"})

		output := builder.Directory(outputPath)
		_, err = output.Export(ctx, filepath.Join(backendAbsolutePath, outputPath))
		Expect(err).To(BeNil())

		backend = client.Container().From("golang:1.21").
			WithDirectory("/backend", backendDirectory).
			WithWorkdir("/backend").
			WithExec([]string{"./out/adapticc"}).
			WithExposedPort(8080).
			AsService()

	})

	Context("healtchcheck", func() {
		BeforeEach(func() {
			testDir = client.Host().Directory(".")
		})

		It("calles healthcheck endpoint", func() {
			out, err := client.Container().From("golang:1.21").
				WithServiceBinding("backend", backend).
				WithDirectory("/httpClient", testDir).
				WithWorkdir("/httpClient").
				WithExec([]string{"go", "run", "./http_tester/client.go", "GET", "http://backend:8080/healthcheck"}).
				Stdout(ctx)

			Expect(err).Should(BeNil())

			splits := strings.Split(out, "\n")
			Expect(splits).To(HaveLen(2))

			Expect(splits[0]).To(Equal("0.0.0"))
		})
	})

	AfterAll(func() {
		client.Close()
	})

})
