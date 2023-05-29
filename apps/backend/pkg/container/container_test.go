package container

import (
	"context"
	"errors"

	"github.com/followthepattern/adapticc/pkg/config"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Container", func() {
	var container *Container

	BeforeEach(func() {
		container = New(context.Background(), config.Config{}, nil, nil)
	})

	Context("Working with struct", func() {
		type TestStruct struct {
			Name string
		}
		var (
			testStructName       = "test1"
			errConstructorString = "err in constructor"
			Constructor          = func(cont *Container) (*TestStruct, error) {
				return &TestStruct{Name: testStructName}, nil
			}
			errConstructor = func(cont *Container) (*TestStruct, error) {
				return nil, errors.New(errConstructorString)
			}
		)

		It("registers TestStruct pointer", func() {
			var value *TestStruct
			dependencyKey := getKey(value)
			err := Register(container, Constructor)
			Expect(err).To(BeNil())

			obj, found := container.dependencies[dependencyKey]
			Expect(found).To(BeTrue())

			testStruct, ok := obj.(*TestStruct)
			Expect(ok).To(BeTrue())

			Expect(testStruct.Name).To(Equal(testStructName))
		})

		It("fails registering TestStruct pointer", func() {
			err := Register(container, errConstructor)
			Expect(err).To(MatchError(errConstructorString))
		})

		It("resolves TestStruct value", func() {
			err := Register(container, Constructor)
			Expect(err).To(BeNil())

			testStruct, err := Resolve[TestStruct](container)
			Expect(err).To(BeNil())
			Expect(testStruct).NotTo(BeNil())
		})

		It("fails resolve TestStruct pointer", func() {
			err := Register(container, Constructor)
			Expect(err).To(BeNil())

			testStruct, err := Resolve[*TestStruct](container)
			Expect(err).To(MatchError("can't resolve *container.TestStruct"))
			Expect(testStruct).To(BeNil())
		})
	})

	Context("Working with channels", func() {
		var (
			testChan    chan string
			Constructor = func(cont *Container) (*chan string, error) {
				result := make(chan string)
				return &result, nil
			}
		)

		It("registers *chan string", func() {
			dependencyKey := getKey(testChan)
			err := Register(container, Constructor)
			Expect(err).To(BeNil())

			obj, found := container.dependencies[dependencyKey]
			Expect(found).To(BeTrue())

			_, ok := obj.(*chan string)
			Expect(ok).To(BeTrue())
		})

		It("resolve *chan string", func() {
			err := Register(container, Constructor)
			Expect(err).To(BeNil())

			testChanResult, err := Resolve[chan string](container)
			Expect(err).To(BeNil())
			Expect(testChanResult).NotTo(BeNil())
		})
	})
})
