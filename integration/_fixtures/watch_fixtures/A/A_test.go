package A_test

import (
	. "github.com/onsi/ginkgo/integration/_fixtures/watch_fixtures/A"

	. "github.com/allenluce/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("A", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
