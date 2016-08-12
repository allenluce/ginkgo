package containernode_test

import (
	. "github.com/allenluce/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestContainernode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Containernode Suite")
}
