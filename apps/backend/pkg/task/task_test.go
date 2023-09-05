package task

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSuite(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Task Test Suite")
}

var _ = Describe("Task", func() {
	var (
		action TaskAction[string] = func() (string, error) {
			return "hello", nil
		}
	)

	Context("Wait", func() {
		It("successful finish", func() {
			value, err := New(action).Run().Wait()

			Expect(err).To(BeNil())
			Expect(value).NotTo(BeNil())
			Expect(*value).To(Equal("hello"))
		})

	})
})
