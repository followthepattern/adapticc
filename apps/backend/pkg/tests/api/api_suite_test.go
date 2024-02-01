package test_api

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func TestSuite(t *testing.T) {
	setupSqlMock()

	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "API Test Suite")
}

// setupSqlMock overrides regexp comparer
func setupSqlMock() {
	sqlmock.QueryMatcherRegexp = sqlCompareFunc
}
