package data_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestData(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Data Suite")
}

var cleaner func()

var _ = BeforeSuite(func() {
})

var _ = AfterSuite(func() {
	cleaner()
})
