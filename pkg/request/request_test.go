package request

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSuite(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Request Test Suite")
}

var _ = Describe("Request", func() {
	var (
		req RequestHandler[string, struct{}]
		ctx context.Context
	)

	BeforeEach(func() {
		ctx = context.Background()
		req = New[string, struct{}](ctx, "test")
	})

	Context("Wait", func() {
		It("successful finish", func() {
			go func() {
				req.Reply(Success)
			}()

			_, err := req.Wait()
			Expect(err).To(BeNil())
		})

		It("receive error", func() {
			go func() {
				req.ReplyError(errors.New("unexpected error"))
			}()

			_, err := req.Wait()
			Expect(err).To(MatchError("unexpected error"))
		})

		It("times out", func() {
			req.SetOptions(TimeoutOption[string, struct{}](time.Second))
			_, err := req.Wait()
			Expect(err).To(MatchError(requestTimedout))
		})

		It("context gets cancelled", func() {
			ctx, cancel := context.WithCancel(ctx)

			req = New(ctx, "test", TimeoutOption[string, struct{}](time.Second*10))
			go func() {
				cancel()
			}()

			_, err := req.Wait()
			Expect(err).To(MatchError(contextCancelled))
		})
	})
})
