package remote_test

import (
	. "github.com/allenluce/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRemote(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Remote Spec Forwarding Suite")
}
