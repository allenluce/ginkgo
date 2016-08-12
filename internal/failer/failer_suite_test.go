package failer_test

import (
	. "github.com/allenluce/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFailer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Failer Suite")
}
